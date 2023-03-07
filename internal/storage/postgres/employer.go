package postgres

import (
	"context"
	"eSearcher/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type EmployersDB struct {
	pool *pgxpool.Pool
}

func NewEmployersStore(pool *pgxpool.Pool) *EmployersDB {
	return &EmployersDB{pool: pool}
}

func (e *EmployersDB) Create(applicant *models.Employer) (string, error) {
	ctx := context.Background()
	conn, err := e.pool.Acquire(ctx)
	if err != nil {
		return "-1", err
	}
	defer conn.Release()
	if _, err = conn.Exec(ctx,
		`INSERT INTO employers (name) VALUES($1)`,
		applicant.Name); err != nil {
		return "-1", err
	}
	return "", nil
}
