---
organization: Turbot
category: ["software development"]
icon_url: "/images/plugins/turbot/grafana.svg"
brand_color: "#e0653b"
display_name: "Grafana"
short_name: "grafana"
description: "Steampipe plugin to query dashboards, data sources and more from Grafana."
og_description: "Query Grafana with SQL! Open source CLI. No DB required."
og_image: "/images/plugins/turbot/grafana-social-graphic.png"
engines: ["steampipe", "sqlite", "postgres", "export"]
---

# Grafana + Steampipe

[Grafana](https://grafana.com) is a cloud hosting company that provides virtual private servers and other infrastructure services.

[Steampipe](https://steampipe.io) is an open source CLI to instantly query cloud APIs using SQL.

List dashboards in your Grafana account:

```sql
select
  id,
  title,
  url
from
  grafana_dashboard
```

```
+----+--------------+---------------------------+
| id | title        | url                       |
+----+--------------+---------------------------+
| 3  | my dashboard | /d/Y4EbrQV7k/my-dashboard |
+----+--------------+---------------------------+
```

## Documentation

- **[Table definitions & examples →](/plugins/turbot/grafana/tables)**

## Get started

### Install

Download and install the latest Grafana plugin:

```bash
steampipe plugin install grafana
```

### Configuration

Installing the latest grafana plugin will create a config file (`~/.steampipe/config/grafana.spc`) with a single connection named `grafana`:

```hcl
connection "grafana" {
  plugin = "grafana"

  # URL of the Grafana installation
  url = "http://localhost:3000"

  # Authentication - API key
  auth = "eyJrIjoidGQ3VlMwVjlFVVc1TVNncjVWNGVYZnNDcaZIQkp2U2giLCJuIjoidGVzdDIsImlkIjoxfQ=="

  # Alternate authentication - username and password
  # auth = "admin:admin"
}
```

- `url` (required) - Root URL of a Grafana server. May alternatively be set via the `GRAFANA_URL` environment variable.
- `auth` (required) - API token or basic auth username:password. May alternatively be set via the `GRAFANA_AUTH` environment variable.
- `ca_cert` - Certificate CA bundle to use to verify the Grafana server's certificate. May alternatively be set via the `GRAFANA_CA_CERT` environment variable.
- `insecure_skip_verify` - Skip TLS certificate verification. May alternatively be set via the `GRAFANA_INSECURE_SKIP_VERIFY` environment variable.
- `org_id` - The organization id to operate on within grafana. May alternatively be set via the `GRAFANA_ORG_ID` environment variable.
- `tls_cert` - Client TLS certificate file to use to authenticate to the Grafana server. May alternatively be set via the `GRAFANA_TLS_CERT` environment variable.
- `tls_key` - Client TLS key file to use to authenticate to the Grafana server. May alternatively be set via the `GRAFANA_TLS_KEY` environment variable.


