package storage

import "eSearcher/internal/models"

type VacancyStorage interface {
	Create(vacancy *models.Vacancy) (string, error)
	Search(params *models.SearchVacancyParams) ([]*models.Vacancy, error)
}

type ApplicantStorage interface {
	Create(applicant *models.Applicant) (string, error)
}

type SpecializationStorage interface {
	Get() ([]*models.Specialization, error)
}

type Storage struct {
	VacancyStorage
	ApplicantStorage
	SpecializationStorage
}
