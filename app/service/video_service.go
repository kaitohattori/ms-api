package service

import (
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
	if filter.SortType == model.VideoSortTypePopular {
		return s.repository.FindAllSortedByViewCount(ctx, filter)
	} else if filter.SortType == model.VideoSortTypeRecommended {
		return s.repository.FindAllRecommended(ctx, filter)
	} else {
		return nil, model.ErrRecordNotFound
	}
}

func (s VideoService) Get(ctx *gin.Context, videoId int) (*model.Video, error) {
	return s.repository.FindOne(ctx, videoId)
}

func (s VideoService) Add(ctx *gin.Context, userId string, addVideo *model.AddVideo) (*model.Video, error) {
	video := &model.Video{
		UserId: userId,
	}
	addVideo.SetParamsTo(video)
	if err := s.repository.Insert(ctx, video); err != nil {
		return nil, err
	}
	return video, nil
}

func (s VideoService) Update(ctx *gin.Context, userId string, videoId int, updateVideo *model.UpdateVideo) (*model.Video, error) {
	video := &model.Video{
		Id:        videoId,
		UserId:    userId,
		UpdatedAt: time.Now(),
	}
	updateVideo.SetParamsTo(video)
	if err := s.repository.Update(ctx, video); err != nil {
		return nil, err
	}
	return video, nil
}

func (s VideoService) Remove(ctx *gin.Context, userId string, videoId int) error {
	return s.repository.Delete(ctx, videoId)
}
