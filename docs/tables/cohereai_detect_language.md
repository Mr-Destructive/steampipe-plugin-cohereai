# Table: cohereai_detect_language

Detect langugage for given input strings(list of strings).

Notes:
* A `texts` is a list of strings to detect the language.

## Examples

### Basic example with default settings

```sql
select
  language_name,
  language_code,
  text
from
  cohereai_detect_language
where
  texts = '["Этот текст на Русском языке", "नमस्ते गुरुजी", "some plain text"]';
```

### Examples for large text input

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
