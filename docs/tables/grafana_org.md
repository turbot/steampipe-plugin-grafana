# Table: grafana_org

Organizations in the Grafana installation.

Warning: The API used by this table requires admin user access via basic authentication (i.e. `admin:password`) in the `auth` config field. [Reference](https://grafana.com/docs/grafana/latest/http_api/org/#search-all-organizations).

## Examples

### List all orgs

```sql
select
  *
from
  grafana_org
```
