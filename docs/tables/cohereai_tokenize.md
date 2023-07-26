# Table: cohereai_tokenize

Get Tokens for a given input string.

Notes:
* A `text` is a string to get the tokens from.
* Maximum length of text is 65536 characters.

## Examples

### Basic example with default settings

```sql
select
  tokens
from
  cohereai_tokenize
where
  text = 'hello world';
```
