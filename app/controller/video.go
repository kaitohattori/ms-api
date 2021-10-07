package controller

import (
	"fmt"
	"ms-api/app/model"
	"ms-api/app/service"
	"ms-api/app/util"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type VideoController struct {
	service *service.VideoService
}

func NewVideoController(service *service.VideoService) *VideoController {
	return &VideoController{service: service}
}

// VideoController Find docs
// @Summary Find videos
// @Description find videos
// @Tags Videos
// @Accept json
// @Produce json
// @Param userId query string false "User ID"
// @Param limit query int false "limit"
// @Param sortType query string true "sort type [popular, recommended]"
// @Success 200 {array} model.Video
// @Failure 400 {object} util.HTTPError
// @Failure 404 {object} util.HTTPError
// @Failure 500 {object} util.HTTPError
// @Router /videos [get]
func (c *VideoController) Find(ctx *gin.Context) {
	sortTypeStr := ctx.Query("sortType")
	userId := ctx.Query("userId")
	limitStr := ctx.Query("limit")
	// limit
	var limit *int = nil
	if limitStr != "" {
		value, err := strconv.Atoi(limitStr)
		if err != nil {
			util.NewError(ctx, http.StatusBadRequest, err)
			return
		} else {
			limit = &value
		}
	}
	// sortType
	sortType := model.VideoSortType(sortTypeStr)
	if err := sortType.Valid(); err != nil {
		fmt.Println(err)
		util.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	filter := model.NewVideoFilter(sortType, limit, &userId)
	videos, err := c.service.Find(ctx, filter)
	if err != nil {
		util.NewError(ctx, http.StatusNotFound, err)
		return
	}
	ctx.JSON(http.StatusOK, videos)

	select {
	case <-ctx.Done():
		fmt.Println(ctx.Err())
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
		util.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	video, err := c.service.Get(ctx, videoId)
	if err != nil {
		util.NewError(ctx, http.StatusNotFound, err)
		return
	}
	ctx.JSON(http.StatusOK, video)
}

// VideoController Add docs
// @Summary Add a video
// @Description add video
// @Tags Videos
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param video body model.AddVideo true "Add video"
// @Success 200 {object} model.Video
// @Failure 400 {object} util.HTTPError
// @Failure 404 {object} util.HTTPError
// @Failure 500 {object} util.HTTPError
// @Router /videos [post]
func (c *VideoController) Add(ctx *gin.Context) {
	userId := util.AuthUtil.GetUserId(util.AuthUtil{}, ctx)
	var addVideo model.AddVideo
	if err := ctx.ShouldBindJSON(&addVideo); err != nil {
		util.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	if err := addVideo.Valid(); err != nil {
		util.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	video, err := c.service.Add(ctx, userId, &addVideo)
	if err != nil {
		util.NewError(ctx, http.StatusNotFound, err)
		return
	}
	ctx.JSON(http.StatusOK, video)
}

// VideoController Update docs
// @Summary Update a video
// @Description update video
// @Tags Videos
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Video ID"
// @Param video body model.UpdateVideo true "Update video"
// @Success 200 {object} model.Video
// @Failure 400 {object} util.HTTPError
// @Failure 404 {object} util.HTTPError
// @Failure 500 {object} util.HTTPError
// @Router /videos/{id} [post]
func (c *VideoController) Update(ctx *gin.Context) {
	userId := util.AuthUtil.GetUserId(util.AuthUtil{}, ctx)
	videoIdStr := ctx.Param("id")
	videoId, err := strconv.Atoi(videoIdStr)
	if err != nil {
		util.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	var updateVideo model.UpdateVideo
	if err := ctx.ShouldBindJSON(&updateVideo); err != nil {
		util.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	if err := updateVideo.Valid(); err != nil {
		util.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	video, err := c.service.Update(ctx, userId, videoId, &updateVideo)
	if err != nil {
		util.NewError(ctx, http.StatusNotFound, err)
		return
	}
	ctx.JSON(http.StatusOK, video)
}

// VideoController Delete docs
// @Summary Delete a video
// @Description delete video
// @Tags Videos
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Video ID"
// @Success 200 {object} util.HTTPMessageResponse
// @Failure 400 {object} util.HTTPError
// @Failure 404 {object} util.HTTPError
// @Failure 500 {object} util.HTTPError
// @Router /videos/{id} [delete]
func (c *VideoController) Delete(ctx *gin.Context) {
	userId := util.AuthUtil.GetUserId(util.AuthUtil{}, ctx)
	videoIdStr := ctx.Param("id")
	videoId, err := strconv.Atoi(videoIdStr)
	if err != nil {
		util.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	if err = c.service.Remove(ctx, userId, videoId); err != nil {
		util.NewError(ctx, http.StatusNotFound, err)
		return
	}
	resp := util.HTTPMessageResponse{
		Message: "success",
	}
	ctx.JSON(http.StatusOK, resp)
}
