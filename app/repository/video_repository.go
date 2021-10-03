package repository

import (
	"ms-api/app/model"

	"github.com/gin-gonic/gin"
)

type VideoRepository struct {
}

func NewVideoRepository() *VideoRepository {
	return &VideoRepository{}
}

func (r VideoRepository) FindAllSortedByViewCount(ctx *gin.Context, filter model.VideoFilter) ([]model.Video, error) {
	return model.Video.FindAllSortedByViewCount(model.Video{}, ctx, filter)
}

func (r VideoRepository) FindOne(ctx *gin.Context, videoId int) (*model.Video, error) {
	return model.Video.FindOne(model.Video{}, ctx, videoId)
}

func (r VideoRepository) Insert(ctx *gin.Context, video *model.Video) error {
	return video.Insert(ctx)
}

func (r VideoRepository) Update(ctx *gin.Context, video *model.Video) error {
	return video.Update(ctx)
}

func (r VideoRepository) Delete(ctx *gin.Context, videoId int) error {
	video, err := r.FindOne(ctx, videoId)
	if err != nil {
		return err
	}
	if err := video.Delete(ctx); err != nil {
		return err
	}
	return nil
}
