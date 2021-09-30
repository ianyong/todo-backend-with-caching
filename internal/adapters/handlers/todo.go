package handlers

import (
	"net/http"

	"github.com/ianyong/todo-backend/internal/api"
	"github.com/ianyong/todo-backend/internal/json"
	"github.com/ianyong/todo-backend/internal/services"
	"github.com/ianyong/todo-backend/internal/views/todoviews"
)

func GetAllTodos(r *http.Request, s *services.Services) (*api.Response, error) {
	todos, err := s.TodoService.GetAllTodos()
	if err != nil {
		return nil, err
	}

	todoListViews := make([]todoviews.ListView, len(todos))
	for i := range todos {
		todo := todos[i]
		todoListViews[i] = todoviews.ListViewFrom(&todo)
	}

	data, err := json.EncodeView(todoListViews)
	if err != nil {
		return nil, err
	}

	return &api.Response{
		Payload: data,
		Code:    http.StatusOK,
	}, nil
}
