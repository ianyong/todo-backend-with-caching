package todohandlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"github.com/ianyong/todo-backend/internal/adapters/userinterface/api"
	"github.com/ianyong/todo-backend/internal/errors/externalerrors"
	"github.com/ianyong/todo-backend/internal/services"
)

func Destroy(r *http.Request, s *services.Services) (*api.Response, error) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		return nil, &externalerrors.InvalidURLError{
			Message: err.Error(),
		}
	}

	todo, err := s.TodoService.DeleteTodo(id)
	if err != nil {
		return nil, err
	}

	return &api.Response{
		Messages: api.StatusMessages{
			api.SuccessMessage(fmt.Sprintf("Successfully deleted '%s'!", todo.Name)),
		},
		Code: http.StatusOK,
	}, nil
}
