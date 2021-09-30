package services

import (
	"github.com/jmoiron/sqlx"

	"github.com/ianyong/todo-backend/internal/adapters/infrastructure/dbrepositories"
	"github.com/ianyong/todo-backend/internal/core/domainservices"
)

type Services struct {
	TodoService *domainservices.TodoService
}

func SetUp(db *sqlx.DB) *Services {
	todoRepo := dbrepositories.NewTodoDatabaseRepository(db)
	todoService := domainservices.NewTodoService(todoRepo)

	return &Services{
		TodoService: todoService,
	}
}
