package api

import "errors"

// ExternalError is implemented by all errors that are meant to be returned in the API response.
type ExternalError interface {
	error
	StatusCode() int
}

// asExternalError finds the first error within the error chain that is an ExternalError. If an ExternalError is found,
// it returns the ExternalError and true. Otherwise, it returns nil and false.
func asExternalError(err error) (ExternalError, bool) {
	var e ExternalError
	ok := errors.As(err, &e)
	return e, ok
}
