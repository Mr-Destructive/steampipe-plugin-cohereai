package cohere

import (
	"context"
	"encoding/json"

	coherego "github.com/cohere-ai/cohere-go"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

// Defines the cohere_embed table
func tableCohereEmbed(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "cohereai_embed",
		Description: "Get embeddings from Cohere AI.",
		List: &plugin.ListConfig{
			Hydrate: embed,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "texts", Require: plugin.Optional},
			},
		},
		Columns: []*plugin.Column{
			// Columns returned from the Cohere API
			{Name: "embeddings", Type: proto.ColumnType_STRING, Transform: transform.FromField("Embeddings"), Description: "Embeddings for the given texts."},
			{Name: "texts", Type: proto.ColumnType_STRING, Transform: transform.FromQual("texts"), Description: "The texts to embed, encoded as a JSON array."},
		},
	}
}

// EmbedRequestQual defines the structure of the settings qual
type EmbedRequestQual struct {
	Texts []string `json:"texts"`
}

// EmbedRow defines the row structure returned from the API
type EmbedRow struct {
	Embeddings [][]float64
	Texts      []string
}

// embed handles querying the Cohere AI API and returning embed data
func embed(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// Create the API client
	client, err := connect(ctx, d)
	if err != nil {
		return nil, err
	}

	// Build the EmbedOptions from the request
	textString := d.EqualsQuals["texts"].GetStringValue()
	var texts []string
	if textString != "" {
		err := json.Unmarshal([]byte(textString), &texts)
		if err != nil {
			plugin.Logger(ctx).Error("cohereai_embed.embed", "unmarshal_error", err)
			return nil, err
		}
	}
	opts := coherego.EmbedOptions{
		Model: "small",
		Texts: texts,
	}

	// Make the Embed API call
	resp, err := client.Embed(opts)
	if err != nil {
		return nil, err
	}

	// Return embed data
	row := EmbedRow{resp.Embeddings, opts.Texts}
	d.StreamListItem(ctx, row)
	return nil, nil
}
