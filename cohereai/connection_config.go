package cohere

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/schema"
)

type cohereConfig struct {
	APIKey *string `cty:"api_key" hcl:"api_key"`
}

var ConfigSchema = map[string]*schema.Attribute{
	"api_key": {
		Type: schema.TypeString,
	},
}

func ConfigInstance() interface{} {
	return &cohereConfig{}
}

func GetConfig(connection *plugin.Connection) cohereConfig {
	if connection == nil || connection.Config == nil {
		return cohereConfig{}
	}
	config, _ := connection.Config.(cohereConfig)
	return config
}
