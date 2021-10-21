package services

import (
	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"

	"github.com/ianyong/todo-backend/internal/adapters/infrastructure/dbrepositories"
	"github.com/ianyong/todo-backend/internal/core/domainservices"
)

type Services struct {
	TodoService *domainservices.TodoService
	CacheDB     *redis.Client
}

func SetUp(db *sqlx.DB, cacheDB *redis.Client) *Services {
	todoRepo := dbrepositories.NewTodoDatabaseRepository(db)
	todoService := domainservices.NewTodoService(todoRepo)

	return &Services{
		TodoService: todoService,
		CacheDB:     cacheDB,
	}
}
