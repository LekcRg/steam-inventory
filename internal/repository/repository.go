package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/LekcRg/steam-inventory/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type Repo struct {
	db     *sqlx.DB
	config *config.Postgres
	pool   *pgxpool.Pool
}

const PostgresDriver = "pgx"

func New(
	ctx context.Context, cfg *config.Postgres, log *zap.Logger,
) (*Repo, error) {
	repo := &Repo{
		config: cfg,
	}

	nativeDB, err := repo.getPgPool(ctx)
	if err != nil {
		return nil, err
	}

	db := sqlx.NewDb(nativeDB, PostgresDriver)

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	log.Info("Success connect to postgres db")

	repo.db = db

	return repo, nil
}

func (r *Repo) getURIFromConfig() string {
	dbURI := r.config.URI
	if r.config.URI == "" {
		dbURI = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?pool_max_conns=%s",
			r.config.User, r.config.Password, r.config.Host,
			r.config.Port, r.config.DB, r.config.MaxConns,
		)
	}

	return dbURI
}

func (r *Repo) getPgPool(ctx context.Context) (*sql.DB, error) {
	connConfig, err := pgxpool.ParseConfig(r.getURIFromConfig())
	if err != nil {
		return nil, err
	}

	connPool, err := pgxpool.NewWithConfig(ctx, connConfig)
	if err != nil {
		return nil, err
	}
	r.pool = connPool

	return stdlib.OpenDBFromPool(connPool), nil
}

func (r *Repo) Close() error {
	r.pool.Close()

	return r.db.Close()
}
