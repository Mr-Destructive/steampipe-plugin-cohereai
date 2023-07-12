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
		Name:        "cohereai_classify",
		Description: "Classification in Cohere AI.",
		List: &plugin.ListConfig{
			Hydrate: listClassification,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "inputs", Require: plugin.Optional},
				{Name: "examples", Require: plugin.Optional},
				{Name: "settings", Require: plugin.Optional},
			},
		},
		Columns: []*plugin.Column{
			// Columns returned from the Cohere API
			{Name: "id", Type: proto.ColumnType_STRING, Transform: transform.FromField("Classification.ID"), Description: "The ID of the classification."},
			{Name: "classification", Type: proto.ColumnType_STRING, Transform: transform.FromField("Classification.Prediction"), Description: "The classification results for the given input text(s)."},
			{Name: "confidence", Type: proto.ColumnType_DOUBLE, Transform: transform.FromField("Classification.Confidence"), Description: "The confidence score of the classification."},
			{Name: "labels", Type: proto.ColumnType_JSON, Transform: transform.FromField("Classification.Labels"), Description: "The labels of the classification."},

			// Qual columns to provide input to the API
			{Name: "inputs", Type: proto.ColumnType_STRING, Transform: transform.FromQual("inputs"), Description: "The input text that was classified."},
			{Name: "examples", Type: proto.ColumnType_STRING, Transform: transform.FromQual("examples"), Description: "The example text classified."},
			{Name: "settings", Type: proto.ColumnType_JSON, Transform: transform.FromQual("settings"), Description: "Settings is a JSONB object that accepts any of the classify API request parameters."},
		},
	}
}

// ClassificationRequestQual defines the structure of the settings qual
type ClassificationRequestQual struct {
	Model    *string             `json:"model,omitempty"`
	Inputs   *[]string           `json:"inputs,omitempty"`
	Examples *[]coherego.Example `json:"examples,omitempty"`
	Preset   *string             `json:"preset,omitempty"`
	Truncate string              `json:"truncate,omitempty"`
}

// ClassificationRow defines the row structure returned from the API
type ClassificationRow struct {
	coherego.Classification
	Input []string
}

// listClassification handles querying the Cohere AI API and returning a list of labels
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
			plugin.Logger(ctx).Error("cohereai_classification.listClassification", "connection_error", err)
			return nil, err
		}
	}
	exampleString := d.EqualsQuals["examples"].GetStringValue()
	if exampleString != "" {
		err := json.Unmarshal([]byte(exampleString), &examples)
		if err != nil {
			plugin.Logger(ctx).Error("cohereai_classification.listClassification", "connection_error", err)
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
			plugin.Logger(ctx).Error("cohereai_classification.listClassification", "connection_error", err)
			return nil, err
		}
		if crQual.Model != nil {
			cr.Model = *crQual.Model
		}
		if crQual.Preset != nil {
			cr.Preset = *crQual.Preset
		}
	}

	// Query the Cohere API
	resp, err := conn.Classify(cr)
	if err != nil {
		plugin.Logger(ctx).Error("cohereai_classification.listClassification", "api_error", err)
		return nil, err
	}

	plugin.Logger(ctx).Trace("cohereai_classification.listClassification", "response", resp)
	// Return tokenize data
	for _, c := range resp.Classifications {
		row := ClassificationRow{c, cr.Inputs}
		d.StreamListItem(ctx, row)
	}
	return nil, nil
}
