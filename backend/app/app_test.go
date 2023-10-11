package app

import (
	"testing"

	"github.com/Volomn/voauth/backend/domain"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type MockDB struct{ mock.Mock }

type MockUUIDGenerator struct{ mock.Mock }

func (gen *MockUUIDGenerator) New() (uuid.UUID, error) {
	args := gen.Called()
	return args.Get(0).(uuid.UUID), args.Error(1)
}

type MockPasswordHasher struct{ mock.Mock }

func (hasher *MockPasswordHasher) HashPassword(password string) (string, error) {
	args := hasher.Called(password)
	return args.String(0), args.Error(1)
}

func (hasher *MockPasswordHasher) IsPasswordMatch(password, hashedPassword string) bool {
	args := hasher.Called(password, hashedPassword)
	return args.Bool(0)
}

type MockUserRepository struct{ mock.Mock }

func (repo *MockUserRepository) GetUserByEmail(db *gorm.DB, email string) *domain.User {
	args := repo.Called(db, email)
	returnValue := args.Get(0)
	if returnValue == nil {
		return nil
	}
	return returnValue.(*domain.User)
}

func (repo *MockUserRepository) Save(db *gorm.DB, user domain.User) error {
	args := repo.Called(db, user)
	return args.Error(0)
}

func TestSignupUser(t *testing.T) {
	//Instantiate mock infra
	var mockUserRepository MockUserRepository
	var mockPasswordHasher MockPasswordHasher
	var mockUUIDGenerator MockUUIDGenerator

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

	app := NewApplication(ApplicationConfig{}, &gorm.DB{}, &mockPasswordHasher, &mockUUIDGenerator, &mockUserRepository)
	user, err := app.SignupUser(testFirstName, testLastName, testEmail, testPassword)
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
	var mockUserRepository MockUserRepository
	var mockPasswordHasher MockPasswordHasher
	var mockUUIDGenerator MockUUIDGenerator

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

	app := NewApplication(ApplicationConfig{}, &gorm.DB{}, &mockPasswordHasher, &mockUUIDGenerator, &mockUserRepository)
	user, err := app.SignupUser(testFirstName, testLastName, testEmail, testPassword)
	assert.Equal(t, UserWithEmailAlreadyExistsError, err)
	assert.Equal(t, domain.User{}, user)
}

func TestSignupUserWithWeakPassword(t *testing.T) {
	//Instantiate mock infra
	var mockUserRepository MockUserRepository
	var mockPasswordHasher MockPasswordHasher
	var mockUUIDGenerator MockUUIDGenerator

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

	app := NewApplication(ApplicationConfig{}, &gorm.DB{}, &mockPasswordHasher, &mockUUIDGenerator, &mockUserRepository)
	user, err := app.SignupUser(testFirstName, testLastName, testEmail, testPassword)
	assert.Equal(t, WeakPasswordError, err)
	assert.Equal(t, domain.User{}, user)
}

func TestAuthenticateWithEmailAndPassword(t *testing.T) {
	//Instantiate mock infra
	var mockUserRepository MockUserRepository
	var mockPasswordHasher MockPasswordHasher
	var mockUUIDGenerator MockUUIDGenerator

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
		app := NewApplication(ApplicationConfig{}, &gorm.DB{}, &mockPasswordHasher, &mockUUIDGenerator, &mockUserRepository)
		authUser, err := app.AuthenticateWithEmailAndPassword(test.Email, test.Password)
		assert.Equal(t, test.ExpectedError, err)
		assert.Equal(t, test.ExpectedUser, authUser)
	}

}
