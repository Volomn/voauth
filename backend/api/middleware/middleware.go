package middleware

import (
	"context"
	"log/slog"
	"net/http"
	"reflect"

	"github.com/Volomn/voauth/backend/api/app"
	"github.com/go-chi/jwtauth"
	"github.com/go-chi/render"
	"github.com/google/uuid"
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

func AuthenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, claims, err := jwtauth.FromContext(r.Context())
		if err != nil {
			slog.Error("Error authenticating token", "error", err.Error())
			render.Status(r, 401)
			render.JSON(w, r, map[string]string{"msg": "Unauthorized"})
			return
		}
		slog.Info("Token claims", "claims", claims)
		slog.Info("Type of authId is ", "type", reflect.TypeOf(claims["authUUID"]))
		authUUIDString, ok := claims["authUUID"].(string)
		if ok == false {
			slog.Info("authUUIDString is not a valid string")
			render.Status(r, 401)
			render.JSON(w, r, map[string]string{"msg": "Unauthorized"})
			return
		}
		authUUID, err := uuid.Parse(authUUIDString)
		if err != nil {
			slog.Info("authUUID is not valid", "error", err.Error())
			render.Status(r, 401)
			render.JSON(w, r, map[string]string{"msg": "Unauthorized"})
			return
		}

		ctx := context.WithValue(r.Context(), "authUserUUID", authUUID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
