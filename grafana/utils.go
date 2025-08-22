package grafana

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/go-openapi/strfmt"
	goapi "github.com/grafana/grafana-openapi-client-go/client"
	"github.com/hashicorp/go-cleanhttp"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

// Client manages the connection configuration.
type Client struct {
	client *goapi.GrafanaHTTPAPI
}

func connect(_ context.Context, d *plugin.QueryData) (*Client, error) {

	// Load connection from cache, which preserves throttling protection etc
	cacheKey := "grafana"
	if cachedData, ok := d.ConnectionManager.Cache.Get(cacheKey); ok {
		return cachedData.(*Client), nil
	}

	// Default to the env var settings
	gurl := os.Getenv("GRAFANA_URL")
	auth := os.Getenv("GRAFANA_AUTH")
	caCert := os.Getenv("GRAFANA_CA_CERT")
	insecureSkipVerify := strings.ToLower(os.Getenv("GRAFANA_INSECURE_SKIP_VERIFY")) == "true"
	orgIDString := os.Getenv("GRAFANA_ORG_ID")
	tlsCert := os.Getenv("GRAFANA_TLS_CERT")
	tlsKey := os.Getenv("GRAFANA_TLS_KEY")

	var orgID int64
	var err error
	if orgIDString != "" {
		orgIDInt, err := strconv.Atoi(orgIDString)
		if err != nil {
			return nil, fmt.Errorf("org_id is not a valid integer")
		}
		orgID = int64(orgIDInt)
	}

	// Prefer config settings
	grafanaConfig := GetConfig(d.Connection)
	if grafanaConfig.URL != nil {
		gurl = *grafanaConfig.URL
	}
	if grafanaConfig.Auth != nil {
		auth = *grafanaConfig.Auth
	}
	if grafanaConfig.CaCert != nil {
		caCert = *grafanaConfig.CaCert
	}
	if grafanaConfig.InsecureSkipVerify != nil {
		insecureSkipVerify = *grafanaConfig.InsecureSkipVerify
	}
	if grafanaConfig.OrgID != nil {
		orgID = int64(*grafanaConfig.OrgID)
	}
	if grafanaConfig.TLSCert != nil {
		tlsCert = *grafanaConfig.TLSCert
	}
	if grafanaConfig.TLSKey != nil {
		tlsKey = *grafanaConfig.TLSKey
	}

	// Error if the minimum config is not set
	if gurl == "" {
		return nil, errors.New("url must be configured")
	}
	if auth == "" {
		return nil, errors.New("auth must be configured")
	}

	// Parse the URL to extract host and scheme
	parsedURL, err := url.Parse(gurl)
	if err != nil {
		return nil, fmt.Errorf("invalid URL: %s", err.Error())
	}

	// Configure TLS
	tlsConfig := &tls.Config{}
	if caCert != "" {
		ca, err := os.ReadFile(caCert)
		if err != nil {
			return nil, fmt.Errorf("ca_cert error: %s", err.Error())
		}
		pool := x509.NewCertPool()
		pool.AppendCertsFromPEM(ca)
		tlsConfig.RootCAs = pool
	}
	if tlsKey != "" && tlsCert != "" {
		cert, err := tls.LoadX509KeyPair(tlsCert, tlsKey)
		if err != nil {
			return nil, fmt.Errorf("tls_key and tls_cert error: %s", err.Error())
		}
		tlsConfig.Certificates = []tls.Certificate{cert}
	}
	if insecureSkipVerify {
		tlsConfig.InsecureSkipVerify = true
	}

	// Create HTTP client with TLS config
	httpClient := cleanhttp.DefaultClient()
	transport := cleanhttp.DefaultTransport()
	transport.TLSClientConfig = tlsConfig
	httpClient.Transport = transport

	// Configure the Grafana OpenAPI client
	cfg := goapi.DefaultTransportConfig()
	cfg.Host = parsedURL.Host
	cfg.BasePath = "/api"
	if parsedURL.Scheme == "https" {
		cfg.Schemes = []string{"https"}
	} else {
		cfg.Schemes = []string{"http"}
	}
	cfg.Client = httpClient
	cfg.NumRetries = 3
	cfg.RetryTimeout = 5 * time.Second

	// Configure authentication
	authParts := strings.SplitN(auth, ":", 2)
	if len(authParts) == 2 {
		// Basic auth
		cfg.BasicAuth = url.UserPassword(authParts[0], authParts[1])
		if orgID != 0 {
			cfg.OrgID = orgID
		}
	} else {
		// API key
		cfg.APIKey = authParts[0]
	}

	// Create the client
	client := goapi.NewHTTPClientWithConfig(strfmt.Default, cfg)

	conn := &Client{
		client: client,
	}

	// Save to cache
	d.ConnectionManager.Cache.Set(cacheKey, conn)

	return conn, nil
}

func isNotFoundError(err error) bool {
	return strings.HasPrefix(err.Error(), "status: 404")
}
