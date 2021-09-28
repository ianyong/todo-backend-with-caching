package domainservices

import (
	"fmt"

	"github.com/ianyong/todo-backend/internal/core/domainmodels"
	"github.com/ianyong/todo-backend/internal/ports/repositories"
)

type TodoService struct {
	todoRepo repositories.Todo
}

func NewTodoService(todoRepo repositories.Todo) *TodoService {
	return &TodoService{
		todoRepo: todoRepo,
	}
}

func (s *TodoService) GetAllTodos() ([]domainmodels.Todo, error) {
	todos, err := s.todoRepo.GetAll()
	if err != nil {
		return nil, fmt.Errorf("unable to get all todos: %w", err)
	}
	return todos, nil
}

func (s *TodoService) GetTodo(id int64) (*domainmodels.Todo, error) {
	todo, err := s.todoRepo.Get(id)
	if err != nil {
		return nil, fmt.Errorf("unable to get todo with id %d: %w", id, err)
	}
	return todo, nil
}

func (s *TodoService) AddTodo(todo *domainmodels.Todo) error {
	err := s.todoRepo.Add(todo)
	if err != nil {
		return fmt.Errorf("unable to add todo: %w", err)
	}
	return nil
}

func (s *TodoService) UpdateTodo(todo *domainmodels.Todo) error {
	err := s.todoRepo.Update(todo)
	if err != nil {
		return fmt.Errorf("unable to update todo with id %d: %w", todo.ID, err)
	}
	return nil
}

func (s *TodoService) DeleteTodo(id int64) error {
	err := s.todoRepo.Delete(id)
	if err != nil {
		return fmt.Errorf("unable to delete todo with id %d: %w", id, err)
	}
	return nil
}
