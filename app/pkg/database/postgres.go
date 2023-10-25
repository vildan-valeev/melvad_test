package database

import (
	"context"
	"github.com/vildan-valeev/melvad_test/pkg/database/zerologadapter"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

// db represents the database connection.
type db struct {
	pool *pgxpool.Pool

	// level of log.
	level string

	// dsn datasource name for connection.
	dsn string
}

// New returns a new instance of db associated with the given datasource name.
func New(dsn, level string) *db {
	return &db{pool: nil, level: level, dsn: dsn}
}

// Open opens the database connection.
func (db *db) Open(ctx context.Context) (err error) {
	l := zerologadapter.NewContextLogger(zerologadapter.WithoutPGXModule())

	// NewPool to the database.
	db.pool, err = NewPool(ctx, db.dsn, WithLogger(l), WithLogLevel(db.level))

	return err
}

// Close closes the database connection.
func (db *db) Close() error {
	if db.pool != nil {
		db.pool.Close()
	}

	return nil
}

func (db *db) Ping(ctx context.Context) error {
	return db.pool.Ping(ctx)
}

func (db *db) Exec(ctx context.Context, sql string, args ...any) (pgconn.CommandTag, error) {
	return db.pool.Exec(ctx, sql, args...)
}

func (db *db) Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error) {
	return db.pool.Query(ctx, sql, args...)
}

func (db *db) QueryRow(ctx context.Context, sql string, args ...any) pgx.Row {
	return db.pool.QueryRow(ctx, sql, args...)
}

func (db *db) CopyFrom(ctx context.Context, identifier pgx.Identifier, strings []string, source pgx.CopyFromSource) (int64, error) {
	return db.pool.CopyFrom(ctx, identifier, strings, source)
}

func (db *db) SendBatch(ctx context.Context, batch *pgx.Batch) pgx.BatchResults {
	return db.pool.SendBatch(ctx, batch)
}

func (db *db) Begin(ctx context.Context) (pgx.Tx, error) {
	return db.pool.Begin(ctx)
}

func (db *db) BeginOpt(ctx context.Context, options pgx.TxOptions) (pgx.Tx, error) {
	return db.pool.BeginTx(ctx, options)
}
