# Table: grafana_folder_permission

Permissions granted on folders in the Grafana installation.

Note: A `folder_uid` must be provided in all queries to this table.

## Examples

### List all permissions for a folder

```sql
select
  *
from
  grafana_folder
where
  folder_uid = 'BtcDlQ97z'
```

### List all folders with their permissions

```sql
select
  f.uid,
  f.title,
  fp.*
from
  grafana_folder as f,
  grafana_folder_permission as fp
where
  f.uid = fp.folder_uid
```
