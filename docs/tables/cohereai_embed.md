# Table: cohereai_embed

Get Embeddings for given input strings(list of strings).

Notes:
* A `texts` is a list of strings to detect the language.

## Examples

### Basic example with default settings

The table `cohereai_embed` can return the embeddings from a given list of strings.

```sql
select
  embeddings
from
  cohereai_embed
where
  texts = '["hello", "world"]'
```

The table `cohereai_embed` has the following columns:

- `embeddings` as the list of `float64` values of the particular text.
- `text` as the string input.

More information can be found about `cohereai_embed` table by using the inspect command or from the [api reference](https://docs.cohere.com/reference/embed)

```
.inspect cohereai_embed
```

