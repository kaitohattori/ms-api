package model

import (
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

type Rate struct {
	Id        int        `json:"id" gorm:"unique;autoIncrement;not null"`
	UserId    string     `json:"userId,omitempty" gorm:"primaryKey"`
	VideoId   int        `json:"videoId,omitempty" gorm:"primaryKey"`
	Value     float32    `json:"value,omitempty"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

func (Rate) FindOne(ctx *gin.Context, videoId int, userId string) (*Rate, error) {
	rate := Rate{}
	ctxDB := DbConnection.WithContext(ctx.Request.Context())
	err := ctxDB.Where("video_id = ? AND user_id = ?", videoId, userId).First(&rate).Error
	if err != nil {
		return nil, err
	}

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		return &rate, nil
	}
}

func (Rate) Average(ctx *gin.Context, videoId int) (*float32, error) {
	result := []struct {
		Average float32
	}{}
	ctxDB := DbConnection.WithContext(ctx)
	err := ctxDB.Model(&Rate{}).Select("avg(value) as average").Group("video_id").Having("video_id = ?", videoId).Find(&result).Error
	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, ErrRecordNotFound
	}

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		return &result[0].Average, nil
	}
}

func (r *Rate) Update(ctx *gin.Context) (*Rate, error) {
	ctxDB := DbConnection.WithContext(ctx)
	ctxDB.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "video_id"}, {Name: "user_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"value", "updated_at"}),
	}).Create(&r)

	rate, err := Rate.FindOne(Rate{}, ctx, r.VideoId, r.UserId)
	if err != nil {
		return nil, err
	}

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		return rate, nil
	}
}
