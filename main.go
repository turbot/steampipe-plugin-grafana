package main

import (
	"github.com/turbot/steampipe-plugin-grafana/grafana"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{PluginFunc: grafana.Plugin})
}
