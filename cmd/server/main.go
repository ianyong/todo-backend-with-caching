package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ianyong/todo-backend/internal/adapters/infrastructure/database"
	"github.com/ianyong/todo-backend/internal/adapters/infrastructure/inmemorydatabase"
	"github.com/ianyong/todo-backend/internal/adapters/userinterface/router"
	"github.com/ianyong/todo-backend/internal/config"
	"github.com/ianyong/todo-backend/internal/services"
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

	cacheDB := inmemorydatabase.SetUp(cfg)

	s := services.SetUp(db, cacheDB)

	addr := fmt.Sprintf(":%d", cfg.ServerPort)
	r := router.SetUp(s, cfg)

	err = http.ListenAndServe(addr, r)
	if err != nil {
		log.Fatalln(err)
	}
}
