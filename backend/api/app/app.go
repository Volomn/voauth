package app

import (
	"context"

	"github.com/Volomn/voauth/backend/domain"
)

type Application interface {
	GetAuthSecretKey(ctx context.Context) string
	SignupUser(ctx context.Context, firstName, lastName, email, password string) (domain.User, error)
	AuthenticateWithEmailAndPassword(ctx context.Context, email, password string) (domain.User, error)
	AddNote(ctx context.Context, title, content string) (domain.Note, error)
}
