package inmemorydatabase

import (
	"github.com/go-redis/redis/v8"

	"github.com/ianyong/todo-backend/internal/config"
)

// SetUp sets up a redis.Client database connection and returns it.
func SetUp(cfg *config.Config) *redis.Client {
	cache := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	return cache
}
