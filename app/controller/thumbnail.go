package controller

import (
	"ms-api/app/service"
	"ms-api/app/util"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ThumbnailController struct {
	service *service.MediaService
}

func NewThumbnailController(service *service.MediaService) *ThumbnailController {
	return &ThumbnailController{service: service}
}

// ThumbnailController Upload docs
// @Summary Upload media
// @Description upload media
// @Tags Media
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
func (c *ThumbnailController) Upload(ctx *gin.Context) {
	userId := util.AuthUtil.GetUserId(util.AuthUtil{}, ctx)
	file, header, err := ctx.Request.FormFile("file")
	if err != nil {
		util.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	title := ctx.PostForm("title")
	video, err := c.service.Upload(ctx, userId, title, file, *header)
	if err != nil {
		util.NewError(ctx, http.StatusNotFound, err)
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
// @Failure 400 {object} util.HTTPError
// @Failure 404 {object} util.HTTPError
// @Failure 500 {object} util.HTTPError
// @Router /videos/{id}/thumbnail [get]
func (c *ThumbnailController) GetThumbnailImage(ctx *gin.Context) {
	videoIdStr := ctx.Param("id")
	videoId, err := strconv.Atoi(videoIdStr)
	if err != nil {
		util.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	thumbnailImage, err := c.service.GetThumbnailImage(ctx, videoId)
	if err != nil {
		util.NewError(ctx, http.StatusNotFound, err)
		return
	}
	ctx.Data(http.StatusOK, "image/jpeg", thumbnailImage.Bytes())
}
