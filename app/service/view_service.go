package service

import (
	"ms-api/app/repository"

	"github.com/gin-gonic/gin"
)

type ViewService struct {
	repository *repository.ViewRepository
}

func NewViewService(repository *repository.ViewRepository) *ViewService {
	return &ViewService{repository: repository}
}

func (s ViewService) Add(ctx *gin.Context, videoId int, userId string) error {
	_, err := s.repository.Insert(ctx, videoId, userId)
	if err != nil {
		return err
	}
	return nil
}

func (s ViewService) Total(ctx *gin.Context, videoId int) (*int, error) {
	value, err := s.repository.Count(ctx, videoId)
	if err != nil {
		return nil, err
	}
	return value, nil
}
