package repositories

import "github.com/ianyong/todo-backend/internal/core/domainmodels"

type Todo interface {
	GetAll() ([]domainmodels.Todo, error)
	Get(id int64) (*domainmodels.Todo, error)
	Add(todo *domainmodels.Todo) (*domainmodels.Todo, error)
	Update(todo *domainmodels.Todo) error
	Delete(id int64) error
}
