package api

import (
	"encoding/json"
	"net/http"

	"github.com/ianyong/todo-backend/internal/services"
)

type Response struct {
	Payload json.RawMessage `json:"payload"`
	Code    int             `json:"-"`
}

type Handler = func(*http.Request, *services.Services) (*Response, error)

// WrapHandler converts the internal Handler type into a standard http.HandlerFunc.
func WrapHandler(s *services.Services, handler Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Set headers
		w.Header().Set("Content-Type", "application/json")

		res, err := handler(r, s)

		// Handle errors
		if err != nil {
			// TODO: Handle errors.
			return
		}

		// Handle response
		serveHTTPResponse(res, w)
	}
}

// serveHTTPResponse takes in a *Response and a http.ResponseWriter and writes
// the appropriate response to the response body.
func serveHTTPResponse(response *Response, w http.ResponseWriter) {
	if response == nil {
		response = &Response{}
	}

	if response.Code > 0 {
		w.WriteHeader(response.Code)
	}

	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		panic(err)
	}
}
