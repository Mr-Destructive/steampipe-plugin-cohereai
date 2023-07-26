# Table: cohereai_embed

Get Embeddings for given input strings(list of strings).

Notes:
* A `texts` is a list of strings to detect the language.

## Examples

### Basic example with default settings

```sql
select
  embeddings
from
  cohereai_embed
where
  texts = '["hello", "world"]';
```

