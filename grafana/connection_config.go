package grafana

import (
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin/schema"
)

type grafanaConfig struct {
	URL                *string `cty:"url"`
	Auth               *string `cty:"auth"`
	CaCert             *string `cty:"ca_cert"`
	InsecureSkipVerify *bool   `cty:"insecure_skip_verify"`
	OrgID              *int    `cty:"org_id"`
	TLSCert            *string `cty:"tls_cert"`
	TLSKey             *string `cty:"tls_key"`
}

var ConfigSchema = map[string]*schema.Attribute{
	"url": {
		Type: schema.TypeString,
	},
	"auth": {
		Type: schema.TypeString,
	},
	"ca_cert": {
		Type: schema.TypeString,
	},
	"insecure_skip_verify": {
		Type: schema.TypeBool,
	},
	"org_id": {
		Type: schema.TypeInt,
	},
	"tls_cert": {
		Type: schema.TypeString,
	},
	"tls_key": {
		Type: schema.TypeString,
	},
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
