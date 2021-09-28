package repository

import "ms-api/app/model"

type ViewRepository struct {
}

func NewViewRepository() *ViewRepository {
	return &ViewRepository{}
}

func (r ViewRepository) Count(filter model.ViewFilter) (*int, error) {
	return model.View.Count(model.View{}, filter)
}

func (r ViewRepository) Insert(videoId int) (int, error) {
	return model.View.Insert(model.View{}, videoId)
}
