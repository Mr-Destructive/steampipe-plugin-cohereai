package cohere

import (
	"context"
	"encoding/json"

	coherego "github.com/cohere-ai/cohere-go"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

// Defines the cohereai_classification table
func tableCohereClassification(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "cohereai_classification",
		Description: "Classification in Cohere AI.",
		List: &plugin.ListConfig{
			Hydrate: listClassification,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "inputs", Require: plugin.Optional},
				{Name: "examples", Require: plugin.Optional},
			},
		},
		Columns: []*plugin.Column{
			// Columns returned from the Cohere API
			{Name: "classification", Type: proto.ColumnType_STRING, Transform: transform.FromField("Classification.Prediction"), Description: "The classification results for the given input text(s)."},

			{Name: "inputs", Type: proto.ColumnType_STRING, Transform: transform.FromQual("inputs"), Description: "The input text that was classified."},
			{Name: "examples", Type: proto.ColumnType_STRING, Transform: transform.FromQual("examples"), Description: "The example text classified."},
			{Name: "settings", Type: proto.ColumnType_JSON, Transform: transform.FromQual("settings"), Description: "Settings is a JSONB object that accepts any of the completion API request parameters."},
		},
	}
}

// ClassificationRequestQual defines the structure of the settings qual
type ClassificationRequestQual struct {
	Model    *string `json:"model"`
	Inputs   *string `json:"inputs"`
	Examples *string `json:"examples"`
}

// ClassificationRow defines the row structure returned from the API
type ClassificationRow struct {
	coherego.Classification
	Input []string
}

// listCompletion handles querying the Cohere AI API and returning tokenize data
func listClassification(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {

	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("cohereai_classification.listClassification", "connection_error", err)
		return nil, err
	}

	// Default settings taken from the Cohere API docs
	// https://docs.cohere.ai/reference/classify
	var inputs []string
	var examples []coherego.Example
	inputsString := d.EqualsQuals["inputs"].GetStringValue()
	if inputsString != "" {
		err := json.Unmarshal([]byte(inputsString), &inputs)
		if err != nil {
			plugin.Logger(ctx).Error("cohereai_classification.listClassification", "unmarshal_error", err)
			return nil, err
		}
	}
	exampleString := d.EqualsQuals["examples"].GetStringValue()
	if exampleString != "" {
		err := json.Unmarshal([]byte(exampleString), &examples)
		if err != nil {
			plugin.Logger(ctx).Error("cohereai_classification.listClassification", "unmarshal_error", err)
			return nil, err
		}
	}
	cr := coherego.ClassifyOptions{
		Model:    "large",
		Inputs:   inputs,
		Examples: examples,
	}

	settingsString := d.EqualsQuals["settings"].GetJsonbValue()
	if settingsString != "" {
		var crQual ClassificationRequestQual
		err := json.Unmarshal([]byte(settingsString), &crQual)
		if err != nil {
			plugin.Logger(ctx).Error("cohereai_classification.listClassification", "unmarshal_error", err)
			return nil, err
		}
	}

	// Query the Cohere API
	resp, err := conn.Classify(cr)
	if err != nil {
		plugin.Logger(ctx).Error("cohereai_classification.listClassification", "api_error", err)
		return nil, err
	}
	// Return tokenize data
	for _, c := range resp.Classifications {
		row := ClassificationRow{c, cr.Inputs}
		d.StreamListItem(ctx, row)
	}
	return nil, nil
}
