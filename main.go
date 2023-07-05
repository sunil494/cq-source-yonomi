package main

import (
	"github.com/sunil494/cq-source-yonomi/plugin"

	"github.com/cloudquery/plugin-sdk/v3/serve"
)

func main() {
	serve.Source(plugin.Plugin())
}
