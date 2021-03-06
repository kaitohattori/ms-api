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

type RateController struct {
}

func NewRateController() *RateController {
	return &RateController{}
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
// @Failure 400 {object} util.HTTPError
// @Failure 404 {object} util.HTTPError
// @Failure 500 {object} util.HTTPError
// @Router /videos/{id}/rate [get]
func (c *RateController) Get(ctx *gin.Context) {
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
	rate, err := model.RateFindOne(videoId, *userId)
	if err != nil {
		log.Println(err)
		util.NewError(ctx, http.StatusNotFound, err)
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
// @Param rate body model.Rate true "Rate"
// @Success 204 ""
// @Failure 400 {object} util.HTTPError
// @Failure 404 {object} util.HTTPError
// @Failure 500 {object} util.HTTPError
// @Router /videos/{id}/rate [patch]
func (c *RateController) Update(ctx *gin.Context) {
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
	rate := &model.Rate{
		UserId:    *userId,
		VideoId:   videoId,
		UpdatedAt: time.Now(),
	}
	if err := ctx.BindJSON(&rate); err != nil {
		log.Println(err)
		util.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	_, err = rate.Update()
	if err != nil {
		log.Println(err)
		util.NewError(ctx, http.StatusNotFound, err)
		return
	}
	ctx.JSON(http.StatusNoContent, nil)
}

// RateController Average docs
// @Summary Get rate average
// @Description get rate average by Video ID
// @Tags Rates
// @Accept json
// @Produce json
// @Param id path int true "Video ID"
// @Success 200 {object} util.HTTPValueResponse
// @Failure 400 {object} util.HTTPError
// @Failure 404 {object} util.HTTPError
// @Failure 500 {object} util.HTTPError
// @Router /videos/{id}/rate/average [get]
func (c *RateController) Average(ctx *gin.Context) {
	videoIdStr := ctx.Param("id")
	videoId, err := strconv.Atoi(videoIdStr)
	if err != nil {
		log.Println(err)
		util.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	value, err := model.RateAverage(videoId)
	if err != nil {
		log.Println(err)
		util.NewError(ctx, http.StatusNotFound, err)
		return
	}
	resp := util.HTTPValueResponse{
		Value: *value,
	}
	ctx.JSON(http.StatusOK, resp)
}
