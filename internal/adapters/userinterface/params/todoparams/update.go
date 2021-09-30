package todoparams

import (
	"gopkg.in/guregu/null.v4"

	"github.com/ianyong/todo-backend/internal/core/domainmodels"
	"github.com/ianyong/todo-backend/internal/errors/externalerrors"
)

type UpdateParams struct {
	ID          null.Int    `json:"id"`
	Name        null.String `json:"name"`
	Description null.String `json:"description"`
	DueDate     null.Time   `json:"dueDate"`
	IsCompleted null.Bool   `json:"isCompleted"`
}

func (params *UpdateParams) Validate(id int64) error {
	if params.ID.IsZero() {
		return &externalerrors.MissingParamError{Param: "id"}
	}
	if !params.ID.Equal(null.IntFrom(id)) {
		return &externalerrors.IDMismatchError{}
	}
	if params.Name.IsZero() {
		return &externalerrors.MissingParamError{Param: "name"}
	}
	if params.Description.IsZero() {
		return &externalerrors.MissingParamError{Param: "description"}
	}
	if params.DueDate.IsZero() {
		return &externalerrors.MissingParamError{Param: "dueDate"}
	}
	if params.IsCompleted.IsZero() {
		return &externalerrors.MissingParamError{Param: "isCompleted"}
	}
	return nil
}

func (params *UpdateParams) ToModel() *domainmodels.Todo {
	return &domainmodels.Todo{
		ID:          params.ID.ValueOrZero(),
		Name:        params.Name.ValueOrZero(),
		Description: params.Description.ValueOrZero(),
		DueDate:     params.DueDate.ValueOrZero(),
		IsCompleted: params.IsCompleted.ValueOrZero(),
	}
}
