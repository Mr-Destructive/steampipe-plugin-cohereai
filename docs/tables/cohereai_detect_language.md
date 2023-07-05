# Table: cohereai_detect_language

Detect langugage for given input strings(list of strings).

Notes:
* A `texts` is a list of strings to detect the language.

## Examples

### Basic example with default settings

```sql
select
  language_name, language_code, text
from
  cohereai_detect_language
where
  texts = '["Этот текст на Русском языке", "नमस्ते गुरुजी", "some plain text"]'
```

