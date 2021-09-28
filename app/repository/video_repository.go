package repository

import "ms-api/app/model"

type VideoRepository struct {
}

func NewVideoRepository() *VideoRepository {
	return &VideoRepository{}
}

func (r VideoRepository) FindAll(filter model.VideoFilter) ([]model.Video, error) {
	return model.Video.FindAll(model.Video{}, filter)
}

func (r VideoRepository) FindOne(videoId int) (*model.Video, error) {
	return model.Video.FindOne(model.Video{}, videoId)
}

func (r VideoRepository) Insert(video *model.Video) (int, error) {
	return video.Insert()
}

func (r VideoRepository) Update(video *model.Video) (bool, error) {
	return video.Update()
}

func (r VideoRepository) Delete(videoId int) (bool, error) {
	video, err := r.FindOne(videoId)
	if err != nil {
		return false, err
	}
	return video.Delete()
}
