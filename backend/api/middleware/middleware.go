package middleware

import (
	"context"
	"net/http"

	"github.com/Volomn/voauth/backend/domain"
)

type Application interface {
	SignupUser(firstName, lastName, email, password string) (domain.User, error)
}

func ApplicationMiddleware(app Application) func(http.Handler) http.Handler {
	// Store applicaton instance in request context
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), "app", app)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}

}
