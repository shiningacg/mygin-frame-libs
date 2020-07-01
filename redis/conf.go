package redis

import "github.com/go-redis/redis/v8"

type Config struct {
	Host     string
	Secret   string
	Identity string
}

func (c *Config) Adapter() *redis.Options {
	return &redis.Options{
		Addr:     c.Host,
		Password: c.Secret,
	}
}

func useOption(config *Config, ops []ConfigOption) {
	for _, option := range ops {
		option(config)
	}
}

type ConfigOption func(config *Config)

func OptionDefault(config *Config) {
	if config.Identity == "" {
		config.Identity = DEFAULT
	}
}
