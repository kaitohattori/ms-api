package model

import (
	"bytes"
	"io"
	"ms-api/app/util"
	"os"
)

type Thumbnail struct {
	Closer io.ReadCloser
}

func NewThumbnail(f *os.File) Thumbnail {
	return Thumbnail{
		Closer: f,
	}
}

func ThumbnailGet(videoId int) (*Thumbnail, error) {
	thumbnailFilePath := util.FileUtilThumbnailFilePath(videoId)
	f, err := os.Open(thumbnailFilePath)
	if err != nil {
		return nil, err
	}
	data := NewThumbnail(f)
	return &data, nil
}

func (t Thumbnail) Bytes() []byte {
	buffer := new(bytes.Buffer)
	io.Copy(buffer, t.Closer)
	return buffer.Bytes()
}
