package dbrepositories

import (
	"github.com/jmoiron/sqlx"

	"github.com/ianyong/todo-backend/internal/core/domainmodels"
)

type TodoDatabaseRepository struct {
	db *sqlx.DB
}

func NewTodoDatabaseRepository(db *sqlx.DB) *TodoDatabaseRepository {
	return &TodoDatabaseRepository{
		db: db,
	}
}

func (r *TodoDatabaseRepository) GetAll() ([]domainmodels.Todo, error) {
	var todos []domainmodels.Todo
	err := r.db.Select(&todos, "SELECT * FROM todos")
	if err != nil {
		return nil, err
	}
	return todos, nil
}

func (r *TodoDatabaseRepository) Get(id int64) (*domainmodels.Todo, error) {
	var todo domainmodels.Todo
	err := r.db.Get(&todo, "SELECT * FROM todos WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	return &todo, nil
}

func (r *TodoDatabaseRepository) Add(todo *domainmodels.Todo) (*domainmodels.Todo, error) {
	var id int64
	err := r.db.Get(
		&id,
		"INSERT INTO todos (name, description, due_date, is_completed) VALUES ($1, $2, $3, $4) RETURNING id",
		todo.Name,
		todo.Description,
		todo.DueDate,
		todo.IsCompleted,
	)
	if err != nil {
		return nil, err
	}

	return r.Get(id)
}

func (r *TodoDatabaseRepository) Update(todo *domainmodels.Todo) (*domainmodels.Todo, error) {
	_, err := r.db.Exec(
		"UPDATE todos SET name = $1, description = $2, due_date = $3, is_completed = $4 WHERE id = $5",
		todo.Name,
		todo.Description,
		todo.DueDate,
		todo.IsCompleted,
		todo.ID,
	)
	if err != nil {
		return nil, err
	}

	return r.Get(todo.ID)
}

func (r *TodoDatabaseRepository) Delete(id int64) error {
	_, err := r.db.Exec("DELETE FROM todos WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
