package model

import (
	"errors"
)

var (
	// ErrNetworkError when something goes wrong with the API call
	ErrNetworkError = errors.New("network error")

	// ErrInvalidParams when the query params are not valid
	ErrInvalidParams = errors.New("params are not valid")
)
