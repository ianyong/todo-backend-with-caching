package tests

import (
	"log"

	"github.com/ianyong/todo-backend/internal/adapters/infrastructure/database"
	"github.com/ianyong/todo-backend/internal/config"
	"github.com/ianyong/todo-backend/internal/services"
)

// SetUp initialises the application and domain services in a test environment.
func SetUp() *services.Services {
	cfg, err := config.LoadTest()
	if err != nil {
		log.Fatalf("failed to load config: %v\n", err)
	}

	db, err := database.SetUp(cfg)
	if err != nil {
		log.Fatalf("failed to connect to database: %v\n", err)
	}

	s := services.SetUp(db)
	return s
}
