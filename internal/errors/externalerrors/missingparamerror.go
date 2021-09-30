package externalerrors

import (
	"fmt"
	"net/http"
)

type MissingParamError struct {
	Param string
}

func (e *MissingParamError) Error() string {
	return fmt.Sprintf("%s must be set to a value", e.Param)
}

func (e *MissingParamError) StatusCode() int {
	return http.StatusBadRequest
}
