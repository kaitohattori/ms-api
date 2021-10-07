package model

import (
	"bytes"
	"io"
	"ms-api/app/util"
	"os"

	"github.com/gin-gonic/gin"
)

type ThumbnailImage struct {
	Closer io.ReadCloser
}

func NewThumbnailImage(f *os.File) ThumbnailImage {
	return ThumbnailImage{
		Closer: f,
	}
}

func (ThumbnailImage) Get(ctx *gin.Context, videoId int) (*ThumbnailImage, error) {
	thumbnailFilePath := util.MediaUtil.ThumbnailFilePath(videoId)
	f, err := os.Open(thumbnailFilePath)
	if err != nil {
		return nil, err
	}
	data := NewThumbnailImage(f)
	return &data, nil
}

func (t ThumbnailImage) Bytes() []byte {
	buffer := new(bytes.Buffer)
	io.Copy(buffer, t.Closer)
	return buffer.Bytes()
}
