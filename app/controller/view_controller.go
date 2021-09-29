package controller

import (
	"ms-api/app/httputil"
	"ms-api/app/model"
	"ms-api/app/service"
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
// @Success 200 {object} httputil.HTTPValueResponse
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /videos/{id}/view/total [get]
func (c *ViewController) Total(ctx *gin.Context) {
	videoIdStr := ctx.Param("id")
	videoId, err := strconv.Atoi(videoIdStr)
	if err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	filter := model.NewViewFilter(&videoId, nil)
	total, err := c.service.Total(ctx, filter)
	if err != nil {
		httputil.NewError(ctx, http.StatusNotFound, err)
		return
	}
	resp := httputil.HTTPValueResponse{
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
// @Param id path int true "Video ID"
// @Success 200 {object} httputil.HTTPMessageResponse
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /videos/{id}/view [post]
func (c *ViewController) Add(ctx *gin.Context) {
	videoIdStr := ctx.Param("id")
	videoId, err := strconv.Atoi(videoIdStr)
	if err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	err = c.service.Add(ctx, videoId)
	if err != nil {
		httputil.NewError(ctx, http.StatusNotFound, err)
		return
	}
	resp := httputil.HTTPMessageResponse{
		Message: "success",
	}
	ctx.JSON(http.StatusOK, resp)
}
