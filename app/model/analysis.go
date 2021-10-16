package model

import (
	"time"
)

type Analysis struct {
	Id        int       `json:"id" gorm:"primaryKey"`
	UserId    string    `json:"userId,omitempty"`
	VideoId   int       `json:"videoId,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}

func AnalysisCount(videoId int) (*int, error) {
	var count int64
	if err := DbConnection.Model(&Analysis{}).Where("video_id = ?", videoId).Count(&count).Error; err != nil {
		return nil, err
	}
	int_count := int(count)
	return &int_count, nil
}

func AnalysisInsert(userId string, videoId int) (*Analysis, error) {
	analysis := &Analysis{
		VideoId: videoId,
		UserId:  userId,
	}
	if err := DbConnection.Create(analysis).Error; err != nil {
		return nil, err
	}
	return analysis, nil
}
