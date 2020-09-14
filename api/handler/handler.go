package handler

import (
	"log"
	"net/http"

	"github.com/sotnikov-s/glossarize-backend/api/errors"
)

// Handler is a wrapper over the http.Handler with possibility to return
// errors which are handled outside the handle function.
type Handler func(w http.ResponseWriter, r *http.Request) error

// ServeHTTP implements the http.Handler interface. It handles HTTP requests.
func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := h(w, r); err != nil {
		var status int
		var context string
		switch err := err.(type) {
		case *errors.Error:
			status = int(err.Kind())
			context = err.Context()
		default:
			status = http.StatusInternalServerError
			context = "Something went wrong!"
		}
		log.Printf("request to %s failed: %v", r.URL.String(), err)
		if err := WriteErrorResponse(w, status, context); err != nil {
			log.Printf("failed to write error response: %v", err)
		}
	}
}
