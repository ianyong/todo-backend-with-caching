package todoviews

import (
	"time"

	"github.com/ianyong/todo-backend/internal/core/domainmodels"
)

type View struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"dueDate"`
	IsCompleted bool      `json:"isCompleted"`
}

func ViewFrom(todo *domainmodels.Todo) View {
	return View{
		ID:          todo.ID,
		Name:        todo.Name,
		Description: todo.Description,
		DueDate:     todo.DueDate,
		IsCompleted: todo.IsCompleted,
	}
}
