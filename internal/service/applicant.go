package service

import (
	"eSearcher/internal/models"
	"eSearcher/internal/storage"
)

type Applicants struct {
	store storage.ApplicantStorage
}

func NewApplicants(storage storage.ApplicantStorage) *Applicants {
	return &Applicants{store: storage}
}

func (a *Applicants) Create(applicant *models.Applicant) error {
	_, err := a.store.Create(applicant)
	if err != nil {
		return err
	}
	return nil
}
