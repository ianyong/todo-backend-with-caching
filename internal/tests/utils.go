package tests

import (
	"fmt"
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

// TruncateTables truncates the specified tables from the test database.
func (c *TestComponents) TruncateTables(tables ...string) error {
	for _, table := range tables {
		// Note: PostgreSQL does not support placeholder arguments for the TRUNCATE command.
		_, err := c.DB.Exec(fmt.Sprintf("TRUNCATE %s RESTART IDENTITY CASCADE", table))
		if err != nil {
			return fmt.Errorf("unable to truncate table '%s': %w", table, err)
		}
	}
	return nil
}
