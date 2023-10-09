package wireguardreceiver

import (
	"context"
	"time"

	"github.com/rogercoll/wireguardreceiver/internal/metadata"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"
	"go.opentelemetry.io/collector/receiver/scraperhelper"
)

func NewFactory() receiver.Factory {
	return receiver.NewFactory(
		metadata.Type,
		createDefaultReceiverConfig,
		receiver.WithMetrics(createMetricsReceiver, metadata.MetricsStability))
}

func createDefaultConfig() *Config {
	cfg := scraperhelper.NewDefaultScraperControllerSettings(metadata.Type)
	cfg.CollectionInterval = 10 * time.Second
	cfg.Timeout = 5 * time.Second
	return &Config{
		ScraperControllerSettings: cfg,
		MetricsBuilderConfig:      metadata.DefaultMetricsBuilderConfig(),
	}
}

func createDefaultReceiverConfig() component.Config {
	return createDefaultConfig()
}

func createMetricsReceiver(
	_ context.Context,
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
