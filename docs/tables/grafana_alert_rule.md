---
title: "Steampipe Table: grafana_alert_rule - Query Grafana Alert Rules using SQL"
description: "Query Grafana alert rules to understand configuration, grouping, state, labels, and folder association."
---

# Table: grafana_alert_rule - Query Grafana Alert Rules using SQL

Grafana alert rules define evaluation conditions that trigger alerts. Rules belong to folders and rule groups, can be paused, and have configurable behaviors for no data and execution errors. Labels help categorize alerts for routing and notification policies.

## Table Usage Guide

Use the `grafana_alert_rule` table to explore alert rule details such as `uid`, `title`, `rule_group`, `condition`, `is_paused`, `labels`, `folder_uid`, and error/no-data states. This helps you audit alerting configuration, find paused rules, and understand rule placement across folders.

## Examples

### Basic info

Retrieve key information about your Grafana alert rules, such as their names, groups, folders, conditions, evaluation intervals, labels, and pause status. This query helps you quickly audit your alert configuration, see which rules are currently paused, how rules are grouped, and what labels are used for routing and categorization.


### List all alert rules
Explore all alert rules in your Grafana instance.

```sql+postgres
select 
  title, 
  rule_group, 
  folder_uid, 
  condition, 
  interval, 
  labels, 
  is_paused 
from 
  grafana_alert_rule;
```

```sql+sqlite
select 
  title, 
  rule_group, 
  folder_uid, 
  condition, 
  interval, 
  labels, 
  is_paused 
from 
  grafana_alert_rule;
```

### List paused alert rules
Identify alert rules that are currently paused.

```sql+postgres
select
  uid,
  title,
  rule_group,
  folder_uid,
  is_paused
from
  grafana_alert_rule
where
  is_paused;
```

```sql+sqlite
select
  uid,
  title,
  rule_group,
  folder_uid,
  is_paused
from
  grafana_alert_rule
where
  is_paused;
```

### List alert rules with their folder titles
See where each alert rule resides by joining to folders.

```sql+postgres
select
  r.uid,
  r.title,
  r.rule_group,
  f.title as folder_title,
  r.folder_uid
from
  grafana_alert_rule as r
  left join grafana_folder as f on r.folder_uid = f.uid;
```

```sql+sqlite
select
  r.uid,
  r.title,
  r.rule_group,
  f.title as folder_title,
  r.folder_uid
from
  grafana_alert_rule as r
  left join grafana_folder as f on r.folder_uid = f.uid;
```