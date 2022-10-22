package wireguardreceiver

import (
	"context"
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver/scraperhelper"
)

const (
	typeStr   = "wireguard_stats"
	stability = component.StabilityLevelInDevelopment
)

func NewFactory() component.ReceiverFactory {
	return component.NewReceiverFactory(
		typeStr,
		createDefaultReceiverConfig,
		component.WithMetricsReceiver(createMetricsReceiver, stability))
}

func createDefaultConfig() *Config {
	return &Config{
		ScraperControllerSettings: scraperhelper.ScraperControllerSettings{
			ReceiverSettings:   config.NewReceiverSettings(config.NewComponentID(typeStr)),
			CollectionInterval: 10 * time.Second,
		},
	}
}

func createDefaultReceiverConfig() config.Receiver {
	return createDefaultConfig()
}

func createMetricsReceiver(
	ctx context.Context,
	params component.ReceiverCreateSettings,
	config config.Receiver,
	consumer consumer.Metrics,
) (component.MetricsReceiver, error) {

	wireguardConfig := config.(*Config)
	dsr, err := newReceiver(wireguardConfig, params, consumer)
	if err != nil {
		return nil, err
	}

	return dsr, nil
}
