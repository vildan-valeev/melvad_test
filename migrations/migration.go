package main

import (
	"context"
	"embed"
	"fmt"
	"github.com/caarlos0/env/v9"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/tern/v2/migrate"
	"io/fs"
)

type Config struct {
	DSN     string `env:"DSN,required" envDefault:"postgres://postgres:postgres@localhost:5432/postgres"`
	Version string `env:"MIGRATION_VERSION" envDefault:"last"`
	Table   string `env:"MIGRATION_VERSION_TABLE" envDefault:"public.schema_version"`
	Auto    bool   `env:"MIGRATION_AUTO" envDefault:"false"`
}

//go:embed sql
var migrationFiles embed.FS

type Migrator struct {
	migrator *migrate.Migrator
}

func NewMigration(ctx context.Context) (Migrator, error) {
	var cfg Config

	if err := env.Parse(&cfg); err != nil {
		return Migrator{}, err
	}
	fmt.Println("MIGRATOR DSN", cfg.DSN)
	conn, err := pgx.Connect(ctx, cfg.DSN)
	if err != nil {
		return Migrator{}, err
	}

	migrator, err := migrate.NewMigratorEx(
		ctx, conn, cfg.Table,
		&migrate.MigratorOptions{
			DisableTx: false,
		})
	if err != nil {
		return Migrator{}, err
	}

	migrationRoot, _ := fs.Sub(migrationFiles, "sql")

	err = migrator.LoadMigrations(migrationRoot)
	if err != nil {
		return Migrator{}, err
	}

	return Migrator{
		migrator: migrator,
	}, nil
}

func (m Migrator) Migrate(ctx context.Context) error {
	err := m.migrator.Migrate(ctx)
	return err
}

func (m Migrator) MigrateTo(ctx context.Context, ver int32) error {
	err := m.migrator.MigrateTo(ctx, ver)
	return err
}

func (m Migrator) Info(ctx context.Context) (int32, int32, string, error) {

	version, err := m.migrator.GetCurrentVersion(ctx)
	if err != nil {
		return 0, 0, "", err
	}
	info := ""

	var last int32
	for _, thisMigration := range m.migrator.Migrations {
		last = thisMigration.Sequence

		cur := version == thisMigration.Sequence
		indicator := "  "
		if cur {
			indicator = "->"
		}
		info = info + fmt.Sprintf(
			"%2s %3d %s\n",
			indicator,
			thisMigration.Sequence, thisMigration.Name)
	}

	return version, last, info, nil
}
