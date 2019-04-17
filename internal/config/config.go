package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
    Host          string `envconfig:"MONGO_HOST" required:"true"`
    Database      string `envconfig:"MONGO_DB" required:"true"`
    User          string `envconfig:"MONGO_USER" default:""`
    Password      string `envconfig:"MONGO_PASSWORD" default:""`
    MetricsPort   string `envconfig:"METRICS_PORT" required:"false" default:"8085"`
    MicroRegistry string `envconfig:"MICRO_REGISTRY" required:"false"`
}

func NewConfig() (*Config, error) {
    cfg := &Config{}
    err := envconfig.Process("", cfg)

    return cfg, err
}
