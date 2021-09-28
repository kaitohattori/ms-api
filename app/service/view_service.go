package service

import (
	"ms-api/app/model"
	"ms-api/app/repository"
)

type ViewService struct {
	repository *repository.ViewRepository
}

func NewViewService(repository *repository.ViewRepository) *ViewService {
	return &ViewService{repository: repository}
}

func (s ViewService) Total(filter model.ViewFilter) (*int, error) {
	value, err := s.repository.Count(filter)
	if err != nil {
		return nil, err
	}
	return value, nil
}

func (s ViewService) Add(videoId int) error {
	_, err := s.repository.Insert(videoId)
	if err != nil {
		return err
	}
	return nil
}
