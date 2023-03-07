package storage

import "eSearcher/internal/models"

type VacancyStorage interface {
	Create(vacancy *models.Vacancy) (string, error)
	Search(params *models.SearchVacancyParams) ([]*models.Vacancy, error)
}

type ApplicantStorage interface {
	Create(applicant *models.Applicant) (int, error)
	Search(params *models.SearchApplicantParams) ([]*models.Applicant, error)
}

type SpecializationStorage interface {
	GetAllSpecializations() ([]*models.Specialization, error)
}

type EmployerStorage interface {
	Create(employer *models.Employer) (string, error)
}

type Storage struct {
	VacancyStorage
	ApplicantStorage
	SpecializationStorage
	EmployerStorage
}
