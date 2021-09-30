package routes

import (
	"github.com/go-chi/chi/v5"

	"github.com/ianyong/todo-backend/internal/adapters/handlers/todohandlers"
	"github.com/ianyong/todo-backend/internal/api"
	"github.com/ianyong/todo-backend/internal/services"
)

func GetTodoRoutes(s *services.Services) func(r chi.Router) {
	return func(r chi.Router) {
		r.Get("/", api.WrapHandler(s, todohandlers.List))
		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", api.WrapHandler(s, todohandlers.Read))
		})
	}
}
