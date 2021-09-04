package grafana

import (
	"context"

	gapi "github.com/grafana/grafana-api-golang-client"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
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
			{Name: "model", Type: proto.ColumnType_JSON, Hydrate: getDashboard, Description: "Full data model representing the dashbaord configuration."},
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
	// NOTE: Search API supports paging, but SDK does not
	items, err := conn.gapi.Dashboards()
	if err != nil {
		plugin.Logger(ctx).Error("grafana_dashboard.listDashboard", "query_error", err)
		return nil, err
	}
	for _, i := range items {
		d.StreamListItem(ctx, i)
	}
	return nil, nil
}

func getDashboard(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("grafana_dashboard.getDashboard", "connection_error", err)
		return nil, err
	}
	dashboard := h.Item.(gapi.FolderDashboardSearchResponse)
	item, err := conn.gapi.DashboardByUID(dashboard.UID)
	if err != nil {
		plugin.Logger(ctx).Error("grafana_dashboard.getDashboard", "query_error", err, "dashboard", dashboard)
		return nil, err
	}
	return item, nil
}
