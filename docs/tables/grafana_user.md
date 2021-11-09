# Table: grafana_user

Users in the Grafana installation.

Warning: The API used by this table requires admin user access via basic authentication (i.e. `admin:password`) in the `auth` config field. [Reference](https://grafana.com/docs/grafana/latest/http_api/user/#search-users).

## Examples

### List all users

```sql
select
  *
from
  grafana_user
```

### List all admin users

```sql
select
  login,
  email,
  is_admin
from
  grafana_user
where
  is_admin
```

### Users who have not been seen for more than 30 days

```sql
select
  login,
  email,
  last_seen_at,
  last_seen_at_age
from
  grafana_user
where
  last_seen_at < current_timestamp - interval '30 days'
```

### Users created in the last 7 days

```sql
select
  login,
  email,
  created_at
from
  grafana_user
where
  created_at > current_timestamp - interval '7 days'
```
