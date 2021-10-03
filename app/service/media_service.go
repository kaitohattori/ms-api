package service

import (
	"mime/multipart"
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

func (s MediaService) Upload(ctx *gin.Context, userId string, title string, file multipart.File, header multipart.FileHeader) (*model.Video, error) {
	media := &model.Media{
		File:       file,
		FileHeader: header,
		Title:      title,
		UserId:     userId,
	}
	return s.repository.Upload(ctx, media)
}

func (s MediaService) GetThumbnailImage(ctx *gin.Context, videoId int) (model.ThumbnailImage, error) {
	return s.repository.GetThumbnailImage(ctx, videoId)
}
