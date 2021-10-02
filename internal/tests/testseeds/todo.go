package testseeds

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"

	"github.com/ianyong/todo-backend/internal/core/domainmodels"
)

var TodoSeeds = []domainmodels.Todo{
	{
		ID:          1,
		Name:        "CS3219 OTOT Task A",
		Description: "Docker & Kubernetes Task",
		DueDate:     time.Date(2021, time.November, 1, 0, 0, 0, 0, time.Local),
		IsCompleted: false,
	},
	{
		ID:          2,
		Name:        "CS3219 OTOT Task B",
		Description: "CRUD Application Task",
		DueDate:     time.Date(2021, time.November, 2, 0, 0, 0, 0, time.Local),
		IsCompleted: true,
	},
	{
		ID:          3,
		Name:        "CS3219 OTOT Task C",
		Description: "Authentication & Authorization Task",
		DueDate:     time.Date(2021, time.November, 3, 0, 0, 0, 0, time.Local),
		IsCompleted: false,
	},
	{
		ID:          4,
		Name:        "CS3219 OTOT Task D",
		Description: "Pub-Sub Messaging",
		DueDate:     time.Date(2021, time.November, 4, 0, 0, 0, 0, time.Local),
		IsCompleted: false,
	},
	{
		ID:          5,
		Name:        "CS3219 OTOT Task E",
		Description: "Blogpost Task",
		DueDate:     time.Date(2021, time.November, 5, 0, 0, 0, 0, time.Local),
		IsCompleted: false,
	},
	{
		ID:          6,
		Name:        "CS3219 OTOT Task F",
		Description: "Caching Task",
		DueDate:     time.Date(2021, time.November, 6, 0, 0, 0, 0, time.Local),
		IsCompleted: false,
	},
	{
		ID:          7,
		Name:        "CS3219 OTOT Task G",
		Description: "Module Content Enhancement Task",
		DueDate:     time.Date(2021, time.November, 7, 0, 0, 0, 0, time.Local),
		IsCompleted: false,
	},
}

func SeedTodos(db *sqlx.DB) error {
	for _, todo := range TodoSeeds {
		_, err := db.Exec(
			"INSERT INTO todos (id, name, description, due_date, is_completed) VALUES ($1, $2, $3, $4, $5)",
			todo.ID,
			todo.Name,
			todo.Description,
			todo.DueDate,
			todo.IsCompleted,
		)
		if err != nil {
			return fmt.Errorf("unable to insert seed data into database: %w", err)
		}
	}
	return nil
}
