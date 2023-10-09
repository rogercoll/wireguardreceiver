package wireguardreceiver

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/component/componenttest"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver/receivertest"
	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
)

func genDevice() (*wgtypes.Device, error) {
	peer, err := getPeer()
	if err != nil {
		return nil, err
	}
	return &wgtypes.Device{
		Name:  "wg0",
		Peers: []wgtypes.Peer{*peer},
	}, nil
}

func TestScrape(t *testing.T) {
	cfg := createDefaultConfig()
	cfg.CollectionInterval = 100 * time.Millisecond

	client := make(mockClient)
	consumer := make(mockConsumer)

	r, err := newReceiver(cfg, receivertest.NewNopCreateSettings(), consumer, client.factory)
	require.NoError(t, err)
	assert.NotNil(t, r)

	device, err := genDevice()
	require.NoError(t, err)

	go func() {
		client <- deviceResult{
			devices: []*wgtypes.Device{device},
			err:     nil,
		}
	}()

	assert.NoError(t, r.Start(context.Background(), componenttest.NewNopHost()))

	md := <-consumer
	assert.Equal(t, 1, md.ResourceMetrics().Len())

	assert.NoError(t, r.Shutdown(context.Background()))
}

type deviceResult struct {
	err     error
	devices []*wgtypes.Device
}

type mockClient chan deviceResult

func (c mockClient) factory() (wireguardClient, error) {
	return c, nil
}

func (c mockClient) Devices() ([]*wgtypes.Device, error) {
	report := <-c
	if report.err != nil {
		return nil, report.err
	}
	return report.devices, nil
}

type mockConsumer chan pmetric.Metrics

func (m mockConsumer) Capabilities() consumer.Capabilities {
	return consumer.Capabilities{}
}

func (m mockConsumer) ConsumeMetrics(ctx context.Context, md pmetric.Metrics) error {
	m <- md
	return nil
}
