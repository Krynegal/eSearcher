package postgres

import (
	"context"
	"database/sql"
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

func (a *ApplicantsDB) Create(applicant *models.Applicant) error {
	ctx := context.Background()
	conn, err := a.pool.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()
	if _, err = conn.Exec(ctx,
		`INSERT INTO applicant_info (user_id, name, lastname, status_id, phone, birthday, description, male)
			VALUES($1, $2, $3, $4, $5, $6, $7, $8)`,
		applicant.ID,
		applicant.Info.Name,
		applicant.Info.Lastname,
		applicant.Info.Status,
		applicant.Info.Phone,
		applicant.Info.Birthday,
		applicant.Info.Description,
		applicant.Info.Male,
	); err != nil {
		return err
	}
	for _, experience := range applicant.Experiences {
		if _, err = conn.Exec(ctx,
			`INSERT INTO experience (user_id, start, finish, organization, position, duties, skills)
			VALUES($1, $2, $3, $4, $5, $6, $7)`,
			applicant.ID,
			experience.Start,
			experience.Finish,
			experience.Organization,
			experience.Position,
			experience.Duties,
			experience.Skills,
		); err != nil {
			return err
		}
	}
	for _, education := range applicant.Educations {
		if _, err = conn.Exec(ctx,
			`INSERT INTO education (user_id, institution_id, grade_id, faculty, specialization, finish)
			VALUES($1, $2, $3, $4, $5, $6)`,
			applicant.ID,
			education.Institution,
			education.Grade,
			education.Faculty,
			education.Specialization,
			education.Finish,
		); err != nil {
			return err
		}
	}
	for _, language := range applicant.Languages {
		if _, err = conn.Exec(ctx,
			`INSERT INTO applicant_id_language_id (user_id, language_id, language_level)
			VALUES($1, $2, $3)`,
			applicant.ID,
			language.Language,
			language.Level,
		); err != nil {
			return err
		}
	}
	for _, specialization := range applicant.Specializations {
		if _, err = conn.Exec(ctx,
			`INSERT INTO applicant_id_specialization_id (user_id, specialization_id, salary)
			VALUES($1, $2, $3)`,
			applicant.ID,
			specialization.Specialization,
			specialization.Salary,
		); err != nil {
			return err
		}
	}
	for _, addedSchedule := range applicant.Schedule.Added {
		if _, err = conn.Exec(ctx,
			`INSERT INTO applicant_id_schedule_id (user_id, schedule_id)
			VALUES($1, $2)`,
			applicant.ID,
			addedSchedule,
		); err != nil {
			return err
		}
	}
	for _, addedBusyness := range applicant.Busyness.Added {
		if _, err = conn.Exec(ctx,
			`INSERT INTO applicant_id_busyness_id (user_id, busyness_id)
			VALUES($1, $2)`,
			applicant.ID,
			addedBusyness,
		); err != nil {
			return err
		}
	}
	return nil
}

func (a *ApplicantsDB) Update(applicant *models.Applicant) error {
	ctx := context.Background()
	conn, err := a.pool.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()
	if _, err = conn.Exec(ctx,
		`UPDATE applicant_info SET
			name=$2,lastname=$3, status_id=$4, phone=$5, birthday=$6, description=$7, male=$8 WHERE user_id=$1`,
		applicant.ID,
		applicant.Info.Name,
		applicant.Info.Lastname,
		applicant.Info.Status,
		applicant.Info.Phone,
		applicant.Info.Birthday,
		applicant.Info.Description,
		applicant.Info.Male); err != nil {
		return err
	}
	for _, experience := range applicant.Experiences {
		if experience.ID != 0 {
			if _, err = conn.Exec(ctx,
				`UPDATE experience SET
					organization=$3, start=$4, finish=$5, position=$6, duties=$7, skills=$8 WHERE user_id=$1 AND id=$2`,
				applicant.ID,
				experience.ID,
				experience.Organization,
				experience.Start,
				experience.Finish,
				experience.Position,
				experience.Duties,
				experience.Skills,
			); err != nil {
				return err
			}
		} else {
			if _, err = conn.Exec(ctx,
				`INSERT INTO experience (user_id, start, finish, organization, position, duties, skills)
					VALUES($1, $2, $3, $4, $5, $6, $7)`,
				applicant.ID,
				experience.Start,
				experience.Finish,
				experience.Organization,
				experience.Position,
				experience.Duties,
				experience.Skills,
			); err != nil {
				return err
			}
		}
	}
	for _, education := range applicant.Educations {
		if education.ID != 0 {
			if _, err = conn.Exec(ctx,
				`UPDATE education SET
					institution_id=$3, grade_id=$4, faculty=$5, specialization=$6, finish=$7 WHERE user_id=$1 AND id=$2`,
				applicant.ID,
				education.ID,
				education.Institution,
				education.Grade,
				education.Faculty,
				education.Specialization,
				education.Finish,
			); err != nil {
				return err
			}
		} else {
			if _, err = conn.Exec(ctx,
				`INSERT INTO education (user_id, institution_id, grade_id, faculty, specialization, finish)
					VALUES($1, $2, $3, $4, $5, $6)`,
				applicant.ID,
				education.Institution,
				education.Grade,
				education.Faculty,
				education.Specialization,
				education.Finish,
			); err != nil {
				return err
			}
		}
	}
	for _, language := range applicant.Languages {
		if language.ID != 0 {
			if _, err = conn.Exec(ctx,
				`UPDATE applicant_id_language_id SET
					language_id=$3, language_level=$4 WHERE user_id=$1 AND id=$2`,
				applicant.ID,
				language.ID,
				language.Language,
				language.Level,
			); err != nil {
				return err
			}
		} else {
			if _, err = conn.Exec(ctx,
				`INSERT INTO applicant_id_language_id (user_id, language_id, language_level)
					VALUES($1, $2, $3)`,
				applicant.ID,
				language.Language,
				language.Level,
			); err != nil {
				return err
			}
		}
	}
	for _, specialization := range applicant.Specializations {
		if specialization.ID != 0 {
			if _, err = conn.Exec(ctx,
				`UPDATE applicant_id_specialization_id SET
					specialization_id=$3, salary=$4 WHERE user_id=$1 AND id=$2`,
				applicant.ID,
				specialization.ID,
				specialization.Specialization,
				specialization.Salary,
			); err != nil {
				return err
			}
		} else {
			if _, err = conn.Exec(ctx,
				`INSERT INTO applicant_id_specialization_id (user_id, specialization_id, salary)
					VALUES($1, $2, $3)`,
				applicant.ID,
				specialization.Specialization,
				specialization.Salary,
			); err != nil {
				return err
			}
		}
	}
	for _, addedSchedule := range applicant.Schedule.Added {
		if _, err = conn.Exec(ctx,
			`INSERT INTO applicant_id_schedule_id (user_id, schedule_id)
			VALUES($1, $2)`,
			applicant.ID,
			addedSchedule,
		); err != nil {
			return err
		}
	}
	for _, deletedSchedule := range applicant.Schedule.Deleted {
		if _, err = conn.Exec(ctx,
			`DELETE FROM applicant_id_schedule_id WHERE user_id=$1 AND schedule_id=$2`,
			applicant.ID,
			deletedSchedule,
		); err != nil {
			return err
		}
	}
	for _, addedBusyness := range applicant.Busyness.Added {
		if _, err = conn.Exec(ctx,
			`INSERT INTO applicant_id_busyness_id (user_id, busyness_id)
			VALUES($1, $2) ON CONFLICT DO NOTHING`,
			applicant.ID,
			addedBusyness,
		); err != nil {
			return err
		}
	}
	for _, deletedBusyness := range applicant.Busyness.Deleted {
		if _, err = conn.Exec(ctx,
			`DELETE FROM applicant_id_busyness_id WHERE user_id=$1 AND busyness_id=$2`,
			applicant.ID,
			deletedBusyness,
		); err != nil {
			return err
		}
	}
	return nil
}

func (a *ApplicantsDB) Get(uid int) (*models.Applicant, error) {
	ctx := context.Background()
	conn, err := a.pool.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	var applicantInfo models.ApplicantInfo
	var birthday sql.NullTime
	var status sql.NullInt64
	var name, lastname, phone, description sql.NullString
	if err = conn.QueryRow(ctx,
		`select name, lastname, status_id, phone, birthday, description, male 
			from applicant_info where user_id = $1`, uid).Scan(
		&name,
		&lastname,
		&status,
		&phone,
		&birthday,
		&description,
		&applicantInfo.Male,
	); err != nil {
		return nil, err
	}
	applicantInfo.Name = name.String
	applicantInfo.Lastname = lastname.String
	applicantInfo.Status = int(status.Int64)
	applicantInfo.Phone = phone.String
	applicantInfo.Description = description.String
	applicantInfo.Birthday = birthday.Time.Format("2006-06-02")

	fmt.Printf("1: %+v\n", applicantInfo)
	applicantExperiences := make([]models.ApplicantExperience, 0)
	var start, expFinish sql.NullTime
	var expID sql.NullInt64
	var organization, position, duties, skills sql.NullString
	aExp, err := conn.Query(ctx, `
		select id, organization, start, finish, position, duties, skills from experience where user_id = $1`, uid)
	for aExp.Next() {
		var experience models.ApplicantExperience
		if err = aExp.Scan(
			&expID,
			&organization,
			&start,
			&expFinish,
			&position,
			&duties,
			&skills); err != nil {
			return nil, err
		}
		if expID.Valid {
			experience.ID = int(expID.Int64)
		} else {
			experience.ID = -1
		}
		experience.Start = start.Time.Format("2006-06-02")
		experience.Finish = expFinish.Time.Format("2006-06-02")
		experience.Organization = organization.String
		experience.Position = position.String
		experience.Duties = duties.String
		experience.Skills = skills.String
		applicantExperiences = append(applicantExperiences, experience)
	}
	fmt.Println("2")
	applicantEducations := make([]models.ApplicantEducation, 0)
	var eduFinish sql.NullTime
	var eduID, institution, grade, specialization sql.NullInt64
	var faculty sql.NullString
	aEdu, err := conn.Query(ctx, `
		select id, institution_id, grade_id, faculty, specialization, finish from education where user_id = $1`, uid)
	for aEdu.Next() {
		var education models.ApplicantEducation
		if err = aEdu.Scan(
			&eduID,
			&institution,
			&grade,
			&faculty,
			&specialization,
			&eduFinish); err != nil {
			return nil, err
		}
		education.ID = int(eduID.Int64)
		education.Institution = int(institution.Int64)
		education.Grade = int(grade.Int64)
		education.Specialization = int(specialization.Int64)
		education.Faculty = faculty.String
		education.Finish = eduFinish.Time.Format("2006-06-02")
		applicantEducations = append(applicantEducations, education)
	}
	fmt.Printf("3: %v\n", applicantEducations)
	applicantLanguages := make([]models.ApplicantLanguage, 0)
	aLang, err := conn.Query(ctx, `
		select id, language_id, language_level from applicant_id_language_id where user_id = $1`, uid)
	for aLang.Next() {
		var language models.ApplicantLanguage
		if err = aLang.Scan(&language.ID, &language.Language, &language.Level); err != nil {
			return nil, err
		}
		applicantLanguages = append(applicantLanguages, language)
	}
	fmt.Println("4")
	applicantSpecializations := make([]models.ApplicantSpecialization, 0)
	var specID, specializationID, salary sql.NullInt64
	aSpec, err := conn.Query(ctx, `
		select id, specialization_id, salary from applicant_id_specialization_id where user_id = $1`, uid)
	for aSpec.Next() {
		var spec models.ApplicantSpecialization
		if err = aSpec.Scan(
			&specID,
			&specializationID,
			&salary); err != nil {
			return nil, err
		}
		spec.ID = int(specID.Int64)
		spec.Specialization = int(specializationID.Int64)
		spec.Salary = int(salary.Int64)
		applicantSpecializations = append(applicantSpecializations, spec)
	}
	fmt.Println("5")
	var applicantBusyness models.ApplicantBusyness
	aBus, err := conn.Query(ctx, `
		select busyness_id from applicant_id_busyness_id where user_id = $1`, uid)
	applicantBusyness.Busyness = make([]int, 0)
	for aBus.Next() {
		fmt.Println("cycle")
		var busyness int
		if err = aBus.Scan(&busyness); err != nil {
			return nil, err
		}
		applicantBusyness.Busyness = append(applicantBusyness.Busyness, busyness)
	}

	var applicantSchedule models.ApplicantSchedule
	aSch, err := conn.Query(ctx, `
		select schedule_id from applicant_id_schedule_id where user_id = $1`, uid)
	applicantSchedule.Schedule = make([]int, 0)
	for aSch.Next() {
		var schedule int
		if err = aSch.Scan(&schedule); err != nil {
			return nil, err
		}
		applicantSchedule.Schedule = append(applicantSchedule.Schedule, schedule)
	}

	applicant := models.Applicant{
		ID:              uid,
		Info:            applicantInfo,
		Experiences:     applicantExperiences,
		Educations:      applicantEducations,
		Languages:       applicantLanguages,
		Specializations: applicantSpecializations,
		Busyness:        applicantBusyness,
		Schedule:        applicantSchedule,
	}
	fmt.Printf("%+v\n", applicant)
	return &applicant, err
}

func (a *ApplicantsDB) Search(params *models.SearchApplicantParams) ([]int, error) {
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
		`select user_id from applicant_info WHERE
		id IN (select user_id from applicant_id_schedule_id WHERE schedule_id = ANY ($1)) 
		AND
		id IN (select user_id from applicant_id_busyness_id WHERE busyness_id = ANY ($2))
		AND
		id IN (select user_id from applicant_id_specialization_id WHERE specialization_id = ANY ($3))
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
	return applicantIDs, nil
}
