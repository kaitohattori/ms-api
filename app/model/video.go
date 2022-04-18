package model

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"mime/multipart"
	"ms-api/app/util"
	"ms-api/config"
	"net/http"
	"strconv"
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

func VideoFind(filter VideoFilter) ([]Video, error) {
	if filter.SortType == VideoSortTypePopular {
		return VideoFindAllSortedByAnalysisCount(filter)
	} else if filter.SortType == VideoSortTypeRecommended {
		return VideoFindAllRecommended(filter)
	} else {
		return nil, errors.New("sort type is invalid")
	}
}

func VideoFindAllSortedByAnalysisCount(filter VideoFilter) ([]Video, error) {
	videos := []Video{}
	query := DbConnection.Model(&Video{}).Joins("left join (select video_id, count(id) as analysis_count from analyses group by video_id) as v on videos.id = v.video_id")
	if filter.UserId != nil && *filter.UserId != "" {
		query.Where("user_id = ?", *filter.UserId)
	}
	query.Order("COALESCE(analysis_count, 0) desc")
	if filter.Limit != nil && *filter.Limit != 0 {
		query.Limit(*filter.Limit)
	}
	query.Find(&videos)
	return videos, nil
}

func VideoFindAllRecommended(filter VideoFilter) ([]Video, error) {
	videos := []Video{}

	url := fmt.Sprintf("%s/videos/recommended/", config.Config.RecommendationAPIURL())
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println(err)
	}

	// Params
	params := request.URL.Query()
	if filter.UserId != nil && *filter.UserId != "" {
		params.Add("user_id", *filter.UserId)
	}
	if filter.Limit != nil {
		params.Add("limit", strconv.Itoa(*filter.Limit))
	}
	request.URL.RawQuery = params.Encode()

	// Request
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	byteArray, _ := ioutil.ReadAll(response.Body)
	err = json.Unmarshal(byteArray, &videos)
	if err != nil {
		return nil, err
	}
	return videos, nil
}

func VideoFindOne(videoId int) (*Video, error) {
	video := Video{}
	if err := DbConnection.First(&video, videoId).Error; err != nil {
		return nil, err
	}
	return &video, nil
}

func VideoInsert(userId string, title string) (*Video, error) {
	v := &Video{
		Title:     title,
		UserId:    userId,
		UpdatedAt: time.Now(),
		CreatedAt: time.Now(),
	}
	if err := DbConnection.Create(&v).Error; err != nil {
		return nil, err
	}
	return v, nil
}

func (v *Video) Update() error {
	if err := DbConnection.Save(&v).Error; err != nil {
		return err
	}
	return nil
}

func (v *Video) Delete() error {
	if err := DbConnection.Delete(&v).Error; err != nil {
		return err
	}
	return nil
}

func VideoUpload(ctx *gin.Context, userId string, title string, file multipart.File, fileHeader multipart.FileHeader) (*Video, error) {
	video := Video{
		Title:     title,
		UserId:    userId,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	tx := DbConnection.Begin()
	if err := tx.Create(&video).Error; err != nil {
		// Rollback
		tx.Rollback()
		return nil, err
	}

	// Make working directory
	dirPath, err := util.FileUtilMakeWorkingDirectory(video.Id)
	if err != nil {
		// Rollback
		tx.Rollback()
		// Delete working directory
		util.FileUtilDeleteWorkingDirectory(*dirPath)
		return nil, err
	}
	// Save video
	videoFilePath, err := util.FileUtilSaveVideo(file, fileHeader, *dirPath)
	if err != nil {
		// Rollback
		tx.Rollback()
		// Delete working directory
		util.FileUtilDeleteWorkingDirectory(*dirPath)
		return nil, err
	}
	// Process video
	if err := util.FileUtilProcessVideo(ctx, video.Id, *videoFilePath); err != nil {
		// Rollback
		tx.Rollback()
		// Delete working directory
		util.FileUtilDeleteWorkingDirectory(*dirPath)
		return nil, err
	}
	// Delete working directory
	util.FileUtilDeleteWorkingDirectory(*dirPath)

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

func (v VideoFilter) Valid() error {
	if v.SortType == VideoSortTypeRecommended {
		if v.UserId == nil {
			return errors.New("login is required when sortType is recommended")
		}
	}
	return nil
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
		return fmt.Errorf("failed: %w get %s", errors.New("invalid type"), v)
	}
}

func VideoSortTypeDefault() VideoSortType {
	return VideoSortTypePopular
}
