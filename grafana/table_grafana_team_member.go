package grafana

import (
	"context"
	"strconv"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
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
	tid := d.EqualsQuals["team_id"].GetInt64Value()

	// Convert int64 to string for the API call
	tidStr := strconv.FormatInt(tid, 10)

	// Use the teams API to get team members
	result, err := conn.client.Teams.GetTeamMembers(tidStr)
	if err != nil {
		if isNotFoundError(err) {
			return nil, nil
		}
		plugin.Logger(ctx).Error("grafana_team.listTeamMember", "query_error", err, "team_id", tid)
		return nil, err
	}


	for _, member := range result.Payload {
		d.StreamListItem(ctx, member)
	}
	return nil, nil
}
