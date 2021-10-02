package main

import (
	"log"

	"github.com/ianyong/todo-backend/internal/adapters/infrastructure/database"
	"github.com/ianyong/todo-backend/internal/config"
)

func main() {
	cfg, err := config.LoadTest()
	if err != nil {
		log.Fatalf("failed to load config: %v\n", err)
	}

	// Connect to the database server without a database specified since we want to drop it.
	dbName := cfg.DBName
	cfg.DBName = ""
	db, err := database.SetUp(cfg)
	if err != nil {
		log.Fatalf("failed to connect to database: %v\n", err)
	}

	_, err = db.Exec("DROP DATABASE " + dbName)
	if err != nil {
		log.Fatalf("failed to drop database: %v\n", err)
	}

	log.Printf("Successfully dropped database '%s'", dbName)
}
