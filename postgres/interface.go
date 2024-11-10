package postgres

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

const TxKey = "PostgresTxKey"

type QueryController interface {
	Exec(ctx context.Context, sql string, args ...any) (commandTag pgconn.CommandTag, err error)
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
}

type TxController interface {
	Begin(ctx context.Context) (TxController, error)
	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error
}
