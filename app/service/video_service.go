package service

import (
	"fmt"
	"ms-api/app/model"
	"ms-api/app/repository"
	"time"

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

func (s VideoService) Add(ctx *gin.Context, addVideo model.AddVideo, userId string) (*model.Video, error) {
	video := &model.Video{
		Title:        addVideo.Title,
		ThumbnailUrl: fmt.Sprintf("http://ms-tv.local/web-api/video/%d/thumbnail", 10),
		UserId:       "test_user_id",
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
	newVideo, err := s.repository.Insert(ctx, video, userId)
	if err != nil {
		return nil, err
	}
	return newVideo, nil
}

func (s VideoService) Update(ctx *gin.Context, videoId int, updateVideo model.UpdateVideo, userId string) (*model.Video, error) {
	video := &model.Video{
		Id:           videoId,
		Title:        updateVideo.Title,
		ThumbnailUrl: updateVideo.ThumbnailUrl,
		UserId:       updateVideo.UserId,
	}
	if err := s.repository.Update(ctx, video, userId); err != nil {
		return nil, err
	}
	return video, nil
}

func (s VideoService) Remove(ctx *gin.Context, videoId int, userId string) error {
	return s.repository.Delete(ctx, videoId, userId)
}
