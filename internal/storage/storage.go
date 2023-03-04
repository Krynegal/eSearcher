package storage

import "eSearcher/internal/models"

type VacancyStorage interface {
	Create(vacancy *models.Vacancy) (string, error)
	Search(params *models.SearchVacancyParams) ([]*models.Vacancy, error)
}

type EmployeeStorage interface {
}

type EmployerStorage interface {
}

type Storage struct {
	VacancyStorage
	EmployeeStorage
	EmployerStorage
}
