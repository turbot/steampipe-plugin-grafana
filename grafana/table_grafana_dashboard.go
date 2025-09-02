package grafana

import (
	"context"

	"github.com/grafana/grafana-openapi-client-go/client/search"
	"github.com/grafana/grafana-openapi-client-go/models"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableGrafanaDashboard(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "grafana_dashboard",
		Description: "Dashboards in the Grafana installation.",
		List: &plugin.ListConfig{
			Hydrate: listDashboard,
		},
		// NOTE: When using the SDK, the Get does not include most of the List
		// information, so is deliberately not implemented. Instead we rely on list
		// with hydration of Get data.
		Columns: []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_INT, Description: "Unique identifier for the dashboard."},
			{Name: "uid", Type: proto.ColumnType_STRING, Transform: transform.FromField("UID"), Description: "Globally unique identifier for the dashboard."},
			{Name: "title", Type: proto.ColumnType_STRING, Description: "Title of the dashboard."},
			// Other columns
			{Name: "dashboard_type", Type: proto.ColumnType_STRING, Transform: transform.FromField("Type"), Description: "Type of the dashboard, e.g. dash-db."},
			{Name: "folder_id", Type: proto.ColumnType_INT, Transform: transform.FromField("FolderID"), Description: "Unique identifier of the folder that contains the dashboard."},
			{Name: "folder_title", Type: proto.ColumnType_STRING, Description: "Title of the folder that contains the dashboard."},
			{Name: "folder_uid", Type: proto.ColumnType_STRING, Description: "Globally unique identifier of the folder that contains the dashboard."},
			{Name: "folder_url", Type: proto.ColumnType_STRING, Description: "URL of the folder that contains the dashboard."},
			{Name: "is_starred", Type: proto.ColumnType_BOOL, Description: "True if the dashboard has been starred."},
			{Name: "model", Type: proto.ColumnType_JSON, Hydrate: getDashboard, Transform: transform.FromField("Dashboard"), Description: "Full data model representing the dashboard configuration."},
			{Name: "slug", Type: proto.ColumnType_STRING, Hydrate: getDashboard, Transform: transform.FromField("Meta.Slug"), Description: "Slug of the dashboard."},
			{Name: "tags", Type: proto.ColumnType_JSON, Description: "List of tags for the dashboard."},
			{Name: "uri", Type: proto.ColumnType_STRING, Description: "URI of the dashboard."},
			{Name: "url", Type: proto.ColumnType_STRING, Description: "URL of the dashboard."},
		},
	}
}

func listDashboard(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("grafana_dashboard.listDashboard", "connection_error", err)
		return nil, err
	}

	// Use the search API to get dashboards with pagination
	params := search.NewSearchParams()
	dashType := "dash-db"
	limit := int64(1000) // Default page size
	page := int64(1)     // Start with first page

	params.WithType(&dashType)
	params.WithLimit(&limit)
	params.WithPage(&page)

	// Continue fetching pages until no more results
	for {
		result, err := conn.client.Search.Search(params)
		if err != nil {
			plugin.Logger(ctx).Error("grafana_dashboard.listDashboard", "query_error", err)
			return nil, err
		}

		// If no results, break the loop
		if len(result.Payload) == 0 {
			break
		}

		// Stream the results from this page
		for _, hit := range result.Payload {
			d.StreamListItem(ctx, hit)
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

func getDashboard(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("grafana_dashboard.getDashboard", "connection_error", err)
		return nil, err
	}

	hit := h.Item.(*models.Hit)

	// Use the dashboards API to get full dashboard details
	result, err := conn.client.Dashboards.GetDashboardByUID(hit.UID)
	if err != nil {
		plugin.Logger(ctx).Error("grafana_dashboard.getDashboard", "query_error", err, "dashboard", hit)
		return nil, err
	}

	// Return the full DashboardFullWithMeta structure which contains both Dashboard and Meta
	return result.Payload, nil
}
