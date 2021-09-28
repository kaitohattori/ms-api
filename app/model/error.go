package model

import "errors"

var (
	// ErrNoRow example
	ErrNoRow = errors.New("no rows in result set")

	ErrInvalidType = errors.New("invalid type")
)
