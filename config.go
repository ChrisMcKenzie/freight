package freight

type Config struct {
	Base     BaseConfig `mapstructure:"freight", hcl:"freight,squash"`
	Projects []*Project `mapstructure:"project" hcl:"project"`
}

// BaseConfig defines the top level or "root" config options available in a freight
// config file
type BaseConfig struct {
	Root string `mapstructure:"root", hcl:"root"`
}
