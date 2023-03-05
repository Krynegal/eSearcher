package postgres

import (
	"context"
	"eSearcher/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ApplicantsDB struct {
	pool *pgxpool.Pool
}

func NewApplicantsStore(pool *pgxpool.Pool) *ApplicantsDB {
	return &ApplicantsDB{pool: pool}
}

func (a *ApplicantsDB) Create(applicant *models.Applicant) (string, error) {
	ctx := context.Background()
	conn, err := a.pool.Acquire(ctx)
	if err != nil {
		return "-1", err
	}
	defer conn.Release()
	if _, err = conn.Exec(ctx,
		`INSERT INTO applicants (name, 'lastname') VALUES($1, $2)`,
		applicant.Name, applicant.Lastname); err != nil {
		return "-1", err
	}
	return "", nil
}
