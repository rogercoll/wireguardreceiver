package wireguardreceiver

import (
	"context"
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"
	"go.opentelemetry.io/collector/receiver/scraperhelper"
)

const (
	typeStr   = "wireguard_stats"
	stability = "terst"
)

func NewFactory() receiver.Factory {
	return receiver.NewFactory(
		typeStr,
		createDefaultReceiverConfig,
		receiver.WithMetrics(createMetricsReceiver, component.StabilityLevelAlpha))
}

func createDefaultConfig() *Config {
	scs := scraperhelper.NewDefaultScraperControllerSettings(typeStr)
	scs.CollectionInterval = 10 * time.Second
	scs.Timeout = 5 * time.Second
	return &Config{
		ScraperControllerSettings: scs,
	}
}

func createDefaultReceiverConfig() component.Config {
	return createDefaultConfig()
}

func createMetricsReceiver(
	ctx context.Context,
	params receiver.CreateSettings,
	config component.Config,
	consumer consumer.Metrics,
) (receiver.Metrics, error) {

	wireguardConfig := config.(*Config)
	dsr, err := newReceiver(wireguardConfig, params, consumer, nil)
	if err != nil {
		return nil, err
	}

	return dsr, nil
}
