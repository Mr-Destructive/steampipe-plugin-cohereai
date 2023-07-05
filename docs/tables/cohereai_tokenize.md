# Table: cohereai_tokenize

Get Tokens for a given input string.

Notes:
* A `text` is a strings to get the tokens from.

## Examples

### Basic example with default settings

```sql
select
  tokens
from
  cohereai_tokenize
where
  text = 'hello world'
```
