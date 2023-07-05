# Table: cohereai_generation

Create generations for a given text prompt.

Notes:
* A `prompt` or `settings -> 'prompt'` where qualifier is required for all queries.

## Examples

### Basic example with default settings

```sql
select
  generation
from
  cohereai_generation
where
  prompt = 'Give suggestions for a title for a science-fiction novel.';
```

### Completion with specific settings

`settings` is a JSONB object that accepts any of the [completion API request
parameters](https://docs.cohere.com/reference/generate).

```sql
select
  generation
from
  cohereai_generation
where
  settings = '{
    "model": "command-light",
    "num_generations": 3,
    "max_tokens": 100,
    "temperature": 0.9,
    "top_p": 1.0,
    "frequency_penalty": 0.0
  }'
  and prompt = 'Give suggestions for a title for a science-fiction novel.';
```

### Prompt through settings

The `prompt` column takes precedence, but you can also provide prompt text
through `settings` if easier.

```sql
select
  generation
from
  cohereai_generation
where
  settings = '{"prompt": "Write app ideas for AI-realted domains."}';
```

