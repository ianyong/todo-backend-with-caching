package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ianyong/todo-backend/internal/config"
	"github.com/ianyong/todo-backend/internal/router"
)

// main is the entry point for the server.
func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %v\n", err)
	}

	addr := fmt.Sprintf(":%d", cfg.ServerPort)
	r := router.SetUp()

	err = http.ListenAndServe(addr, r)
	if err != nil {
		log.Fatalln(err)
	}
}
