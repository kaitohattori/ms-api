package service

import (
	"ms-api/app/model"
	"ms-api/app/repository"

	"github.com/gin-gonic/gin"
)

type VideoService struct {
	repository *repository.VideoRepository
}

func NewVideoService(repository *repository.VideoRepository) *VideoService {
	return &VideoService{repository: repository}
}

func (s VideoService) Find(ctx *gin.Context, filter model.VideoFilter) ([]model.Video, error) {
	return s.repository.FindAll(ctx, filter)
}

func (s VideoService) Get(ctx *gin.Context, videoId int) (*model.Video, error) {
	return s.repository.FindOne(ctx, videoId)
}

func (s VideoService) Add(ctx *gin.Context, video *model.Video) error {
	return s.repository.Insert(ctx, video)
}

func (s VideoService) Update(ctx *gin.Context, video *model.Video) error {
	return s.repository.Update(ctx, video)
}

func (s VideoService) Remove(ctx *gin.Context, videoId int) error {
	return s.repository.Delete(ctx, videoId)
}
