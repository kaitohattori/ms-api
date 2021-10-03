package model

import (
	"io"
	"mime/multipart"
	"ms-api/app/util"
	"time"

	"github.com/gin-gonic/gin"
)

type Media struct {
	File       multipart.File
	FileHeader multipart.FileHeader
	Title      string
	UserId     string
}

type ThumbnailImage *io.ReadCloser

func (m *Media) Upload(ctx *gin.Context) (*Video, error) {
	video := Video{
		Title:     m.Title,
		UserId:    m.UserId,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if err := video.Insert(ctx); err != nil {
		return nil, err
	}

	// Make working directory
	dirPath, err := util.MediaUtil.MakeWorkingDirectory(video.Id)
	if err != nil {
		return nil, err
	}
	// Save video
	videoFilePath, err := util.MediaUtil.SaveVideo(m.File, m.FileHeader, *dirPath)
	if err != nil {
		return nil, err
	}
	// Process video
	if err := util.MediaUtil.ProcessVideo(ctx, video.Id, *videoFilePath); err != nil {
		return nil, err
	}
	// Delete working directory
	util.MediaUtil.DeleteWorkingDirectory(*dirPath)

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		return &video, nil
	}
}

func (Media) GetThumbnailImage(ctx *gin.Context, videoId int) (ThumbnailImage, error) {
	return nil, nil
}
