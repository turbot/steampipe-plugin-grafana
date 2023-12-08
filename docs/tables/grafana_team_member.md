---
title: "Steampipe Table: grafana_team_member - Query Grafana Team Members using SQL"
description: "Allows users to query Grafana Team Members, providing a comprehensive view of all the team members and their associated details."
---

# Table: grafana_team_member - Query Grafana Team Members using SQL

Grafana Team Members are the users associated with a specific team in Grafana, a multi-platform analytics and visualization software. Team Members have specific roles and permissions within the team, which determine their access and capabilities in Grafana. Understanding the members and their roles is crucial for managing access and permissions effectively within Grafana.

## Table Usage Guide

The `grafana_team_member` table provides insights into the team members within Grafana. As a system administrator, you can explore member-specific details through this table, including their role, email, and associated metadata. Use it to manage and monitor user access and permissions, ensuring appropriate security and functionality within your Grafana teams.

**Important Notes**
- You must specify the `team_id` in the `where` clause to query this table.

## Examples

### List all members for a team
Explore which users are part of a specific team in Grafana, useful for understanding team composition and managing user access rights.

```sql+postgres
select
  *
from
  grafana_team_member
where
  team_id = 1;
```

```sql+sqlite
select
  *
from
  grafana_team_member
where
  team_id = 1;
```

### List all members of all teams
Explore which members belong to which teams in Grafana to understand team composition and facilitate effective communication. This is useful for managers to keep track of team structures and ensure the right teams are working on the right projects.

```sql+postgres
select
  t.name,
  tm.login,
  tm.email
from
  grafana_team as t,
  grafana_team_member as tm
where
  tm.team_id = t.id;
```

```sql+sqlite
select
  t.name,
  tm.login,
  tm.email
from
  grafana_team as t,
  grafana_team_member as tm
where
  tm.team_id = t.id;
```