# Table: cohereai_classify

Get classification for a given input strings and examples.

Notes:
* A `inputs` is a list of strings to classify.
* A `examples` is a list of {"text": "apple", "label": "fruit"} structure of type [Example](https://docs.cohere.com/reference/classify)

## Examples

### Basic example with default settings

```sql
select
  classification
from
  cohereai_classify
where
  inputs = '["apple", "blue", "pineapple"]' and
  examples = '[{"text": "apple", "label": "fruit"}, {"text": "green", "label": "color"}, {"text": "grapes", "label": "fruit"}, {"text": "purple", "label": "color"}]'
```

### Classification with specific settings

`settings` is a JSONB object that accepts any of the [completion API request
parameters](https://docs.cohere.com/reference/classify).

```sql
select
  classification
from
  cohereai_classify
where
  settings = '{
    "model": "embed-multilingual-v2.0"
  }'
  and 
  inputs = '["Help!", "Call me when you can"]'
  and 
  examples = '[{"text": "Help!", "label": "urgent"}, {"text": "SOS", "label": "urgent"}, {"text": "Call me when you can", "label": "not urgent"}, {"text": "Talk later?", "label": "not urgent"}]'
```

