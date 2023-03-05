package service

import "eSearcher/internal/models"

type VacancyService interface {
	CreateVacancy(vacancy *models.Vacancy) error
	SearchVacancy(params *models.SearchVacancyParams) ([]*models.Vacancy, error)
}

type ApplicantsService interface {
	Create(applicant *models.Applicant) error
}

type SpecializationsService interface {
	Get() ([]*models.Specialization, error)
}

type Services struct {
	VacancyService
	ApplicantsService
	SpecializationsService
}
