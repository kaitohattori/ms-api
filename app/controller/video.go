package controller

import (
	"log"
	"ms-api/app/model"
	"ms-api/app/util"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type VideoController struct {
}

func NewVideoController() *VideoController {
	return &VideoController{}
}

// VideoController Find docs
// @Summary Find videos
// @Description find videos. login is required when sortType is recommended
// @Tags Videos
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param limit query int false "limit"
// @Param sortType query string true "sort type [popular, recommended]" default(popular)
// @Success 200 {array} model.Video
// @Failure 400 {object} util.HTTPError
// @Failure 404 {object} util.HTTPError
// @Failure 500 {object} util.HTTPError
// @Router /videos [get]
func (c *VideoController) Find(ctx *gin.Context) {
	sortTypeStr := ctx.Query("sortType")
	limitStr := ctx.Query("limit")
	// limit
	var limit *int = nil
	if limitStr != "" {
		value, err := strconv.Atoi(limitStr)
		if err != nil {
			log.Println(err)
			util.NewError(ctx, http.StatusBadRequest, err)
			return
		} else {
			limit = &value
		}
	}
	// sortType
	sortType := model.VideoSortType(sortTypeStr)
	if err := sortType.Valid(); err != nil {
		log.Println(err)
		log.Println("Set sortType to default value [popular]")
		sortType = model.VideoSortTypeDefault()
	}
	// filter
	filter := model.VideoFilter{}
	if sortType == model.VideoSortTypeRecommended {
		userId, _ := util.AuthUtilGetUserId(ctx)
		filter = model.VideoFilter{SortType: sortType, Limit: limit, UserId: userId}
	} else {
		filter = model.VideoFilter{SortType: sortType, Limit: limit, UserId: nil}
	}
	// check userId is set. login is required when sortType is recommended
	if err := filter.Valid(); err != nil {
		log.Println(err)
		util.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	videos, err := model.VideoFind(filter)
	if err != nil {
		log.Println(err)
		util.NewError(ctx, http.StatusNotFound, err)
		return
	}
	ctx.JSON(http.StatusOK, videos)

	select {
	case <-ctx.Done():
		log.Println(ctx.Err())
		return
	default:
		return
	}

}

// VideoController Get docs
// @Summary Get a video
// @Description get video by ID
// @Tags Videos
// @Accept json
// @Produce json
// @Param id path int true "Video ID"
// @Success 200 {object} model.Video
// @Failure 400 {object} util.HTTPError
// @Failure 404 {object} util.HTTPError
// @Failure 500 {object} util.HTTPError
// @Router /videos/{id} [get]
func (c *VideoController) Get(ctx *gin.Context) {
	videoIdStr := ctx.Param("id")
	videoId, err := strconv.Atoi(videoIdStr)
	if err != nil {
		log.Println(err)
		util.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	video, err := model.VideoFindOne(videoId)
	if err != nil {
		log.Println(err)
		util.NewError(ctx, http.StatusNotFound, err)
		return
	}
	ctx.JSON(http.StatusOK, video)
}

// VideoController Add docs
// @Summary Add a video
// @Description add video
// @Tags Videos
// @Accept mpfd
// @Produce json
// @Security ApiKeyAuth
// @Param title formData string true "Video Title"
// @Success 200 {object} model.Video
// @Failure 400 {object} util.HTTPError
// @Failure 404 {object} util.HTTPError
// @Failure 500 {object} util.HTTPError
// @Router /videos [post]
func (c *VideoController) Add(ctx *gin.Context) {
	userId, err := util.AuthUtilGetUserId(ctx)
	if err != nil {
		log.Println(err)
		util.NewError(ctx, http.StatusUnauthorized, err)
		return
	}
	title := ctx.PostForm("title")
	video, err := model.VideoInsert(*userId, title)
	if err != nil {
		log.Println(err)
		util.NewError(ctx, http.StatusNotFound, err)
		return
	}
	ctx.JSON(http.StatusOK, video)
}

// VideoController Update docs
// @Summary Update a video
// @Description update video
// @Tags Videos
// @Accept mpfd
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Video ID"
// @Param title formData string true "Video Title"
// @Success 200 {object} model.Video
// @Failure 400 {object} util.HTTPError
// @Failure 404 {object} util.HTTPError
// @Failure 500 {object} util.HTTPError
// @Router /videos/{id} [post]
func (c *VideoController) Update(ctx *gin.Context) {
	userId, err := util.AuthUtilGetUserId(ctx)
	if err != nil {
		log.Println(err)
		util.NewError(ctx, http.StatusUnauthorized, err)
		return
	}
	videoIdStr := ctx.Param("id")
	videoId, err := strconv.Atoi(videoIdStr)
	if err != nil {
		log.Println(err)
		util.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	title := ctx.PostForm("title")
	v := &model.Video{
		Id:        videoId,
		UserId:    *userId,
		Title:     title,
		UpdatedAt: time.Now(),
	}
	if err := v.Update(); err != nil {
		log.Println(err)
		util.NewError(ctx, http.StatusNotFound, err)
		return
	}
	ctx.JSON(http.StatusOK, v)
}

// VideoController Delete docs
// @Summary Delete a video
// @Description delete video
// @Tags Videos
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Video ID"
// @Success 204 ""
// @Failure 400 {object} util.HTTPError
// @Failure 404 {object} util.HTTPError
// @Failure 500 {object} util.HTTPError
// @Router /videos/{id} [delete]
func (c *VideoController) Delete(ctx *gin.Context) {
	videoIdStr := ctx.Param("id")
	videoId, err := strconv.Atoi(videoIdStr)
	if err != nil {
		log.Println(err)
		util.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	video, err := model.VideoFindOne(videoId)
	if err != nil {
		log.Println(err)
		util.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	if err := video.Delete(); err != nil {
		log.Println(err)
		util.NewError(ctx, http.StatusNotFound, err)
		return
	}
	ctx.JSON(http.StatusNoContent, nil)
}

// VideoController Upload docs
// @Summary Upload video
// @Description upload video
// @Tags Videos
// @Accept mpfd
// @Produce json
// @Security ApiKeyAuth
// @Param file formData file true "Video File"
// @Param title formData string true "Video Title"
// @Success 200 {object} model.Video
// @Failure 400 {object} util.HTTPError
// @Failure 404 {object} util.HTTPError
// @Failure 500 {object} util.HTTPError
// @Router /videos/upload [post]
func (c *VideoController) Upload(ctx *gin.Context) {
	userId, err := util.AuthUtilGetUserId(ctx)
	if err != nil {
		log.Println(err)
		util.NewError(ctx, http.StatusUnauthorized, err)
		return
	}
	file, header, err := ctx.Request.FormFile("file")
	if err != nil {
		log.Println(err)
		util.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	title := ctx.PostForm("title")
	video, err := model.VideoUpload(ctx, *userId, title, file, *header)
	if err != nil {
		log.Println(err)
		util.NewError(ctx, http.StatusNotFound, err)
		return
	}
	ctx.JSON(http.StatusOK, video)
}
