package grafana

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"strconv"
	"strings"

	gapi "github.com/grafana/grafana-api-golang-client"
	//smapi "github.com/grafana/synthetic-monitoring-api-go-client"
	"github.com/hashicorp/go-cleanhttp"

	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
)

// Client manages the connection configuration.
type Client struct {
	gapi *gapi.Client
	// Not used yet, but reserved for similar use to Terraform provider
	//smapi *smapi.Client
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

	var orgID int
	var err error
	if orgIDString != "" {
		orgID, err = strconv.Atoi(orgIDString)
		if err != nil {
			return nil, fmt.Errorf("org_id is not a valid integer")
		}
	}

	// Prefer config settings
	grafanaConfig := GetConfig(d.Connection)
	if &grafanaConfig != nil {
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
			orgID = *grafanaConfig.OrgID
		}
		if grafanaConfig.TLSCert != nil {
			tlsCert = *grafanaConfig.TLSCert
		}
		if grafanaConfig.TLSKey != nil {
			tlsKey = *grafanaConfig.TLSKey
		}
	}

	// Error if the minimum config is not set
	if gurl == "" {
		return nil, errors.New("url must be configured")
	}
	if auth == "" {
		return nil, errors.New("auth must be configured")
	}

	conn := &Client{}
	cli := cleanhttp.DefaultClient()
	transport := cleanhttp.DefaultTransport()
	transport.TLSClientConfig = &tls.Config{}

	if caCert != "" {
		ca, err := ioutil.ReadFile(caCert)
		if err != nil {
			return nil, fmt.Errorf("ca_cert error: %s", err.Error())
		}
		pool := x509.NewCertPool()
		pool.AppendCertsFromPEM(ca)
		transport.TLSClientConfig.RootCAs = pool
	}
	if tlsKey != "" && tlsCert != "" {
		cert, err := tls.LoadX509KeyPair(tlsCert, tlsKey)
		if err != nil {
			return nil, fmt.Errorf("tls_key and tls_cert error: %s", err.Error())
		}
		transport.TLSClientConfig.Certificates = []tls.Certificate{cert}
	}
	if insecureSkipVerify {
		transport.TLSClientConfig.InsecureSkipVerify = true
	}

	//cli.Transport = logging.NewTransport("Grafana", transport)
	cfg := gapi.Config{
		Client: cli,
	}
	// If the org is set, then put it on the config
	if orgID != 0 {
		cfg.OrgID = int64(orgID)
	}
	authParts := strings.SplitN(auth, ":", 2)
	if len(authParts) == 2 {
		cfg.BasicAuth = url.UserPassword(authParts[0], authParts[1])
	} else {
		cfg.APIKey = authParts[0]
	}
	gclient, err := gapi.New(gurl, cfg)
	if err != nil {
		return nil, err
	}

	conn.gapi = gclient

	// Save to cache
	d.ConnectionManager.Cache.Set(cacheKey, conn)

	return conn, nil
}

func isNotFoundError(err error) bool {
	return strings.HasPrefix(err.Error(), "status: 404")
}
