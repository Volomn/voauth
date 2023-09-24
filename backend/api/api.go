package api

import (
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
)

type HealthResponse struct {
	Status string `json:"status"`
}

func (rd HealthResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func GetApiRouter() chi.Router {
	// create api router
	router := chi.NewRouter()

	// configure middlewares
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.URLFormat)
	router.Use(render.SetContentType(render.ContentTypeJSON))
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
		Debug:          true,

		// AllowCredentials indicates whether the request can include user credentials like
		// cookies, HTTP authentication or client side SSL certificates.
		// AllowCredentials bool

	}))

	router.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		render.Status(r, http.StatusOK)
		render.Render(w, r, &HealthResponse{Status: "OK"})
	})
	return router
}
