---
title: "Steampipe Table: grafana_org - Query Grafana Organizations using SQL"
description: "Allows users to query Grafana Organizations, providing details such as organization id, name, and address."
---

# Table: grafana_org - Query Grafana Organizations using SQL

Grafana Organizations are a way to manage access to resources in Grafana. Each user can belong to one or many organizations, and each dashboard belongs to a particular organization. Organizations allow for grouping and isolation of resources.

## Table Usage Guide

The `grafana_org` table allows users to gain insights into Grafana Organizations. As a system administrator or a DevOps engineer, you can use this table to fetch details about different organizations, their associated users, and dashboards. This can be particularly beneficial for managing access to resources and ensuring the isolation of resources across different organizations.

**Important Notes**
- The API used by this table requires admin user access via basic authentication (i.e. `admin:password`) in the `auth` config field. [Reference](https://grafana.com/docs/grafana/latest/http_api/org/#search-all-organizations).

## Examples

### List all orgs
Explore the various organizations within your Grafana setup to understand their settings and configurations. This aids in managing resources and maintaining an overview of your organizational structure.

```sql+postgres
select
  *
from
  grafana_org;
```

```sql+sqlite
select
  *
from
  grafana_org;
```