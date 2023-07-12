package cohere

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name: "steampipe-plugin-cohereai",
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
			Schema:      ConfigSchema,
		},
		DefaultTransform: transform.FromGo().NullIfZero(),
		DefaultGetConfig: &plugin.GetConfig{
			ShouldIgnoreError: isNotFoundError,
		},
		TableMap: map[string]*plugin.Table{
			"cohereai_classify":        tableCohereClassification(ctx),
			"cohereai_detect_language": tableCohereDetectLanguage(ctx),
			"cohereai_detokenize":      tableCohereDetokenize(ctx),
			"cohereai_embed":           tableCohereEmbed(ctx),
			"cohereai_generation":      tableCohereGeneration(ctx),
			"cohereai_summarize":       tableCohereSummarize(ctx),
			"cohereai_tokenize":        tableCohereTokenize(ctx),
		},
	}
	return p
}
