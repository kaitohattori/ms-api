package service

import (
	"fmt"
	"ms-api/app/model"
	"ms-api/app/repository"
	"time"
)

type VideoService struct {
	repository *repository.VideoRepository
}

func NewVideoService(repository *repository.VideoRepository) *VideoService {
	return &VideoService{repository: repository}
}

func (s VideoService) Find(filter *model.VideoFilter) ([]model.Video, error) {
	return s.repository.FindAll(filter)
}

func (s VideoService) Get(videoId int) (*model.Video, error) {
	return s.repository.FindOne(videoId)
}

func (s VideoService) Add(addVideo model.AddVideo) (*model.Video, error) {
	video := &model.Video{
		Title:        addVideo.Title,
		ThumbnailUrl: fmt.Sprintf("http://ms-tv.local/web-api/video/%d/thumbnail", 10),
		UserId:       "test_user_id",
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
	lastId, err := s.repository.Insert(video)
	if err != nil {
		return nil, err
	}
	video.Id = lastId
	return video, nil
}

func (s VideoService) Update(videoId int, updateVideo model.UpdateVideo) (*model.Video, error) {
	video := &model.Video{
		Id:           videoId,
		Title:        updateVideo.Title,
		ThumbnailUrl: updateVideo.ThumbnailUrl,
		UserId:       updateVideo.UserId,
	}
	_, err := s.repository.Update(video)
	if err != nil {
		return nil, err
	}
	return video, nil
}

func (s VideoService) Remove(videoId int) (bool, error) {
	return s.repository.Delete(videoId)
}
