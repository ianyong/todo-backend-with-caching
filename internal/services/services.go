package services

import (
	"context"

	"github.com/go-redis/cache/v8"
	"github.com/jmoiron/sqlx"

	"github.com/ianyong/todo-backend/internal/adapters/infrastructure/dbrepositories"
	"github.com/ianyong/todo-backend/internal/core/domainservices"
)

type Services struct {
	TodoService *domainservices.TodoService
	Cache       *cache.Cache
	CacheCtx    context.Context
}

func SetUp(db *sqlx.DB, cache *cache.Cache) *Services {
	todoRepo := dbrepositories.NewTodoDatabaseRepository(db)
	todoService := domainservices.NewTodoService(todoRepo)

	return &Services{
		TodoService: todoService,
		Cache:       cache,
		CacheCtx:    context.TODO(),
	}
}
