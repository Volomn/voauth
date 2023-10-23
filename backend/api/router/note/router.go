package note

import (
	"net/http"

	"github.com/Volomn/voauth/backend/api/middleware"
	"github.com/go-chi/chi/v5"
)

func GetNoteRouter() chi.Router {
	var router = chi.NewRouter()
	router.Post("/", middleware.AuthenticationMiddleware(http.HandlerFunc(AddNoteHandler)).(http.HandlerFunc))
	return router
}
