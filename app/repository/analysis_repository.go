package repository

import (
	"ms-api/app/model"

	"github.com/gin-gonic/gin"
)

type AnalysisRepository struct {
}

func NewAnalysisRepository() *AnalysisRepository {
	return &AnalysisRepository{}
}

func (r AnalysisRepository) Count(ctx *gin.Context, videoId int) (*int, error) {
	return model.Analysis.Count(model.Analysis{}, ctx, videoId)
}

func (r AnalysisRepository) Insert(ctx *gin.Context, analysis *model.Analysis) error {
	return analysis.Insert(ctx)
}
