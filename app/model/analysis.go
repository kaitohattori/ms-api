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

func AnalysisCount(ctx *gin.Context, videoId int) (*int, error) {
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

func AnalysisInsert(ctx *gin.Context, userId string, videoId int) (*Analysis, error) {
	analysis := &Analysis{
		VideoId: videoId,
		UserId:  userId,
	}
	ctxDB := DbConnection.WithContext(ctx.Request.Context())
	if err := ctxDB.Create(analysis).Error; err != nil {
		return nil, err
	}

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		return analysis, nil
	}
}
