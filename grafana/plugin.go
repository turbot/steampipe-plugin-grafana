package grafana

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name: "steampipe-plugin-grafana",
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
		},
		DefaultTransform: transform.FromGo().NullIfZero(),
		DefaultGetConfig: &plugin.GetConfig{
			ShouldIgnoreError: isNotFoundError,
		},
		TableMap: map[string]*plugin.Table{
			"grafana_alert_rule":        tableGrafanaAlertRule(ctx),
			"grafana_dashboard":         tableGrafanaDashboard(ctx),
			"grafana_datasource":        tableGrafanaDatasource(ctx),
			"grafana_folder":            tableGrafanaFolder(ctx),
			"grafana_folder_permission": tableGrafanaFolderPermission(ctx),
			"grafana_org":               tableGrafanaOrg(ctx),
			"grafana_team":              tableGrafanaTeam(ctx),
			"grafana_team_member":       tableGrafanaTeamMember(ctx),
			"grafana_user":              tableGrafanaUser(ctx),
		},
	}
	return p
}
