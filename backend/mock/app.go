package mock

import (
	"context"

	"github.com/Volomn/voauth/backend/domain"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type MockApplication struct{ mock.Mock }

func (app *MockApplication) GetAuthSecretKey(ctx context.Context) string {
	args := app.Called(ctx)
	return args.String(0)
}

func (app *MockApplication) SignupUser(ctx context.Context, firstName, lastName, email, password string) (domain.User, error) {
	args := app.Called(ctx, firstName, lastName, email, password)
	return args.Get(0).(domain.User), args.Error(1)
}

func (app *MockApplication) AuthenticateWithEmailAndPassword(ctx context.Context, email, password string) (domain.User, error) {
	args := app.Called(ctx, email, password)
	return args.Get(0).(domain.User), args.Error(1)
}

func (app *MockApplication) AddNote(ctx context.Context, title, content string) (domain.Note, error) {
	args := app.Called(ctx, title, content)
	return args.Get(0).(domain.Note), args.Error(1)
}

func (app *MockApplication) UpdateNote(ctx context.Context, noteUUID uuid.UUID, title, content string) (domain.Note, error) {
	args := app.Called(ctx, noteUUID, title, content)
	return args.Get(0).(domain.Note), args.Error(1)
}

func (app *MockApplication) DeleteNote(ctx context.Context, noteUUID uuid.UUID) error {
	args := app.Called(ctx, noteUUID)
	return args.Error(0)
}
