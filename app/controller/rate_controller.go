package controller

import (
	"ms-api/app/httputil"
	"ms-api/app/service"
	"ms-api/app/util"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RateController struct {
	service *service.RateService
}

func NewRateController(service *service.RateService) *RateController {
	return &RateController{service: service}
}

// RateController Get docs
// @Summary Get rate
// @Description get rate by Video ID
// @Tags Rates
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Video ID"
// @Success 200 {object} model.Rate
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /videos/{id}/rate [get]
func (c *RateController) Get(ctx *gin.Context) {
	userId := util.AuthUtil.GetUserId(util.AuthUtil{}, ctx)
	videoIdStr := ctx.Param("id")
	videoId, err := strconv.Atoi(videoIdStr)
	if err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	rate, err := c.service.Get(ctx, videoId, userId)
	if err != nil {
		httputil.NewError(ctx, http.StatusNotFound, err)
		return
	}
	ctx.JSON(http.StatusOK, rate)
}

// RateController Update docs
// @Summary Update rate
// @Description update Rate
// @Tags Rates
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Video ID"
// @Param value query float32 true "Rate value"
// @Success 200 {object} model.Rate
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /videos/{id}/rate [post]
func (c *RateController) Update(ctx *gin.Context) {
	userId := util.AuthUtil.GetUserId(util.AuthUtil{}, ctx)
	videoIdStr := ctx.Param("id")
	videoId, err := strconv.Atoi(videoIdStr)
	if err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	valueStr := ctx.Query("value")
	value, err := strconv.ParseFloat(valueStr, 32)
	if err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	rate, err := c.service.Update(ctx, userId, videoId, float32(value))
	if err != nil {
		httputil.NewError(ctx, http.StatusNotFound, err)
		return
	}
	ctx.JSON(http.StatusOK, rate)
}

// RateController Average docs
// @Summary Get rate average
// @Description get rate average by Video ID
// @Tags Rates
// @Accept json
// @Produce json
// @Param id path int true "Video ID"
// @Success 200 {object} httputil.HTTPValueResponse
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /videos/{id}/rate/average [get]
func (c *RateController) Average(ctx *gin.Context) {
	videoIdStr := ctx.Param("id")
	videoId, err := strconv.Atoi(videoIdStr)
	if err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	value, err := c.service.Average(ctx, videoId)
	if err != nil {
		httputil.NewError(ctx, http.StatusNotFound, err)
		return
	}
	resp := httputil.HTTPValueResponse{
		Value: *value,
	}
	ctx.JSON(http.StatusOK, resp)
}
