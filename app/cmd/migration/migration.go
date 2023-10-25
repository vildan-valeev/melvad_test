package main

import (
	"context"
	"embed"

	tern "github.com/jackc/tern/migrate"

	"github.com/caarlos0/env/v9"
)

//go:embed sql
var MgrsFS embed.FS

type Migration struct{}

type Config struct {
	DSN     string `env:"TN_SIGN_PG_DSN,required"`
	Version string `env:"TN_SIGN_MIGRATION_VERSION" envDefault:"last"`
	Table   string `env:"TN_SIGN_MIGRATION_VERSION_TABLE" envDefault:"public.schema_version"`
	Auto    bool   `env:"TN_SIGN_MIGRATION_AUTO" envDefault:"false"`
}

func NewMigration() *Migration {
	return &Migration{}
}

func (m Migration) Run(ctx context.Context) error {
	var cfg Config

	if err := env.Parse(&cfg); err != nil {
		return err
	}

	migrate, err := tern.New(
		MgrsFS,
		tern.WithMigrationsPath("sql"),
		tern.WithInternalConn(cfg.DSN),
		tern.WithDestinationVersion(cfg.Version),
		tern.WithVersionTable(cfg.Table),
	)
	if err != nil {
		return err
	}

	return migrate.Run(ctx)
}
