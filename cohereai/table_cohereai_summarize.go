package cohere

import (
	"context"
	"encoding/json"

	coherego "github.com/cohere-ai/cohere-go"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

// Defines the cohere_summarize table
func tableCohereSummarize(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "cohereai_summarize",
		Description: "Summarize text using Cohere AI.",
		List: &plugin.ListConfig{
			Hydrate: summarize,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "text", Require: plugin.Optional},
				{Name: "settings", Require: plugin.Optional},
			},
		},
		Columns: []*plugin.Column{
			// Columns returned from the Cohere API
			{Name: "id", Type: proto.ColumnType_STRING, Transform: transform.FromField("SummarizeResponse.ID"), Description: "ID for the given text."},
			{Name: "summary", Type: proto.ColumnType_STRING, Transform: transform.FromField("SummarizeResponse.Summary"), Description: "Summary for the given text."},

			// Qual columns to provide input to the API
			{Name: "text", Type: proto.ColumnType_STRING, Transform: transform.FromQual("text"), Description: "The text to summarize, encoded as a string."},
			{Name: "settings", Type: proto.ColumnType_JSON, Transform: transform.FromQual("settings"), Description: "Settings is a JSONB object that accepts any of the summarize API request parameters."},
		},
	}
}

// SummarizeRequestQual defines the structure of the settings qual
type SummarizeRequestQual struct {
	Text              *string  `json:"text,omitempty"`
	Format            *string  `json:"format,omitempty"`
	Length            *string  `json:"length,omitempty"`
	Extractiveness    *string  `json:"extractiveness,omitempty"`
	Temperature       *float64 `json:"temperature,omitempty"`
	AdditionalCommand *string  `json:"additional_command,omitempty"`
	Model             *string  `json:"model,omitempty"`
}

// SummarizeRow defines the row structure returned from the API
type SummarizeRow struct {
	coherego.SummarizeResponse
	Text string
}

// summarize handles querying the Cohere AI API and returning summarized text as summary
func summarize(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// Create the API client
	client, err := connect(ctx, d)
	if err != nil {
		return nil, err
	}

	// Build the SummarizeOptions from the request
	opts := coherego.SummarizeOptions{
		Text:   d.EqualsQuals["text"].GetStringValue(),
		Model:  "summarize-xlarge",
		Format: "paragraph",
	}
	settingString := d.EqualsQuals["settings"].GetJsonbValue()
	if settingString != "" {
		var crQual SummarizeRequestQual
		err := json.Unmarshal([]byte(settingString), &crQual)
		if err != nil {
			plugin.Logger(ctx).Error("cohereai_summarize.summarize", "connection_error", err)
			return nil, err
		}

		if crQual.Length != nil {
			opts.Length = *crQual.Length
		}
		if crQual.Extractiveness != nil {
			opts.Extractiveness = *crQual.Extractiveness
		}
		if crQual.Temperature != nil {
			opts.Temperature = crQual.Temperature
		}
		if crQual.Format != nil {
			opts.Format = *crQual.Format
		}
		if crQual.Model != nil {
			opts.Model = *crQual.Model
		}
		if crQual.AdditionalCommand != nil {
			opts.AdditionalCommand = *crQual.AdditionalCommand
		}
		if crQual.Text != nil {
			opts.Text = *crQual.Text
		}
	}

	// Make the Summarize API call
	resp, err := client.Summarize(opts)
	if err != nil {
		plugin.Logger(ctx).Error("cohereai_summarize.summarize", "api_error", err)
		return nil, err
	}

	plugin.Logger(ctx).Debug("cohereai_summarize.summarize", "summary", resp.Summary)
	// Return summarize data
	row := SummarizeRow{*resp, opts.Text}
	d.StreamListItem(ctx, row)
	return nil, nil
}
