package cohere

import (
	"context"
	"encoding/json"

	coherego "github.com/cohere-ai/cohere-go"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

// Defines the cohere_detect_language table
func tableCohereDetectLanguage(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "cohereai_detect_language",
		Description: "Detect languages of texts using Cohere AI.",
		List: &plugin.ListConfig{
			Hydrate: detectLanguage,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "texts", Require: plugin.Optional},
			},
		},
		Columns: []*plugin.Column{
			// Columns returned from the Cohere API
			{Name: "language_name", Type: proto.ColumnType_STRING, Transform: transform.FromField("LanguageDetectResult.LanguageName"), Description: "The name of the detected language."},
			{Name: "language_code", Type: proto.ColumnType_STRING, Transform: transform.FromField("LanguageDetectResult.LanguageCode"), Description: "The ISO 639-1 code for the detected language."},

			// Qual columns to provide input to the API
			{Name: "text", Type: proto.ColumnType_STRING, Transform: transform.FromField("Text"), Description: "Text from which to detect it's language."},
			{Name: "texts", Type: proto.ColumnType_STRING, Transform: transform.FromQual("texts"), Description: "The texts to detect languages for, encoded as a JSON array."},
		},
	}
}

// DetectLanguageRequestQual defines the structure of the settings qual
type DetectLanguageRequestQual struct {
	Texts []string `json:"texts"`
}

// DetectLanguageRow defines the row structure returned from the API
type DetectLanguageRow struct {
	coherego.LanguageDetectResult
	Texts []string
	Text  string
}

// detectLanguage handles querying the Cohere AI API and returning detected language name and code
func detectLanguage(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// Create the API client
	client, err := connect(ctx, d)
	if err != nil {
		return nil, err
	}

	textString := d.EqualsQuals["texts"].GetStringValue()
	var texts []string
	if textString != "" {
		err := json.Unmarshal([]byte(textString), &texts)
		if err != nil {
			plugin.Logger(ctx).Error("cohereai_detect_language.detectLanguage", "unmarshal_error", err)
			return nil, err
		}
	}
	// Build the DetectLanguageOptions from the request
	opts := coherego.DetectLanguageOptions{
		Texts: texts,
	}

	// Make the DetectLanguage API call
	resp, err := client.DetectLanguage(opts)
	if err != nil {
		plugin.Logger(ctx).Error("cohereai_detect_language.detectLanguage", "api_error", err)
		return nil, err
	}

	plugin.Logger(ctx).Debug("cohereai_detect_language.detectLanguage", "response", resp)
	// Return detect language data
	for i, result := range resp.Results {
		rows := DetectLanguageRow{
			result,
			opts.Texts,
			opts.Texts[i],
		}
		d.StreamListItem(ctx, rows)
	}
	return nil, nil
}
