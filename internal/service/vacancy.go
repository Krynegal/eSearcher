package service

import (
	"eSearcher/internal/models"
	"eSearcher/internal/storage"
	"fmt"
)

type Vacancies struct {
	store storage.VacancyStorage
}

func NewVacancies(storage storage.VacancyStorage) *Vacancies {
	return &Vacancies{store: storage}
}

func (v *Vacancies) CreateVacancy(vacancy *models.Vacancy) error {
	id, err := v.store.Create(vacancy)
	if err != nil {
		return err
	}
	fmt.Printf("id: %v", id)
	return nil
}

func (v *Vacancies) SearchVacancy(params *models.SearchVacancyParams) ([]*models.Vacancy, error) {
	vacancies, err := v.store.Search(params)
	if err != nil {
		return nil, err
	}
	return vacancies, nil
}
