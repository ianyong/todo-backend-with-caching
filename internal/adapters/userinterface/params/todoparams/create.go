package todoparams

import (
	"gopkg.in/guregu/null.v4"

	"github.com/ianyong/todo-backend/internal/core/domainmodels"
)

type CreateParams struct {
	Name        null.String `json:"name"`
	Description null.String `json:"description"`
	DueDate     null.Time   `json:"dueDate"`
}

func (params *CreateParams) ToModel() *domainmodels.Todo {
	return &domainmodels.Todo{
		Name:        params.Name.ValueOrZero(),
		Description: params.Description.ValueOrZero(),
		DueDate:     params.DueDate.ValueOrZero(),
		IsCompleted: false,
	}
}
