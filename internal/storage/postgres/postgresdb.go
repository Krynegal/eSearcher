package postgres

import (
	"context"
	"eSearcher/configs"
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPostgresDB(cfg *configs.Config) (*pgxpool.Pool, error) {
	postgresURL := fmt.Sprintf("postgres://%s:%s@%s:%s/postgres?sslmode=disable",
		cfg.PostgresUser, cfg.PostgresPass, cfg.PostgresHost, cfg.PostgresPort)
	if err := migrateDB(postgresURL); err != nil {
		return nil, err
	}
	pool, err := pgxpool.New(context.Background(), postgresURL)
	if err != nil {
		return nil, err
	}
	if err = pool.Ping(context.Background()); err != nil {
		return nil, err
	}
	return pool, nil
}

func migrateDB(postgresURL string) error {
	m, err := migrate.New("file://schema", postgresURL)
	if err != nil {
		return err
	}
	if err = m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}
	return nil
}
