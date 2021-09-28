package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"

	"github.com/ianyong/todo-backend/internal/adapters/dbrepositories"
	"github.com/ianyong/todo-backend/internal/adapters/handlers"
	"github.com/ianyong/todo-backend/internal/core/domainservices"
)

func GetTodoRoutes(db *sqlx.DB) func(r chi.Router) {
	todoRepo := dbrepositories.NewTodoDatabaseRepository(db)
	todoService := domainservices.NewTodoService(todoRepo)
	return func(r chi.Router) {
		r.Get("/", handlers.GetAllTodos(todoService))
	}
}
