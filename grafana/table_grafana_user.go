package grafana

import (
	"context"

	gapi "github.com/grafana/grafana-api-golang-client"

	"github.com/turbot/steampipe-plugin-sdk/v3/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
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
	// NOTE: API supports paging, but SDK does not
	items, err := conn.gapi.Users()
	if err != nil {
		plugin.Logger(ctx).Error("grafana_user.listUser", "query_error", err)
		return nil, err
	}
	for _, i := range items {
		d.StreamListItem(ctx, i)
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
		user := h.Item.(gapi.UserSearch)
		item, err := conn.gapi.User(user.ID)
		if err != nil {
			plugin.Logger(ctx).Error("grafana_user.getUser", "query_error", err, "user", user)
			return nil, err
		}
		return item, nil
	}

	// Prefer to get by ID
	if d.KeyColumnQuals["id"] != nil {
		id := d.KeyColumnQuals["id"].GetInt64Value()
		item, err := conn.gapi.User(id)
		if err != nil {
			plugin.Logger(ctx).Error("grafana_user.getUser", "query_error", err, "id", id)
			return nil, err
		}
		return item, nil
	}

	// Otherwise, get by email
	email := d.KeyColumnQuals["email"].GetStringValue()
	item, err := conn.gapi.UserByEmail(email)
	if err != nil {
		plugin.Logger(ctx).Error("grafana_user.getUser", "query_error", err, "email", email)
		return nil, err
	}
	return item, nil
}
