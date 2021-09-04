# Table: grafana_team

Teams in the Grafana installation.

## Examples

### List all teams

```sql
select
  *
from
  grafana_team
```

### List teams with the most members

```sql
select
  name,
  member_count
from
  grafana_team
order by
  member_count desc
limit 5
```

### List teams with no members (e.g. to clean up)

```sql
select
  name,
  member_count
from
  grafana_team
where
  member_count = 0
```
