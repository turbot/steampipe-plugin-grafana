---
title: "Steampipe Table: grafana_dashboard_permission - Query Grafana Dashboard Permissions using SQL"
description: "Allows users to query Grafana  Permissions, specifically to retrieve information on the permissions assigned to dashboards within a Grafana instance."
---

# Table: grafana_dashboard_permission - Query Grafana Dashboard Permissions using SQL

Grafana Dashboard Permissions is a feature within Grafana that allows the assignment of user permissions to specific dashboards. This functionality enables the control of user access to dashboards, ensuring that only authorized users can view and edit specific dashboards. It is an integral part of managing user accessibility and security in Grafana.

## Table Usage Guide

The `grafana_dashboard_permission` table provides insights into the permissions assigned to dashboards within a Grafana instance. As a system administrator or security analyst, explore dashboard-specific permission details through this table, including the role, user, team, and permission level assigned to each dashboard. Utilize it to uncover information about user access, such as who can view or edit certain dashboards, and to ensure the proper implementation of access control policies.

## Examples

### List all permissions for a dashboard
Explore which permissions are granted for a specific dashboard in Grafana to manage access control effectively. This can help in maintaining security and ensuring only authorized users can make changes.

```sql+postgres
select
  *
from
  grafana_dashboard_permission
where
  dashboard_uid = 'BtcDlQ97z';
```

```sql+sqlite
select
  *
from
  grafana_dashboard_permission
where
  dashboard_uid = 'BtcDlQ97z';
```