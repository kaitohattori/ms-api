package model

import (
	"time"

	"github.com/gin-gonic/gin"
)

type Analysis struct {
	Id        int       `json:"id" gorm:"primaryKey"`
	UserId    string    `json:"userId,omitempty"`
	VideoId   int       `json:"videoId,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}

func (Analysis) Count(ctx *gin.Context, videoId int) (*int, error) {
	var count int64
	ctxDB := DbConnection.WithContext(ctx.Request.Context())
	if err := ctxDB.Model(&Analysis{}).Where("video_id = ?", videoId).Count(&count).Error; err != nil {
		return nil, err
	}
	int_count := int(count)

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		return &int_count, nil
	}
}

func (v *Analysis) Insert(ctx *gin.Context) error {
	ctxDB := DbConnection.WithContext(ctx.Request.Context())
	if err := ctxDB.Create(&v).Error; err != nil {
		return err
	}

	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		return nil
	}
}
