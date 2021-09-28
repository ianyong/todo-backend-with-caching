package repositories

import "github.com/ianyong/todo-backend/internal/core/domainmodels"

type TodoRepository interface {
	Get(id int64) (domainmodels.Todo, error)
	Add(todo domainmodels.Todo) error
	Update(todo domainmodels.Todo) error
	Delete(id int64) error
}
