package note

import (
	"net/http"

	"github.com/Volomn/voauth/backend/api/middleware"
	"github.com/go-chi/chi/v5"
)

func GetNoteRouter() chi.Router {
	var router = chi.NewRouter()
	router.Post("/", middleware.AuthenticationMiddleware(http.HandlerFunc(AddNoteHandler)).(http.HandlerFunc))
	router.Put("/{noteUUID}", middleware.AuthenticationMiddleware(http.HandlerFunc(UpdateNoteHandler)).(http.HandlerFunc))
	router.Delete("/{noteUUID}", middleware.AuthenticationMiddleware(http.HandlerFunc(DeleteNoteHandler)).(http.HandlerFunc))
	router.Get("/", middleware.AuthenticationMiddleware(http.HandlerFunc(FetchNotesHandler)).(http.HandlerFunc))
	router.Get("/{noteUUID}", middleware.AuthenticationMiddleware(http.HandlerFunc(GetNoteHandler)).(http.HandlerFunc))
	return router
}
