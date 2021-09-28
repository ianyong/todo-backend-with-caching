package domainmodels

import "time"

// Todo represents a task that needs to be done.
type Todo struct {
	name        string
	description string
	dueDate     time.Time
	isCompleted bool
}

func NewTodo(name string, description string, dueDate time.Time) *Todo {
	return &Todo{
		name:        name,
		description: description,
		dueDate:     dueDate,
		isCompleted: false,
	}
}

func (t *Todo) SetComplete() {
	t.isCompleted = true
}

func (t *Todo) SetIncomplete() {
	t.isCompleted = false
}
