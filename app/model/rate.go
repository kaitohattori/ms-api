package model

import (
	"errors"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Rate struct {
	Id        int       `json:"id" gorm:"unique;autoIncrement;not null"`
	UserId    string    `json:"userId,omitempty" gorm:"primaryKey"`
	VideoId   int       `json:"videoId,omitempty" gorm:"primaryKey"`
	Value     float32   `json:"value,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}

func RateFindOne(videoId int, userId string) (*Rate, error) {
	rate := Rate{}
	err := DbConnection.Where("video_id = ? AND user_id = ?", videoId, userId).First(&rate).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		rate = Rate{
			Value: 0,
		}
		return &rate, nil
	}
	if err != nil {
		return nil, err
	}
	return &rate, nil
}

func RateAverage(videoId int) (*float32, error) {
	result := []struct {
		Average float32
	}{}
	err := DbConnection.Model(&Rate{}).Select("avg(value) as average").Group("video_id").Having("video_id = ?", videoId).Find(&result).Error
	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		emptyValue := float32(0)
		return &emptyValue, nil
	}
	return &result[0].Average, nil
}

func (r *Rate) Update() (*Rate, error) {
	DbConnection.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "video_id"}, {Name: "user_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"value", "updated_at"}),
	}).Create(&r)

	rate, err := RateFindOne(r.VideoId, r.UserId)
	if err != nil {
		return nil, err
	}
	return rate, nil
}
