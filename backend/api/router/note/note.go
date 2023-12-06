package note

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/Volomn/voauth/backend/api/app"
	"github.com/Volomn/voauth/backend/api/queryservice"
	a "github.com/Volomn/voauth/backend/app"
	q "github.com/Volomn/voauth/backend/queryservice"
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

func DeleteNoteHandler(w http.ResponseWriter, r *http.Request) {
	noteUUIDString := chi.URLParam(r, "noteUUID")
	slog.Info("Note uuid from request", "noteUUID", noteUUIDString)
	noteUUID, err := uuid.Parse(noteUUIDString)
	if err != nil {
		render.Status(r, 422)
		render.JSON(w, r, map[string]string{"msg": "Invalid note uuid"})
		return
	}

	application := r.Context().Value("app").(app.Application)
	authUserUUID := r.Context().Value("authUserUUID").(uuid.UUID)

	ctx := context.Background()
	ctx = context.WithValue(ctx, "auth", a.Auth{UserUUID: authUserUUID})

	err = application.DeleteNote(ctx, noteUUID)

	if err != nil {
		slog.Info("Unable to delete note", "error", err.Error())
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
	render.JSON(w, r, map[string]string{"msg": "Note deleted successfully"})
}

func FetchNotesHandler(w http.ResponseWriter, r *http.Request) {
	svc := r.Context().Value("noteQueryService").(queryservice.NoteQueryService)
	authUserUUID := r.Context().Value("authUserUUID").(uuid.UUID)

	ctx := context.Background()
	ctx = context.WithValue(ctx, "auth", q.Auth{UserUUID: authUserUUID})

	notes, err := svc.FetchUserNotes(ctx)

	if err != nil {
		slog.Info("Unable fetch notes", "error", err.Error())
		var authError *q.AuthenticationError
		var authorizationError *q.AuthorizationError
		var notFoundError *q.EntityNotFoundError

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
	if err := render.RenderList(w, r, NewNoteListResponse(notes)); err != nil {
		slog.Error("Unable to render response for fetch notes", "error", err.Error())
		panic(err.Error())
	}
}

func GetNoteHandler(w http.ResponseWriter, r *http.Request) {
	noteUUIDString := chi.URLParam(r, "noteUUID")
	slog.Info("Note uuid from request", "noteUUID", noteUUIDString)
	noteUUID, err := uuid.Parse(noteUUIDString)
	if err != nil {
		render.Status(r, 422)
		render.JSON(w, r, map[string]string{"msg": "Invalid note uuid"})
		return
	}

	svc := r.Context().Value("noteQueryService").(queryservice.NoteQueryService)
	authUserUUID := r.Context().Value("authUserUUID").(uuid.UUID)

	ctx := context.Background()
	ctx = context.WithValue(ctx, "auth", q.Auth{UserUUID: authUserUUID})

	note, err := svc.GetUserNote(ctx, noteUUID)

	if err != nil {
		slog.Info("Unable get note", "noteUUID", noteUUID.String(), "authUserUUID", authUserUUID.String(), "error", err.Error())
		var authError *q.AuthenticationError
		var authorizationError *q.AuthorizationError
		var notFoundError *q.EntityNotFoundError

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
	if err := render.Render(w, r, NewNoteResponse(note)); err != nil {
		slog.Error("Unable to render response for get note", "error", err.Error())
		panic(err.Error())
	}
}




func MarkNoteHandler(w http.ResponseWriter, r *http.Request) {
    // Extract the "noteUUID" parameter from the URL
    noteUUIDString := chi.URLParam(r, "noteUUID")
    slog.Info("Note uuid from request", "noteUUID", noteUUIDString)

    // Parse the noteUUIDString into a UUID (Universally Unique Identifier)
    noteUUID, err := uuid.Parse(noteUUIDString)
    if err != nil {
        // If parsing fails, respond with a 422 Unprocessable Entity status
        render.Status(r, 422)
        render.JSON(w, r, map[string]string{"msg": "Invalid note uuid"})
        return
    }

    // Retrieve the "app" and "authUserUUID" values from the request context
    application := r.Context().Value("app").(app.Application)
    authUserUUID := r.Context().Value("authUserUUID").(uuid.UUID)

    // Create a new context with authentication information
    ctx := context.Background()
    ctx = context.WithValue(ctx, "auth", a.Auth{UserUUID: authUserUUID})

    // Call the MarkNote method of the application with the noteUUID
    err = application.MarkNote(ctx, noteUUID)

    if err != nil {
        // Handle different types of errors that may occur during marking the note
        slog.Info("Unable to mark note", "error", err.Error())

        var authError *a.AuthenticationError
        var authorizationError *a.AuthorizationError
        var notFoundError *a.EntityNotFoundError

        if errors.As(err, &authError) {
            // If it's an authentication error, respond with a 401 Unauthorized status
            render.Status(r, 401)
            render.JSON(w, r, map[string]string{"msg": authError.Message})
            return
        } else if errors.As(err, &authorizationError) {
            // If it's an authorization error, respond with a 403 Forbidden status
            render.Status(r, 403)
            render.JSON(w, r, map[string]string{"msg": "Not allowed"})
            return
        } else if errors.As(err, &notFoundError) {
            // If it's an entity not found error, respond with a 404 Not Found status
            render.Status(r, 404)
            render.JSON(w, r, map[string]string{"msg": notFoundError.Message})
            return
        } else {
            // For other errors, respond with a 400 Bad Request status
            render.Status(r, 400)
            render.JSON(w, r, map[string]string{"msg": err.Error()})
            return
        }
    }

    // If there are no errors, respond with a 200 OK status
    render.Status(r, 200)
    render.JSON(w, r, map[string]string{"msg": "Note marked successfully"})
}






	      

func UnarchiveNoteHandler(w http.ResponseWriter, r *http.Request) {
    // Extract the "noteUUID" parameter from the URL
    noteUUIDString := chi.URLParam(r, "noteUUID")
    slog.Info("Note uuid from request", "noteUUID", noteUUIDString)

    // Parse the noteUUIDString into a UUID (Universally Unique Identifier)
    noteUUID, err := uuid.Parse(noteUUIDString)
    if err != nil {
        // If parsing fails, respond with a 422 Unprocessable Entity status
        render.Status(r, 422)
        render.JSON(w, r, map[string]string{"msg": "Invalid note uuid"})
        return
    }

    // Retrieve the "app" and "authUserUUID" values from the request context
    application := r.Context().Value("app").(app.Application)
    authUserUUID := r.Context().Value("authUserUUID").(uuid.UUID)

    // Create a new context with authentication information
    ctx := context.Background()
    ctx = context.WithValue(ctx, "auth", a.Auth{UserUUID: authUserUUID})

    // Call the UnarchiveNote method of the application with the noteUUID
    err = application.UnarchiveNote(ctx, noteUUID)

    // Handle different types of errors that may occur during unarchiving the note
    if errors.As(err, &authError) {
        // If it's an authentication error, respond with a 401 Unauthorized status
        render.Status(r, 401)
        render.JSON(w, r, map[string]string{"msg": authError.Message})
        return
    } else if errors.As(err, &authorizationError) {
        // If it's an authorization error, respond with a 403 Forbidden status
        render.Status(r, 403)
        render.JSON(w, r, map[string]string{"msg": "Not allowed"})
        return
    } else if errors.As(err, &notFoundError) {
        // If it's an entity not found error, respond with a 404 Not Found status
        render.Status(r, 404)
        render.JSON(w, r, map[string]string{"msg": notFoundError.Message})
        return
    } else {
        // For other errors, respond with a 400 Bad Request status
        render.Status(r, 400)
        render.JSON(w, r, map[string]string{"msg": err.Error()})
        return
    }

    // If there are no errors, respond with a 200 OK status
    render.Status(r, 200)
    render.JSON(w, r, map[string]string{"msg": "Note unarchived successfully"})
}



