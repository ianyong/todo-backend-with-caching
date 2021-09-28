package domainmodels

import "time"

// Todo represents a task that needs to be done.
type Todo struct {
	ID          int64
	Name        string
	Description string
	DueDate     time.Time
	IsCompleted bool
}
