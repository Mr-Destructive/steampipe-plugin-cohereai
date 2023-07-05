# Table: cohereai_summarize

Create a summary for a given text prompt.

## Examples

### Basic example with default settings

```sql
select
  summary
from
  cohereai_summarize
where
  text = 'In the vast expanse of the cosmos, where stars twinkle like scattered diamonds, a sense of awe and wonder takes hold. The universe, with its countless galaxies and celestial bodies, holds secrets yet to be fully unraveled. From the intricate dance of planets within our own solar system to the majestic swirls of galaxies millions of light-years away, there is a profound beauty in the cosmic symphony that unfolds before our eyes.';
```

### Completion with specific settings

`settings` is a JSONB object that accepts any of the [completion API request
parameters](https://docs.cohere.com/reference/summarize-2).

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
    "format": "bullets"
  }'
  and text = 'In the vast expanse of the cosmos, where stars twinkle like scattered diamonds, a sense of awe and wonder takes hold. The universe, with its countless galaxies and celestial bodies, holds secrets yet to be fully unraveled. From the intricate dance of planets within our own solar system to the majestic swirls of galaxies millions of light-years away, there is a profound beauty in the cosmic symphony that unfolds before our eyes.';
```


