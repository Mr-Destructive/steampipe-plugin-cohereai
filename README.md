![image](https://hub.steampipe.io/images/plugins/Mr-Destructive/cohereai-social-graphic.png)

# Cohere-AI Plugin for Steampipe

Use SQL to generate, classify, summarize text and more from CohereAI.

- **[Get started →](https://hub.steampipe.io/plugins/mr-destructive/cohereai)**
- Documentation: [Table definitions & examples](https://hub.steampipe.io/plugins/mr-destructive/cohereai/tables)
- Community: [Slack Channel](https://steampipe.io/community/join)
- Get involved: [Issues](https://github.com/mr-destructive/steampipe-plugin-cohereai/issues)

## Quick start

Install the plugin with [Steampipe](https://steampipe.io):

```shell
steampipe plugin install mr-destructive/cohereai
```

Configure your API key in `~/.steampipe/config/cohereai.spc`:

```hcl
connection "cohereai" {
  plugin = "mr-destructive/cohereai"

  # API key for requests. Required.
  # Get your API key at https://dashboard.cohere.ai/api-keys
  # This can also be set via the `COHEREAI_API_KEY` environment variable.
  api_key = "asLGEMKWMfkeFKENW038493fnWeng"
}
```

Or through environment variables:

```
export COHEREAI_API_KEY=asLGEMKWMfkeFKENW038493fnWeng
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
git clone https://github.com/mr-destructive/steampipe-plugin-cohereai.git
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
- [Cohere AI Plugin](https://github.com/mr-destructive/steampipe-plugin-cohereai/labels/help%20wanted)
