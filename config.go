package wireguardreceiver

//var _ config.Receiver = (*Config)(nil)

type Config struct {
	Exclude ExcludeConfig `mapstructure:"exclude"`
}

type ExcludeConfig struct {
	Interface ExcludeInterfaceConfig `mapstructure:"interface"`
}

type ExcludeInterfaceConfig struct {
	Names []string `mapstructure:"names"`
}
