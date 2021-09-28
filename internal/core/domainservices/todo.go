package domainservices

import (
	"fmt"
	"time"

	"github.com/ianyong/todo-backend/internal/core/domainmodels"
	"github.com/ianyong/todo-backend/internal/ports/repositories"
)

type TodoService struct {
	repo repositories.Todo
}

func (s *TodoService) AddTodo(name string, description string, dueDate time.Time) error {
	todo := domainmodels.NewTodo(name, description, dueDate)
	err := s.repo.Add(todo)
	if err != nil {
		return fmt.Errorf("unable to add todo: %w", err)
	}
	return nil
}

func (s *TodoService) DeleteTodo(id int64) error {
	err := s.repo.Delete(id)
	if err != nil {
		return fmt.Errorf("unable to delete todo with id %d: %w", id, err)
	}
	return nil
}
