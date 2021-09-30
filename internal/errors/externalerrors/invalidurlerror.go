package externalerrors

import (
	"fmt"
	"net/http"
)

type InvalidURLError struct {
	Message string
}

func (e *InvalidURLError) Error() string {
	return fmt.Sprintf("Invalid parameter in URL: %s", e.Message)
}

func (e *InvalidURLError) StatusCode() int {
	return http.StatusBadRequest
}
