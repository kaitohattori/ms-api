package model

import (
	"fmt"
)

type VideoFilter struct {
	SortType VideoSortType `json:"sortType,omitempty"`
	Limit    *int          `json:"limit,omitempty"`
	UserId   *string       `json:"userId,omitempty"`
}

func NewVideoFilter(sortType VideoSortType, limit *int, userId *string) VideoFilter {
	return VideoFilter{SortType: sortType, Limit: limit, UserId: userId}
}

type VideoSortType string

const (
	Popular     VideoSortType = "popular"
	Recommended               = "recommended"
)

func (v VideoSortType) Valid() error {
	switch v {
	case Popular, Recommended:
		return nil
	default:
		return fmt.Errorf("failed: %w get %s", ErrInvalidType, v)
	}
}
