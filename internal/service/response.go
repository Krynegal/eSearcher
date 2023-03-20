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

func (rs *Responses) GetUIDsByVacancyID(vacancyID string) ([]int, error) {
	UIDs, err := rs.store.GetUIDsByVacancyID(vacancyID)
	if err != nil {
		return nil, err
	}
	return UIDs, nil
}

func (rs *Responses) ChangeStatus(response *models.Response) error {
	if err := rs.store.ChangeStatus(response); err != nil {
		return err
	}
	return nil
}

func (rs *Responses) GetVacancyIDsByUID(uid int) ([]string, error) {
	vacancyIDs, err := rs.store.GetVacancyIDsByUID(uid)
	if err != nil {
		return nil, err
	}
	return vacancyIDs, nil
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
