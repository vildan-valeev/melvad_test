package infra

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type (
	Tx             = pgx.Tx
	TxOptions      = pgx.TxOptions
	CommandTag     = pgconn.CommandTag
	Row            = pgx.Row
	Rows           = pgx.Rows
	Identifier     = pgx.Identifier
	CopyFromSource = pgx.CopyFromSource
	Batch          = pgx.Batch
	BatchResults   = pgx.BatchResults
)

type Database interface {
	Ping(context.Context) error

	Beginner
	TableOperator
}

// DB is the minimal representation of a relational database that powers repositories.
type DB interface {
	Beginner
	TableOperator
}

// TableOperator represents the methods required to interact with database
// tables. Typically this is satisfied by both databases and transactions.
type TableOperator interface {
	Execer
	Queryer
	Batcher
}

// Beginner represents an object that can begin a transaction.
type Beginner interface {
	Begin(context.Context) (Tx, error)
	BeginOpt(context.Context, TxOptions) (Tx, error)
}

// Execer executes a query that returns no result.
type Execer interface {
	Exec(context.Context, string, ...any) (CommandTag, error)
}

// Queryer executes a query that populates dest with the returned rows.
type Queryer interface {
	Query(context.Context, string, ...any) (Rows, error)
	QueryRow(context.Context, string, ...any) Row
}

type Batcher interface {
	CopyFrom(context.Context, Identifier, []string, CopyFromSource) (int64, error)
	SendBatch(context.Context, *Batch) BatchResults
}

// NotFound is a helper function to check if an error is `pgx.ErrNoRows`.
func NotFound(err error) bool {
	return errors.Is(err, pgx.ErrNoRows)
}
