package cohere

import (
	"context"
	"errors"
	"os"
	"strings"

	coherego "github.com/cohere-ai/cohere-go"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func connect(ctx context.Context, d *plugin.QueryData) (*coherego.Client, error) {
	cacheKey := "cohereai"
	if cachedData, ok := d.ConnectionManager.Cache.Get(cacheKey); ok {
		return cachedData.(*coherego.Client), nil
	}

	conn, err := connectCached(ctx, d, nil)
	if err != nil {
		return nil, err
	}
	return conn.(*coherego.Client), nil
}

var connectCached = plugin.HydrateFunc(connectUncached).Memoize()

func connectUncached(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (any, error) {

	var conn *coherego.Client

	// Default to the env var settings
	apiKey := os.Getenv("COHERE_API_KEY")

	// Prefer config settings
	cohereConfig := GetConfig(d.Connection)
	if cohereConfig.APIKey != nil {
		apiKey = *cohereConfig.APIKey
	}

	// Error if the minimum config is not set
	if apiKey == "" {
		return conn, errors.New("api_key must be configured")
	}

	conn, err := coherego.CreateClient(apiKey)
	if err != nil {
		return conn, err
	}

	return conn, nil
}

func isNotFoundError(err error) bool {
	return strings.Contains(err.Error(), "status code: 404")
}
