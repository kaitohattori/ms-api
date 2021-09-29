package service

import (
	"ms-api/app/model"
	"ms-api/app/repository"

	"github.com/gin-gonic/gin"
)

type ViewService struct {
	repository *repository.ViewRepository
}

func NewViewService(repository *repository.ViewRepository) *ViewService {
	return &ViewService{repository: repository}
}

func (s ViewService) Total(ctx *gin.Context, filter model.ViewFilter) (*int, error) {
	value, err := s.repository.Count(ctx, filter)
	if err != nil {
		return nil, err
	}
	return value, nil
}

func (s ViewService) Add(ctx *gin.Context, videoId int) error {
	_, err := s.repository.Insert(ctx, videoId)
	if err != nil {
		return err
	}
	return nil
}
