## v0.5.0 [2023-10-02]

_Dependencies_

- Upgraded to [steampipe-plugin-sdk v5.6.1](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v561-2023-09-29) with support for rate limiters.
- Recompiled plugin with Go version `1.21`.

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
