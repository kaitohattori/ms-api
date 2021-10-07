package service

import (
	"ms-api/app/model"
	"ms-api/app/repository"

	"github.com/gin-gonic/gin"
)

type AnalysisService struct {
	repository *repository.AnalysisRepository
}

func NewAnalysisService(repository *repository.AnalysisRepository) *AnalysisService {
	return &AnalysisService{repository: repository}
}

func (s AnalysisService) Total(ctx *gin.Context, videoId int) (*int, error) {
	return s.repository.Count(ctx, videoId)
}

func (s AnalysisService) Add(ctx *gin.Context, userId string, videoId int) (*model.Analysis, error) {
	analysis := &model.Analysis{
		VideoId: videoId,
		UserId:  userId,
	}
	if err := s.repository.Insert(ctx, analysis); err != nil {
		return nil, err
	}
	return analysis, nil
}
