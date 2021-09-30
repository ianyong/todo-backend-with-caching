package internalerrors

import "fmt"

type JSONEncodeError struct {
	Message string
}

func (e *JSONEncodeError) Error() string {
	return fmt.Sprintf("error encoding JSON: %s", e.Message)
}
