package database

import (
	"context"
	"fmt"
	"github.com/vildan-valeev/melvad_test/pkg/database/uuid"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/tracelog"
)

type Options struct {
	log   tracelog.Logger
	level tracelog.LogLevel
}

// NewPool returns a connection pool with server.
func NewPool(ctx context.Context, dataSourceName string, opts ...Option) (*pgxpool.Pool, error) {
	o := &Options{log: nil, level: tracelog.LogLevelError}

	// Custom options.
	for _, opt := range opts {
		opt(o)
	}

	poolConfig, err := pgxpool.ParseConfig(dataSourceName)
	if err != nil {
		return nil, fmt.Errorf("parse DSN '%s': %w", dataSourceName, err)
	}

	poolConfig.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		uuid.Register(conn.TypeMap())
		return nil
	}

	poolConfig.ConnConfig.Tracer = &tracelog.TraceLog{
		Logger:   o.log,
		LogLevel: tracelog.LogLevelTrace,
	}

	return pgxpool.NewWithConfig(ctx, poolConfig)
}

// NewConn establishes a connection with server.
func NewConn(ctx context.Context, dataSourceName string, opts ...Option) (*pgx.Conn, error) {
	o := &Options{log: nil, level: tracelog.LogLevelError}

	// Custom options.
	for _, opt := range opts {
		opt(o)
	}

	connConfig, err := pgx.ParseConfig(dataSourceName)
	if err != nil {
		return nil, fmt.Errorf("parse DSN '%s': %w", dataSourceName, err)
	}

	connConfig.Tracer = &tracelog.TraceLog{
		Logger:   o.log,
		LogLevel: tracelog.LogLevelTrace,
	}

	return pgx.ConnectConfig(ctx, connConfig)
}
