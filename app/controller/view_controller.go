package controller

import (
	"ms-api/app/service"
	"ms-api/app/util"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ViewController struct {
	service *service.ViewService
}

func NewViewController(service *service.ViewService) *ViewController {
	return &ViewController{service: service}
}

// ViewController Total docs
// @Summary Total View
// @Description total View
// @Tags Views
// @Accept json
// @Produce json
// @Param id path int true "Video ID"
// @Success 200 {object} util.HTTPValueResponse
// @Failure 400 {object} util.HTTPError
// @Failure 404 {object} util.HTTPError
// @Failure 500 {object} util.HTTPError
// @Router /videos/{id}/view/total [get]
func (c *ViewController) Total(ctx *gin.Context) {
	videoIdStr := ctx.Param("id")
	videoId, err := strconv.Atoi(videoIdStr)
	if err != nil {
		util.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	total, err := c.service.Total(ctx, videoId)
	if err != nil {
		util.NewError(ctx, http.StatusNotFound, err)
		return
	}
	resp := util.HTTPValueResponse{
		Value: *total,
	}
	ctx.JSON(http.StatusOK, resp)
}

// ViewController Add docs
// @Summary Add a View
// @Description add View
// @Tags Views
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Video ID"
// @Success 200 {object} model.View
// @Failure 400 {object} util.HTTPError
// @Failure 404 {object} util.HTTPError
// @Failure 500 {object} util.HTTPError
// @Router /videos/{id}/view [post]
func (c *ViewController) Add(ctx *gin.Context) {
	userId := util.AuthUtil.GetUserId(util.AuthUtil{}, ctx)
	videoIdStr := ctx.Param("id")
	videoId, err := strconv.Atoi(videoIdStr)
	if err != nil {
		util.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	view, err := c.service.Add(ctx, userId, videoId)
	if err != nil {
		util.NewError(ctx, http.StatusNotFound, err)
		return
	}
	ctx.JSON(http.StatusOK, view)
}
