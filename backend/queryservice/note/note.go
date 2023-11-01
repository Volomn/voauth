package note

import (
	"context"
	"errors"
	"log/slog"

	"github.com/Volomn/voauth/backend/infra/repository"
	"github.com/Volomn/voauth/backend/queryservice"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type NoteQueryService struct {
	db *gorm.DB
}

func NewNoteQueryService(db *gorm.DB) *NoteQueryService {
	return &NoteQueryService{db: db}
}

func (svc *NoteQueryService) FetchUserNotes(ctx context.Context) ([]NoteListItem, error) {
	var notes []repository.DbNote
	var result []NoteListItem = []NoteListItem{}

	auth, ok := ctx.Value("auth").(queryservice.Auth)
	if ok == false {
		return []NoteListItem{}, &queryservice.AuthenticationError{Message: "Authentication not provided"}
	}

	slog.Info("Fetching user notes", "userUUID", auth.UserUUID.String())
	queryResult := svc.db.Where(&repository.DbNote{OwnerUUID: auth.UserUUID}).Find(&notes)
	if queryResult.Error != nil {
		slog.Error("Fetching user notes fail", "error", queryResult.Error.Error())
		return result, nil
	}
	slog.Info("Notes after query", "notes", notes, "result", result)
	for _, note := range notes {
		result = append(result, NoteListItem{
			UUID:       note.UUID,
			Title:      note.Title,
			Content:    note.Content,
			IsPublic:   note.IsPublic,
			IsFavorite: note.IsFavorite,
			IsArchived: note.IsArchived,
		})
	}
	slog.Info("Notes after loop", "notes", notes, "result", result)
	return result, nil
}

func (svc *NoteQueryService) GetUserNote(ctx context.Context, noteUUID uuid.UUID) (Note, error) {
	var note repository.DbNote
	var result Note

	auth, ok := ctx.Value("auth").(queryservice.Auth)
	if ok == false {
		return Note{}, &queryservice.AuthenticationError{Message: "Authentication not provided"}
	}

	slog.Info("Getting user note", "authUserUUID", auth.UserUUID.String(), "noteUUID", noteUUID.String())
	queryResult := svc.db.Where(&repository.DbNote{OwnerUUID: auth.UserUUID, UUID: noteUUID}).First(&note)
	if errors.Is(queryResult.Error, gorm.ErrRecordNotFound) {
		return result, &queryservice.EntityNotFoundError{Message: "Note not found"}
	}

	return Note{
		UUID:       note.UUID,
		Title:      note.Title,
		Content:    note.Content,
		IsPublic:   note.IsPublic,
		IsFavorite: note.IsFavorite,
		IsArchived: note.IsArchived,
	}, nil
}
