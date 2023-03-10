package service

import "eSearcher/internal/models"

type VacancyService interface {
	CreateVacancy(vacancy *models.Vacancy) error
	SearchVacancy(params *models.SearchVacancyParams) ([]*models.Vacancy, error)
}

type ApplicantsService interface {
	Create(applicant *models.Applicant) error
	Get(id string) (*models.Applicant, error)
	SearchApplicant(params *models.SearchApplicantParams) ([]*models.Applicant, error)
}

type SpecializationsService interface {
	GetAllSpecializations() ([]*models.Specialization, error)
}

type EmployersService interface {
	Create(applicant *models.Employer) error
}

type Services struct {
	VacancyService
	ApplicantsService
	SpecializationsService
	EmployersService
}
