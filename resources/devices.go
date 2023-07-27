package resources

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/transformers"
	"github.com/sunil494/cq-source-yonomi/client"
	"github.com/sunil494/cq-source-yonomi/internal/yonomi"
)

func YonomiDevicesTable() *schema.Table {
	return &schema.Table{
		Name:      "yonomi_devices",
		Resolver:  fetchYonomiDevicesData,
		Multiplex: client.DeviceMultiplex,
		Transform: transformers.TransformWithStruct(&yonomi.DevicesDataBlock{}),
	}
}

func fetchYonomiDevicesData(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	device := c.Device
	data, err := yonomi.GetDevicesData(c.Spec.Authorization, device.DeviceId)
	if err != nil {
		return err
	}
	res <- data
	return nil
}
