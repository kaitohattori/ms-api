package util

import "errors"

var (
	ErrNoRow = errors.New("no rows in result set")

	ErrInvalidType = errors.New("invalid type")

	ErrNameInvalid = errors.New("name is empty")

	ErrRecordNotFound = errors.New("not found")

	ErrAuthInvalidAudience = errors.New("invalid audience")

	ErrAuthInvalidIssuer = errors.New("invalid issuer")

	ErrAuthPemCertNotFound = errors.New("unable to find appropriate key")
)
