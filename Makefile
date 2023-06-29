install:
	go build -o ~/.steampipe/plugins/hub.steampipe.io/plugins/turbot/cohereai@latest/steampipe-plugin-cohereai.plugin *.go
local:
	/home/meet/code/playground/github/steampipe/go/bin/go build -o ~/.steampipe/plugins/hub.steampipe.io/plugins/turbot/cohereai@latest/steampipe-plugin-cohereai.plugin *.go
