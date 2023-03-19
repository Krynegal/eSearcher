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

func (v *Vacancies) GetEmployerVacancies(uid int) ([]*models.Vacancy, error) {
	vacancies, err := v.store.GetEmployerVacancies(uid)
	if err != nil {
		return nil, err
	}
	return vacancies, nil
}

func (v *Vacancies) GetByIDs(ids []string) ([]*models.Vacancy, error) {
	var vacancies []*models.Vacancy
	for _, id := range ids {
		vacancy, err := v.store.GetByID(id)
		if err != nil {
			return nil, err
		}
		vacancies = append(vacancies, vacancy)
	}
	return vacancies, nil
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
