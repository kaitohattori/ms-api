package model

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"ms-api/config"
	"net/http"
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

func (Video) FindAllSortedByViewCount(ctx *gin.Context, filter VideoFilter) ([]Video, error) {
	videos := []Video{}
	ctxDB := DbConnection.WithContext(ctx)
	// subQuery := ctxDB.Select("video_id", "count(id) as view_count").Table("views").Group("video_id")
	// subQuery := fmt.Sprintf("select video_id, count(id) as view_count from views group by video_id")
	// subQuery := "select video_id, count(id) as view_count from views group by video_id"
	// query := ctxDB.Model(&Video{}).Joins("left join (?) as v on videos.id = v.video_id", subQuery)
	// TODO: もうちょっとかっこよく書きたい
	query := ctxDB.Model(&Video{}).Joins("left join (select video_id, count(id) as view_count from views group by video_id) as v on videos.id = v.video_id")
	if filter.UserId != nil && *filter.UserId != "" {
		query.Where("user_id = ?", filter.UserId)
	}
	query.Order("COALESCE(view_count, 0) desc")
	if filter.Limit != nil && *filter.Limit != 0 {
		query.Limit(*filter.Limit)
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

func (Video) FindAllRecommended(ctx *gin.Context, filter VideoFilter) ([]Video, error) {
	videos := []Video{}
	url := fmt.Sprintf("%s/videos/recommended?userId=%s&limit=%d", config.Config.RecommendationAPIURL(), *filter.UserId, *filter.Limit)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx.Request.Context())
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	byteArray, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(byteArray, &videos)
	if err != nil {
		return nil, err
	}

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
	if err := ctxDB.First(&video, videoId).Error; err != nil {
		return nil, err
	}

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		return &video, nil
	}
}

func (v *Video) Insert(ctx *gin.Context) error {
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

func (v *Video) Update(ctx *gin.Context) error {
	ctxDB := DbConnection.WithContext(ctx.Request.Context())
	if err := ctxDB.Save(&v).Error; err != nil {
		return err
	}

	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		return nil
	}
}

func (v *Video) Delete(ctx *gin.Context) error {
	ctxDB := DbConnection.WithContext(ctx.Request.Context())
	if err := ctxDB.Delete(&v).Error; err != nil {
		return err
	}

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
