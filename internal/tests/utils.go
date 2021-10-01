package tests

import (
	"log"

	"github.com/jmoiron/sqlx"

	"github.com/ianyong/todo-backend/internal/adapters/infrastructure/database"
	"github.com/ianyong/todo-backend/internal/config"
	"github.com/ianyong/todo-backend/internal/services"
)

type TestComponents struct {
	DB       *sqlx.DB
	Services *services.Services
}

// SetUp initialises the database and services in a test environment.
func SetUp() *TestComponents {
	cfg, err := config.LoadTest()
	if err != nil {
		log.Fatalf("failed to load config: %v\n", err)
	}

	db, err := database.SetUp(cfg)
	if err != nil {
		log.Fatalf("failed to connect to database: %v\n", err)
	}

	s := services.SetUp(db)
	return &TestComponents{
		DB:       db,
		Services: s,
	}
}
