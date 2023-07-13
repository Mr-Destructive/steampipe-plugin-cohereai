# Table: cohereai_summarize

Create a summary for a given text prompt.

Notes:
* A `text` is a string to get the summary 
* The maximum length of text is 100,000 characters.

## Examples

### Basic example with default settings

The `cohereai_summarize` table can return the summary from a given text prompt.

```sql
select
  summary
from
  cohereai_summarize
where
  text = 'In the vast expanse of the cosmos, where stars twinkle like scattered diamonds, a sense of awe and wonder takes hold. The universe, with its countless galaxies and celestial bodies, holds secrets yet to be fully unraveled. From the intricate dance of planets within our own solar system to the majestic swirls of galaxies millions of light-years away, there is a profound beauty in the cosmic symphony that unfolds before our eyes.';
```

The table `cohereai_summarize` has the following columns:

- `summary` as a string summary text for the given input text.
- `id` as a unique identifier for the generated summary.
- `text` as the string input.

### Summarization with tweaked settings

`settings` is a JSONB object that accepts any of the [Summarize API request parameters](https://docs.cohere.com/reference/summarize-2).

```sql
select
  summary
from
  cohereai_summarize
where
  settings = '{
   "model": "summarize-medium",
   "length": "short",
   "temperature": 0.9,
   "format": "bullets"}' 
  and text = 'In the vast expanse of the cosmos, where stars twinkle like scattered diamonds, a sense of awe and wonder takes hold. The universe, with its countless galaxies and celestial bodies, holds secrets yet to be fully unraveled. From the intricate dance of planets within our own solar system to the majestic swirls of galaxies millions of light-years away, there is a profound beauty in the cosmic symphony that unfolds before our eyes.';
```

The settings can be tweaked by specifying `settings` as json string in the query. The `format` of the generated summary can be changed from `bullets`, `paragraph` or leave it to `auto` by default, change the model as either `summarize-medium` or `summarize-xlarge`(default), tweak the temperature in the range of 0 to 5, etc.
