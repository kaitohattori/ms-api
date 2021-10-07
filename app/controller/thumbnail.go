package controller

import (
	"ms-api/app/model"
	"ms-api/app/util"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ThumbnailController struct {
}

func NewThumbnailController() *ThumbnailController {
	return &ThumbnailController{}
}

// ThumbnailController GetThumbnailImage docs
// @Summary Get a video thumbnail image
// @Description get video thumbnail image by Video ID
// @Tags Thumbnail
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
	thumbnailImage, err := model.ThumbnailImage.Get(model.ThumbnailImage{}, ctx, videoId)
	if err != nil {
		util.NewError(ctx, http.StatusNotFound, err)
		return
	}
	ctx.Data(http.StatusOK, "image/jpeg", thumbnailImage.Bytes())
}
