package model

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

type Video struct {
	Id        int       `json:"id"`
	UserId    string    `json:"userId,omitempty"`
	Title     string    `json:"title,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}

func (Video) FindAll(ctx *gin.Context, filter VideoFilter) ([]Video, error) {
	video := Video{}
	video.Title = fmt.Sprintf("video %d", 10)
	video.CreatedAt = time.Now()
	video.UpdatedAt = time.Now()

	videos := []Video{video}
	return videos, nil
}

func (Video) FindOne(ctx *gin.Context, videoId int) (*Video, error) {
	video := &Video{}
	video.Id = videoId
	video.Title = fmt.Sprintf("video %d", 10)
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
	Title string `json:"title" example:"video title"`
}

func (v UpdateVideo) Valid() error {
	switch {
	case len(v.Title) == 0:
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
	return nil
}
