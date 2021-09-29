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

func (r ViewRepository) Insert(ctx *gin.Context, videoId int, userId string) (*model.View, error) {
	view := model.View{}
	_, err := model.View.Insert(view, ctx, videoId, userId)
	if err != nil {
		return nil, err
	}
	return &view, nil
}

func (r ViewRepository) Count(ctx *gin.Context, videoId int) (*int, error) {
	return model.View.Count(model.View{}, ctx, videoId)
}
