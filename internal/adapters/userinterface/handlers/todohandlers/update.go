package todohandlers

import (
	"fmt"
	"net/http"

	"github.com/ianyong/todo-backend/internal/adapters/userinterface/api"
	"github.com/ianyong/todo-backend/internal/adapters/userinterface/json"
	"github.com/ianyong/todo-backend/internal/adapters/userinterface/params/todoparams"
	"github.com/ianyong/todo-backend/internal/adapters/userinterface/views/todoviews"
	"github.com/ianyong/todo-backend/internal/services"
)

func Update(r *http.Request, s *services.Services) (*api.Response, error) {
	var updateParams todoparams.UpdateParams
	err := json.DecodeParams(r.Body, &updateParams)
	if err != nil {
		return nil, fmt.Errorf("unable to decode request body into params: %w", err)
	}

	todo, err := s.TodoService.UpdateTodo(updateParams.ToModel())
	if err != nil {
		return nil, err
	}

	todoView := todoviews.ViewFrom(todo)

	data, err := json.EncodeView(todoView)
	if err != nil {
		return nil, err
	}

	return &api.Response{
		Payload: data,
		Code:    http.StatusOK,
	}, nil
}
