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

func (r RateRepository) Insert(ctx *gin.Context, videoId int, userId string) (int, error) {
	return model.Rate.Insert(model.Rate{}, ctx, videoId, userId)
}

func (r RateRepository) Average(ctx *gin.Context, videoId int) (*float32, error) {
	return model.Rate.Average(model.Rate{}, ctx, videoId)
}
