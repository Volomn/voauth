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

	// New routes for marking as favorite and unarchiving
	router.Put("/{noteUUID}/mark-as-favorite", middleware.AuthenticationMiddleware(http.HandlerFunc(MarkAsFavoriteHandler)).(http.HandlerFunc))
	router.Put("/{noteUUID}/unarchive", middleware.AuthenticationMiddleware(http.HandlerFunc(UnarchiveNoteHandler)).(http.HandlerFunc))
	return router
}
