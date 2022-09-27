package grafana

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
)

func tableGrafanaOrg(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "grafana_org",
		Description: "Orgs in the Grafana installation.",
		List: &plugin.ListConfig{
			Hydrate: listOrg,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AnyColumn([]string{"id", "name"}),
			Hydrate:    getOrg,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_INT, Description: "Unique identifier for the org."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name of the org."},
		},
	}
}

func listOrg(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("grafana_org.listOrg", "connection_error", err)
		return nil, err
	}
	// NOTE: API supports paging, but SDK does not
	items, err := conn.gapi.Orgs()
	if err != nil {
		plugin.Logger(ctx).Error("grafana_org.listOrg", "query_error", err)
		return nil, err
	}
	for _, i := range items {
		d.StreamListItem(ctx, i)
	}
	return nil, nil
}

func getOrg(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("grafana_org.getOrg", "connection_error", err)
		return nil, err
	}
	// Prefer to get by ID
	if d.KeyColumnQuals["id"] != nil {
		id := d.KeyColumnQuals["id"].GetInt64Value()
		item, err := conn.gapi.Org(id)
		if err != nil {
			plugin.Logger(ctx).Error("grafana_org.getOrg", "query_error", err, "id", id)
			return nil, err
		}
		return item, nil
	}
	// Otherwise, get by name
	name := d.KeyColumnQuals["name"].GetStringValue()
	item, err := conn.gapi.OrgByName(name)
	if err != nil {
		plugin.Logger(ctx).Error("grafana_org.getOrg", "query_error", err, "name", name)
		return nil, err
	}
	return item, nil
}
