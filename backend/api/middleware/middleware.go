package middleware

import (
	"context"
	"net/http"

	"github.com/Volomn/voauth/backend/api/app"
)

func ApplicationMiddleware(app app.Application) func(http.Handler) http.Handler {
	// Store applicaton instance in request context
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), "app", app)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}

}

func QueryServiceMiddleWare(name string, queryService any) func(http.Handler) http.Handler {
	// Store query esrvice instance in request context
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), name, queryService)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
