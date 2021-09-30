package routes

import (
	"github.com/go-chi/chi/v5"

	"github.com/ianyong/todo-backend/internal/adapters/userinterface/api"
	todohandlers2 "github.com/ianyong/todo-backend/internal/adapters/userinterface/handlers/todohandlers"
	"github.com/ianyong/todo-backend/internal/services"
)

func GetTodoRoutes(s *services.Services) func(r chi.Router) {
	return func(r chi.Router) {
		r.Get("/", api.WrapHandler(s, todohandlers2.List))
		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", api.WrapHandler(s, todohandlers2.Read))
			r.Delete("/", api.WrapHandler(s, todohandlers2.Destroy))
		})
	}
}
