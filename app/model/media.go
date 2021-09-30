package model

import (
	"io"

	"github.com/gin-gonic/gin"
)

type Media struct {
	FileName string
	Title    string
	UserId   string
}

type ThumbnailImage *io.ReadCloser

func (v Media) Upload(ctx *gin.Context) error {
	return nil
}

func (Media) GetThumbnailImage(ctx *gin.Context, videoId int) (ThumbnailImage, error) {
	return nil, nil
}
