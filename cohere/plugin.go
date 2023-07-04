package cohere

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name: "steampipe-plugin-cohere",
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
		},
		DefaultTransform: transform.FromGo().NullIfZero(),
		DefaultGetConfig: &plugin.GetConfig{
			ShouldIgnoreError: isNotFoundError,
		},
		TableMap: map[string]*plugin.Table{
			"cohereai_tokenize": tableCohereTokenize(ctx),
			"cohereai_generation": tableCohereGeneration(ctx),
            "cohereai_classification": tableCohereClassification(ctx),
		},
	}
	return p
}
