package todoparams

import (
	"time"

	"github.com/ianyong/todo-backend/internal/core/domainmodels"
)

type UpdateParams struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"dueDate"`
	IsCompleted bool      `json:"isCompleted"`
}

func (params *UpdateParams) ToModel() *domainmodels.Todo {
	return &domainmodels.Todo{
		ID:          params.ID,
		Name:        params.Name,
		Description: params.Description,
		DueDate:     params.DueDate,
		IsCompleted: params.IsCompleted,
	}
}
