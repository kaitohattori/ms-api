package service

import (
	"ms-api/app/model"
	"ms-api/app/repository"

	"github.com/gin-gonic/gin"
)

type RateService struct {
	repository *repository.RateRepository
}

func NewRateService(repository *repository.RateRepository) *RateService {
	return &RateService{repository: repository}
}

func (s RateService) Get(ctx *gin.Context, videoId int, userId string) (*model.Rate, error) {
	value, err := s.repository.FindOne(ctx, videoId, userId)
	if err != nil {
		return nil, err
	}
	return value, nil
}

func (s RateService) Add(ctx *gin.Context, videoId int, userId string) error {
	_, err := s.repository.Insert(ctx, videoId, userId)
	if err != nil {
		return err
	}
	return nil
}

func (s RateService) Average(ctx *gin.Context, videoId int) (*float32, error) {
	value, err := s.repository.Average(ctx, videoId)
	if err != nil {
		return nil, err
	}
	return value, nil
}
