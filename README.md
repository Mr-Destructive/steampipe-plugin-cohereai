# Cohere-AI Plugin for Steampipe

Use SQL to query models, completions and more from CohereAI.

- **[Get started →](https://hub.steampipe.io/plugins/turbot/cohereai)**
- Documentation: [Table definitions & examples](https://hub.steampipe.io/plugins/turbot/cohereai/tables)
- Community: [Slack Channel](https://steampipe.io/community/join)
- Get involved: [Issues](https://github.com/turbot/steampipe-plugin-cohereai/issues)

## Quick start

Install the plugin with [Steampipe](https://steampipe.io):

```shell
steampipe plugin install cohereai
```

Configure your API key in `~/.steampipe/config/cohereai.spc`:

```hcl
connection "cohereai" {
  plugin  = "cohereai"
  api_key = "YOUR_API_KEY"
}
```

Run steampipe:

```shell
steampipe query
```

Run a query:

```sql
select
  generation
from
  cohereai_generation
where
  prompt = '';
```

## Developing

Prerequisites:

- [Steampipe](https://steampipe.io/downloads)
- [Golang](https://golang.org/doc/install)

Clone:

```sh
git clone https://github.com/turbot/steampipe-plugin-cohereai.git
cd steampipe-plugin-cohereai
```

Build, which automatically installs the new version to your `~/.steampipe/plugins` directory:

```
make
```

Configure the plugin:

```
cp config/* ~/.steampipe/config
vi ~/.steampipe/config/cohereai.spc
```

Try it!

```
steampipe query
> .inspect cohereai
```

Further reading:

- [Writing plugins](https://steampipe.io/docs/develop/writing-plugins)
- [Writing your first table](https://steampipe.io/docs/develop/writing-your-first-table)

## Contributing

Please see the [contribution guidelines](https://github.com/turbot/steampipe/blob/main/CONTRIBUTING.md) and our [code of conduct](https://github.com/turbot/steampipe/blob/main/CODE_OF_CONDUCT.md). All contributions are subject to the [Apache 2.0 open source license](https://github.com/turbot/steampipe-plugin-cohereai/blob/main/LICENSE).

`help wanted` issues:

- [Steampipe](https://github.com/turbot/steampipe/labels/help%20wanted)
- [Cohere AI Plugin](https://github.com/turbot/steampipe-plugin-cohereai/labels/help%20wanted)
