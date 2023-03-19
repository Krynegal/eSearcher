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
		`INSERT INTO applicant_info (id, name, lastname) VALUES($1, $2) RETURNING id`,
		applicant.Info.Name, applicant.Info.Lastname).Scan(&id); err != nil {
		return -1, err
	}
	return id, nil
}

func (a *ApplicantsDB) Get(id string) (*models.Applicant, error) {
	ctx := context.Background()
	conn, err := a.pool.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	var applicantInfo models.ApplicantInfo
	if err = conn.QueryRow(ctx,
		`select name, lastname, status_id, phone, birthday, description, male 
			from applicant_info where user_id = $1`, id).Scan(
		&applicantInfo.Name,
		&applicantInfo.Lastname,
		&applicantInfo.Status,
		&applicantInfo.Phone,
		&applicantInfo.Birthday,
		&applicantInfo.Description,
		&applicantInfo.Male,
	); err != nil {
		return nil, err
	}

	fmt.Printf("1: %+v", applicantInfo)
	var applicantExperiences []models.ApplicantExperience
	aExp, err := conn.Query(ctx, `
		select organization, start, finish, position, duties, skills from experience where user_id = $1`, id)
	for aExp.Next() {
		var experience models.ApplicantExperience
		if err = aExp.Scan(&experience.Organization,
			&experience.Start,
			&experience.Finish,
			&experience.Position,
			&experience.Duties,
			&experience.Skills); err != nil {
			return nil, err
		}
		applicantExperiences = append(applicantExperiences, experience)
	}
	fmt.Println("2")
	var applicantEducations []models.ApplicantEducation
	aEdu, err := conn.Query(ctx, `
		select institution_id, grade_id, faculty, specialization, finish from education where user_id = $1`, id)
	for aEdu.Next() {
		var education models.ApplicantEducation
		if err = aEdu.Scan(&education.Institution,
			&education.Grade,
			&education.Faculty,
			&education.Specialization,
			&education.Finish); err != nil {
			return nil, err
		}
		applicantEducations = append(applicantEducations, education)
	}
	fmt.Printf("3: %v", applicantEducations)
	var applicantLanguages []models.ApplicantLanguage
	aLang, err := conn.Query(ctx, `
		select language_id, language_level from applicant_id_language_id where user_id = $1`, id)
	for aLang.Next() {
		var language models.ApplicantLanguage
		if err = aLang.Scan(&language.Language, &language.Level); err != nil {
			return nil, err
		}
		applicantLanguages = append(applicantLanguages, language)
	}
	fmt.Println("4")
	var applicantSpecializations []models.ApplicantSpecialization
	aSpec, err := conn.Query(ctx, `
		select specialization_id, salary from applicant_id_specialization_id where user_id = $1`, id)
	for aSpec.Next() {
		var specialization models.ApplicantSpecialization
		if err = aSpec.Scan(&specialization.Specialization, &specialization.Salary); err != nil {
			return nil, err
		}
		applicantSpecializations = append(applicantSpecializations, specialization)
	}
	fmt.Println("5")
	var applicantBusyness []int
	aBus, err := conn.Query(ctx, `
		select busyness_id from applicant_id_busyness_id where user_id = $1`, id)
	for aBus.Next() {
		var busyness int
		if err = aBus.Scan(&busyness); err != nil {
			return nil, err
		}
		applicantBusyness = append(applicantBusyness, busyness)
	}

	var applicantSchedule []int
	aSch, err := conn.Query(ctx, `
		select schedule_id from applicant_id_schedule_id where user_id = $1`, id)
	for aSch.Next() {
		var schedule int
		if err = aSch.Scan(&schedule); err != nil {
			return nil, err
		}
		applicantSchedule = append(applicantSchedule, schedule)
	}
	applicant := models.Applicant{
		ID:              "",
		Info:            applicantInfo,
		Experiences:     applicantExperiences,
		Educations:      applicantEducations,
		Languages:       applicantLanguages,
		Specializations: applicantSpecializations,
		Busyness:        applicantBusyness,
		Schedule:        applicantSchedule,
	}
	return &applicant, err
}

func (a *ApplicantsDB) Search(params *models.SearchApplicantParams) ([]string, error) {
	ctx := context.Background()
	conn, err := a.pool.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()

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
	var applicantIDs []string
	for rows.Next() {
		var id string
		if err = rows.Scan(&id); err != nil {
			return nil, err
		}
		applicantIDs = append(applicantIDs, id)
	}
	fmt.Printf("applicantIDs: %v", applicantIDs)
	return applicantIDs, nil
}
