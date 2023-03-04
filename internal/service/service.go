package service

import "eSearcher/internal/models"

type VacancyService interface {
	CreateVacancy(vacancy *models.Vacancy) error
	SearchVacancy(params *models.SearchVacancyParams) ([]*models.Vacancy, error)
}

type EmployeeService interface {
}

type EmployerService interface {
}

type Services struct {
	VacancyService
	EmployeeService
	EmployerService
}
