package database

import (
	"github.com/jackc/pgx/v5/tracelog"
	"github.com/vildan-valeev/melvad_test/pkg/database/zerologadapter"
)

// Option -.
type Option func(*Options)

// WithLogger .
func WithLogger(log *zerologadapter.Logger) Option {
	return func(c *Options) {
		c.log = log
	}
}

// WithLogLevel .
func WithLogLevel(lvl string) Option {
	return func(c *Options) {
		level, err := tracelog.LogLevelFromString(lvl)
		if err == nil {
			c.level = level
		}
	}
}
