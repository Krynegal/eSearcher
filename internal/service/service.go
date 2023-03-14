package service

import (
	"eSearcher/internal/models"
)

type AuthService interface {
	CreateUser(login, password string, role int) (int, error)
	AuthUser(login, password string) (*models.User, error)
	GenerateToken(uid int, roles int) (string, error)
}

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

type ResponsesService interface {
	Add(response *models.Response) error
	Delete(response *models.Response) error
}

type Services struct {
	AuthService
	VacancyService
	ApplicantsService
	SpecializationsService
	EmployersService
	ResponsesService
}
