package wireguardreceiver

import (
	"errors"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/receiver/scraperhelper"
)

var _ component.Config = (*Config)(nil)

type Config struct {
	scraperhelper.ScraperControllerSettings `mapstructure:",squash"`
	// TODO: implement exclude option
	Exclude ExcludeConfig `mapstructure:"exclude"`
}

type ExcludeConfig struct {
	Interface ExcludeInterfaceConfig `mapstructure:"interface"`
}

type ExcludeInterfaceConfig struct {
	Names []string `mapstructure:"names"`
}

func (c Config) Validate() error {
	if c.CollectionInterval == 0 {
		return errors.New("config.CollectorInterval must be specified")
	}
	return nil
}
