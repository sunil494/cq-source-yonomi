package plugin

import (
	"github.com/sunil494/cq-source-yonomi/client"
	"github.com/sunil494/cq-source-yonomi/resources"

	"github.com/cloudquery/plugin-sdk/v3/plugins/source"
	"github.com/cloudquery/plugin-sdk/v3/schema"
)

var (
	Version = "development"
)

func Plugin() *source.Plugin {
	return source.NewPlugin(
		"sunil494-yonomi",
		Version,
		schema.Tables{
			resources.YonomiDevicesTable(),
		},
		client.New,
	)
}
