package config

import (
	"github.com/caarlos0/env/v9"
	"github.com/rs/zerolog/log"
)

type Config struct {
	LogLevel string `env:"LOGLEVEL,required" envDefault:"debug"` // debug, info, warn, error, fatal, ""
	IP       string `env:"IP,required" envDefault:"0.0.0.0"`
	HTTPPort string `env:"HTTP_PORT,required" envDefault:"8000"`
	DSN      string `env:"TN_SIGN_PG_DSN,required"  envDefault:"8000"`
}

// NewConfig создает экземпляр Config и заполняет его значениями переменных окружения.
func NewConfig() *Config {
	log.Info().Msg("Init config")
	cfg := &Config{}

	if err := env.Parse(cfg); err != nil {
		log.Fatal().Err(err).Msg("parse env")
	}

	return cfg
}
