package user

import (
	"github.com/go-chi/chi/v5"
)

func GetUserRouter() chi.Router {
	var router = chi.NewRouter()
	router.Post("/", SignupUserHandler)
	return router
}
