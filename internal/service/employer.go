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

func (e *Employers) Create(employer *models.Employer) error {
	_, err := e.store.Create(employer)
	if err != nil {
		return err
	}
	return nil
}
