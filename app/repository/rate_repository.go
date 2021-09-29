package repository

import (
	"ms-api/app/model"

	"github.com/gin-gonic/gin"
)

type RateRepository struct {
}

func NewRateRepository() *RateRepository {
	return &RateRepository{}
}

func (r RateRepository) FindOne(ctx *gin.Context, videoId int, userId string) (*model.Rate, error) {
	return model.Rate.FindOne(model.Rate{}, ctx, videoId, userId)
}

func (r RateRepository) Insert(ctx *gin.Context, videoId int, userId string) (*model.Rate, error) {
	rate := model.Rate{}
	_, err := model.Rate.Insert(rate, ctx, videoId, userId)
	if err != nil {
		return nil, err
	}
	return &rate, nil
}

func (r RateRepository) Average(ctx *gin.Context, videoId int) (*float32, error) {
	return model.Rate.Average(model.Rate{}, ctx, videoId)
}
