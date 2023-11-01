package note_test

import (
	"context"
	"log/slog"
	"testing"

	"github.com/Volomn/voauth/backend/domain"
	"github.com/Volomn/voauth/backend/infra"
	"github.com/Volomn/voauth/backend/infra/repository"
	"github.com/Volomn/voauth/backend/queryservice"
	notequery "github.com/Volomn/voauth/backend/queryservice/note"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestFetchUserNotes(t *testing.T) {
	// initialize repositories
	var userRepository repository.UserRepository
	var noteRepository repository.NoteRepository

	// initialize db
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	assert.Equal(t, nil, err)

	// drop all tables
	infra.DropAllTables(db)

	// create all tables
	infra.AutoMigrateDB(db)

	// create user uuid
	userUUID, err := uuid.NewUUID()
	assert.Equal(t, nil, err)

	// create another user uuid
	anotherUserUUID, err := uuid.NewUUID()
	assert.Equal(t, nil, err)

	// create new user
	user, err := domain.NewUser(userUUID, "John", "Doe", "johndoe@test.com", "hashpassword", "somewhere", "", "")
	assert.Equal(t, nil, err)

	// save user in db
	err = userRepository.Save(db, *user)
	assert.Equal(t, nil, err)

	// create another new user
	anotherUser, err := domain.NewUser(anotherUserUUID, "Jane", "Doe", "janedoe@test.com", "hashpassword", "somewhere", "", "")
	assert.Equal(t, nil, err)

	// save another user in db
	err = userRepository.Save(db, *anotherUser)
	assert.Equal(t, nil, err)

	// create note1 uuid
	noteUUID1, err := uuid.NewUUID()
	assert.Equal(t, nil, err)

	// create first note
	note1, err := domain.NewNote(noteUUID1, userUUID, false, false, false, "Title one", "Lorem ipsum lorem ipsum")
	assert.Equal(t, nil, err)

	// create note2 uuid
	noteUUID2, err := uuid.NewUUID()
	assert.Equal(t, nil, err)

	//create second note
	note2, err := domain.NewNote(noteUUID2, userUUID, false, false, false, "Title two", "Lorem ipsum lorem ipsum ipsum")
	assert.Equal(t, nil, err)

	// save note1
	err = noteRepository.Save(db, *note1)
	assert.Equal(t, nil, err)

	// save note2
	err = noteRepository.Save(db, *note2)
	assert.Equal(t, nil, err)

	ctx := context.Background()

	// fetch user notes
	svc := notequery.NewNoteQueryService(db)

	tests := []struct {
		Ctx    context.Context
		Error  error
		Result []notequery.NoteListItem
	}{
		{
			Ctx:   context.WithValue(ctx, "auth", queryservice.Auth{UserUUID: userUUID}),
			Error: nil,
			Result: []notequery.NoteListItem{
				{
					UUID:       noteUUID1,
					IsPublic:   false,
					IsFavorite: false,
					IsArchived: false,
					Title:      "Title one",
					Content:    "Lorem ipsum lorem ipsum",
				},
				{
					UUID:       noteUUID2,
					IsPublic:   false,
					IsFavorite: false,
					IsArchived: false,
					Title:      "Title two",
					Content:    "Lorem ipsum lorem ipsum ipsum",
				},
			},
		},
		{
			Ctx:    ctx,
			Error:  &queryservice.AuthenticationError{Message: "Authentication not provided"},
			Result: []notequery.NoteListItem{},
		},
		{
			Ctx:    context.WithValue(ctx, "auth", queryservice.Auth{UserUUID: anotherUserUUID}),
			Error:  nil,
			Result: []notequery.NoteListItem{},
		},
	}
	for _, test := range tests {
		notes, err := svc.FetchUserNotes(test.Ctx)
		slog.Info("Form test", "notes", notes, "expected", test.Result)
		assert.Equal(t, test.Error, err)
		assert.Equal(t, test.Result, notes)
	}

}

func TestGetUserNote(t *testing.T) {
	// initialize repositories
	var userRepository repository.UserRepository
	var noteRepository repository.NoteRepository

	// initialize db
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	assert.Equal(t, nil, err)

	// drop all tables
	infra.DropAllTables(db)

	// create all tables
	infra.AutoMigrateDB(db)

	// create user uuid
	userUUID, err := uuid.NewUUID()
	assert.Equal(t, nil, err)

	// create new user
	user, err := domain.NewUser(userUUID, "John", "Doe", "johndoe@test.com", "hashpassword", "somewhere", "", "")
	assert.Equal(t, nil, err)

	// create another user uuid
	anotherUserUUID, err := uuid.NewUUID()
	assert.Equal(t, nil, err)

	// create another user
	anotherUser, err := domain.NewUser(anotherUserUUID, "Jane", "Doe", "janedoe@test.com", "hashpassword", "somewhere", "", "")
	assert.Equal(t, nil, err)

	// save user in db
	err = userRepository.Save(db, *user)
	assert.Equal(t, nil, err)

	// save anotherUser in db
	err = userRepository.Save(db, *anotherUser)
	assert.Equal(t, nil, err)

	// create note1 uuid
	noteUUID1, err := uuid.NewUUID()
	assert.Equal(t, nil, err)

	// create first note
	note1, err := domain.NewNote(noteUUID1, userUUID, false, false, false, "Title one", "Lorem ipsum lorem ipsum")
	assert.Equal(t, nil, err)

	// create note2 uuid
	noteUUID2, err := uuid.NewUUID()
	assert.Equal(t, nil, err)

	//create second note
	note2, err := domain.NewNote(noteUUID2, userUUID, false, false, false, "Title two", "Lorem ipsum lorem ipsum ipsum")
	assert.Equal(t, nil, err)

	// save note1
	err = noteRepository.Save(db, *note1)
	assert.Equal(t, nil, err)

	// save note2
	err = noteRepository.Save(db, *note2)
	assert.Equal(t, nil, err)

	notFoundNoteUUID, err := uuid.NewUUID()
	assert.Equal(t, nil, err)

	ctx := context.Background()

	// fetch user notes
	svc := notequery.NewNoteQueryService(db)

	tests := []struct {
		Ctx      context.Context
		NoteUUID uuid.UUID
		Error    error
		Result   notequery.Note
	}{
		{
			Ctx:      context.WithValue(ctx, "auth", queryservice.Auth{UserUUID: userUUID}),
			NoteUUID: noteUUID1,
			Error:    nil,
			Result: notequery.Note{
				UUID:       noteUUID1,
				IsPublic:   false,
				IsFavorite: false,
				IsArchived: false,
				Title:      "Title one",
				Content:    "Lorem ipsum lorem ipsum",
			},
		},
		{
			Ctx:      context.WithValue(ctx, "auth", queryservice.Auth{UserUUID: userUUID}),
			NoteUUID: noteUUID2,
			Error:    nil,
			Result: notequery.Note{
				UUID:       noteUUID2,
				IsPublic:   false,
				IsFavorite: false,
				IsArchived: false,
				Title:      "Title two",
				Content:    "Lorem ipsum lorem ipsum ipsum",
			},
		},
		{
			Ctx:      context.WithValue(ctx, "auth", queryservice.Auth{UserUUID: userUUID}),
			NoteUUID: notFoundNoteUUID,
			Error:    &queryservice.EntityNotFoundError{Message: "Note not found"},
			Result:   notequery.Note{},
		},
		{
			Ctx:      context.WithValue(ctx, "auth", queryservice.Auth{UserUUID: anotherUserUUID}),
			NoteUUID: noteUUID1,
			Error:    &queryservice.EntityNotFoundError{Message: "Note not found"},
			Result:   notequery.Note{},
		},
	}
	for _, test := range tests {
		note, err := svc.GetUserNote(test.Ctx, test.NoteUUID)
		assert.Equal(t, test.Error, err)
		assert.Equal(t, test.Result, note)
	}
}
