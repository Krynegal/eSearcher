package postgres

import (
	"context"
	"eSearcher/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ResponsesDB struct {
	pool *pgxpool.Pool
}

func NewResponsesStore(pool *pgxpool.Pool) *ResponsesDB {
	return &ResponsesDB{pool: pool}
}

func (r *ResponsesDB) Add(response *models.Response) error {
	ctx := context.Background()
	conn, err := r.pool.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()
	if _, err = conn.Exec(ctx,
		`INSERT INTO responses (applicant_id, vacancy_id) VALUES($1, $2)`,
		response.ApplicantID, response.VacancyID); err != nil {
		return err
	}
	return nil
}

func (r *ResponsesDB) Delete(response *models.Response) error {
	ctx := context.Background()
	conn, err := r.pool.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()
	if _, err = conn.Exec(ctx,
		`DELETE FROM responses WHERE applicant_id=$1 AND vacancy_id=$2`,
		response.ApplicantID, response.VacancyID); err != nil {
		return err
	}
	return nil
}
