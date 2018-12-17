package internal

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Host     string `envconfig:"MONGO_HOST" required:"true"`
	Database string `envconfig:"MONGO_DB" required:"true"`
	User     string `envconfig:"MONGO_USER" required:"true"`
	Password string `envconfig:"MONGO_PASSWORD" required:"true"`
}

func NewConfig() (*Config, error) {
	var err error

	config := &Config{}

	if err = envconfig.Process("", config); err != nil {
		return nil, err
	}

	return config, nil
}
