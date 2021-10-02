package model

import (
	"io"
	"ms-api/app/util"
	"time"

	"github.com/gin-gonic/gin"
)

type Media struct {
	FileName string
	Title    string
	UserId   string
}

type ThumbnailImage *io.ReadCloser

func (m *Media) Upload(ctx *gin.Context) (*Video, error) {
	video := Video{
		Title:     m.Title,
		UserId:    m.UserId,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	video.Insert(ctx)

	// Make working directory
	dirPath, err := util.MediaUtil.MakeDirForVideoProcess(video.Id)
	if err != nil {
		return nil, err
	}
	// Save video
	if err := util.MediaUtil.SaveVideo(m.FileName, *dirPath); err != nil {
		return nil, err
	}
	// Process video
	if err := util.MediaUtil.ProcessVideo(ctx, video.Id, *dirPath); err != nil {
		return nil, err
	}
	// Delete working directory
	util.MediaUtil.DeleteDirForVideoProcess(*dirPath)

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
