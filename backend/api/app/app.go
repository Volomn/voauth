package app

import (
	"context"

	"github.com/Volomn/voauth/backend/domain"
	"github.com/google/uuid"
)

type Application interface {
	GetAuthSecretKey(ctx context.Context) string
	SignupUser(ctx context.Context, firstName, lastName, email, password string) (domain.User, error)
	AuthenticateWithEmailAndPassword(ctx context.Context, email, password string) (domain.User, error)
	AddNote(ctx context.Context, title, content string) (domain.Note, error)
	UpdateNote(ctx context.Context, noteUUID uuid.UUID, title, content string) (domain.Note, error)
	DeleteNote(ctx context.Context, noteUUID uuid.UUID) error
}
