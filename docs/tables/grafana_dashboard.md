# Table: grafana_dashboard

Dashboards in the Grafana installation.

## Examples

### List all dashboards

```sql
select
  id,
  title,
  url
from
  grafana_dashboard
```

### List all dashboards with a specific tag

```sql
select
  id,
  title,
  url,
  tags
from
  grafana_dashboard
where
  tags ? 'my-app'
```

### List all panels for a specific dashboard

```sql
select
  p->>'title' as panel_title,
  p->>'type' as panel_type
from
  grafana_dashboard as d,
  jsonb_array_elements(model->'panels') as p
where
  d.id = 3
```
