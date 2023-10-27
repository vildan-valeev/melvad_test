package config

import (
	"github.com/caarlos0/env/v9"
	"github.com/rs/zerolog/log"
)

type Config struct {
	LogLevel string `env:"LOGLEVEL,required" envDefault:"debug"` // debug, info, warn, error, fatal, ""
	IP       string `env:"IP,required" envDefault:"0.0.0.0"`
	HTTPPort string `env:"HTTP_PORT,required" envDefault:"8000"`
	DSN      string `env:"PG_DSN,required"  envDefault:"postgres://postgres:postgres@localhost:5432/postgres"`

	RedisHost string `env:"REDIS_HOST,required"  envDefault:"0.0.0.0"`
	RedisPort string `env:"REDIS_PORT,required"  envDefault:"6379"`
	RedisDB   int    `env:"REDIS_DB,required"  envDefault:"0"`
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
