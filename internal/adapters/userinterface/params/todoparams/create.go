package todoparams

import (
	"gopkg.in/guregu/null.v4"

	"github.com/ianyong/todo-backend/internal/core/domainmodels"
	"github.com/ianyong/todo-backend/internal/errors/externalerrors"
)

type CreateParams struct {
	Name        null.String `json:"name"`
	Description null.String `json:"description"`
	DueDate     null.Time   `json:"dueDate"`
}

func (params *CreateParams) Validate() error {
	if params.Name.IsZero() {
		return &externalerrors.MissingParamError{Param: "name"}
	}
	if params.Description.IsZero() {
		return &externalerrors.MissingParamError{Param: "description"}
	}
	if params.DueDate.IsZero() {
		return &externalerrors.MissingParamError{Param: "dueDate"}
	}
	return nil
}

func (params *CreateParams) ToModel() *domainmodels.Todo {
	return &domainmodels.Todo{
		Name:        params.Name.ValueOrZero(),
		Description: params.Description.ValueOrZero(),
		DueDate:     params.DueDate.ValueOrZero(),
		IsCompleted: false,
	}
}
