package grafana

import (
	"context"

	"github.com/grafana/grafana-openapi-client-go/client/users"
	"github.com/grafana/grafana-openapi-client-go/models"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableGrafanaUser(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "grafana_user",
		Description: "Users in the Grafana installation.",
		List: &plugin.ListConfig{
			Hydrate: listUser,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AnyColumn([]string{"id", "email"}),
			Hydrate:    getUser,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_INT, Description: "Unique identifier for the user."},
			{Name: "email", Type: proto.ColumnType_STRING, Description: "Email of the user."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name of the user."},
			{Name: "login", Type: proto.ColumnType_STRING, Description: "Login name of the user."},
			{Name: "theme", Type: proto.ColumnType_STRING, Hydrate: getUser, Description: "UI theme preferred by the user."},
			{Name: "org_id", Type: proto.ColumnType_INT, Hydrate: getUser, Description: "Org the user is a member of."},
			{Name: "is_admin", Type: proto.ColumnType_BOOL, Description: "True if the user is an administrator."},
			{Name: "is_disabled", Type: proto.ColumnType_BOOL, Description: "True if the user has been disabled."},
			{Name: "is_external", Type: proto.ColumnType_BOOL, Hydrate: getUser, Description: "True if the user is external to the system."},
			{Name: "updated_at", Type: proto.ColumnType_TIMESTAMP, Hydrate: getUser, Description: "Last time the user was updated."},
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Hydrate: getUser, Description: "Time when the user was created."},
			{Name: "last_seen_at", Type: proto.ColumnType_TIMESTAMP, Description: "Last time the user logged in."},
			{Name: "last_seen_at_age", Type: proto.ColumnType_STRING, Description: "Display string for when the user last logged in, e.g. 2m."},
			{Name: "auth_labels", Type: proto.ColumnType_JSON, Description: "Auth labels for the user, e.g. OAuth."},
			{Name: "avatar_url", Type: proto.ColumnType_STRING, Description: "URL of the avatar for the user."},
		},
	}
}

func listUser(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("grafana_user.listUser", "connection_error", err)
		return nil, err
	}

	// Use the users API to search for users with pagination
	params := users.NewSearchUsersParams()
	page := int64(1)      // Start with first page
	perpage := int64(100) // Default page size

	params.WithPage(&page)
	params.WithPerpage(&perpage)

	// Continue fetching pages until no more results
	for {
		result, err := conn.client.Users.SearchUsers(params)
		if err != nil {
			plugin.Logger(ctx).Error("grafana_user.listUser", "query_error", err)
			return nil, err
		}

		// If no results, break the loop
		if len(result.Payload) == 0 {
			break
		}

		// Stream the results from this page
		for _, user := range result.Payload {
			d.StreamListItem(ctx, user)
		}

		// If we got fewer results than the perpage limit, we've reached the end
		if len(result.Payload) < int(perpage) {
			break
		}

		// Move to next page
		page++
		params.WithPage(&page)
	}

	return nil, nil
}

func getUser(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("grafana_user.getUser", "connection_error", err)
		return nil, err
	}

	// Hydrate data into the list row
	if h.Item != nil {
		user := h.Item.(*models.UserSearchHitDTO)
		result, err := conn.client.Users.GetUserByID(user.ID)
		if err != nil {
			plugin.Logger(ctx).Error("grafana_user.getUser", "query_error", err, "user", user)
			return nil, err
		}
		return result.Payload, nil
	}

	// Prefer to get by ID
	if d.EqualsQuals["id"] != nil {
		id := d.EqualsQuals["id"].GetInt64Value()
		result, err := conn.client.Users.GetUserByID(id)
		if err != nil {
			plugin.Logger(ctx).Error("grafana_user.getUser", "query_error", err, "id", id)
			return nil, err
		}
		return result.Payload, nil
	}

	// Otherwise, get by email
	email := d.EqualsQuals["email"].GetStringValue()
	result, err := conn.client.Users.GetUserByLoginOrEmail(email)
	if err != nil {
		plugin.Logger(ctx).Error("grafana_user.getUser", "query_error", err, "email", email)
		return nil, err
	}
	return result.Payload, nil
}
