# Table: cohereai_tokenize

Get Tokens for a given input string.

Notes:
* A `text` is a string to get the tokens from.
* Maximum length of text is 65536 characters.

More information can be found about `cohereai_tokenize` table by using the inspect command or from the [api reference](https://docs.cohere.com/reference/tokenize)

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
