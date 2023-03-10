package service

import (
	"eSearcher/internal/models"
	"eSearcher/internal/storage"
)

type Responses struct {
	store storage.ResponsesStorage
}

func NewResponses(storage storage.ResponsesStorage) *Responses {
	return &Responses{store: storage}
}

func (rs *Responses) Add(response *models.Response) error {
	if err := rs.store.Add(response); err != nil {
		return err
	}
	return nil
}

func (rs *Responses) Delete(response *models.Response) error {
	if err := rs.store.Delete(response); err != nil {
		return err
	}
	return nil
}
