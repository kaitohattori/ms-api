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

func (r ViewRepository) Count(ctx *gin.Context, videoId int) (*int, error) {
	return model.View.Count(model.View{}, ctx, videoId)
}

func (r ViewRepository) Insert(ctx *gin.Context, view *model.View) error {
	return view.Insert(ctx)
}
