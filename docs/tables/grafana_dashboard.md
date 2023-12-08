---
title: "Steampipe Table: grafana_dashboard - Query Grafana Dashboards using SQL"
description: "Allows users to query Grafana Dashboards, specifically the details of each dashboard, providing insights into the visualization and analytics settings."
---

# Table: grafana_dashboard - Query Grafana Dashboards using SQL

Grafana is a multi-platform open-source analytics and interactive visualization web application. It provides charts, graphs, and alerts for the web when connected to supported data sources, including Prometheus, InfluxDB, and many others. Grafana is most commonly used for visualizing time series data for infrastructure and application analytics.

## Table Usage Guide

The `grafana_dashboard` table provides insights into Grafana Dashboards within Grafana. As a Data Analyst or DevOps engineer, explore dashboard-specific details through this table, including the layout, panels, variables, and associated metadata. Utilize it to uncover information about dashboards, such as the data sources used, the types of panels and visualizations, and the overall layout and structure of the dashboards.

## Examples

### List all dashboards
Discover the segments that contain all your dashboards with this query. It can be used to quickly access the details of each dashboard, such as its ID, title, and URL, without having to navigate through each one individually.

```sql+postgres
select
  id,
  title,
  url
from
  grafana_dashboard;
```

```sql+sqlite
select
  id,
  title,
  url
from
  grafana_dashboard;
```

### List all dashboards with a specific tag
Analyze the settings to understand which dashboards are associated with a specific tag. This can be useful in identifying and organizing dashboards relevant to a particular application or project.

```sql+postgres
select
  id,
  title,
  url,
  tags
from
  grafana_dashboard
where
  tags ? 'my-app';
```

```sql+sqlite
Error: SQLite does not support the '?' operator for JSON arrays.
```

### List all panels for a specific dashboard
Gain insights into the different panels within a specific Grafana dashboard. This allows you to understand and manage the types and titles of panels for a selected dashboard.

```sql+postgres
select
  p->>'title' as panel_title,
  p->>'type' as panel_type
from
  grafana_dashboard as d,
  jsonb_array_elements(model->'panels') as p
where
  d.id = 3;
```

```sql+sqlite
select
  json_extract(p.value, '$.title') as panel_title,
  json_extract(p.value, '$.type') as panel_type
from
  grafana_dashboard as d,
  json_each(d.model, '$.panels') as p
where
  d.id = 3;
```