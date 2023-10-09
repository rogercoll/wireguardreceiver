package wireguardreceiver

import (
	"errors"

	"github.com/rogercoll/wireguardreceiver/internal/metadata"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/receiver/scraperhelper"
)

var _ component.Config = (*Config)(nil)

type Config struct {
	scraperhelper.ScraperControllerSettings `mapstructure:",squash"`
	// TODO: implement exclude option
	Exclude ExcludeConfig `mapstructure:"exclude"`

	// MetricsBuilderConfig config. Enable or disable stats by name.
	metadata.MetricsBuilderConfig `mapstructure:",squash"`
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
