package mysql

import "fmt"

type Config struct {
	Host     string
	User     string
	Password string
	Database string
	Charset  string
	Identity string
}

func (c *Config) DSN() string {
	template := "%v:%v@tcp(%v)/%v?charset=%v"
	return fmt.Sprintf(template, c.User, c.Password, c.Host, c.Database, c.Charset)
}

type ConfigOption func(config *Config)

func useOption(config *Config, ops []ConfigOption) {
	for _, option := range ops {
		option(config)
	}
}

func OptionDefault(config *Config) {
	if config.Charset == "" {
		config.Charset = "utf8"
	}
	if config.Identity == "" {
		config.Identity = DEFAULT
	}
}

func OptionIdentity(identity string) ConfigOption {
	return func(config *Config) {
		config.Identity = identity
	}
}
