package cohere

import (
	"context"

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
			},
		},
		Columns: []*plugin.Column{
			// Columns returned from the Cohere API
			{Name: "summary", Type: proto.ColumnType_STRING, Transform: transform.FromField("Summary"), Description: "Summary for the given text."},
			{Name: "text", Type: proto.ColumnType_STRING, Transform: transform.FromQual("text"), Description: "The text to summarize, encoded as a string."},
		},
	}
}

// SummarizeRequestQual defines the structure of the settings qual
type SummarizeRequestQual struct {
	Text string `json:"text"`
}

// SummarizeRow defines the row structure returned from the API
type SummarizeRow struct {
	Summary string
	Text    string
}

// summarize handles querying the Cohere AI API and returning summarize data
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

	// Make the Summarize API call
	resp, err := client.Summarize(opts)
	if err != nil {
		return nil, err
	}

	// Return summarize data
	row := SummarizeRow{resp.Summary, opts.Text}
	d.StreamListItem(ctx, row)
	return nil, nil
}
