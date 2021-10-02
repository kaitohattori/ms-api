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
	return s.repository.FindOne(ctx, videoId, userId)
}

func (s RateService) Average(ctx *gin.Context, videoId int) (*float32, error) {
	return s.repository.Average(ctx, videoId)
}

func (s RateService) Add(ctx *gin.Context, userId string, videoId int) (*model.Rate, error) {
	rate := &model.Rate{
		VideoId: videoId,
		UserId:  userId,
	}
	if err := s.repository.Insert(ctx, rate); err != nil {
		return nil, err
	}
	return rate, nil
}
