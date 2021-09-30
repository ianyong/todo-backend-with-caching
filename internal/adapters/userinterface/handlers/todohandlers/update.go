package todohandlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"github.com/ianyong/todo-backend/internal/adapters/userinterface/api"
	"github.com/ianyong/todo-backend/internal/adapters/userinterface/json"
	"github.com/ianyong/todo-backend/internal/adapters/userinterface/params/todoparams"
	"github.com/ianyong/todo-backend/internal/adapters/userinterface/views/todoviews"
	"github.com/ianyong/todo-backend/internal/errors/externalerrors"
	"github.com/ianyong/todo-backend/internal/services"
)

func Update(r *http.Request, s *services.Services) (*api.Response, error) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		return nil, &externalerrors.InvalidURLError{
			Message: err.Error(),
		}
	}

	var updateParams todoparams.UpdateParams
	err = json.DecodeParams(r.Body, &updateParams)
	if err != nil {
		return nil, fmt.Errorf("unable to decode request body into params: %w", err)
	}

	err = updateParams.Validate(id)
	if err != nil {
		return nil, fmt.Errorf("params failed validation: %w", err)
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
