package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/ianyong/todo-backend/internal/core/domainservices"
)

func GetAllTodos(todoService *domainservices.TodoService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO: Perform error handling
		todos, _ := todoService.GetAllTodos()

		w.Header().Set("Content-Type", "application/json")
		// TODO: Perform error handling
		_ = json.NewEncoder(w).Encode(todos)
	}
}
