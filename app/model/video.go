package model

import (
	"time"

	"github.com/gin-gonic/gin"
)

type Video struct {
	Id        int       `json:"id" gorm:"primaryKey"`
	UserId    string    `json:"userId,omitempty"`
	Title     string    `json:"title,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}

func (Video) FindAll(ctx *gin.Context, filter VideoFilter) ([]Video, error) {
	videos := []Video{}
	ctxDB := DbConnection.WithContext(ctx)
	subQuery := ctxDB.Model(&View{}).Select("video_id", "count(id) as view_count").Group("video_id")
	query := ctxDB.Model(&Video{}).Joins("left join (?) as views on videos.id = views.video_id", subQuery)
	if filter.UserId != nil && *filter.UserId != "" {
		query = query.Where("UserId = ?", filter.UserId)
	}
	query = query.Order("COALESCE(view_count, 0) desc")
	if filter.Limit != nil && *filter.Limit != 0 {
		query = query.Limit(*filter.Limit)
	}
	query.Find(&videos)

	// ctxDB.Model(&Video{}).Find(&videos)

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		return videos, nil
	}
}

func (Video) FindOne(ctx *gin.Context, videoId int) (*Video, error) {
	video := Video{}
	ctxDB := DbConnection.WithContext(ctx.Request.Context())
	ctxDB.First(&video, videoId)

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		return &video, nil
	}
}

func (v *Video) Insert(ctx *gin.Context) error {
	ctxDB := DbConnection.WithContext(ctx.Request.Context())
	ctxDB.Create(&v)

	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		return nil
	}
}

func (v *Video) Update(ctx *gin.Context) error {
	ctxDB := DbConnection.WithContext(ctx.Request.Context())
	ctxDB.Save(&v)

	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		return nil
	}
}

func (v *Video) Delete(ctx *gin.Context) error {
	ctxDB := DbConnection.WithContext(ctx.Request.Context())
	ctxDB.Delete(&v)

	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		return nil
	}
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
