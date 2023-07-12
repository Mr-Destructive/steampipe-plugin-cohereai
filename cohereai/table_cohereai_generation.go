package cohere

import (
	"context"
	"encoding/json"

	coherego "github.com/cohere-ai/cohere-go"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

// Defines the cohereai_generation table
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
			{Name: "likelihood", Type: proto.ColumnType_DOUBLE, Transform: transform.FromField("Generation.Likelihood"), Description: "The likelihood of the generated text."},
			{Name: "token_likelihoods", Type: proto.ColumnType_JSON, Transform: transform.FromField("Generation.TokenLikelihoods"), Description: "The likelihood of the generated tokens/prompt."},

			// Qual columns to provide input to the API
			{Name: "prompt", Type: proto.ColumnType_STRING, Transform: transform.FromQual("prompt"), Description: "The prompt to get generations for, encoded as a string."},
			{Name: "settings", Type: proto.ColumnType_JSON, Transform: transform.FromQual("settings"), Description: "Settings is a JSONB object that accepts any of the generate API request parameters."},
		},
	}
}

// GenerationRequestQual defines the structure of the settings qual
type GenerationRequestQual struct {
	Model             *string  `json:"model,omitempty"`
	Prompt            *string  `json:"prompt,omitempty"`
	MaxTokens         *uint    `json:"max_tokens,omitempty"`
	NumGenerations    *int     `json:"num_generations,omitempty"`
	Temperature       *float64 `json:"temperature,omitempty"`
	TopP              *float64 `json:"top_p,omitempty"`
	FrequencyPenalty  *float64 `json:"frequency_penalty,omitempty"`
	PresencePenalty   *float64 `json:"presence_penalty,omitempty"`
	Stop              []string `json:"stop,omitempty"`
	Preset            string   `json:"preset,omitempty"`
	ReturnLikelihoods *string  `json:"return_likelihoods,omitempty"`
}

// GenerationRow defines the row structure returned from the API
type GenerationRow struct {
	coherego.Generation
	Prompt string
}

// listGeneration handles querying the Cohere AI API and returning generation data
func listGeneration(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {

	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("cohereai_generation.listCompletion", "connection_error", err)
		return nil, err
	}

	// Default settings taken from the Cohere API docs
	// https://docs.cohere.ai/reference/generate
	var maxTokens uint = 100
	var numGenerations int = 3
	cr := coherego.GenerateOptions{
		Prompt:            d.EqualsQuals["prompt"].GetStringValue(),
		MaxTokens:         &maxTokens,
		NumGenerations:    &numGenerations,
		ReturnLikelihoods: "GENERATION",
	}

	settingsString := d.EqualsQuals["settings"].GetJsonbValue()
	if settingsString != "" {
		// Overwrite any settings provided in the settings qual. If a field
		// is not passed in the settings, then default to the settings above.
		var crQual GenerationRequestQual
		err := json.Unmarshal([]byte(settingsString), &crQual)
		if err != nil {
			plugin.Logger(ctx).Error("cohereai_generation.listGeneration", "connection_error", err)
			return nil, err
		}
		if crQual.Model != nil {
			cr.Model = *crQual.Model
		}
		if crQual.Prompt != nil {
			cr.Prompt = *crQual.Prompt
		}
		if crQual.MaxTokens != nil {
			cr.MaxTokens = crQual.MaxTokens
		}
		if crQual.Temperature != nil {
			cr.Temperature = crQual.Temperature
		}
		if crQual.TopP != nil {
			cr.P = crQual.TopP
		}
		if crQual.FrequencyPenalty != nil {
			cr.FrequencyPenalty = crQual.FrequencyPenalty
		}
		if crQual.PresencePenalty != nil {
			cr.PresencePenalty = crQual.PresencePenalty
		}
		if crQual.Stop != nil {
			cr.StopSequences = crQual.Stop
		}
		if crQual.NumGenerations != nil {
			cr.NumGenerations = crQual.NumGenerations
		}
		if crQual.Preset != "" {
			cr.Preset = crQual.Preset
		}
		if cr.Prompt == "" {
			cr.Prompt = *crQual.Prompt
		}
		if crQual.ReturnLikelihoods != nil {
			cr.ReturnLikelihoods = *crQual.ReturnLikelihoods
		}
	}

	// Query the Cohere API
	resp, err := conn.Generate(cr)
	if err != nil {
		plugin.Logger(ctx).Error("cohereai_generation.listGeneration", "api_error", err)
		return nil, err
	}

	plugin.Logger(ctx).Trace("cohereai_generation.listGeneration", "response", resp)
	// Return generation data
	for _, c := range resp.Generations {
		row := GenerationRow{c, cr.Prompt}
		d.StreamListItem(ctx, row)
	}
	return nil, nil
}
