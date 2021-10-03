package router

import (
	"net/http"

	"github.com/go-chi/cors"
)

var productionOrigins = []string{}
var developmentOrigins = []string{"http://localhost:3000"}

func corsMiddleware(environment string) func(http.Handler) http.Handler {
	options := cors.Options{
		AllowedOrigins:   productionOrigins,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}

	if environment == "development" {
		options.AllowedOrigins = developmentOrigins
	} else if environment == "production" {
		options.AllowedOrigins = productionOrigins
	}

	return cors.Handler(options)
}
