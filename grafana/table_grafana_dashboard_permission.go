package grafana

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableGrafanaDashboardPermission(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "grafana_dashboard_permission",
		Description: "Permissions for a given Dashboard in the Grafana installation.",
		List: &plugin.ListConfig{
			KeyColumns: plugin.SingleColumn("dashboard_uid"),
			Hydrate:    listDashboardPermission,
		},
		Columns: []*plugin.Column{
			// Top columns
			// NOTE - Always zero? {Name: "id", Type: proto.ColumnType_INT, Description: "Unique identifier for the dashboard."},
			{Name: "dashboard_uid", Type: proto.ColumnType_STRING, Transform: transform.FromField("UID"), Description: "Globally unique identifier for the dashboard."},
			{Name: "user_id", Type: proto.ColumnType_INT, Description: "ID of the user granted the permission."},
			{Name: "team_id", Type: proto.ColumnType_INT, Description: "ID of the team granted the permission."},
			{Name: "role", Type: proto.ColumnType_STRING, Description: "Role granted in the permission."},
			{Name: "is_folder", Type: proto.ColumnType_BOOL, Description: "True if the permission was granted to a dashboard."},
			{Name: "permission", Type: proto.ColumnType_INT, Description: "Permission level granted: 1 = View, 2 = Edit, 4 = Admin."},
			{Name: "permission_name", Type: proto.ColumnType_STRING, Description: "Name of the permission level: View, Edit, Admin."},
			{Name: "dashboard_id", Type: proto.ColumnType_INT, Description: "Unique identifier of the dashboard."},
		},
	}
}

func listDashboardPermission(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("grafana_dashboard_permission.listDashboardPermission", "connection_error", err)
		return nil, err
	}
	duid := d.EqualsQuals["dashboard_uid"].GetStringValue()

	// Use the dashboard permissions API to get permissions for the dashboard
	result, err := conn.client.DashboardPermissions.GetDashboardPermissionsListByUID(duid)
	if err != nil {
		if isNotFoundError(err) {
			return nil, nil
		}
		plugin.Logger(ctx).Error("grafana_dashboard_permission.listDashboardPermission", "query_error", err, "dashboard_uid", duid)
		return nil, err
	}

	for _, permission := range result.Payload {
		d.StreamListItem(ctx, permission)
	}
	return nil, nil
}
