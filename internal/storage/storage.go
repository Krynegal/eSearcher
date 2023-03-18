package storage

import "eSearcher/internal/models"

type AuthStorage interface {
	CreateUser(login, password string, role int) (int, error)
	GetUser(login, password string) (*models.User, error)
}

type OptionsStorage interface {
	GetAll(option string) ([]*models.Option, error)
}

type VacancyStorage interface {
	Create(vacancy *models.Vacancy) (string, error)
	Search(params *models.SearchVacancyParams) ([]*models.Vacancy, error)
}

type ApplicantStorage interface {
	Create(applicant *models.Applicant) (int, error)
	Get(id string) (*models.Applicant, error)
	Search(params *models.SearchApplicantParams) ([]string, error)
}

type EmployerStorage interface {
	Create(employer *models.Employer) (string, error)
}

type ResponsesStorage interface {
	Add(response *models.Response) error
	Delete(response *models.Response) error
}

type Storage struct {
	AuthStorage
	OptionsStorage
	VacancyStorage
	ApplicantStorage
	EmployerStorage
	ResponsesStorage
}
