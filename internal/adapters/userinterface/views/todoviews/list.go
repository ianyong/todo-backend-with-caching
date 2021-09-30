package todoviews

import (
	"time"

	"github.com/ianyong/todo-backend/internal/core/domainmodels"
)

type ListView struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	DueDate     time.Time `json:"dueDate"`
	IsCompleted bool      `json:"isCompleted"`
}

func ListViewFrom(todo *domainmodels.Todo) ListView {
	return ListView{
		ID:          todo.ID,
		Name:        todo.Name,
		DueDate:     todo.DueDate,
		IsCompleted: todo.IsCompleted,
	}
}
