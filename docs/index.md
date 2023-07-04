---
organization: Turbot
category: ["ai"]
icon_url: "/images/plugins/turbot/cohereai.svg"
brand_color: "#000000"
display_name: "CohereAI"
short_name: "cohereai"
description: "Steampipe plugin to query generations, classifications and more from CohereAI."
og_description: "Query CohereAI with SQL! Open source CLI. No DB required."
og_image: "/images/plugins/turbot/cohereai-social-graphic.png"
---

# CohereAI + Steampipe

[CohereAI](https://cohere.com) is an Artificial Intelligence research and development company that provides APIs for general models.

[Steampipe](https://steampipe.io) is an open source CLI to instantly query cloud APIs using SQL.

Generate completions for a given text prompt in your OpenAI account:

```sql
select
  generation 
from
  cohereai_generation 
where
  prompt = 'Write an novel title for a magical world.';
```

```
+---------------------------------------------+
| generation                                  |
+---------------------------------------------+
|                                             |
| The Magically Mysterious World of Enchantia |
|                                             |
| The Magic World of Ooze                     |
|                                             |
| The Magicians of Xylar                      |
+---------------------------------------------+
```

## Documentation

- **[Table definitions & examples →](/plugins/turbot/coherai/tables)**

## Get started

### Install

Download and install the latest CohereAI plugin:

```bash
steampipe plugin install cohereai 
```

### Credentials

| Item        | Description                                                                                                                                                                                                                                                                                 |
|-------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| Credentials | OpenAI requires an [API Key](https://dashboard.cohere.ai/api-keys) for all requests.                                                                                                                                                                                 |
| Permissions | API Keys have the same permissions as the user who creates them, and if the user permissions change, the API key permissions also change.                                                                                                                                               |
| Radius      | Each connection represents a single OpenAI Installation.                                                                                                                                                                                                                                   |
| Resolution  | 1. Credentials explicitly set in a steampipe config file (`~/.steampipe/config/cohereai.spc`)<br />2. Credentials specified in environment variables, e.g., `COHEREAI_API_KEY`. |

### Configuration

Installing the latest openai plugin will create a config file (`~/.steampipe/config/cohereai.spc`) with a single connection named `cohereai`:

```hcl
connection "cohereai" {
  plugin = "cohereai"

  # Get your API key at https://dashboard.cohere.ai/api-keys
  # This can also be set via the `COHEREAI_API_KEY` environment variable.
  api_key = "asLGEMKWMfkeFKENW038493fnWeng"
}
```

### Credentials from Environment Variables

The CohereAI plugin will use the standard CohereAI environment variables to obtain credentials **only if other arguments (`api_key`) are not specified** in the connection:

```sh
export COHEREAI_API_KEY=asLGEMKWMfkeFKENW038493fnWeng
```

## Get involved

- Open source: https://github.com/turbot/steampipe-plugin-cohereai
- Community: [Slack Channel](https://steampipe.io/community/join)
