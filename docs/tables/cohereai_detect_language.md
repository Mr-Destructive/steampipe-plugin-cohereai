# Table: cohereai_detect_language

Detect langugage for given input strings(list of strings).

Notes:
* A `texts` is a list of strings to detect the language.

## Examples

### Basic example with default settings

The table `cohereai_detect_language` can return the language name from a given list of strings.

```sql
select
  language_name,
  language_code,
  text
from
  cohereai_detect_language
where
  texts = '["Этот текст на Русском языке", "नमस्ते गुरुजी", "some plain text"]'
```

The table `cohereai_detect_language` has the following columns:

- `language_name` as the language name.
- `language_code` as the language code.
- `text` as the input text.
- `texts` as a list of all the input strings.

More information can be found about `cohereai_detect_language` table by using the inspect command or from the [api reference](https://docs.cohere.com/reference/detect-language)

```
.inspect cohereai_detect_language
```

### Examples for large text input

A large paragraph of text can be parsed to detect the language in a better way, also multiple such strings could be added to the list of strings in the input.

```sql
select
  language_name,
  language_code,
  text
from
  cohereai_detect_language
where
  texts = '["Lorem ipsum dolor sit amet, consectetur adipiscing elit. Mauris vitae ex vitae enim aliquam feugiat ac vel arcu. Nunc pretium nisi sed finibus fermentum.", "Привет! Как дела? Я надеюсь, что у тебя все хорошо. Это тестовый текст на русском языке.", "Bonjour à tous! Jespère que vous allez bien. Ceci est un texte de test en français.", "¡Hola a todos! Espero que estén bien. Este es un texto de prueba en español.", "こんにちは！元気ですか？これは日本語のテストテキストです。"
]';
```
