package postgres

import (
	"context"
	"eSearcher/internal/models"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
)

type OptionsDB struct {
	pool *pgxpool.Pool
}

func NewOptionsStore(pool *pgxpool.Pool) *OptionsDB {
	return &OptionsDB{pool: pool}
}

func (s *OptionsDB) GetAll(option string) ([]*models.Option, error) {
	ctx := context.Background()
	conn, err := s.pool.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()
	query := fmt.Sprintf(`SELECT * FROM %s`, option)
	rows, err := conn.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	var opts []*models.Option
	for rows.Next() {
		var opt models.Option
		if err = rows.Scan(&opt.ID, &opt.Name); err != nil {
			return nil, err
		}
		opts = append(opts, &opt)
	}
	return opts, nil
}
