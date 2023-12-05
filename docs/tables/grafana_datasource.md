---
title: "Steampipe Table: grafana_datasource - Query Grafana Datasources using SQL"
description: "Allows users to query Grafana Datasources, providing detailed information about each datasource configured in Grafana."
---

# Table: grafana_datasource - Query Grafana Datasources using SQL

Grafana is an open-source platform for data visualization, monitoring, and analysis. It allows you to query, visualize, alert on, and understand your metrics no matter where they are stored. A Grafana Datasource is a database or service that stores the data you want to visualize in Grafana.

## Table Usage Guide

The `grafana_datasource` table provides insights into Datasources within Grafana. As a Data Analyst or DevOps engineer, explore datasource-specific details through this table, including its type, access mode, database name, and other configuration details. Utilize it to manage and monitor your data sources, ensuring they are correctly configured and functioning as expected.

**Important Notes**
- You must specify the `id` in the `where` clause to query this table.

## Examples

### Get information for a data source
Explore which type of data source is being used in your Grafana setup by identifying it through a specific identifier. This can help in managing and troubleshooting your data visualization configurations.

```sql
select
  id,
  name,
  datasource_type
from
  grafana_datasource
where
  id = 1
```

### Get configuration of a data source
Analyze the settings to understand the configuration of a specific data source in Grafana. This could be used to troubleshoot issues or optimize data source utilization.

```sql
select
  name,
  jsonb_pretty(json_data)
from
  grafana_datasource
where
  id = 1
```