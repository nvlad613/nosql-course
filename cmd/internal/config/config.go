package config

import (
	"fmt"
	"nosql-course/cmd/internal/logger"

	"github.com/pkg/errors"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Port   int           `required:"true"`
	Host   string        `required:"true"`
	Logger logger.Config `envconfig:"LOG"`
}

func Read() (*Config, error) {
	var cfg Config
	if err := envconfig.Process("app", &cfg); err != nil {
		return nil, errors.Wrap(err, "cannot read config")
	}
	return &cfg, nil
}

func (c *Config) Addr() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}
