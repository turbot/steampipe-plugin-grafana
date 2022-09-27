package grafana

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"
)

func tableGrafanaFolderPermission(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "grafana_folder_permission",
		Description: "Permissions for a given Folder in the Grafana installation.",
		List: &plugin.ListConfig{
			KeyColumns: plugin.SingleColumn("folder_uid"),
			Hydrate:    listFolderPermission,
		},
		Columns: []*plugin.Column{
			// Top columns
			// NOTE - Always zero? {Name: "id", Type: proto.ColumnType_INT, Description: "Unique identifier for the folder."},
			{Name: "folder_uid", Type: proto.ColumnType_STRING, Transform: transform.FromField("FolderUID"), Description: "Globally unique identifier for the folder."},
			{Name: "user_id", Type: proto.ColumnType_INT, Description: "ID of the user granted the permission."},
			{Name: "team_id", Type: proto.ColumnType_INT, Description: "ID of the team granted the permission."},
			{Name: "role", Type: proto.ColumnType_STRING, Description: "Role granted in the permission."},
			{Name: "is_folder", Type: proto.ColumnType_BOOL, Description: "True if the permission was granted to a folder."},
			{Name: "permission", Type: proto.ColumnType_INT, Description: "Permission level granted: 1 = View, 2 = Edit, 4 = Admin."},
			{Name: "permission_name", Type: proto.ColumnType_STRING, Description: "Name of the permission level: View, Edit, Admin."},
			{Name: "folder_id", Type: proto.ColumnType_INT, Description: "Unique identifier of the folder."},
			{Name: "dashboard_id", Type: proto.ColumnType_INT, Description: "Unique identifier of the dashboard."},
		},
	}
}

func listFolderPermission(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("grafana_folder.listFolderPermission", "connection_error", err)
		return nil, err
	}
	fuid := d.KeyColumnQuals["folder_uid"].GetStringValue()
	items, err := conn.gapi.FolderPermissions(fuid)
	if err != nil {
		if isNotFoundError(err) {
			return nil, nil
		}
		plugin.Logger(ctx).Error("grafana_folder.listFolderPermission", "query_error", err, "folder_uid", fuid)
		return nil, err
	}
	for _, i := range items {
		d.StreamListItem(ctx, i)
	}
	return nil, nil
}
