package client

import (
	"github.com/cloudquery/plugin-sdk/v3/schema"
)

func DeviceMultiplex(meta schema.ClientMeta) []schema.ClientMeta {
	var l = make([]schema.ClientMeta, 0)
	client := meta.(*Client)
	for index := range client.Spec.Devices {
		l = append(l, client.WithDevice(client.Spec.Devices[index]))
	}
	return l
}
