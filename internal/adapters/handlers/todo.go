package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/ianyong/todo-backend/internal/api"
	"github.com/ianyong/todo-backend/internal/services"
)

func GetAllTodos(r *http.Request, s *services.Services) (*api.Response, error) {
	todos, err := s.TodoService.GetAllTodos()
	if err != nil {
		return nil, err
	}

	data, err := json.Marshal(todos)
	if err != nil {
		return nil, err
	}

	return &api.Response{
		Payload: data,
		Code:    http.StatusOK,
	}, nil
}
