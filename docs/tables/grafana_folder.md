---
title: "Steampipe Table: grafana_folder - Query Grafana Folders using SQL"
description: "Allows users to query Grafana Folders, providing insights into the structure and organization of dashboards within a Grafana instance."
---

# Table: grafana_folder - Query Grafana Folders using SQL

Grafana Folders are a feature within Grafana that allows users to organize and group dashboards. This feature helps in managing large numbers of dashboards, providing a hierarchical structure for better navigation and search. Grafana Folders can also be used to apply permissions at a folder level, controlling user access to a group of dashboards.

## Table Usage Guide

The `grafana_folder` table provides insights into the organization of dashboards within a Grafana instance. As a DevOps engineer, use this table to explore folder-specific details, including the title, unique identifier, version, and associated metadata. Utilize it to manage and understand the hierarchical structure of dashboards, their grouping, and permissions applied at the folder level.

## Examples

### List all folders
Explore all folders within your Grafana setup to understand their structure and organization. This can be particularly useful for managing and navigating your data visualization projects.

```sql
select
  *
from
  grafana_folder
```

### List all folders with their permissions
Explore which Grafana folders have specific permissions to determine areas in which access may need to be revised or updated. This is useful for managing access control and ensuring appropriate levels of data security.

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