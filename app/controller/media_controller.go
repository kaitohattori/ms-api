package controller

import (
	"ms-api/app/httputil"
	"ms-api/app/model"
	"ms-api/app/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MediaController struct {
	service *service.MediaService
}

func NewMediaController(service *service.MediaService) *MediaController {
	return &MediaController{service: service}
}

// MediaController Upload docs
// @Summary Upload media
// @Description upload media
// @Tags Media
// @Accept mpfd
// @Produce json
// @Param file formData file true "Video File"
// @Param title formData string true "Video Title"
// @Success 200 {object} model.Video
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /videos/upload [post]
func (c *MediaController) Upload(ctx *gin.Context) {
	_, header, err := ctx.Request.FormFile("file")
	if err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	title := ctx.PostForm("title")
	userId := "user_id"
	Media := &model.Media{
		FileName: header.Filename,
		Title:    title,
		UserId:   userId,
	}
	video, err := c.service.Upload(ctx, Media)
	if err != nil {
		httputil.NewError(ctx, http.StatusNotFound, err)
		return
	}
	ctx.JSON(http.StatusOK, video)
}

// VideoController GetThumbnailImage docs
// @Summary Get a video thumbnail image
// @Description get video thumbnail image by Video ID
// @Tags Media
// @Accept json
// @Produce jpeg
// @Param id path int true "Video ID"
// @Success 200 {string} string ""
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /videos/{id}/thumbnail [get]
func (c *MediaController) GetThumbnailImage(ctx *gin.Context) {
	videoIdStr := ctx.Param("id")
	videoId, err := strconv.Atoi(videoIdStr)
	if err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	thumbnailImage, err := c.service.GetThumbnailImage(ctx, videoId)
	if err != nil {
		httputil.NewError(ctx, http.StatusNotFound, err)
		return
	}
	ctx.JSON(http.StatusOK, thumbnailImage)
}
