# Table: cohereai_detokenize

Get string for a given list of tokens.

Notes:
* A `tokens` is a list of tokens(int64) to get the string.

More information can be found about `cohereai_detokenize` table by using the inspect command or from the [api reference](https://docs.cohere.com/reference/detokenize)

## Examples

### Basic example with default settings

```sql
select
  text
from
  cohereai_detokenize
where
  tokens = '[33555, 1114]'
```

