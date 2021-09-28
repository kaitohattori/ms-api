package model

import (
	"errors"
	"fmt"
	"time"
)

type Video struct {
	Id           int       `json:"id"`
	Title        string    `json:"title,omitempty"`
	ThumbnailUrl string    `json:"thumbnailUrl,omitempty"`
	UserId       string    `json:"userId,omitempty"`
	CreatedAt    time.Time `json:"createdAt,omitempty"`
	UpdatedAt    time.Time `json:"updatedAt,omitempty"`
}

//  example
var (
	ErrNameInvalid = errors.New("name is empty")
)

type AddVideo struct {
	Title string `json:"title" example:"video title"`
}

func (v AddVideo) Validation() error {
	switch {
	case len(v.Title) == 0:
		return ErrNameInvalid
	default:
		return nil
	}
}

type UpdateVideo struct {
	Title        string `json:"title" example:"video title"`
	ThumbnailUrl string `json:"thumbnailUrl" example:"video thumbnailUrl"`
	UserId       string `json:"userId" example:"video userId"`
}

func (v UpdateVideo) Validation() error {
	switch {
	case len(v.Title) == 0:
		return ErrNameInvalid
	case len(v.ThumbnailUrl) == 0:
		return ErrNameInvalid
	case len(v.UserId) == 0:
		return ErrNameInvalid
	default:
		return nil
	}
}

func (Video) FindAll(filter *VideoFilter) ([]Video, error) {
	video := Video{}
	video.Title = fmt.Sprintf("video %d", 10)
	video.ThumbnailUrl = fmt.Sprintf("http://ms-tv.local/web-api/video/%d/thumbnail", 10)
	video.CreatedAt = time.Now()
	video.UpdatedAt = time.Now()

	videos := []Video{video}
	fmt.Println("hello world")
	return videos, nil
}

func (Video) FindOne(videoId int) (*Video, error) {
	video := &Video{}
	video.Title = fmt.Sprintf("video %d", 10)
	video.ThumbnailUrl = fmt.Sprintf("http://ms-tv.local/web-api/video/%d/thumbnail", 10)
	video.CreatedAt = time.Now()
	video.UpdatedAt = time.Now()

	return video, nil
}

func (v Video) Insert() (int, error) {
	return 10, nil
}

func (v Video) Update() (bool, error) {
	return true, nil
}

func (v Video) Delete() (bool, error) {
	return true, nil
}
