package json

import (
	"encoding/json"

	"github.com/ianyong/todo-backend/internal/errors/internalerrors"
)

// EncodeView converts a view into JSON.
func EncodeView(view interface{}) ([]byte, error) {
	data, err := json.Marshal(view)
	if err != nil {
		return nil, &internalerrors.JSONEncodeError{
			Message: err.Error(),
		}
	}
	return data, nil
}
