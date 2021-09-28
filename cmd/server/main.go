package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ianyong/todo-backend/internal/config"
	"github.com/ianyong/todo-backend/internal/database"
	"github.com/ianyong/todo-backend/internal/router"
)

// main is the entry point for the server.
func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %v\n", err)
	}

	db, err := database.SetUp(cfg)
	if err != nil {
		log.Fatalf("failed to connect to database: %v\n", err)
	}

	addr := fmt.Sprintf(":%d", cfg.ServerPort)
	r := router.SetUp(db)

	err = http.ListenAndServe(addr, r)
	if err != nil {
		log.Fatalln(err)
	}
}
