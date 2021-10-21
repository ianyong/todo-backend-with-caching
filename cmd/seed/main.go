package main

import (
	"fmt"
	"log"
	"time"

	"github.com/ianyong/todo-backend/internal/adapters/infrastructure/database"
	"github.com/ianyong/todo-backend/internal/adapters/infrastructure/inmemorydatabase"
	"github.com/ianyong/todo-backend/internal/config"
	"github.com/ianyong/todo-backend/internal/core/domainmodels"
	"github.com/ianyong/todo-backend/internal/services"
)

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

	// Very slow to run but quick & easy solution to seed the database.
	for i := 1; i <= 100000; i++ {
		_, err := s.TodoService.AddTodo(&domainmodels.Todo{
			Name:        fmt.Sprintf("Todo %d", i),
			Description: "Lorem ipsum",
			DueDate:     time.Date(2021, time.November, 12, 23, 59, 59, 0, time.Local),
			IsCompleted: false,
		})
		if err != nil {
			log.Fatalf("error when seeding database: %v\n", err)
		}
	}
}
