# Table: cohereai_classify

Get classification for a given input strings and examples.

Notes:
* A `inputs` is a list of strings to classify.(max 96 strings)
* A `examples` is a list of {"text": "apple", "label": "fruit"} structure of type [Example](https://docs.cohere.com/reference/classify)
* Minimum 2 `examples` should be provided and the maximum value is 2500 with each example of maximum of 512 tokens.

## Examples

### Basic classification with given set of inputs and examples

```sql
select
  classification
from
  cohereai_classify
where
  inputs = '["apple", "blue", "pineapple"]'
  and examples = '[{"text": "apple", "label": "fruit"}, {"text": "green", "label": "color"}, {"text": "grapes", "label": "fruit"}, {"text": "purple", "label": "color"}]';
```

### Classification with specific settings(model, preset)

```sql
select
  classification 
from
  cohereai_classify 
where
  settings = '{
 "model": "embed - multilingual - v2.0" }'
 
  and inputs = '["Help!", "Call me when you can"]' 
  and examples = '[{"text": "Help!", "label": "urgent"}, {"text": "SOS", "label": "urgent"}, {"text": "Call me when you can", "label": "not urgent"}, {"text": "Talk later?", "label": "not urgent"}]';
```

### Email Spam Classification

```sql
select 
  classification 
from 
  cohereai_classify 
where 
  inputs = '["Confirm your email address", "hey i need u to send some $"]' 
  and examples = '[{"label": "Spam", "text": "Dermatologists don''t like her!"}, {"label": "Spam", "text": "Hello, open to this?"}, {"label": "Spam", "text": "I need help please wire me $1000 right now"}, {"label": "Spam", "text": "Hot new investment, don’t miss this!"}, {"label": "Spam", "text": "Nice to know you ;)"}, {"label": "Spam", "text": "Please help me?"}, {"label": "Not spam", "text": "Your parcel will be delivered today"}, {"label": "Not spam", "text": "Review changes to our Terms and Conditions"}, {"label": "Not spam", "text": "Weekly sync notes"}, {"label": "Not spam", "text": "Re: Follow up from today’s meeting"}, {"label": "Not spam", "text": "Pre-read for tomorrow"}]';
```
