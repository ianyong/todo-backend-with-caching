package router

import (
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/ianyong/todo-backend/internal/adapters/handlers"
)

// SetUp sets up the middleware stack and routes for a chi.Router and returns it.
func SetUp() chi.Router {
	r := chi.NewRouter()
	setUpMiddleware(r)
	setUpRoutes(r)
	return r
}

// setUpMiddleware sets up the middleware stack for a chi.Router.
func setUpMiddleware(r chi.Router) {
	// Injects a request ID in the context of each request.
	r.Use(middleware.RequestID)
	// Sets a http.Request's RemoteAddr to that of either the X-Forwarded-For or X-Real-IP header.
	r.Use(middleware.RealIP)
	// Logs the start and end of each request.
	r.Use(middleware.Logger)
	// Recovers from panics and return a 500 Internal Service Error.
	r.Use(middleware.Recoverer)
	// Returns a 504 Gateway Timeout after 1 min.
	r.Use(middleware.Timeout(time.Minute))
}

// setUpRoutes sets up the routes for a chi.Router. All API routes are namespaced with '/api/v1'.
func setUpRoutes(r chi.Router) {
	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/", handlers.HelloWorld)
	})
}
