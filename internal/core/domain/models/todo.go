package models

import "time"

// Todo represents a task that needs to be done.
type Todo struct {
	Name        string
	Description string
	DueDate     time.Time
	IsCompleted bool
}
