package cache

import (
	"fmt"

	"github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"

	"github.com/ianyong/todo-backend/internal/config"
)

// SetUp sets up a cache.Cache connection and returns it.
func SetUp(cfg *config.Config) *cache.Cache {
	client := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", cfg.InMemoryDBHost, cfg.InMemoryDBPort),
	})
	return cache.New(&cache.Options{
		Redis: client,
	})
}
