# Table: cohereai_detokenize

Get string for a given list of tokens.

Notes:
* A `tokens` is a list of tokens(int64) to get the string.

## Examples

### Basic example with default settings

```sql
select
  text
from
  cohereai_detokenize
where
  tokens = '[33555, 1114]';
```

