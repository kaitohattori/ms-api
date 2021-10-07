package model

import (
	"bytes"
	"io"
	"mime/multipart"
	"ms-api/app/util"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

type Media struct {
	File       multipart.File
	FileHeader multipart.FileHeader
	Title      string
	UserId     string
}

func (m *Media) Upload(ctx *gin.Context) (*Video, error) {
	ctxDB := DbConnection.WithContext(ctx.Request.Context())
	tx := ctxDB.Begin()

	video := Video{
		Title:     m.Title,
		UserId:    m.UserId,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if err := tx.Create(&video).Error; err != nil {
		// Rollback
		tx.Rollback()
		return nil, err
	}

	// Make working directory
	dirPath, err := util.MediaUtil.MakeWorkingDirectory(video.Id)
	if err != nil {
		// Rollback
		tx.Rollback()
		// Delete working directory
		util.MediaUtil.DeleteWorkingDirectory(*dirPath)
		return nil, err
	}
	// Save video
	videoFilePath, err := util.MediaUtil.SaveVideo(m.File, m.FileHeader, *dirPath)
	if err != nil {
		// Rollback
		tx.Rollback()
		// Delete working directory
		util.MediaUtil.DeleteWorkingDirectory(*dirPath)
		return nil, err
	}
	// Process video
	if err := util.MediaUtil.ProcessVideo(ctx, video.Id, *videoFilePath); err != nil {
		// Rollback
		tx.Rollback()
		// Delete working directory
		util.MediaUtil.DeleteWorkingDirectory(*dirPath)
		return nil, err
	}
	// Delete working directory
	util.MediaUtil.DeleteWorkingDirectory(*dirPath)

	select {
	case <-ctx.Done():
		tx.Rollback()
		return nil, ctx.Err()
	default:
		tx.Commit()
		return &video, nil
	}
}

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
