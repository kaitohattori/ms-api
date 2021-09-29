package model

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

type Video struct {
	Id           int       `json:"id"`
	Title        string    `json:"title,omitempty"`
	ThumbnailUrl string    `json:"thumbnailUrl,omitempty"`
	UserId       string    `json:"userId,omitempty"`
	CreatedAt    time.Time `json:"createdAt,omitempty"`
	UpdatedAt    time.Time `json:"updatedAt,omitempty"`
}

func (Video) FindAll(ctx *gin.Context, filter VideoFilter) ([]Video, error) {
	video := Video{}
	video.Title = fmt.Sprintf("video %d", 10)
	video.ThumbnailUrl = fmt.Sprintf("http://ms-tv.local/web-api/video/%d/thumbnail", 10)
	video.CreatedAt = time.Now()
	video.UpdatedAt = time.Now()

	videos := []Video{video}
	return videos, nil
}

func (Video) FindOne(ctx *gin.Context, videoId int) (*Video, error) {
	video := &Video{}
	video.Title = fmt.Sprintf("video %d", 10)
	video.ThumbnailUrl = fmt.Sprintf("http://ms-tv.local/web-api/video/%d/thumbnail", 10)
	video.CreatedAt = time.Now()
	video.UpdatedAt = time.Now()

	return video, nil
}

func (v Video) Insert(ctx *gin.Context) error {
	return nil
}

func (v Video) Update(ctx *gin.Context) error {
	return nil
}

func (v Video) Delete(ctx *gin.Context) error {
	return nil
}

type AddVideo struct {
	Title string `json:"title" example:"video title"`
}

func (v AddVideo) Valid() error {
	switch {
	case len(v.Title) == 0:
		return ErrNameInvalid
	default:
		return nil
	}
}

func (v AddVideo) SetParamsTo(video *Video) error {
	if err := v.Valid(); err != nil {
		return err
	}
	if v.Title != "" {
		video.Title = v.Title
	}
	return nil
}

type UpdateVideo struct {
	Title        string `json:"title" example:"video title"`
	ThumbnailUrl string `json:"thumbnailUrl" example:"video thumbnailUrl"`
}

func (v UpdateVideo) Valid() error {
	switch {
	case len(v.Title) == 0:
		return ErrNameInvalid
	case len(v.ThumbnailUrl) == 0:
		return ErrNameInvalid
	default:
		return nil
	}
}

func (v UpdateVideo) SetParamsTo(video *Video) error {
	if err := v.Valid(); err != nil {
		return err
	}
	if v.Title != "" {
		video.Title = v.Title
	}
	if v.ThumbnailUrl != "" {
		video.ThumbnailUrl = v.ThumbnailUrl
	}
	return nil
}
