package cohere

import (
	"context"
	"encoding/json"

	coherego "github.com/cohere-ai/cohere-go"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

// Defines the cohereai_completion table
func tableCohereGeneration(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "cohereai_generation",
		Description: "Generation in Cohere AI.",
		List: &plugin.ListConfig{
			Hydrate: listGeneration,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "prompt", Require: plugin.Optional},
				{Name: "settings", Require: plugin.Optional},
			},
		},
		Columns: []*plugin.Column{
			// Columns returned from the Cohere API
			{Name: "generation", Type: proto.ColumnType_STRING, Transform: transform.FromField("Generation.Text"), Description: "Generation for a given text prompt."},

			// Qual columns to provide input to the API
			{Name: "prompt", Type: proto.ColumnType_STRING, Transform: transform.FromQual("prompt"), Description: "The prompt to generate completions for, encoded as a string."},
			{Name: "settings", Type: proto.ColumnType_JSON, Transform: transform.FromQual("settings"), Description: "Settings is a JSONB object that accepts any of the completion API request parameters."},
		},
	}
}

// CompletionRequestQual defines the structure of the settings qual
type GenerationRequestQual struct {
	Model            *string  `json:"model"`
	Prompt           *string  `json:"prompt,omitempty"`
	MaxTokens        *int     `json:"max_tokens,omitempty"`
	Temperature      *float32 `json:"temperature,omitempty"`
	TopP             *float32 `json:"top_p,omitempty"`
	FrequencyPenalty *float32 `json:"frequency_penalty,omitempty"`
	PresencePenalty  *float32 `json:"presence_penalty,omitempty"`
	Stop             []string `json:"stop,omitempty"`
	Echo             *bool    `json:"echo,omitempty"`
	LogProbs         *int     `json:"logprobs,omitempty"`
	User             *string  `json:"user,omitempty"`
}

// CompletionRow defines the row structure returned from the API
type GenerationRow struct {
	coherego.Generation
	Prompt string
}

// listCompletion handles querying the Cohere AI API and returning completion data
func listGeneration(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {

	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("cohereai_generation.listCompletion", "connection_error", err)
		return nil, err
	}

	// Default settings taken from the Cohere API docs
	// https://docs.cohere.ai/reference/generate
	cr := coherego.GenerateOptions{
		Prompt:         d.EqualsQuals["prompt"].GetStringValue(),
		MaxTokens:      100,
		NumGenerations: 3,
	}

	settingsString := d.EqualsQuals["settings"].GetJsonbValue()
	if settingsString != "" {
		// Overwrite any settings provided in the settings qual. If a field
		// is not passed in the settings, then default to the settings above.
		var crQual GenerationRequestQual
		err := json.Unmarshal([]byte(settingsString), &crQual)
		if err != nil {
			plugin.Logger(ctx).Error("cohereai_generation.listGeneration", "unmarshal_error", err)
			return nil, err
		}
		if crQual.Model != nil {
			cr.Model = *crQual.Model
		}
		if crQual.Prompt != nil {
			cr.Prompt = *crQual.Prompt
		}
	}

	// Query the Cohere API
	resp, err := conn.Generate(cr)
	if err != nil {
		plugin.Logger(ctx).Error("cohereai_generation.listGeneration", "api_error", err)
		return nil, err
	}

	plugin.Logger(ctx).Trace("cohereai_generation.listGeneration", "response", resp)
	// Return completion data
	for _, c := range resp.Generations {
		row := GenerationRow{c, cr.Prompt}
		d.StreamListItem(ctx, row)
	}
	return nil, nil
}
