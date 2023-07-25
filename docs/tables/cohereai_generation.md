# Table: cohereai_generation

Create generations for a given text prompt.

Notes:
* A `prompt` or `settings -> 'prompt'` where qualifier is required for all queries.
* The fields `likelihood` and `token_likelihoods` will only return values if `return_likelihoods` is set either as `GENERATION` or `ALL`.*

The `return_likelihoods` can be set to `GENERATION` or `ALL`. If the former is selected, the API would respond with the likelihood for only the generation text, else for the later, the likelihood will br given for both generated and prompt text (Default in the plugin is `GENERATION`).

## Examples

### Basic example with simple prompt

```sql
select
  generation
from
  cohereai_generation
where
  prompt = 'Give suggestions for a title for a science-fiction novel.';
```

### Generation with specific settings(model, number of responses, etc.)

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
   "frequency_penalty": 0.0 }'
  and prompt = 'Give suggestions for a title for a science-fiction novel.';
```

### Pass the prompt string through settings

```sql
select 
  generation 
from 
  cohereai_generation 
where 
  settings = '{
   "prompt": "Write app ideas for AI-realted domains."}';
```

### Spell check a piece of text

```sql
select
  generation
from
  cohereai_generation
where
  settings = '{"num_generations": 1}'
  and prompt = 'Check the smaple. Sample: "The impotance of effictive comunication. This is an exmaple artcile abot missplelled wrds."';
```
