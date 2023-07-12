# Table: cohereai_classify

Get classification for a given input strings and examples.

Notes:
* A `inputs` is a list of strings to classify.
* A `examples` is a list of {"text": "apple", "label": "fruit"} structure of type [Example](https://docs.cohere.com/reference/classify)

## Examples

### Basic classification with given set of inputs and examples

```sql
select
  classification
from
  cohereai_classify
where
  inputs = '["apple", "blue", "pineapple"]'
  and examples = '[{"text": "apple", "label": "fruit"}, {"text": "green", "label": "color"}, {"text": "grapes", "label": "fruit"}, {"text": "purple", "label": "color"}]'
```

The table returns columns such as:

- `classification` as the class label associated to the particular input.
- `id` as a unique identifier for the classification label.
- `confidence` is the confidence score for the top predicted class as a floating point number.
- `labels` as a map containing each label with its confidence score.
- `inputs` as the list of input strings.
- `examples` as a list of text and label pairs.

You can get more information about `cohereai_classify` table by using the inspect command or from the [api reference](https://docs.cohere.com/reference/classify)

```
.inspect cohereai_classify
```

### Classification with specific settings(model, preset)

`settings` is a JSONB object that accepts any of the [Classification API request parameters](https://docs.cohere.com/reference/classify).

```sql
select
   classification
from
   cohereai_classify
where
   settings = '{
 "model": "embed - multilingual - v2.0" }'
 
   and inputs = '["Help!", "Call me when you can"]'
   and examples = '[{"text": "Help!", "label": "urgent"}, {"text": "SOS", "label": "urgent"}, {"text": "Call me when you can", "label": "not urgent"}, {"text": "Talk later?", "label": "not urgent"}]'
```

