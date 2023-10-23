package note

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/Volomn/voauth/backend/api/app"
	a "github.com/Volomn/voauth/backend/app"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/google/uuid"
)

func AddNoteHandler(w http.ResponseWriter, r *http.Request) {
	data := &AddNoteRequestModel{}
	if err := render.Bind(r, data); err != nil {
		slog.Info("binding input data failed", "error", err.Error())
		render.Status(r, 422)
		render.JSON(w, r, map[string]string{"msg": fmt.Sprintf("Invalid request payload, %s", err.Error())})
		return
	}

	application := r.Context().Value("app").(app.Application)
	authUserUUID := r.Context().Value("authUserUUID").(uuid.UUID)

	ctx := context.Background()
	ctx = context.WithValue(ctx, "auth", a.Auth{UserUUID: authUserUUID})

	_, err := application.AddNote(ctx, data.Title, data.Content)

	if err != nil {
		slog.Info("Unable to add note", "error", err.Error())
		var authError *a.AuthenticationError
		if errors.As(err, &authError) {
			slog.Info("Error is authentication error")
			render.Status(r, 401)
			render.JSON(w, r, map[string]string{"msg": authError.Message})
			return
		} else {
			slog.Info("Error is not authentication error", "error", err.Error())
			render.Status(r, 400)
			render.JSON(w, r, map[string]string{"msg": err.Error()})
			return
		}
	}

	render.Status(r, 201)
	render.JSON(w, r, map[string]string{"msg": "Note added successfully"})
}

func UpdateNoteHandler(w http.ResponseWriter, r *http.Request) {
	noteUUIDString := chi.URLParam(r, "noteUUID")
	slog.Info("Note uuid from request", "noteUUID", noteUUIDString)
	noteUUID, err := uuid.Parse(noteUUIDString)
	if err != nil {
		render.Status(r, 422)
		render.JSON(w, r, map[string]string{"msg": "Invalid note uuid"})
		return
	}

	data := &UpdateNoteRequestModel{}
	if err := render.Bind(r, data); err != nil {
		slog.Info("binding input data failed", "error", err.Error())
		render.Status(r, 422)
		render.JSON(w, r, map[string]string{"msg": fmt.Sprintf("Invalid request payload, %s", err.Error())})
		return
	}

	application := r.Context().Value("app").(app.Application)
	authUserUUID := r.Context().Value("authUserUUID").(uuid.UUID)

	ctx := context.Background()
	ctx = context.WithValue(ctx, "auth", a.Auth{UserUUID: authUserUUID})

	_, err = application.UpdateNote(ctx, noteUUID, data.Title, data.Content)

	if err != nil {
		slog.Info("Unable to update note", "error", err.Error())
		var authError *a.AuthenticationError
		var authorizationError *a.AuthorizationError
		var notFoundError *a.EntityNotFoundError

		if errors.As(err, &authError) {
			render.Status(r, 401)
			render.JSON(w, r, map[string]string{"msg": authError.Message})
			return
		} else if errors.As(err, &authorizationError) {
			render.Status(r, 403)
			render.JSON(w, r, map[string]string{"msg": "Not allowed"})
			return
		} else if errors.As(err, &notFoundError) {
			render.Status(r, 404)
			render.JSON(w, r, map[string]string{"msg": notFoundError.Message})
			return

		} else {
			render.Status(r, 400)
			render.JSON(w, r, map[string]string{"msg": err.Error()})
			return
		}
	}

	render.Status(r, 200)
	render.JSON(w, r, map[string]string{"msg": "Note updated successfully"})
}
