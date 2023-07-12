package cohere

import (
	"context"
	"encoding/json"

	coherego "github.com/cohere-ai/cohere-go"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

// Defines the cohere_tokenize table
func tableCohereTokenize(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "cohereai_tokenize",
		Description: "Tokenize in Cohere AI.",
		List: &plugin.ListConfig{
			Hydrate: tokenize,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "text", Require: plugin.Optional},
			},
		},
		Columns: []*plugin.Column{
			// Columns returned from the Cohere API
			{Name: "tokens", Type: proto.ColumnType_STRING, Transform: transform.FromField("Tokens.Tokens"), Description: "Tokens for a given text prompt."},
			{Name: "token_strings", Type: proto.ColumnType_STRING, Transform: transform.FromField("Tokens.TokenStrings"), Description: "Tokens in the form of input string."},

			// Qual columns to provide input to the API
			{Name: "text", Type: proto.ColumnType_STRING, Transform: transform.FromQual("text"), Description: "The text to tokenize for, encoded as a string."},
		},
	}
}

// CompletionRequestQual defines the structure of the settings qual
type TokenizeRequestQual struct {
	Text *string `json:"text,omitempty"`
}

// SummarizeRow defines the row structure returned from the API
type TokenizeRow struct {
	Tokens coherego.TokenizeResponse
	Text   string
}

// tokenize handles querying the Cohere AI API and returning tokens from provided text
func tokenize(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {

	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("cohereai_tokenize.tokenize", "connection_error", err)
		return nil, err
	}

	// Default settings taken from the Cohere API docs
	// https://docs.cohere.ai/reference/generate
	cr := coherego.TokenizeOptions{
		Text: d.EqualsQuals["text"].GetStringValue(),
	}

	settingsString := d.EqualsQuals["settings"].GetJsonbValue()
	if settingsString != "" {
		// Overwrite any settings provided in the settings qual. If a field
		// is not passed in the settings, then default to the settings above.
		var crQual TokenizeRequestQual
		err := json.Unmarshal([]byte(settingsString), &crQual)
		if err != nil {
			plugin.Logger(ctx).Error("cohereai_tokenize.tokenize", "connection_error", err)
			return nil, err
		}
	}

	// Query the Cohere API
	resp, err := conn.Tokenize(cr)
	if err != nil {
		plugin.Logger(ctx).Error("cohereai_tokenize.tokenize", "api_error", err)
		return nil, err
	}

	plugin.Logger(ctx).Trace("cohereai_tokenize.tokenize", "response", resp)
	// Return tokenize data
	row := TokenizeRow{*resp, cr.Text}
	d.StreamListItem(ctx, row)
	return nil, nil
}
