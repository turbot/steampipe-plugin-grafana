package grafana

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

func tableGrafanaFolder(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "grafana_folder",
		Description: "Folders in the Grafana installation.",
		List: &plugin.ListConfig{
			Hydrate: listFolder,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getFolder,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_INT, Description: "Unique identifier for the folder."},
			{Name: "uid", Type: proto.ColumnType_STRING, Description: "Globally unique identifier for the folder."},
			{Name: "title", Type: proto.ColumnType_STRING, Description: "Title of the folder."},
		},
	}
}

func listFolder(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("grafana_folder.listFolder", "connection_error", err)
		return nil, err
	}
	// NOTE: API supports paging, but SDK does not
	items, err := conn.gapi.Folders()
	if err != nil {
		plugin.Logger(ctx).Error("grafana_folder.listFolder", "query_error", err)
		return nil, err
	}
	for _, i := range items {
		d.StreamListItem(ctx, i)
	}
	return nil, nil
}

func getFolder(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("grafana_folder.getFolder", "connection_error", err)
		return nil, err
	}
	id := d.KeyColumnQuals["id"].GetInt64Value()
	item, err := conn.gapi.Folder(id)
	if err != nil {
		plugin.Logger(ctx).Error("grafana_folder.getFolder", "query_error", err, "id", id)
		return nil, err
	}
	return item, nil
}
