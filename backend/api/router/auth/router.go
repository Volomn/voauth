package auth

import (
	"github.com/go-chi/chi/v5"
)

func GetAuthRouter() chi.Router {
	var router = chi.NewRouter()
	router.Post("/", EmailAndPasswordAuthenticationHandler)
	return router
}
