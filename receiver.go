package wireguardreceiver

import (
	"context"
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver"

	"go.opentelemetry.io/collector/receiver/scraperhelper"
)

type wgreceiver struct {
	config        *Config
	wgClient      wireguardClient
	clientFactory clientFactory
}

func newReceiver(config *Config, set receiver.CreateSettings, nextConsumer consumer.Metrics, clientFactory clientFactory) (receiver.Metrics, error) {
	err := config.Validate()
	if err != nil {
		return nil, err
	}

	if clientFactory == nil {
		clientFactory = newWireguardClient
	}

	recv := &wgreceiver{
		config:        config,
		clientFactory: clientFactory,
	}

	scrp, err := scraperhelper.NewScraper(typeStr, recv.scrape, scraperhelper.WithStart(recv.start))
	if err != nil {
		return nil, err
	}
	return scraperhelper.NewScraperControllerReceiver(&recv.config.ScraperControllerSettings, set, nextConsumer, scraperhelper.AddScraper(scrp))
}

func (r *wgreceiver) start(_ context.Context, _ component.Host) error {
	var err error
	r.wgClient, err = r.clientFactory()
	if err != nil {
		return err
	}

	return nil
}

func (r *wgreceiver) scrape(ctx context.Context) (pmetric.Metrics, error) {
	md := pmetric.NewMetrics()

	devices, err := r.wgClient.Devices()
	if err != nil {
		return md, err
	}

	for _, d := range devices {
		for _, peer := range d.Peers {
			peerToMetrics(time.Now(), d.Name, &peer).ResourceMetrics().CopyTo(md.ResourceMetrics())
		}
	}

	return md, nil
}
