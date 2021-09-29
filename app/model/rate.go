package model

import (
	"time"

	"github.com/gin-gonic/gin"
)

type Rate struct {
	Id        int       `json:"id"`
	UserId    string    `json:"userId,omitempty"`
	VideoId   int       `json:"videoId,omitempty"`
	Value     float32   `json:"value,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}

func (Rate) FindOne(ctx *gin.Context, videoId int, userId string) (*Rate, error) {
	rate := &Rate{}
	rate.Id = 1
	rate.UserId = userId
	rate.VideoId = videoId
	rate.Value = 3.0
	rate.CreatedAt = time.Now()
	rate.UpdatedAt = time.Now()
	return rate, nil
}

func (Rate) Average(ctx *gin.Context, videoId int) (*float32, error) {
	var value float32 = 3.0
	return &value, nil
}

func (r Rate) Insert(ctx *gin.Context) error {
	return nil
}
