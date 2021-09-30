package externalerrors

import (
	"fmt"
	"net/http"
)

type JSONDecodeError struct {
	Message string
}

func (e *JSONDecodeError) Error() string {
	return fmt.Sprintf("Error decoding JSON: %s", e.Message)
}

func (e *JSONDecodeError) StatusCode() int {
	return http.StatusBadRequest
}
