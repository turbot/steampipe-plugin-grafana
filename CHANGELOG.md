## v1.0.0 [2024-10-22]

There are no significant changes in this plugin version; it has been released to align with [Steampipe's v1.0.0](https://steampipe.io/changelog/steampipe-cli-v1-0-0) release. This plugin adheres to [semantic versioning](https://semver.org/#semantic-versioning-specification-semver), ensuring backward compatibility within each major version.

_Dependencies_

- Recompiled plugin with Go version `1.22`. ([#38](https://github.com/turbot/steampipe-plugin-grafana/pull/38))
- Recompiled plugin with [steampipe-plugin-sdk v5.10.4](https://github.com/turbot/steampipe-plugin-sdk/blob/develop/CHANGELOG.md#v5104-2024-08-29) that fixes logging in the plugin export tool. ([#38](https://github.com/turbot/steampipe-plugin-grafana/pull/38))

## v0.6.0 [2023-12-12]

_What's new?_

- The plugin can now be downloaded and used with the [Steampipe CLI](https://steampipe.io/docs), as a [Postgres FDW](https://steampipe.io/docs/steampipe_postgres/overview), as a [SQLite extension](https://steampipe.io/docs//steampipe_sqlite/overview) and as a standalone [exporter](https://steampipe.io/docs/steampipe_export/overview). ([#34](https://github.com/turbot/steampipe-plugin-grafana/pull/34))
- The table docs have been updated to provide corresponding example queries for Postgres FDW and SQLite extension. ([#34](https://github.com/turbot/steampipe-plugin-grafana/pull/34))
- Docs license updated to match Steampipe [CC BY-NC-ND license](https://github.com/turbot/steampipe-plugin-grafana/blob/main/docs/LICENSE). ([#34](https://github.com/turbot/steampipe-plugin-grafana/pull/34))

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.8.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v580-2023-12-11) that includes plugin server encapsulation for in-process and GRPC usage, adding Steampipe Plugin SDK version to `_ctx` column, and fixing connection and potential divide-by-zero bugs. ([#33](https://github.com/turbot/steampipe-plugin-grafana/pull/33))

## v0.5.1 [2023-10-05]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.6.2](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v562-2023-10-03) which prevents nil pointer reference errors for implicit hydrate configs. ([#22](https://github.com/turbot/steampipe-plugin-grafana/pull/22))

## v0.5.0 [2023-10-02]

_Dependencies_

- Upgraded to [steampipe-plugin-sdk v5.6.1](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v561-2023-09-29) with support for rate limiters. ([#20](https://github.com/turbot/steampipe-plugin-grafana/pull/20))
- Recompiled plugin with Go version `1.21`. ([#20](https://github.com/turbot/steampipe-plugin-grafana/pull/20))

## v0.4.0 [2023-04-10]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.3.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v530-2023-03-16) which includes fixes for query cache pending item mechanism and aggregator connections not working for dynamic tables. ([#14](https://github.com/turbot/steampipe-plugin-grafana/pull/14))
- Recompiled plugin with [grafana-api-golang-client v0.12.0](https://github.com/grafana/grafana-api-golang-client/releases/tag/v0.12.0)

## v0.3.0 [2022-09-27]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v4.1.7](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v417-2022-09-08) which includes several caching and memory management improvements. ([#11](https://github.com/turbot/steampipe-plugin-grafana/pull/11))ugin with Go version `1.19`. ([#11](https://github.com/turbot/steampipe-plugin-grafana/pull/11))

## v0.2.1 [2022-05-23]

_Bug fixes_

- Fixed the Slack community links in README and docs/index.md files. ([#7](https://github.com/turbot/steampipe-plugin-grafana/pull/7))

## v0.2.0 [2022-04-27]

_Enhancements_

- Added support for native Linux ARM and Mac M1 builds.([#5](https://github.com/turbot/steampipe-plugin-grafana/pull/5))
- Recompiled plugin with [steampipe-plugin-sdk v3.1.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v310--2022-03-30) and Go version `1.18`. ([#4](https://github.com/turbot/steampipe-plugin-grafana/pull/4))

## v0.1.0 [2021-12-15]

_Enhancements_

- Recompiled plugin with [steampipe-plugin-sdk-v1.8.2](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v182--2021-11-22) and Go version 1.17 ([#2](https://github.com/turbot/steampipe-plugin-grafana/pull/2))

## v0.0.2 [2021-10-06]

_Bug fixes_

- Fix brand color in Steampipe Hub icons.

## v0.0.1 [2021-10-06]

_What's new?_

- New tables added
  - [grafana_dashboard](https://hub.steampipe.io/plugins/turbot/grafana/tables/grafana_dashboard)
  - [grafana_datasource](https://hub.steampipe.io/plugins/turbot/grafana/tables/grafana_datasource)
  - [grafana_folder](https://hub.steampipe.io/plugins/turbot/grafana/tables/grafana_folder)
  - [grafana_folder_permission](https://hub.steampipe.io/plugins/turbot/grafana/tables/grafana_folder_permission)
  - [grafana_org](https://hub.steampipe.io/plugins/turbot/grafana/tables/grafana_org)
  - [grafana_team](https://hub.steampipe.io/plugins/turbot/grafana/tables/grafana_team)
  - [grafana_team_member](https://hub.steampipe.io/plugins/turbot/grafana/tables/grafana_team_member)
  - [grafana_user](https://hub.steampipe.io/plugins/turbot/grafana/tables/grafana_user)
