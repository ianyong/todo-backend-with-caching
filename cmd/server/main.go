package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintf(w, "Hello World! The time now is %s.", time.Now().Format("2006-01-02 15:04:05"))
		if err != nil {
			log.Fatalln(err)
		}
	})
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
