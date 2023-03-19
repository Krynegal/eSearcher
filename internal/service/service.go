package service

import (
	"eSearcher/internal/models"
)

type AuthService interface {
	CreateUser(login, password string, role int) (int, error)
	AuthUser(login, password string) (*models.User, error)
	GenerateToken(uid int, roles int) (string, error)
}

type OptionsService interface {
	GetAll(options []string) (map[string][]*models.Option, error)
}

type VacancyService interface {
	GetEmployerVacancies(uid int) ([]*models.Vacancy, error)
	GetByIDs(id []string) ([]*models.Vacancy, error)
	CreateVacancy(vacancy *models.Vacancy) error
	SearchVacancy(params *models.SearchVacancyParams) ([]*models.Vacancy, error)
}

type ApplicantsService interface {
	Create(applicant *models.Applicant) error
	Get(id int) (*models.Applicant, error)
	SearchApplicant(params *models.SearchApplicantParams) ([]*models.Applicant, error)
}

type EmployersService interface {
	Create(applicant *models.Employer) error
}

type ResponsesService interface {
	ChangeStatus(response *models.Response) error
	GetUsersVacancyIDs(uid int) ([]string, error)
	Add(response *models.Response) error
	Delete(response *models.Response) error
}

type Services struct {
	AuthService
	OptionsService
	VacancyService
	ApplicantsService
	EmployersService
	ResponsesService
}
