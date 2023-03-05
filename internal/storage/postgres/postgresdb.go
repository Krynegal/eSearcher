package postgres

import (
	"context"
	"eSearcher/configs"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPostgresDB(cfg *configs.Config) (*pgxpool.Pool, error) {
	postgresURL := fmt.Sprintf("postgres://%s:%s@%s:%s/postgres?sslmode=disable",
		cfg.PostgresUser, cfg.PostgresPass, cfg.PostgresHost, cfg.PostgresPort)
	pool, err := pgxpool.New(context.Background(), postgresURL)
	if err != nil {
		return nil, err
	}
	if err = pool.Ping(context.Background()); err != nil {
		return nil, err
	}
	return pool, nil
}
