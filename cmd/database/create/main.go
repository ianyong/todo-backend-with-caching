package main

import (
	"log"

	"github.com/omeid/pgerror"

	"github.com/ianyong/todo-backend/internal/config"
	"github.com/ianyong/todo-backend/internal/database"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %v\n", err)
	}

	// Connect to the database server without a database specified since it has not been created.
	dbName := cfg.DBName
	cfg.DBName = ""
	db, err := database.SetUp(cfg)
	if err != nil {
		log.Fatalf("failed to connect to database: %v\n", err)
	}

	_, err = db.Exec("CREATE DATABASE " + dbName)
	if err != nil {
		if err := pgerror.DuplicateDatabase(err); err != nil {
			log.Printf("Database '%s' already exists\n", dbName)
			return
		}
		log.Fatalf("failed to create database: %v\n", err)
	}

	log.Printf("Successfully created database '%s'\n", dbName)
}
