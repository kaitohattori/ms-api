package model

import (
	"bytes"
	"io"
	"ms-api/app/util"
	"os"
)

type Thumbnail struct {
	File io.ReadCloser
}

func NewThumbnail(file *os.File) Thumbnail {
	return Thumbnail{
		File: file,
	}
}

func ThumbnailGet(videoId int) (*Thumbnail, error) {
	thumbnailFilePath := util.FileUtilThumbnailFilePath(videoId)
	file, err := os.Open(thumbnailFilePath)
	if err != nil {
		return nil, err
	}
	data := NewThumbnail(file)
	return &data, nil
}

func (t Thumbnail) Bytes() []byte {
	buffer := new(bytes.Buffer)
	io.Copy(buffer, t.File)
	return buffer.Bytes()
}
