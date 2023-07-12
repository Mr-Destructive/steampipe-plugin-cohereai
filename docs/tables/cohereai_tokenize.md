# Table: cohereai_tokenize

Get Tokens for a given input string.

Notes:
* A `text` is a strings to get the tokens from.

## Examples

### Basic example with default settings

The table `cohereai_tokenize` can return the tokens from a given list of strings.

```sql
select
  tokens
from
  cohereai_tokenize
where
  text = 'hello world'
```

The table `cohereai_tokenize` has the following columns:

- `tokens` as the list of tokens(int64).
- `text` as the string input.
- `token_strings` as a list of all the input strings.

More information can be found about `cohereai_tokenize` table by using the inspect command or from the [api reference](https://docs.cohere.com/reference/tokenize)

```
.inspect cohereai_tokenize
```
