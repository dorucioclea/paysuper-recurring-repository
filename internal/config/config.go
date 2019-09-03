package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	MetricsPort   string `envconfig:"METRICS_PORT" required:"false" default:"8085"`
	MicroSelector string `envconfig:"MICRO_SELECTOR" default:""`
}

func NewConfig() (*Config, error) {
	cfg := &Config{}
	err := envconfig.Process("", cfg)

	return cfg, err
}
