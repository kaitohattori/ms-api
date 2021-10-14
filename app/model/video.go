package model

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"ms-api/app/util"
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

func VideoFind(ctx *gin.Context, filter VideoFilter) ([]Video, error) {
	if filter.SortType == VideoSortTypePopular {
		return VideoFindAllSortedByAnalysisCount(ctx, filter)
	} else if filter.SortType == VideoSortTypeRecommended {
		return VideoFindAllRecommended(ctx, filter)
	} else {
		return nil, ErrRecordNotFound
	}
}

func VideoFindAllSortedByAnalysisCount(ctx *gin.Context, filter VideoFilter) ([]Video, error) {
	videos := []Video{}
	ctxDB := DbConnection.WithContext(ctx)
	// subQuery := ctxDB.Select("video_id", "count(id) as analysis_count").Table("analyses").Group("video_id")
	// subQuery := fmt.Sprintf("select video_id, count(id) as analysis_count from analyses group by video_id")
	// subQuery := "select video_id, count(id) as analysis_count from analyses group by video_id"
	// query := ctxDB.Model(&Video{}).Joins("left join (?) as v on videos.id = v.video_id", subQuery)
	// TODO: もうちょっとかっこよく書きたい
	query := ctxDB.Model(&Video{}).Joins("left join (select video_id, count(id) as analysis_count from analyses group by video_id) as v on videos.id = v.video_id")
	if filter.UserId != nil && *filter.UserId != "" {
		query.Where("user_id = ?", filter.UserId)
	}
	query.Order("COALESCE(analysis_count, 0) desc")
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

func VideoFindAllRecommended(ctx *gin.Context, filter VideoFilter) ([]Video, error) {
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

func VideoFindOne(ctx *gin.Context, videoId int) (*Video, error) {
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

func VideoInsert(ctx *gin.Context, userId string, title string) (*Video, error) {
	v := &Video{
		Title:     title,
		UserId:    userId,
		UpdatedAt: time.Now(),
		CreatedAt: time.Now(),
	}

	ctxDB := DbConnection.WithContext(ctx.Request.Context())
	if err := ctxDB.Create(&v).Error; err != nil {
		return nil, err
	}

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		return v, nil
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

func VideoUpload(ctx *gin.Context, userId string, title string, file multipart.File, fileHeader multipart.FileHeader) (*Video, error) {
	ctxDB := DbConnection.WithContext(ctx.Request.Context())
	tx := ctxDB.Begin()

	video := Video{
		Title:     title,
		UserId:    userId,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if err := tx.Create(&video).Error; err != nil {
		// Rollback
		tx.Rollback()
		return nil, err
	}

	// Make working directory
	dirPath, err := util.MediaUtil.MakeWorkingDirectory(video.Id)
	if err != nil {
		// Rollback
		tx.Rollback()
		// Delete working directory
		util.MediaUtil.DeleteWorkingDirectory(*dirPath)
		return nil, err
	}
	// Save video
	videoFilePath, err := util.MediaUtil.SaveVideo(file, fileHeader, *dirPath)
	if err != nil {
		// Rollback
		tx.Rollback()
		// Delete working directory
		util.MediaUtil.DeleteWorkingDirectory(*dirPath)
		return nil, err
	}
	// Process video
	if err := util.MediaUtil.ProcessVideo(ctx, video.Id, *videoFilePath); err != nil {
		// Rollback
		tx.Rollback()
		// Delete working directory
		util.MediaUtil.DeleteWorkingDirectory(*dirPath)
		return nil, err
	}
	// Delete working directory
	util.MediaUtil.DeleteWorkingDirectory(*dirPath)

	select {
	case <-ctx.Done():
		tx.Rollback()
		return nil, ctx.Err()
	default:
		tx.Commit()
		return &video, nil
	}
}

type VideoFilter struct {
	SortType VideoSortType `json:"sortType,omitempty"`
	Limit    *int          `json:"limit,omitempty"`
	UserId   *string       `json:"userId,omitempty"`
}

func NewVideoFilter(sortType VideoSortType, limit *int, userId *string) VideoFilter {
	return VideoFilter{SortType: sortType, Limit: limit, UserId: userId}
}

type VideoSortType string

const (
	VideoSortTypePopular     VideoSortType = "popular"
	VideoSortTypeRecommended VideoSortType = "recommended"
)

func (v VideoSortType) Valid() error {
	switch v {
	case VideoSortTypePopular, VideoSortTypeRecommended:
		return nil
	default:
		return fmt.Errorf("failed: %w get %s", ErrInvalidType, v)
	}
}
