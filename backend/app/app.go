package app

import (
	"context"
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

func (app *Application) GetAuthSecretKey(ctx context.Context) string {
	return app.config.AuthSecretKey
}

func (app *Application) SignupUser(ctx context.Context, firstName, lastName, email, password string) (domain.User, error) {
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

func (app *Application) AuthenticateWithEmailAndPassword(ctx context.Context, email, password string) (domain.User, error) {
	slog.Info("About to authenticate user", "email", email)
	user := app.userRepository.GetUserByEmail(app.db, strings.ToLower(email))
	if user == nil {
		slog.Info("User not found", "email", email)
		return domain.User{}, InvalidLoginCredentialsError
	}
	if app.passwordHasher.IsPasswordMatch(password, user.HashedPassword) == false {
		slog.Info("User password does not match")
		return domain.User{}, InvalidLoginCredentialsError
	}
	return *user, nil
}

func (app *Application) AddNote(ctx context.Context, title, content string) (domain.Note, error) {
	auth, ok := ctx.Value("auth").(Auth)
	if ok == false {
		return domain.Note{}, &AuthenticationError{"Authentication not provided"}
	}
	slog.Info("About to add note", "authUserUUID", auth.UserUUID.String(), "title", title, "content", content)
	user := app.userRepository.GetUserByUUID(app.db, auth.UserUUID)
	if user == nil {
		slog.Info("User not found", "userUUID", auth.UserUUID.String())
		return domain.Note{}, &AuthenticationError{Message: "User not found"}
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

func (app *Application) UpdateNote(ctx context.Context, noteUUID uuid.UUID, title, content string) (domain.Note, error) {
	auth, ok := ctx.Value("auth").(Auth)
	if ok == false {
		return domain.Note{}, &AuthenticationError{Message: "Authentication not provided"}
	}
	slog.Info("About to add note", "authUserUUID", auth.UserUUID.String(), "title", title, "content", content)
	user := app.userRepository.GetUserByUUID(app.db, auth.UserUUID)
	if user == nil {
		slog.Info("User not found", "userUUID", auth.UserUUID.String())
		return domain.Note{}, &AuthenticationError{Message: "User not found"}
	}
	note := app.noteRepository.GetNoteByUUID(app.db, noteUUID)
	if note == nil {
		slog.Info("Note not found", "noteUUID", noteUUID.String())
		return domain.Note{}, &EntityNotFoundError{Message: "Note not found"}
	}
	if note.OwnerUUID != user.UUID {
		slog.Info("User trying to update note not belonging to it.", "authUserUUID", auth.UserUUID, "ownerUUID", note.OwnerUUID)
		return domain.Note{}, &AuthorizationError{Message: "User is not permitted to update this note"}
	}
	if err := note.SetTitle(title); err != nil {
		slog.Info("Unable to update note", "error", err.Error())
		return domain.Note{}, err
	}
	if err := note.SetContent(content); err != nil {
		slog.Info("Unable to update note", "error", err.Error())
		return domain.Note{}, err
	}
	slog.Info("Updated note is", "note", note)
	if err := app.noteRepository.Save(app.db, *note); err != nil {
		slog.Info("Error saving note to db", "error", err.Error())
		return domain.Note{}, SomethingWentWrongError
	}
	return *note, nil
}

func (app *Application) DeleteNote(ctx context.Context, noteUUID uuid.UUID) error {
	auth, ok := ctx.Value("auth").(Auth)
	if ok == false {
		return &AuthenticationError{"Authentication not provided"}
	}
	slog.Info("About to delete note", "authUserUUID", auth.UserUUID.String(), "noteUUID", noteUUID.String())
	user := app.userRepository.GetUserByUUID(app.db, auth.UserUUID)
	if user == nil {
		slog.Info("User not found", "userUUID", auth.UserUUID.String())
		return &AuthenticationError{Message: "User not found"}
	}
	note := app.noteRepository.GetNoteByUUID(app.db, noteUUID)
	if note == nil {
		slog.Info("Note not found", "noteUUID", noteUUID.String())
		return &EntityNotFoundError{Message: "Note not found"}
	}
	if note.OwnerUUID != user.UUID {
		slog.Info("User trying to update note not belonging to it.", "authUserUUID", auth.UserUUID, "ownerUUID", note.OwnerUUID)
		return &AuthorizationError{Message: "User is not permitted to delete this note"}
	}
	if err := app.noteRepository.Delete(app.db, *note); err != nil {
		slog.Info("Error deleting note in db", "error", err.Error())
		return SomethingWentWrongError
	}
	return nil
}
