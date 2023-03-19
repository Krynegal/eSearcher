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

func (a *Applicants) Get(id int) (*models.Applicant, error) {
	applicant, err := a.store.Get(id)
	if err != nil {
		return nil, err
	}
	return applicant, nil
}

func (a *Applicants) SearchApplicant(params *models.SearchApplicantParams) ([]*models.Applicant, error) {
	applicantIDs, err := a.store.Search(params)
	if err != nil {
		return nil, err
	}
	var applicants []*models.Applicant
	for _, id := range applicantIDs {
		var applicant *models.Applicant
		applicant, err = a.store.Get(id)
		if err != nil {
			return nil, err
		}
		applicants = append(applicants, applicant)
	}
	return applicants, nil
}
