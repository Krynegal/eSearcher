package service

import (
	"eSearcher/internal/models"
	"eSearcher/internal/storage"
	"fmt"
)

type Applicants struct {
	store storage.ApplicantStorage
}

func NewApplicants(storage storage.ApplicantStorage) *Applicants {
	return &Applicants{store: storage}
}

func (a *Applicants) Create(applicant *models.Applicant) error {
	id, err := a.store.Create(applicant)
	if err != nil {
		return err
	}
	fmt.Printf("applicant id: %v", id)
	return nil
}

func (a *Applicants) SearchApplicant(params *models.SearchApplicantParams) ([]*models.Applicant, error) {
	applicants, err := a.store.Search(params)
	if err != nil {
		return nil, err
	}
	return applicants, nil
}
