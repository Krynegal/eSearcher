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

func (s *Specializations) Get() ([]*models.Specialization, error) {
	_, err := s.store.Get()
	if err != nil {
		return nil, err
	}
	return nil, nil
}
