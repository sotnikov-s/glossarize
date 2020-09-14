package security

import (
	"net/http"
	"strings"

	"github.com/sotnikov-s/glossarize-backend/api/errors"
	"github.com/sotnikov-s/glossarize-backend/api/handler"
)

// NewAuthMiddleware creates a new http middleware
// which is used for handling the user authorization.
func NewAuthMiddleware(identityBuilder *IdentityBuilder, tokenService TokenService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return handler.Handler(func(w http.ResponseWriter, r *http.Request) error {
			header := r.Header.Get("Authorization")
			if !strings.HasPrefix(header, "Bearer") {
				return errors.Unauthorized.New("The bearer token is required")
			}
			id := strings.TrimPrefix(header, "Bearer ")
			token, err := tokenService.GetByID(id)
			if err != nil {
				return errors.Unauthorized.Wrap(err, http.StatusText(http.StatusUnauthorized))
			}
			ident, err := identityBuilder.Build(token.UserID)
			if err != nil {
				return errors.Unauthorized.Wrap(err, http.StatusText(http.StatusUnauthorized))
			}
			next.ServeHTTP(w, r.WithContext(setIdentityInContext(r.Context(), ident)))
			return nil
		})
	}
}
