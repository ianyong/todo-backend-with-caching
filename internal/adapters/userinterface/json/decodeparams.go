package json

import (
	"encoding/json"
	"io"

	"github.com/ianyong/todo-backend/internal/errors/externalerrors"
)

func DecodeParams(r io.Reader, view interface{}) error {
	decoder := json.NewDecoder(r)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(view)
	if err != nil {
		return &externalerrors.JSONDecodeError{
			Message: err.Error(),
		}
	}
	return nil
}
