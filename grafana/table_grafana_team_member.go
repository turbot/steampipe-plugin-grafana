package grafana

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v3/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
)

func tableGrafanaTeamMember(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "grafana_team_member",
		Description: "Members for a given Team in the Grafana installation.",
		List: &plugin.ListConfig{
			KeyColumns: plugin.SingleColumn("team_id"),
			Hydrate:    listTeamMember,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "team_id", Type: proto.ColumnType_INT, Description: "ID of the team."},
			{Name: "user_id", Type: proto.ColumnType_INT, Description: "ID of the user."},
			{Name: "email", Type: proto.ColumnType_STRING, Description: "Email of the user."},
			// Other columns
			{Name: "avatar_url", Type: proto.ColumnType_STRING, Description: "URL of the avatar for the user."},
			{Name: "login", Type: proto.ColumnType_STRING, Description: "Login of the user."},
			{Name: "org_id", Type: proto.ColumnType_INT, Description: "ID of the org."},
		},
	}
}

func listTeamMember(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("grafana_team.listTeamMember", "connection_error", err)
		return nil, err
	}
	tid := d.KeyColumnQuals["team_id"].GetInt64Value()
	items, err := conn.gapi.TeamMembers(tid)
	if err != nil {
		if isNotFoundError(err) {
			return nil, nil
		}
		plugin.Logger(ctx).Error("grafana_team.listTeamMember", "query_error", err, "team_id", tid)
		return nil, err
	}
	for _, i := range items {
		d.StreamListItem(ctx, i)
	}
	return nil, nil
}
