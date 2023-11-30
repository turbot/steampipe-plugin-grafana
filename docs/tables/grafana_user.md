---
title: "Steampipe Table: grafana_user - Query Grafana Users using SQL"
description: "Allows users to query Grafana Users, specifically user details and their respective roles, providing insights into user management and access control."
---

# Table: grafana_user - Query Grafana Users using SQL

Grafana is a multi-platform open-source analytics and interactive visualization web application. It provides charts, graphs, and alerts for the web when connected to supported data sources. Users in Grafana represent accounts with login credentials that can be granted permissions to access resources within Grafana.

## Table Usage Guide

The `grafana_user` table provides insights into user accounts within Grafana. As an administrator or security analyst, explore user-specific details through this table, including their roles, permissions, and associated metadata. Utilize it to manage user access control, review user permissions, and ensure adherence to security policies.

## Examples

### List all users
Explore all the users within your Grafana platform to manage access and permissions effectively. This helps to maintain security and control over who can access and modify your data.

```sql
select
  *
from
  grafana_user
```

### List all admin users
Identify instances where users have administrative privileges. This can be useful for ensuring proper access controls and identifying potential security risks.

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
Discover the segments of users who have not logged in for over a month. This can be useful to identify inactive users and potentially reach out to them to re-engage them with the platform.

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
Explore which Grafana users were created in the past week. This is useful to keep track of new user activity and growth within your system.

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