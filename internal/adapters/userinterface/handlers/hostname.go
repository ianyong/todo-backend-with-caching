package handlers

import (
	"net/http"
	"os"

	"github.com/ianyong/todo-backend/internal/adapters/userinterface/api"
	"github.com/ianyong/todo-backend/internal/services"
)

func GetHostName(r *http.Request, s *services.Services) (*api.Response, error) {
	hostname, err := os.Hostname()
	if err != nil {
		return nil, err
	}

	return &api.Response{
		Messages: api.StatusMessages{
			api.InformationMessage(hostname),
		},
		Code: http.StatusOK,
	}, nil
}
