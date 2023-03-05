package postgres

import (
	"context"
	"eSearcher/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type SpecializationsDB struct {
	pool *pgxpool.Pool
}

func NewSpecializationsStore(pool *pgxpool.Pool) *SpecializationsDB {
	return &SpecializationsDB{pool: pool}
}

func (a *SpecializationsDB) Get() ([]*models.Specialization, error) {
	ctx := context.Background()
	conn, err := a.pool.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()
	if _, err = conn.Query(ctx, `SELECT * FROM specializations`); err != nil {
		return nil, err
	}
	return nil, nil
}
