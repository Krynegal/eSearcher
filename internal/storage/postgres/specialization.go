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

func (s *SpecializationsDB) GetAllSpecializations() ([]*models.Specialization, error) {
	ctx := context.Background()
	conn, err := s.pool.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()
	rows, err := conn.Query(ctx, `SELECT * FROM specializations`)
	if err != nil {
		return nil, err
	}
	var specializations []*models.Specialization
	for rows.Next() {
		var specialization models.Specialization
		if err = rows.Scan(&specialization.ID, &specialization.Name); err != nil {
			return nil, err
		}
		specializations = append(specializations, &specialization)
	}
	return specializations, nil
}
