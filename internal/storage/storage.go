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
	GetEmployerVacancies(uid int) ([]*models.Vacancy, error)
	GetByID(id string) (*models.Vacancy, error)
	Create(vacancy *models.Vacancy) (string, error)
	Search(params *models.SearchVacancyParams) ([]*models.Vacancy, error)
}

type ApplicantStorage interface {
	Create(applicant *models.Applicant) error
	Get(id int) (*models.Applicant, error)
	Update(applicant *models.Applicant) error
	Search(params *models.SearchApplicantParams) ([]int, error)
}

type EmployerStorage interface {
	Get(id int) (*models.Employer, error)
	Create(employer *models.Employer) error
	Update(employer *models.Employer) error
}

type ResponsesStorage interface {
	GetUIDsByVacancyID(vacancyID string) ([]int, error)
	ChangeStatus(response *models.Response) error
	GetVacancyIDsByUID(uid int) ([]string, error)
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
