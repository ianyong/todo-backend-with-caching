package domainmodels

import "time"

// Todo represents a task that needs to be done.
type Todo struct {
	id          int64
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

func (t *Todo) GetID() int64 {
	return t.id
}

func (t *Todo) GetName() string {
	return t.name
}

func (t *Todo) GetDescription() string {
	return t.description
}

func (t *Todo) GetDueDate() time.Time {
	return t.dueDate
}

func (t *Todo) GetIsCompleted() bool {
	return t.isCompleted
}

func (t *Todo) SetID(id int64) {
	t.id = id
}

func (t *Todo) SetComplete() {
	t.isCompleted = true
}

func (t *Todo) SetIncomplete() {
	t.isCompleted = false
}
