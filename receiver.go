package wireguardreceiver

import (
	"context"
	"sync"
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"

	"go.opentelemetry.io/collector/receiver/scraperhelper"
)

type receiver struct {
	config   *Config
	wgClient wireguardClient
}

func newReceiver(config *Config, set component.ReceiverCreateSettings, nextConsumer consumer.Metrics) (component.MetricsReceiver, error) {
	err := config.Validate()
	if err != nil {
		return nil, err
	}

	client, err := newWireguardClient()
	if err != nil {
		return nil, err
	}

	recv := &receiver{
		config:   config,
		wgClient: client,
	}

	scrp, err := scraperhelper.NewScraper("wireguardReceiver", recv.scrape)
	if err != nil {
		return nil, err
	}
	return scraperhelper.NewScraperControllerReceiver(&recv.config.ScraperControllerSettings, set, nextConsumer, scraperhelper.AddScraper(scrp))
}

func (r *receiver) scrape(ctx context.Context) (pmetric.Metrics, error) {
	devices := r.wgClient.Devices()
	results := make(chan pmetric.Metrics)

	wg := &sync.WaitGroup{}
	wg.Add(len(devices))
	for _, d := range devices {
		go func(d wgtypes.Device) {
			defer wg.Done()
			for _, peer := range d.Peers {
				results <- peerToMetrics(time.Now(), d.Name, &peer)
			}
		}(d)
	}

	wg.Wait()
	close(results)

	md := pmetric.NewMetrics()
	for res := range results {
		res.md.ResourceMetrics().CopyTo(md.ResourceMetrics())
	}
	return md, nil
}
