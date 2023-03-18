package service

import (
	"eSearcher/internal/models"
	"eSearcher/internal/storage"
)

type Options struct {
	store storage.OptionsStorage
}

func NewOptions(storage storage.OptionsStorage) *Options {
	return &Options{store: storage}
}

func (s *Options) GetAll(options []string) (map[string][]*models.Option, error) {
	res := make(map[string][]*models.Option)
	for _, opt := range options {
		opts, err := s.store.GetAll(opt)
		if err != nil {
			return nil, err
		}
		res[opt] = opts
	}
	return res, nil
}
