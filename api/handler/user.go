package handler

import (
	"fmt"
	"net/http"

	"github.com/sotnikov-s/glossarize-backend/api/errors"
	"github.com/sotnikov-s/glossarize-backend/backend/repository"

	"github.com/go-chi/chi"
)

// UserHandler handles requests regarding users register and auth.
type UserHandler struct {
	userRepo *repository.UserRepository
}

// RegisterHandlers registers the UserHandler routes to the router.
func (h *UserHandler) RegisterHandlers(router chi.Router, authMiddleware func(http.Handler) http.Handler) {
	router.Group(func(r chi.Router) {
		r.Use(authMiddleware)
		r.Method(http.MethodPost, "/user/register", Handler(h.register))
		r.Method(http.MethodPost, "/user/auth", Handler(h.auth))
	})
}

// register handles users registration.
func (h *UserHandler) register(w http.ResponseWriter, r *http.Request) error {
	return errors.Internal.Wrap(fmt.Errorf("the register method is not yet implemented"), "request failed")
}

// auth handles users authentication.
func (h *UserHandler) auth(w http.ResponseWriter, r *http.Request) error {
	return errors.Internal.Wrap(fmt.Errorf("the auth method is not yet implemented"), "request failed")
}
