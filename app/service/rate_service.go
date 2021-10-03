package service

import (
	"ms-api/app/model"
	"ms-api/app/repository"
	"time"

	"github.com/gin-gonic/gin"
)

type RateService struct {
	repository *repository.RateRepository
}

func NewRateService(repository *repository.RateRepository) *RateService {
	return &RateService{repository: repository}
}

func (s RateService) Get(ctx *gin.Context, videoId int, userId string) (*model.Rate, error) {
	return s.repository.FindOne(ctx, videoId, userId)
}

func (s RateService) Average(ctx *gin.Context, videoId int) (*float32, error) {
	return s.repository.Average(ctx, videoId)
}

func (s RateService) Update(ctx *gin.Context, userId string, videoId int, value float32) (*model.Rate, error) {
	now := time.Now()
	rate := &model.Rate{
		VideoId:   videoId,
		UserId:    userId,
		Value:     value,
		UpdatedAt: &now,
	}
	updatedRate, err := s.repository.Update(ctx, rate)
	if err != nil {
		return nil, err
	}
	return updatedRate, nil
}
