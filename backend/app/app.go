package app

import (
	"log/slog"
	"strings"

	"github.com/Volomn/voauth/backend/app/repository"
	"github.com/Volomn/voauth/backend/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PasswordHasher interface {
	HashPassword(password string) (string, error)
	IsPasswordMatch(password, hashedPassword string) bool
}

type UUIDGenerator interface {
	New() (uuid.UUID, error)
}

type Application struct {
	config         ApplicationConfig
	db             *gorm.DB
	userRepository repository.UserRepository
	noteRepository repository.NoteRepository
	passwordHasher PasswordHasher
	uuidGenerator  UUIDGenerator
}

func NewApplication(config ApplicationConfig, db *gorm.DB, passwordHasher PasswordHasher, uuidGenerator UUIDGenerator, userRepository repository.UserRepository, noteRepository repository.NoteRepository) *Application {
	return &Application{
		config:         config,
		db:             db,
		passwordHasher: passwordHasher,
		uuidGenerator:  uuidGenerator,
		userRepository: userRepository,
		noteRepository: noteRepository,
	}
}

func (app *Application) GetAuthSecretKey() string {
	return app.config.AuthSecretKey
}

func (app *Application) SignupUser(firstName, lastName, email, password string) (domain.User, error) {
	slog.Info("About to sign up new user", "firstName", firstName, "lastName", lastName, "email", email)
	existingUser := app.userRepository.GetUserByEmail(app.db, strings.ToLower(email))
	if existingUser != nil {
		slog.Info("User with email address already exists", "email", email, "user", existingUser)
		return domain.User{}, UserWithEmailAlreadyExistsError
	}
	newUserUUID, _ := app.uuidGenerator.New()
	if len(strings.TrimSpace(password)) < 8 {
		slog.Info("User sign up, password is weak", "email", email)
		return domain.User{}, WeakPasswordError
	}
	hashedPassword, err := app.passwordHasher.HashPassword(password)
	if err != nil {
		slog.Error("Unable to hash password: %w", err)
		return domain.User{}, SomethingWentWrongError
	}
	user, err := domain.NewUser(newUserUUID, firstName, lastName, email, hashedPassword, "", "", "")
	if err != nil {
		slog.Error("Unable to create new user", "error", err.Error())
		return domain.User{}, err
	}

	if err := app.userRepository.Save(app.db, *user); err != nil {
		slog.Error("Unable to save new user: %w", err)
		return domain.User{}, SomethingWentWrongError
	}
	return *user, nil
}

func (app *Application) AuthenticateWithEmailAndPassword(email, password string) (domain.User, error) {
	slog.Info("About to authenticate user", "email", email)
	user := app.userRepository.GetUserByEmail(app.db, strings.ToLower(email))
	if user == nil {
		slog.Info("User not found", "email", email)
		return domain.User{}, InvalidLoginCredentialsError
	}
	if app.passwordHasher.IsPasswordMatch(password, user.HashedPassword) == false {
		return domain.User{}, InvalidLoginCredentialsError
	}
	return *user, nil
}

func (app *Application) AddNote(userUUID uuid.UUID, title, content string) (domain.Note, error) {
	slog.Info("About to add note", "userUUID", userUUID.String(), "title", title, "content", content)
	user := app.userRepository.GetUserByUUID(app.db, userUUID)
	if user == nil {
		slog.Info("User not found", "userUUID", userUUID.String())
		return domain.Note{}, EntityNotFoundError
	}
	noteUUID, _ := app.uuidGenerator.New()
	note, err := domain.NewNote(noteUUID, user.UUID, false, false, false, title, content)
	if err != nil {
		slog.Info("Error creating note", "error", err.Error())
		return domain.Note{}, err
	}
	if err = app.noteRepository.Save(app.db, *note); err != nil {
		slog.Info("Error saving note to db", "error", err.Error())
		return domain.Note{}, SomethingWentWrongError
	}
	return *note, nil
}
