---
title: "Steampipe Table: grafana_team - Query Grafana Teams using SQL"
description: "Allows users to query Teams in Grafana, specifically the team details, providing insights into team members, permissions, and associated metadata."
---

# Table: grafana_team - Query Grafana Teams using SQL

Grafana is a platform for analytics and visualization that allows you to query, visualize, alert on, and understand your metrics. Teams in Grafana are groups of users that reflect the organization in your system. Teams allow you to grant permissions for managing dashboards and data sources to specific groups of users.

## Table Usage Guide

The `grafana_team` table provides insights into Teams within Grafana. As a DevOps engineer, explore team-specific details through this table, including team members, permissions, and associated metadata. Utilize it to manage and organize user access to dashboards and data sources.

## Examples

### List all teams
Explore all the teams available in your Grafana instance to manage permissions and access controls more effectively.

```sql+postgres
select
  *
from
  grafana_team;
```

```sql+sqlite
select
  *
from
  grafana_team;
```

### List teams with the most members
Discover the teams that have the highest number of members. This information can be useful in understanding team dynamics and resource allocation within the organization.

```sql+postgres
select
  name,
  member_count
from
  grafana_team
order by
  member_count desc
limit 5;
```

```sql+sqlite
select
  name,
  member_count
from
  grafana_team
order by
  member_count desc
limit 5;
```

### List teams with no members (e.g. to clean up)
Determine the teams that currently have no members, which may be useful for organizational cleanup or restructuring.

```sql+postgres
select
  name,
  member_count
from
  grafana_team
where
  member_count = 0;
```

```sql+sqlite
select
  name,
  member_count
from
  grafana_team
where
  member_count = 0;
```