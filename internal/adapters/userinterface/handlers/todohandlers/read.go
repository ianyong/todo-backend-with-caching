package todohandlers

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"github.com/ianyong/todo-backend/internal/adapters/userinterface/api"
	"github.com/ianyong/todo-backend/internal/adapters/userinterface/json"
	"github.com/ianyong/todo-backend/internal/adapters/userinterface/views/todoviews"
	"github.com/ianyong/todo-backend/internal/errors/externalerrors"
	"github.com/ianyong/todo-backend/internal/services"
)

func Read(r *http.Request, s *services.Services) (*api.Response, error) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		return nil, &externalerrors.InvalidURLError{
			Message: err.Error(),
		}
	}

	todo, err := s.TodoService.GetTodo(id)
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
