package handlers

import (
	"net/http"

	"github.com/ianyong/todo-backend/internal/adapters/userinterface/api"
	"github.com/ianyong/todo-backend/internal/services"
)

func NotFound(r *http.Request, s *services.Services) (*api.Response, error) {
	return &api.Response{
		Messages: api.StatusMessages{
			api.ErrorMessage("Route not found"),
		},
		Code: http.StatusNotFound,
	}, nil
}
