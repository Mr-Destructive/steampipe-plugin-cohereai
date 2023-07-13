# Table: cohereai_embed

Get Embeddings for given input strings(list of strings).

Notes:
* A `texts` is a list of strings to detect the language.

More information can be found about `cohereai_embed` table by using the inspect command or from the [api reference](https://docs.cohere.com/reference/embed)

## Examples

### Basic example with default settings

```sql
select
  embeddings
from
  cohereai_embed
where
  texts = '["hello", "world"]'
```

