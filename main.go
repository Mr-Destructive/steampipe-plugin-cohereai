package main

import (
	"github.com/steampipe/steampipe-plugin-cohereai/cohere"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{PluginFunc: cohere.Plugin})
}
