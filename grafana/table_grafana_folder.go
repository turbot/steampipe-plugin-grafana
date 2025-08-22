package grafana

import (
	"context"

	"github.com/grafana/grafana-openapi-client-go/client/folders"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
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

	// Use the folders API to get all folders with pagination
	params := folders.NewGetFoldersParams()
	limit := int64(100) // Default page size
	page := int64(1)    // Start with first page

	params.WithLimit(&limit)
	params.WithPage(&page)

	// Continue fetching pages until no more results
	for {
		result, err := conn.client.Folders.GetFolders(params)
		if err != nil {
			plugin.Logger(ctx).Error("grafana_folder.listFolder", "query_error", err)
			return nil, err
		}

		// If no results, break the loop
		if len(result.Payload) == 0 {
			break
		}

		// Stream the results from this page
		for _, folder := range result.Payload {
			d.StreamListItem(ctx, folder)
		}

		// If we got fewer results than the limit, we've reached the end
		if len(result.Payload) < int(limit) {
			break
		}

		// Move to next page
		page++
		params.WithPage(&page)
	}

	return nil, nil
}

func getFolder(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("grafana_folder.getFolder", "connection_error", err)
		return nil, err
	}
	id := d.EqualsQuals["id"].GetInt64Value()

	// Use the folders API to get folder by ID
	result, err := conn.client.Folders.GetFolderByID(id)
	if err != nil {
		plugin.Logger(ctx).Error("grafana_folder.getFolder", "query_error", err, "id", id)
		return nil, err
	}

	return result.Payload, nil
}
