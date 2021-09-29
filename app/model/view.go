package model

import (
	"time"

	"github.com/gin-gonic/gin"
)

type View struct {
	Id        int       `json:"id"`
	UserId    string    `json:"userId,omitempty"`
	VideoId   int       `json:"videoId,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}

func (View) Count(ctx *gin.Context, filter ViewFilter) (*int, error) {
	var value int = 15
	return &value, nil
}

func (View) Insert(ctx *gin.Context, videoId int) (int, error) {
	return 10, nil
}
