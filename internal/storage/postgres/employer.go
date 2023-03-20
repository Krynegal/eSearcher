package postgres

import (
	"context"
	"database/sql"
	"eSearcher/internal/models"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type EmployersDB struct {
	pool *pgxpool.Pool
}

func NewEmployersStore(pool *pgxpool.Pool) *EmployersDB {
	return &EmployersDB{pool: pool}
}

func (e *EmployersDB) Get(id int) (*models.Employer, error) {
	ctx := context.Background()
	conn, err := e.pool.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	var employerInfo models.EmployerInfo
	var name, phone sql.NullString
	if err = conn.QueryRow(ctx,
		`select name, phone from employers where user_id = $1`, id).Scan(
		&name,
		&phone,
	); err != nil {
		return nil, err
	}
	employerInfo.Name = name.String
	employerInfo.Phone = phone.String

	var employerSphere models.EmployerSphere
	eSphere, err := conn.Query(ctx, `
		select sphere_id from employer_id_sphere_id where user_id = $1`, id)
	for eSphere.Next() {
		var sphere int
		if err = eSphere.Scan(&sphere); err != nil {
			if err == pgx.ErrNoRows {
				employerSphere.Sphere = []int{}
				break
			}
			return nil, err
		}
		employerSphere.Sphere = append(employerSphere.Sphere, sphere)
	}

	employer := models.Employer{
		ID:     id,
		Info:   employerInfo,
		Sphere: employerSphere,
	}
	return &employer, nil
}

func (e *EmployersDB) Create(employer *models.Employer) error {
	ctx := context.Background()
	conn, err := e.pool.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()
	if _, err = conn.Exec(ctx,
		`INSERT INTO employers (user_id, name, phone) VALUES($1, $2, $3)`,
		employer.ID, employer.Info.Name, employer.Info.Phone); err != nil {
		return err
	}
	if len(employer.Sphere.Added) != 0 {
		for _, addedSphereID := range employer.Sphere.Added {
			if _, err = conn.Exec(ctx,
				`INSERT INTO employer_id_sphere_id (user_id, sphere_id) VALUES ($1, $2)`,
				employer.ID,
				addedSphereID,
			); err != nil {
				return err
			}
		}
	}
	return nil
}

func (e *EmployersDB) Update(employer *models.Employer) error {
	ctx := context.Background()
	conn, err := e.pool.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()
	if _, err = conn.Exec(ctx,
		`UPDATE employers SET name=$2, phone=$3 WHERE user_id=$1`,
		employer.ID,
		employer.Info.Name,
		employer.Info.Phone,
	); err != nil {
		return err
	}
	if len(employer.Sphere.Added) != 0 {
		for _, addedSphereID := range employer.Sphere.Added {
			if _, err = conn.Exec(ctx,
				`INSERT INTO employer_id_sphere_id (user_id, sphere_id) VALUES ($1, $2)`,
				employer.ID,
				addedSphereID,
			); err != nil {
				return err
			}
		}
	}
	if len(employer.Sphere.Deleted) != 0 {
		for _, deletedSphereID := range employer.Sphere.Deleted {
			if _, err = conn.Exec(ctx,
				`DELETE FROM employer_id_sphere_id WHERE user_id=$1 AND sphere_id=$2`,
				employer.ID,
				deletedSphereID,
			); err != nil {
				return err
			}
		}
	}
	return nil
}
