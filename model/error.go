package model

import (
	"errors"
)

var (
	// ErrNetworkError when something goes wrong with the API call
	ErrNetworkError = errors.New("network error")

	// ErrInvalidParams when the query params are not valid
	ErrInvalidParams = errors.New("params are not valid")

	// ErrUpstreamHTML for when an endpoint returns an html response instead of json
	ErrUpstreamHTML = errors.New("upstream returned HTML response")
)
