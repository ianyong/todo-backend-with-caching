package main

import (
	"github.com/ianyong/todo-backend/internal/router"
	"log"
	"net/http"
)

// main is the entry point for the server.
func main() {
	r := router.SetUp()

	err := http.ListenAndServe(":8000", r)
	if err != nil {
		log.Fatalln(err)
	}
}
