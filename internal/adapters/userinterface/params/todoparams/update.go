package todoparams

import (
	"gopkg.in/guregu/null.v4"

	"github.com/ianyong/todo-backend/internal/core/domainmodels"
)

type UpdateParams struct {
	ID          null.Int    `json:"id"`
	Name        null.String `json:"name"`
	Description null.String `json:"description"`
	DueDate     null.Time   `json:"dueDate"`
	IsCompleted null.Bool   `json:"isCompleted"`
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
