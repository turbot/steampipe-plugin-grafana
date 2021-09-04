# Table: grafana_datasource

Data sources in the Grafana installation.

Note: An `id` must be provided in all queries to this table.

## Examples

### Get information for a data source

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

```sql
select
  name,
  jsonb_pretty(json_data)
from
  grafana_datasource
where
  id = 1
```
