package grafana

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

type grafanaConfig struct {
	URL                *string `hcl:"url"`
	Auth               *string `hcl:"auth"`
	CaCert             *string `hcl:"ca_cert"`
	InsecureSkipVerify *bool   `hcl:"insecure_skip_verify"`
	OrgID              *int    `hcl:"org_id"`
	TLSCert            *string `hcl:"tls_cert"`
	TLSKey             *string `hcl:"tls_key"`
}

func ConfigInstance() interface{} {
	return &grafanaConfig{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) grafanaConfig {
	if connection == nil || connection.Config == nil {
		return grafanaConfig{}
	}
	config, _ := connection.Config.(grafanaConfig)
	return config
}
