package grafana

import (
	"context"

	gapi "github.com/grafana/grafana-api-golang-client"

	"github.com/turbot/steampipe-plugin-sdk/v3/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin/transform"
)

func tableGrafanaTeam(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "grafana_team",
		Description: "Teams in the Grafana installation.",
		List: &plugin.ListConfig{
			KeyColumns: plugin.OptionalColumns([]string{"query"}),
			Hydrate:    listTeam,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getTeam,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name of the team."},
			{Name: "email", Type: proto.ColumnType_STRING, Description: "Email of the team."},
			{Name: "member_count", Type: proto.ColumnType_INT, Transform: transform.FromField("MemberCount"), Description: "Number of members in the team."},
			// Other columns
			{Name: "avatar_url", Type: proto.ColumnType_STRING, Description: "URL of the avatar for the team."},
			{Name: "home_dashboard_id", Type: proto.ColumnType_INT, Hydrate: getTeamPreferences, Description: "Home dashboard for the team."},
			{Name: "id", Type: proto.ColumnType_INT, Description: "Unique identifier for the team."},
			{Name: "org_id", Type: proto.ColumnType_INT, Description: "Org the team is a member of."},
			{Name: "query", Type: proto.ColumnType_STRING, Transform: transform.FromQual("query"), Description: "Query term for searching teams."},
			{Name: "theme", Type: proto.ColumnType_STRING, Hydrate: getTeamPreferences, Description: "UI theme for the team."},
			{Name: "timezone", Type: proto.ColumnType_STRING, Hydrate: getTeamPreferences, Description: "Timezone for the team."},
		},
	}
}

func listTeam(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("grafana_team.listTeam", "connection_error", err)
		return nil, err
	}
	query := ""
	if d.KeyColumnQuals["query"] != nil {
		query = d.KeyColumnQuals["query"].GetStringValue()
	}
	result, err := conn.gapi.SearchTeam(query)
	if err != nil {
		plugin.Logger(ctx).Error("grafana_team.listTeam", "query_error", err, "query", query)
		return nil, err
	}
	for _, i := range result.Teams {
		d.StreamListItem(ctx, i)
	}
	return nil, nil
}

func getTeam(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("grafana_team.getTeam", "connection_error", err)
		return nil, err
	}
	id := d.KeyColumnQuals["id"].GetInt64Value()
	item, err := conn.gapi.Team(id)
	if err != nil {
		plugin.Logger(ctx).Error("grafana_team.getTeam", "query_error", err, "id", id)
		return nil, err
	}
	return item, nil
}

func getTeamPreferences(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("grafana_team.getTeam", "connection_error", err)
		return nil, err
	}
	team := h.Item.(*gapi.Team)
	item, err := conn.gapi.TeamPreferences(team.ID)
	if err != nil {
		plugin.Logger(ctx).Error("grafana_team.getTeamPreferences", "query_error", err, "team", team)
		return nil, err
	}
	return item, nil
}
