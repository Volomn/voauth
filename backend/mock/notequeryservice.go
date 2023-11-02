package mock

import (
	"context"

	"github.com/Volomn/voauth/backend/queryservice/note"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type MockNoteQueryService struct{ mock.Mock }

func (svc *MockNoteQueryService) FetchUserNotes(ctx context.Context) ([]note.NoteListItem, error) {
	args := svc.Called(ctx)
	return args.Get(0).([]note.NoteListItem), args.Error(1)
}

func (svc *MockNoteQueryService) GetUserNote(ctx context.Context, noteUUID uuid.UUID) (note.Note, error) {
	args := svc.Called(ctx, noteUUID)
	return args.Get(0).(note.Note), args.Error(1)
}
