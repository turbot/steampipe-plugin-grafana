---
title: "Steampipe Table: grafana_folder_permission - Query Grafana Folder Permissions using SQL"
description: "Allows users to query Grafana Folder Permissions, specifically to retrieve information on the permissions assigned to folders within a Grafana instance."
---

# Table: grafana_folder_permission - Query Grafana Folder Permissions using SQL

Grafana Folder Permissions is a feature within Grafana that allows the assignment of user permissions to specific folders. This functionality enables the control of user access to dashboards, ensuring that only authorized users can view and edit specific dashboards. It is an integral part of managing user accessibility and security in Grafana.

## Table Usage Guide

The `grafana_folder_permission` table provides insights into the permissions assigned to folders within a Grafana instance. As a system administrator or security analyst, explore folder-specific permission details through this table, including the role, user, team, and permission level assigned to each folder. Utilize it to uncover information about user access, such as who can view or edit certain dashboards, and to ensure the proper implementation of access control policies.

**Important Notes**
- You must specify the `folder_uid` in the `where` clause to query this table.

## Examples

### List all permissions for a folder
Explore which permissions are granted for a specific folder in Grafana to manage access control effectively. This can help in maintaining security and ensuring only authorized users can make changes.

```sql
select
  *
from
  grafana_folder
where
  folder_uid = 'BtcDlQ97z'
```

### List all folders with their permissions
Explore which folders have specific permissions in your Grafana setup. This can help in managing access controls and ensuring proper security protocols.

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