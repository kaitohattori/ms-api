package model

import (
	"time"
)

type View struct {
	Id        int       `json:"id"`
	UserId    string    `json:"userId,omitempty"`
	VideoId   int       `json:"videoId,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}

func (View) FindAll(filter *ViewFilter) ([]View, error) {
	view := View{}
	view.Id = 10
	view.UserId = "user_id"
	view.VideoId = 20
	view.CreatedAt = time.Now()
	view.UpdatedAt = time.Now()

	views := []View{view}
	return views, nil
}

func (View) FindOne(viewId int) (*View, error) {
	view := &View{}
	view.Id = 10
	view.UserId = "user_id"
	view.VideoId = 20
	view.CreatedAt = time.Now()
	view.UpdatedAt = time.Now()

	return view, nil
}

func (View) Count(filter ViewFilter) (*int, error) {
	var value int = 15
	return &value, nil
}

func (View) Average(videoId int) (int, error) {
	return 10, nil
}

func (View) Insert(videoId int) (int, error) {
	return 10, nil
}

func (v View) Update() (bool, error) {
	return true, nil
}

func (v View) Delete() (bool, error) {
	return true, nil
}
