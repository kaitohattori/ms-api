package controller

import (
	"ms-api/app/model"
	"ms-api/app/util"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AnalysisController struct {
}

func NewAnalysisController() *AnalysisController {
	return &AnalysisController{}
}

// AnalysisController Total docs
// @Summary Total Analysis
// @Description total Analysis
// @Tags Analysis
// @Accept json
// @Produce json
// @Param id path int true "Video ID"
// @Success 200 {object} util.HTTPValueResponse
// @Failure 400 {object} util.HTTPError
// @Failure 404 {object} util.HTTPError
// @Failure 500 {object} util.HTTPError
// @Router /videos/{id}/analysis/total [get]
func (c *AnalysisController) Total(ctx *gin.Context) {
	videoIdStr := ctx.Param("id")
	videoId, err := strconv.Atoi(videoIdStr)
	if err != nil {
		util.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	total, err := model.AnalysisCount(videoId)
	if err != nil {
		util.NewError(ctx, http.StatusNotFound, err)
		return
	}
	resp := util.HTTPValueResponse{
		Value: *total,
	}
	ctx.JSON(http.StatusOK, resp)
}

// AnalysisController Add docs
// @Summary Add a Analysis
// @Description add Analysis
// @Tags Analysis
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Video ID"
// @Success 200 {object} model.Analysis
// @Failure 400 {object} util.HTTPError
// @Failure 404 {object} util.HTTPError
// @Failure 500 {object} util.HTTPError
// @Router /videos/{id}/analysis [post]
func (c *AnalysisController) Add(ctx *gin.Context) {
	userId := util.AuthUtilGetUserId(ctx)
	videoIdStr := ctx.Param("id")
	videoId, err := strconv.Atoi(videoIdStr)
	if err != nil {
		util.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	analysis, err := model.AnalysisInsert(userId, videoId)
	if err != nil {
		util.NewError(ctx, http.StatusNotFound, err)
		return
	}
	ctx.JSON(http.StatusOK, analysis)
}
