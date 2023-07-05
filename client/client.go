package client

import (
	"context"
	"fmt"

	"github.com/cloudquery/plugin-pb-go/specs"
	"github.com/cloudquery/plugin-sdk/v3/plugins/source"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/rs/zerolog"
	"github.com/sunil494/cq-source-yonomi/internal/yonomi"
)

type Client struct {
	Logger zerolog.Logger
	Yonomi *yonomi.Client
	Spec   *Spec

	Device YonomiDeviceConfigBlock
}

func (c *Client) ID() string {
	return fmt.Sprintf("yonomi:%s", c.Device.Name)
}

func (c *Client) WithDevice(device YonomiDeviceConfigBlock) *Client {
	newC := *c
	newC.Logger = c.Logger.With().Str("device", device.Name).Logger()
	newC.Device = device
	return &newC
}

func New(ctx context.Context, logger zerolog.Logger, s specs.Source, opts source.Options) (schema.ClientMeta, error) {
	var pluginSpec Spec

	if err := s.UnmarshalSpec(&pluginSpec); err != nil {
		return nil, fmt.Errorf("failed to unmarshal plugin spec: %w", err)
	}

	c, err := yonomi.NewClient()
	if err != nil {
		return nil, err
	}

	return &Client{
		Logger: logger,
		Yonomi: c,
		Spec:   &pluginSpec,
	}, nil
}
