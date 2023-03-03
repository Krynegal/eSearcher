package service

import (
	"eSearcher/internal/storage"
	"fmt"
)

type Vacancies struct {
	store storage.VacancyStorage
}

func NewVacancies(storage storage.VacancyStorage) *Vacancies {
	return &Vacancies{store: storage}
}

func (v *Vacancies) CreateVacancy(name, desc string) error {
	id, err := v.store.Create(name, desc)
	if err != nil {
		return err
	}
	fmt.Printf("id: %v", id)
	return nil
}
