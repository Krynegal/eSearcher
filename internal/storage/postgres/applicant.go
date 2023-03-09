package postgres

import (
	"context"
	"eSearcher/internal/models"
	"fmt"
	"github.com/jackc/pgx/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ApplicantsDB struct {
	pool *pgxpool.Pool
}

func NewApplicantsStore(pool *pgxpool.Pool) *ApplicantsDB {
	return &ApplicantsDB{pool: pool}
}

func (a *ApplicantsDB) Create(applicant *models.Applicant) (int, error) {
	ctx := context.Background()
	conn, err := a.pool.Acquire(ctx)
	if err != nil {
		return -1, err
	}
	defer conn.Release()
	var id int
	if err = conn.QueryRow(ctx,
		`INSERT INTO applicants (name, lastname) VALUES($1, $2) RETURNING id`,
		applicant.Name, applicant.Lastname).Scan(&id); err != nil {
		return -1, err
	}
	return id, nil
}

func (a *ApplicantsDB) Search(params *models.SearchApplicantParams) ([]*models.Applicant, error) {
	ctx := context.Background()
	conn, err := a.pool.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()
	//rows, err := conn.Query(ctx,
	//	`select id, name, lastname, schedule_id, busyness_id from applicants
	//	JOIN applicant_id_schedule_id on applicants.id=applicant_id_schedule_id.applicant_id
	//	JOIN applicant_id_busyness_id on applicants.id=applicant_id_busyness_id.applicant_id
	//	WHERE applicant_id_schedule_id.schedule_id IN $1 OR applicant_id_busyness_id.busyness_id IN $2`,
	//	params.Schedule, params.Busyness)

	FillEmpty(params)
	fmt.Printf("params: %+v", params)
	schedule := &pgtype.Int4Array{}
	if err = schedule.Set(params.Schedule); err != nil {
		return nil, err
	}
	busyness := &pgtype.Int4Array{}
	if err = busyness.Set(params.Busyness); err != nil {
		return nil, err
	}
	specialization := &pgtype.Int4Array{}
	if err = specialization.Set(params.Specialization); err != nil {
		return nil, err
	}
	rows, err := conn.Query(ctx,
		`select id from applicants WHERE
		id IN (select applicant_id from applicant_id_schedule_id WHERE schedule_id = ANY ($1)) 
		AND
		id IN (select applicant_id from applicant_id_busyness_id WHERE busyness_id = ANY ($2))
		AND
		id IN (select applicant_id from applicant_id_specialization_id WHERE specialization_id = ANY ($3))
		`,
		schedule, busyness, specialization)

	if err != nil {
		return nil, err
	}
	var applicantIDs []int
	for rows.Next() {
		var id int
		if err = rows.Scan(&id); err != nil {
			return nil, err
		}
		applicantIDs = append(applicantIDs, id)
	}
	fmt.Printf("applicantIDs: %v", applicantIDs)
	var applicants []*models.Applicant
	for _, id := range applicantIDs {
		applicant := models.Applicant{}
		if err = conn.QueryRow(ctx,
			`select id, name, lastname from applicants where id = $1`, id).Scan(
			&applicant.ID, applicant.Name, applicant.Lastname); err != nil {
			return nil, err
		}
		aSchedule, err := conn.Query(ctx, `
		select schedule_id from applicant_id_schedule_id where applicant_id = $1`, id)
		for aSchedule.Next() {
			var s int
			if err = rows.Scan(&s); err != nil {
				return nil, err
			}
			applicant.Schedule = append(applicant.Schedule, s)
		}
		aBusyness, err := conn.Query(ctx, `
		select busyness_id from applicant_id_busyness_id where applicant_id = $1`, id)
		for aBusyness.Next() {
			var b int
			if err = rows.Scan(&b); err != nil {
				return nil, err
			}
			applicant.Busyness = append(applicant.Busyness, b)
		}
		aSpecializations, err := conn.Query(ctx, `
		select specialization_id from applicant_id_specialization_id where applicant_id = $1`, id)
		for aSpecializations.Next() {
			var s int
			if err = rows.Scan(&s); err != nil {
				return nil, err
			}
			applicant.Specialization = append(applicant.Specialization, s)
		}
		applicants = append(applicants, &applicant)
	}
	return nil, nil
}
