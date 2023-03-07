package service

import (
	"eSearcher/internal/models"
	"eSearcher/internal/storage"
)

type Specializations struct {
	store storage.SpecializationStorage
}

func NewSpecializations(storage storage.SpecializationStorage) *Specializations {
	return &Specializations{store: storage}
}

func (s *Specializations) GetAllSpecializations() ([]*models.Specialization, error) {
	specializations, err := s.store.GetAllSpecializations()
	if err != nil {
		return nil, err
	}
	return specializations, nil
}
