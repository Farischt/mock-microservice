package types

import (
	"net/http"
	"time"
)

type ApiError struct {
	Status    int       `json:"status"`
	Err       string    `json:"error"`
	Timestamp time.Time `json:"timestamp"`
}

func NewApiError(s int, e string) ApiError {
	return ApiError{
		Status:    s,
		Err:       e,
		Timestamp: time.Now(),
	}
}

func (a ApiError) Error() string {
	return a.Err
}

type ApiResponse[T any] struct {
	Status    int       `json:"status"`
	Timestamp time.Time `json:"timestamp"`
	Method    string    `json:"method"`
	Path      string    `json:"path"`
	Data      T         `json:"data"`
}

func NewApiResponse[T any](s int, d T, r *http.Request) ApiResponse[T] {
	return ApiResponse[T]{
		Status:    s,
		Timestamp: time.Now(),
		Data:      d,
		Method:    r.Method,
		Path:      r.URL.Path,
	}
}
