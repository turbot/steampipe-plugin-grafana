# Table: grafana_team_member

Members of a team in the Grafana installation.

Note: A `team_uid` must be provided in all queries to this table.

## Examples

### List all members for a team

```sql
select
  *
from
  grafana_team_member
where
  team_id = 1
```

### List all members of all teams

```sql
select
  t.name,
  tm.login,
  tm.email
from
  grafana_team as t,
  grafana_team_member as tm
where
  tm.team_id = t.id
```
