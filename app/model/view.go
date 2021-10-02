package model

import (
	"time"

	"github.com/gin-gonic/gin"
)

type View struct {
	Id        int       `json:"id" gorm:"primaryKey"`
	UserId    string    `json:"userId,omitempty"`
	VideoId   int       `json:"videoId,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}

func (View) Count(ctx *gin.Context, videoId int) (*int, error) {
	var count int64
	ctxDB := DbConnection.WithContext(ctx.Request.Context())
	ctxDB.Model(&View{}).Where("video_id = ?", videoId).Count(&count)
	count2 := int(count)
	return &count2, nil
}

func (v *View) Insert(ctx *gin.Context) error {
	ctxDB := DbConnection.WithContext(ctx.Request.Context())
	ctxDB.Create(&v)

	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		return nil
	}
}
