package externalerrors

import (
	"net/http"
)

type IDMismatchError struct{}

func (e *IDMismatchError) Error() string {
	return "ID in request body does not match ID in URL"
}

func (e *IDMismatchError) StatusCode() int {
	return http.StatusBadRequest
}
