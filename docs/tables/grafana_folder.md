# Table: grafana_folder

Folders in the Grafana installation.

## Examples

### List all folders

```sql
select
  *
from
  grafana_folder
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
