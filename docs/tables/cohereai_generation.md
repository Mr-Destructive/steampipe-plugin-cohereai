# Table: cohereai_generation

Create generations for a given text prompt.

Notes:
* A `prompt` or `settings -> 'prompt'` where qualifier is required for all queries.

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

The table `cohereai_generation` has the following columns:

- `generation` as the generated text.
- `likelihood` as the float64 value of the likelihood of the generated text.
- `token_likelihoods` as the list of object containing likelihood of a token as a float64, and token as a string.
- `prompt` as the string input to generate response.

*NOTE: The fields `likelihood` and `token_likelihoods` will only return values if `return_likelihoods` is set either as `GENERATION` or `ALL`.*

The `return_likelihoods` can be set to `GENERATION` or `ALL`. If the former is selected, the API would respond with the likelihood for only the generation text, else for the later, the likelihood will br given for both generated and prompt text (Default in the plugin is `GENERATION`).

More information can be found about `cohereai_generation` table by using the inspect command or from the [api reference](https://docs.cohere.com/reference/generate)

```
.inspect cohereai_generation
```

### Generation with specific settings(model, number of responses, etc.)

By specifying the settings as json string in the query, the generation can be tweaked. Like the number of generations can be changed (default in plugin as 3, between 1 and 5 only), the temperature can be altered between `0` to `5` as a decimal value, etc.

`settings` is a JSONB object that accepts any of the [Generation API request parameters](https://docs.cohere.com/reference/generate).

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
  and prompt = 'Give suggestions for a title for a science-fiction novel.
```

### Pass the prompt string through settings

The `prompt` column takes precedence, but you can also provide prompt text through `settings` if easier with other settings as well.

```sql
select 
  generation 
from 
  cohereai_generation 
where 
  settings = '{
   "prompt": "Write app ideas for AI-realted domains."}';
```
