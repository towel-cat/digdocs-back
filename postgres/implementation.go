package postgres

import (
	"context"
	"digdocs-back/cfg"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"os"
	"sync"
)

type PgController struct {
	queryController QueryController
	txController    TxController
}

func NewController(controller QueryController) *PgController {
	return &PgController{queryController: controller}
}

func (controller *PgController) Begin(ctx context.Context) (TxController, error) {
	return controller.txController.Begin(ctx)
}

func (controller *PgController) Commit(ctx context.Context) error {
	return controller.txController.Commit(ctx)
}

func (controller *PgController) Rollback(ctx context.Context) error {
	return controller.txController.Rollback(ctx)
}

func (controller *PgController) Exec(ctx context.Context, sql string, args ...any) (commandTag pgconn.CommandTag, err error) {
	return controller.queryController.Exec(ctx, sql, args...)
}

func (controller *PgController) Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error) {
	return controller.queryController.Query(ctx, sql, args...)
}

func (controller *PgController) QueryRow(ctx context.Context, sql string, args ...any) pgx.Row {
	return controller.queryController.QueryRow(ctx, sql, args...)
}

type PgPool struct {
	pool sync.Map
}

func (p *PgPool) GetPool(name string) *pgxpool.Pool {
	if pool, ok := p.pool.Load(name); ok {
		return pool.(*pgxpool.Pool)
	}

	config, err := pgxpool.ParseConfig(
		fmt.Sprintf("user=%s password=%s host=%s port=%s database=%s sslmode=disable",
			cfg.Env.PgUser, cfg.Env.PgPass, cfg.Env.PgHost, cfg.Env.PgPort, cfg.Env.PgDbName))
	if err != nil {
		fmt.Printf("parse config error: %v", err)
		os.Exit(1)
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		fmt.Printf("create pool error: %v", err)
		os.Exit(1)
	}

	return pool
}
