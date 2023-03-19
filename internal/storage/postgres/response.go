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

func (r *ResponsesDB) GetUsersVacancyIDs(uid int) ([]string, error) {
	ctx := context.Background()
	conn, err := r.pool.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()
	var vacancyIDs []string
	rows, err := conn.Query(ctx, `
		select vacancy_id from responses where user_id = $1`, uid)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var vacancyID string
		if err = rows.Scan(&vacancyID); err != nil {
			return nil, err
		}
		vacancyIDs = append(vacancyIDs, vacancyID)
	}
	return vacancyIDs, nil
}

func (r *ResponsesDB) Add(response *models.Response) error {
	ctx := context.Background()
	conn, err := r.pool.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()
	if _, err = conn.Exec(ctx,
		`INSERT INTO responses (user_id, vacancy_id, status_id) VALUES($1, $2, $3)`,
		response.ApplicantID, response.VacancyID, response.StatusID); err != nil {
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
		`DELETE FROM responses WHERE user_id=$1 AND vacancy_id=$2`,
		response.ApplicantID, response.VacancyID); err != nil {
		return err
	}
	return nil
}
