package grafana

import (
	"context"
	"strconv"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableGrafanaDatasource(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "grafana_datasource",
		Description: "Data sources in the Grafana installation.",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getDatasource,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_INT, Description: "Unique identifier for the data source."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name of the data source."},
			{Name: "datasource_type", Type: proto.ColumnType_STRING, Transform: transform.FromField("Type"), Description: "Type of the data source, e.g. postgres."},
			// Other columns
			{Name: "access", Type: proto.ColumnType_STRING, Description: "Method of access to the source, e.g. proxy."},
			{Name: "basic_auth", Type: proto.ColumnType_BOOL, Description: "True if the data source uses basic authentication."},
			{Name: "basic_auth_user", Type: proto.ColumnType_STRING, Description: "Username for basic authentication."},
			{Name: "database", Type: proto.ColumnType_STRING, Description: "Database name for the data source, e.g. mydb."},
			{Name: "is_default", Type: proto.ColumnType_BOOL, Description: "True if this is the default data source."},
			{Name: "json_data", Type: proto.ColumnType_JSON, Description: "Detailed configuration for the data source."},
			{Name: "org_id", Type: proto.ColumnType_INT, Description: "Unique identifier of the Grafana organization for this data source."},
			{Name: "secure_json_data", Type: proto.ColumnType_JSON, Transform: transform.FromField("SecureJSONData"), Description: "Secure configuration for the data source."},
			{Name: "url", Type: proto.ColumnType_STRING, Description: "URL of the data source."},
			{Name: "user", Type: proto.ColumnType_STRING, Description: "Username for the data source, e.g. myuser."},
		},
	}
}

func getDatasource(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("grafana_datasource.getDatasource", "connection_error", err)
		return nil, err
	}
	id := d.EqualsQuals["id"].GetInt64Value()

	// Convert int64 to string for the API call
	idStr := strconv.FormatInt(id, 10)

	// Use the datasources API to get datasource by ID
	result, err := conn.client.Datasources.GetDataSourceByID(idStr)
	if err != nil {
		plugin.Logger(ctx).Error("grafana_datasource.getDatasource", "query_error", err, "id", id)
		return nil, err
	}

	return result.Payload, nil
}
