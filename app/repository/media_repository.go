package repository

import (
	"ms-api/app/model"

	"github.com/gin-gonic/gin"
)

type MediaRepository struct {
}

func NewMediaRepository() *MediaRepository {
	return &MediaRepository{}
}

func (r MediaRepository) Upload(ctx *gin.Context, media *model.Media) (*model.Video, error) {
	return media.Upload(ctx)
}

func (r MediaRepository) GetThumbnailImage(ctx *gin.Context, videoId int) (model.ThumbnailImage, error) {
	return model.Media.GetThumbnailImage(model.Media{}, ctx, videoId)
}
