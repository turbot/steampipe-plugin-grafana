package grafana

import (
	"context"

	"github.com/grafana/grafana-openapi-client-go/client/orgs"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
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

	// Use the orgs API to get all orgs with pagination
	params := orgs.NewSearchOrgsParams()
	page := int64(1)      // Start with first page
	perpage := int64(100) // Default page size

	params.WithPage(&page)
	params.WithPerpage(&perpage)

	// Continue fetching pages until no more results
	for {
		result, err := conn.client.Orgs.SearchOrgs(params)
		if err != nil {
			plugin.Logger(ctx).Error("grafana_org.listOrg", "query_error", err)
			return nil, err
		}

		// If no results, break the loop
		if len(result.Payload) == 0 {
			break
		}

		// Stream the results from this page
		for _, org := range result.Payload {
			d.StreamListItem(ctx, org)
		}

		// If we got fewer results than the perpage limit, we've reached the end
		if len(result.Payload) < int(perpage) {
			break
		}

		// Move to next page
		page++
		params.WithPage(&page)
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
	if d.EqualsQuals["id"] != nil {
		id := d.EqualsQuals["id"].GetInt64Value()
		result, err := conn.client.Orgs.GetOrgByID(id)
		if err != nil {
			plugin.Logger(ctx).Error("grafana_org.getOrg", "query_error", err, "id", id)
			return nil, err
		}
		return result.Payload, nil
	}

	// Otherwise, get by name
	name := d.EqualsQuals["name"].GetStringValue()
	result, err := conn.client.Orgs.GetOrgByName(name)
	if err != nil {
		plugin.Logger(ctx).Error("grafana_org.getOrg", "query_error", err, "name", name)
		return nil, err
	}
	return result.Payload, nil
}
