package config

import (
	"fmt"
	"log/slog"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Server ServerConfig
	Logger LoggerConfig
}

type ServerConfig struct {
	Address string `envconfig:"EXCHANGES_SERVER_ADDR" default:":8000"`
}

type LoggerConfig struct {
	Level slog.Level `envconfig:"EXCHANGES_LOGGER_LEVEL" default:"Info"`
}

func Load() (*Config, error) {
	cfg := new(Config)

	if err := envconfig.Process("", cfg); err != nil {
		return nil, fmt.Errorf("error occured while parsing env config: %w", err)
	}
	return cfg, nil
}
