package cohere

import (
	"context"
	"encoding/json"

	coherego "github.com/cohere-ai/cohere-go"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

// Defines the cohere_detokenize table
func tableCohereDetokenize(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "cohereai_detokenize",
		Description: "Detokenize tokens into text using Cohere AI.",
		List: &plugin.ListConfig{
			Hydrate: detokenize,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "tokens", Require: plugin.Optional},
			},
		},
		Columns: []*plugin.Column{
			// Columns returned from the Cohere API
			{Name: "text", Type: proto.ColumnType_STRING, Transform: transform.FromField("Text"), Description: "The detokenized text."},

			// Qual columns to provide input to the API
			{Name: "tokens", Type: proto.ColumnType(proto.ColumnType_STRING.Number()), Transform: transform.FromQual("tokens"), Description: "The tokens to detokenize, encoded as a JSON array."},
		},
	}
}

// DetokenizeRequestQual defines the structure of the settings qual
type DetokenizeRequestQual struct {
	Tokens []int64 `json:"tokens"`
}

// DetokenizeRow defines the row structure returned from the API
type DetokenizeRow struct {
	Text   string
	Tokens []int64
}

// detokenize handles querying the Cohere AI API and returning detokenized data as text
func detokenize(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// Create the API client
	client, err := connect(ctx, d)
	if err != nil {
		return nil, err
	}

	// Build the DetokenizeOptions from the request
	tokenList := d.EqualsQuals["tokens"].GetStringValue()
	var tokens []int64
	if tokenList != "" {
		err := json.Unmarshal([]byte(tokenList), &tokens)
		if err != nil {
			plugin.Logger(ctx).Error("cohereai_detokenize.detokenize", "connection_error", err)
			return nil, err
		}
	}
	opts := coherego.DetokenizeOptions{
		Tokens: tokens,
	}

	// Make the Detokenize API call
	resp, err := client.Detokenize(opts)
	if err != nil {
		plugin.Logger(ctx).Error("cohereai_detokenize.detokenize", "api_error", err)
		return nil, err
	}

	plugin.Logger(ctx).Debug("cohereai_detokenize.detokenize", "text", resp.Text)
	// Return detokenize data
	row := DetokenizeRow{
		Text:   resp.Text,
		Tokens: opts.Tokens,
	}
	d.StreamListItem(ctx, row)
	return nil, nil
}
