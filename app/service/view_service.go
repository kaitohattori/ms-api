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

func (s ViewService) Total(ctx *gin.Context, videoId int) (*int, error) {
	return s.repository.Count(ctx, videoId)
}

func (s ViewService) Add(ctx *gin.Context, view *model.View) error {
	return s.repository.Insert(ctx, view)
}
