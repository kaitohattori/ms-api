package service

import (
	"ms-api/app/model"
	"ms-api/app/repository"

	"github.com/gin-gonic/gin"
)

type MediaService struct {
	repository *repository.MediaRepository
}

func NewMediaService(repository *repository.MediaRepository) *MediaService {
	return &MediaService{repository: repository}
}

func (s MediaService) Upload(ctx *gin.Context, media *model.Media) error {
	return s.repository.Upload(ctx, media)
}

func (s MediaService) GetThumbnailImage(ctx *gin.Context, videoId int) (model.ThumbnailImage, error) {
	return s.repository.GetThumbnailImage(ctx, videoId)
}
