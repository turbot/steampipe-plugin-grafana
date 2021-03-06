![image](https://hub.steampipe.io/images/plugins/turbot/grafana-social-graphic.png)

# Grafana Plugin for Steampipe

Use SQL to query instances, domains and more from Grafana.

- **[Get started →](https://hub.steampipe.io/plugins/turbot/grafana)**
- Documentation: [Table definitions & examples](https://hub.steampipe.io/plugins/turbot/grafana/tables)
- Community: [Slack Channel](https://steampipe.io/community/join)
- Get involved: [Issues](https://github.com/turbot/steampipe-plugin-grafana/issues)

## Quick start

Install the plugin with [Steampipe](https://steampipe.io):

```shell
steampipe plugin install grafana
```

Run a query:

```sql
select
  id,
  title,
  url
from
  grafana_dashboard
```

## Developing

Prerequisites:

- [Steampipe](https://steampipe.io/downloads)
- [Golang](https://golang.org/doc/install)

Clone:

```sh
git clone https://github.com/turbot/steampipe-plugin-grafana.git
cd steampipe-plugin-grafana
```

Build, which automatically installs the new version to your `~/.steampipe/plugins` directory:

```
make
```

Configure the plugin:

```
cp config/* ~/.steampipe/config
vi ~/.steampipe/config/grafana.spc
```

Try it!

```
steampipe query
> .inspect grafana
```

Further reading:

- [Writing plugins](https://steampipe.io/docs/develop/writing-plugins)
- [Writing your first table](https://steampipe.io/docs/develop/writing-your-first-table)

## Contributing

Please see the [contribution guidelines](https://github.com/turbot/steampipe/blob/main/CONTRIBUTING.md) and our [code of conduct](https://github.com/turbot/steampipe/blob/main/CODE_OF_CONDUCT.md). All contributions are subject to the [Apache 2.0 open source license](https://github.com/turbot/steampipe-plugin-grafana/blob/main/LICENSE).

`help wanted` issues:

- [Steampipe](https://github.com/turbot/steampipe/labels/help%20wanted)
- [Grafana Plugin](https://github.com/turbot/steampipe-plugin-grafana/labels/help%20wanted)
