package externalerrors

import (
	"fmt"
	"net/http"
)

type RecordNotFoundError struct {
	Model string
	ID    int64
}

func (e *RecordNotFoundError) Error() string {
	return fmt.Sprintf("%s with ID %d could not be found.", e.Model, e.ID)
}

func (e *RecordNotFoundError) StatusCode() int {
	return http.StatusNotFound
}
