package grafana

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableGrafanaAlertRule(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "grafana_alert_rule",
		Description: "Alert rules in the Grafana installation.",
		List: &plugin.ListConfig{
			Hydrate: listAlertRule,
		},
		// NOTE: When using the SDK, the Get does not include most of the List
		// information, so is deliberately not implemented. Instead we rely on list
		// with hydration of Get data.
		Columns: []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_INT, Description: "Unique identifier for the alert rule."},
			{Name: "uid", Type: proto.ColumnType_STRING, Transform: transform.FromField("UID"), Description: "Globally unique identifier for the alert rule."},
			{Name: "title", Type: proto.ColumnType_STRING, Description: "Title of the alert rule."},
			// Other columns
			{Name: "condition", Type: proto.ColumnType_STRING, Description: "Condition of the alert rule."},
			{Name: "interval", Type: proto.ColumnType_STRING, Transform: transform.FromField("For"), Description: "Evaluation interval for the alert rule."},
			{Name: "is_paused", Type: proto.ColumnType_BOOL, Description: "True if the alert rule has been paused."},
			{Name: "no_data_state", Type: proto.ColumnType_STRING, Description: "No data state of the alert rule."},
			{Name: "execution_error_state", Type: proto.ColumnType_STRING, Description: "Execution error state of the alert rule."},
			{Name: "folder_uid", Type: proto.ColumnType_STRING, Description: "Globally unique identifier of the folder that contains the alert rule."},
			{Name: "rule_group", Type: proto.ColumnType_STRING, Description: "Rule group of the alert rule."},
			{Name: "updated", Type: proto.ColumnType_STRING, Description: "Last time the alert rule was updated."},
			{Name: "labels", Type: proto.ColumnType_JSON, Description: "Labels of the alert rule."},
		},
	}
}

func listAlertRule(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("grafana_alert_rule.listAlertRule", "connection_error", err)
		return nil, err
	}

	// Use the Provisioning API to get permissions for the folder
	result, err := conn.client.Provisioning.GetAlertRules()
	if err != nil {
		if isNotFoundError(err) {
			return nil, nil
		}
		plugin.Logger(ctx).Error("grafana_alert_rule.listAlertRule", "query_error", err)
		return nil, err
	}

	for _, permission := range result.Payload {
		d.StreamListItem(ctx, permission)
	}
	return nil, nil
}
