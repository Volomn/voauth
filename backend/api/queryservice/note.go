package queryservice

import (
	"context"

	"github.com/Volomn/voauth/backend/queryservice/note"
	"github.com/google/uuid"
)

type NoteQueryService interface {
	FetchUserNotes(ctx context.Context) ([]note.NoteListItem, error)
	GetUserNote(ctx context.Context, noteUUID uuid.UUID) (note.Note, error)
}
