package app

import (
	"context"
	"testing"

	"github.com/Volomn/voauth/backend/domain"
	"github.com/Volomn/voauth/backend/mock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestSignupUser(t *testing.T) {
	//Instantiate mock infra
	var mockUserRepository mock.MockUserRepository
	var mockNoteRepository mock.MockNoteRepository
	var mockPasswordHasher mock.MockPasswordHasher
	var mockUUIDGenerator mock.MockUUIDGenerator

	testEmail := "johndoe@test.com"
	testFirstName := "John"
	testLastName := "Doe"
	testPassword := "password"
	hashedPassword := "hashpasswordhash"
	userUUID, err := uuid.NewUUID()

	// return hashedPassword when HashPassword method is called with testPassword
	mockPasswordHasher.On("HashPassword", testPassword).Return(hashedPassword, nil)

	// return userUUID and err when we call New
	mockUUIDGenerator.On("New").Return(userUUID, err)
	mockUserRepository.On("GetUserByEmail", &gorm.DB{}, testEmail).Return(nil)
	mockUserRepository.On("Save", &gorm.DB{}, domain.User{
		Aggregate:      domain.Aggregate{UUID: userUUID},
		FirstName:      testFirstName,
		LastName:       testLastName,
		HashedPassword: hashedPassword,
		Email:          testEmail,
		Bio:            "",
		Address:        "",
		PhotoURL:       "",
	}).Return(nil)

	app := NewApplication(ApplicationConfig{}, &gorm.DB{}, &mockPasswordHasher, &mockUUIDGenerator, &mockUserRepository, &mockNoteRepository)
	user, err := app.SignupUser(context.Background(), testFirstName, testLastName, testEmail, testPassword)
	assert.Equal(t, nil, err)
	assert.Equal(t, testEmail, user.Email)
	assert.Equal(t, testFirstName, user.FirstName)
	assert.Equal(t, testLastName, user.LastName)
	assert.Equal(t, hashedPassword, user.HashedPassword)
	assert.Equal(t, "", user.Bio)
	assert.Equal(t, "", user.Address)
	assert.Equal(t, "", user.PhotoURL)
}

func TestSignupUserWithAlreadyExistingEmail(t *testing.T) {
	//Instantiate mock infra
	var mockUserRepository mock.MockUserRepository
	var mockNoteRepository mock.MockNoteRepository
	var mockPasswordHasher mock.MockPasswordHasher
	var mockUUIDGenerator mock.MockUUIDGenerator

	testEmail := "johndoe@test.com"
	testFirstName := "John"
	testLastName := "Doe"
	testPassword := "password"
	hashedPassword := "hashpasswordhash"
	userUUID, err := uuid.NewUUID()

	// return hashedPassword when HashPassword method is called with testPassword
	mockPasswordHasher.On("HashPassword", testPassword).Return(hashedPassword, nil)

	// return userUUID and err when we call New
	mockUUIDGenerator.On("New").Return(userUUID, err)

	// return existing user when getting user with testEmail
	mockUserRepository.On("GetUserByEmail", &gorm.DB{}, testEmail).Return(&domain.User{
		Aggregate:      domain.Aggregate{UUID: userUUID},
		FirstName:      testFirstName,
		LastName:       testLastName,
		Email:          testEmail,
		HashedPassword: hashedPassword,
		Address:        "",
		Bio:            "",
		PhotoURL:       "",
	})

	app := NewApplication(ApplicationConfig{}, &gorm.DB{}, &mockPasswordHasher, &mockUUIDGenerator, &mockUserRepository, &mockNoteRepository)
	user, err := app.SignupUser(context.Background(), testFirstName, testLastName, testEmail, testPassword)
	assert.Equal(t, UserWithEmailAlreadyExistsError, err)
	assert.Equal(t, domain.User{}, user)
}

func TestSignupUserWithWeakPassword(t *testing.T) {
	//Instantiate mock infra
	var mockUserRepository mock.MockUserRepository
	var mockNoteRepository mock.MockNoteRepository
	var mockPasswordHasher mock.MockPasswordHasher
	var mockUUIDGenerator mock.MockUUIDGenerator

	testEmail := "johndoe@test.com"
	testFirstName := "John"
	testLastName := "Doe"
	testPassword := "  1234  "
	hashedPassword := "hashpasswordhash"
	userUUID, err := uuid.NewUUID()

	// return hashedPassword when HashPassword method is called with testPassword
	mockPasswordHasher.On("HashPassword", testPassword).Return(hashedPassword, nil)

	// return userUUID and err when we call New
	mockUUIDGenerator.On("New").Return(userUUID, err)

	// return existing user when getting user with testEmail
	mockUserRepository.On("GetUserByEmail", &gorm.DB{}, testEmail).Return(nil)

	app := NewApplication(ApplicationConfig{}, &gorm.DB{}, &mockPasswordHasher, &mockUUIDGenerator, &mockUserRepository, &mockNoteRepository)
	user, err := app.SignupUser(context.Background(), testFirstName, testLastName, testEmail, testPassword)
	assert.Equal(t, WeakPasswordError, err)
	assert.Equal(t, domain.User{}, user)
}

func TestAuthenticateWithEmailAndPassword(t *testing.T) {
	//Instantiate mock infra
	var mockUserRepository mock.MockUserRepository
	var mockNoteRepository mock.MockNoteRepository
	var mockPasswordHasher mock.MockPasswordHasher
	var mockUUIDGenerator mock.MockUUIDGenerator

	userUUID, err := uuid.NewUUID()
	assert.Equal(t, nil, err)

	validEmail := "johndoe@test.com"
	validPassword := "password"
	invalidEmail := "invalidemail@test.com"
	invalidPassword := "invalidPassword"

	user := &domain.User{
		Aggregate:      domain.Aggregate{UUID: userUUID},
		Email:          validEmail,
		HashedPassword: "hashpasswordhash",
		FirstName:      "John",
		LastName:       "Doe",
	}

	mockPasswordHasher.On("IsPasswordMatch", validPassword, user.HashedPassword).Return(true)
	mockPasswordHasher.On("IsPasswordMatch", invalidPassword, user.HashedPassword).Return(false)
	mockUserRepository.On("GetUserByEmail", &gorm.DB{}, validEmail).Return(user)
	mockUserRepository.On("GetUserByEmail", &gorm.DB{}, invalidEmail).Return(nil)

	tests := []struct {
		Email         string
		Password      string
		ExpectedUser  domain.User
		ExpectedError error
	}{
		{Email: validEmail, Password: validPassword, ExpectedUser: *user, ExpectedError: nil},
		{Email: validEmail, Password: "invalidPassword", ExpectedUser: domain.User{}, ExpectedError: InvalidLoginCredentialsError},
		{Email: invalidEmail, Password: validPassword, ExpectedUser: domain.User{}, ExpectedError: InvalidLoginCredentialsError},
	}

	for _, test := range tests {
		// return hashedPassword when HashPassword method is called with testPassword
		app := NewApplication(ApplicationConfig{}, &gorm.DB{}, &mockPasswordHasher, &mockUUIDGenerator, &mockUserRepository, &mockNoteRepository)
		authUser, err := app.AuthenticateWithEmailAndPassword(context.Background(), test.Email, test.Password)
		assert.Equal(t, test.ExpectedError, err)
		assert.Equal(t, test.ExpectedUser, authUser)
	}

}

func TestAddNote(t *testing.T) {
	//Instantiate mock infra
	var mockUserRepository mock.MockUserRepository
	var mockNoteRepository mock.MockNoteRepository
	var mockUUIDGenerator mock.MockUUIDGenerator
	var mockPasswordHasher mock.MockPasswordHasher

	userUUID, err := uuid.NewUUID()
	assert.Equal(t, nil, err)

	notFoundUserUUID, err := uuid.NewUUID()
	assert.Equal(t, nil, err)

	auth := Auth{UserUUID: userUUID}
	ctx := context.Background()
	contextWithValidAuth := context.WithValue(ctx, "auth", auth)
	contextWithNoAuth := ctx
	contextWithInvalidAuth := context.WithValue(ctx, "auth", Auth{UserUUID: notFoundUserUUID})

	noteUUID, err := uuid.NewUUID()
	assert.Equal(t, nil, err)

	mockUUIDGenerator.On("New").Return(noteUUID, err)
	mockUserRepository.On("GetUserByUUID", &gorm.DB{}, userUUID).Return(&domain.User{
		Aggregate:      domain.Aggregate{UUID: userUUID},
		FirstName:      "John",
		LastName:       "Doe",
		Email:          "johndoe@test.com",
		HashedPassword: "hashpasswordhash",
	})
	mockUserRepository.On("GetUserByUUID", &gorm.DB{}, notFoundUserUUID).Return(nil)

	tests := []struct {
		Ctx           context.Context
		Title         string
		Content       string
		ExpectedNote  domain.Note
		ExpectedError error
	}{
		{Ctx: contextWithValidAuth, Title: "Title one", Content: "Content one", ExpectedNote: domain.Note{Aggregate: domain.Aggregate{UUID: noteUUID}, OwnerUUID: userUUID, Title: "Title one", Content: "Content one"}, ExpectedError: nil},
		{Ctx: contextWithValidAuth, Title: "", Content: "Content two", ExpectedNote: domain.Note{}, ExpectedError: domain.EmptyNoteTitleError},
		{Ctx: contextWithValidAuth, Title: "Title three", Content: "", ExpectedNote: domain.Note{}, ExpectedError: domain.EmptyNoteContentError},
		{Ctx: contextWithInvalidAuth, Title: "Title four", Content: "Content four", ExpectedNote: domain.Note{}, ExpectedError: &AuthenticationError{Message: "User not found"}},
		{Ctx: contextWithNoAuth, Title: "Title five", Content: "Content five", ExpectedNote: domain.Note{}, ExpectedError: &AuthenticationError{Message: "Authentication not provided"}},
	}

	for _, test := range tests {
		// return hashedPassword when HashPassword method is called with testPassword
		mockNoteRepository.On("Save", &gorm.DB{}, domain.Note{
			Aggregate: domain.Aggregate{UUID: noteUUID},
			OwnerUUID: userUUID,
			Title:     test.Title,
			Content:   test.Content,
		}).Return(nil)
		app := NewApplication(ApplicationConfig{}, &gorm.DB{}, &mockPasswordHasher, &mockUUIDGenerator, &mockUserRepository, &mockNoteRepository)
		note, err := app.AddNote(test.Ctx, test.Title, test.Content)
		assert.Equal(t, test.ExpectedError, err)
		assert.Equal(t, test.ExpectedNote, note)
	}
}
