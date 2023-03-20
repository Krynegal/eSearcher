package service

import (
	"eSearcher/internal/models"
	"eSearcher/internal/storage"
)

type Employers struct {
	store storage.EmployerStorage
}

func NewEmployers(storage storage.EmployerStorage) *Employers {
	return &Employers{store: storage}
}

func (e *Employers) Get(id int) (*models.Employer, error) {
	employer, err := e.store.Get(id)
	if err != nil {
		return nil, err
	}
	return employer, nil
}

func (e *Employers) Create(employer *models.Employer) error {
	if err := e.store.Create(employer); err != nil {
		return err
	}
	return nil
}

func (e *Employers) Update(employer *models.Employer) error {
	if err := e.store.Update(employer); err != nil {
		return err
	}
	return nil
}
