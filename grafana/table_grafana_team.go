package grafana

import (
	"context"
	"strconv"

	"github.com/grafana/grafana-openapi-client-go/client/teams"
	"github.com/grafana/grafana-openapi-client-go/models"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
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

	// Use the teams API to search for teams with pagination
	params := teams.NewSearchTeamsParams()
	if d.EqualsQuals["query"] != nil {
		query := d.EqualsQuals["query"].GetStringValue()
		params.WithQuery(&query)
	}

	// Set pagination parameters
	page := int64(1)      // Start with first page
	perpage := int64(100) // Default page size
	params.WithPage(&page)
	params.WithPerpage(&perpage)

	// Continue fetching pages until no more results
	for {
		result, err := conn.client.Teams.SearchTeams(params)
		if err != nil {
			plugin.Logger(ctx).Error("grafana_team.listTeam", "query_error", err)
			return nil, err
		}

		// If no results, break the loop
		if len(result.Payload.Teams) == 0 {
			break
		}

		// Stream the results from this page
		for _, team := range result.Payload.Teams {
			d.StreamListItem(ctx, team)
		}

		// If we got fewer results than the perpage limit, we've reached the end
		if len(result.Payload.Teams) < int(perpage) {
			break
		}

		// Move to next page
		page++
		params.WithPage(&page)
	}

	return nil, nil
}

func getTeam(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("grafana_team.getTeam", "connection_error", err)
		return nil, err
	}
	id := d.EqualsQuals["id"].GetInt64Value()

	// Convert int64 to string for the API call
	idStr := strconv.FormatInt(id, 10)

	// Use the teams API to get team by ID
	result, err := conn.client.Teams.GetTeamByID(idStr)
	if err != nil {
		plugin.Logger(ctx).Error("grafana_team.getTeam", "query_error", err, "id", id)
		return nil, err
	}

	return result.Payload, nil
}

func getTeamPreferences(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("grafana_team.getTeam", "connection_error", err)
		return nil, err
	}

	team := h.Item.(*models.TeamDTO)
	idStr := strconv.FormatInt(team.ID, 10)

	// Use the teams API to get team preferences
	result, err := conn.client.Teams.GetTeamPreferences(idStr)
	if err != nil {
		plugin.Logger(ctx).Error("grafana_team.getTeamPreferences", "query_error", err, "team", team)
		return nil, err
	}

	return result.Payload, nil
}
