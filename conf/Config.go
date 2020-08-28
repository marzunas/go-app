package conf

type Config struct {
	Port           int
	HashWaitPeriod int
}

var cfg *Config

func GetConfig() *Config {
	if cfg == nil {
		cfg = &Config{}
		// load the config properties from a file
		cfg.Port = 8080
		cfg.HashWaitPeriod = 5
	}
	return cfg
}
