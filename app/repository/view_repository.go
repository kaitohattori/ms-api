package repository

import (
	"ms-api/app/model"

	"github.com/gin-gonic/gin"
)

type ViewRepository struct {
}

func NewViewRepository() *ViewRepository {
	return &ViewRepository{}
}

func (r ViewRepository) Insert(ctx *gin.Context, videoId int, userId string) (int, error) {
	return model.View.Insert(model.View{}, ctx, videoId, userId)
}

func (r ViewRepository) Count(ctx *gin.Context, videoId int) (*int, error) {
	return model.View.Count(model.View{}, ctx, videoId)
}
