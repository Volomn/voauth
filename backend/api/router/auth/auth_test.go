package auth

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Volomn/voauth/backend/app"
	"github.com/Volomn/voauth/backend/domain"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockApplication struct{ mock.Mock }

func (app *MockApplication) SignupUser(firstName, lastName, email, password string) (domain.User, error) {
	args := app.Called(firstName, lastName, email, password)
	return args.Get(0).(domain.User), args.Error(1)
}

func (app *MockApplication) AuthenticateWithEmailAndPassword(email, password string) (domain.User, error) {
	args := app.Called(email, password)
	return args.Get(0).(domain.User), args.Error(1)
}

func (app *MockApplication) GetAuthSecretKey() string {
	args := app.Called()
	return args.String(0)
}
func TestEmailAndPasswordAuthenticationHandler(t *testing.T) {
	var mockApplication MockApplication

	mockApplication.On("GetAuthSecretKey").Return("secret")

	newUserUUID, err := uuid.NewUUID()
	assert.Equal(t, nil, err)

	validEmail := "johndoe@test.com"
	invalidEmail := "invalid@test.com"
	validPassword := "password"
	invalidPassword := "invalidpassword"
	hashedValidPassword := "hashpasswordhash"

	newUser := &domain.User{
		Aggregate:      domain.Aggregate{UUID: newUserUUID},
		FirstName:      "John",
		LastName:       "Doe",
		Email:          validEmail,
		HashedPassword: hashedValidPassword,
		Address:        "",
		Bio:            "",
	}

	tests := []struct {
		Email         string
		Password      string
		ExpectedUser  domain.User
		ExpectedError error
	}{
		{Email: validEmail, Password: validPassword, ExpectedUser: *newUser, ExpectedError: nil},
		{Email: validEmail, Password: invalidPassword, ExpectedUser: domain.User{}, ExpectedError: app.InvalidLoginCredentialsError},
		{Email: invalidEmail, Password: validPassword, ExpectedUser: domain.User{}, ExpectedError: app.InvalidLoginCredentialsError},
	}

	for _, test := range tests {
		mockApplication.On("AuthenticateWithEmailAndPassword", test.Email, test.Password).Return(test.ExpectedUser, test.ExpectedError)
		var jsonData = []byte(fmt.Sprintf(`{
			"email": "%s",
			"password": "%s"
		}`, test.Email, test.Password))

		req, err := http.NewRequest("POST", "/api/auth/", bytes.NewBuffer(jsonData))
		assert.Equal(t, nil, err)

		req.Header.Set("Content-Type", "application/json; charset=UTF-8")

		rr := httptest.NewRecorder()

		handler := http.HandlerFunc(EmailAndPasswordAuthenticationHandler)

		// Create a new context.Context and populate it with data.
		ctx := context.Background()
		ctx = context.WithValue(ctx, "app", &mockApplication)
		handler.ServeHTTP(rr, req.WithContext(ctx))

		responseBody, err := io.ReadAll(rr.Body)
		assert.Equal(t, nil, err)

		if test.ExpectedError != nil {
			// Check the status code is what we expect.
			assert.Equal(t, http.StatusBadRequest, rr.Code)

			type Response struct {
				Msg string `json:"msg"`
			}
			var response Response

			err = json.Unmarshal(responseBody, &response)
			assert.Equal(t, nil, err)

			assert.Equal(t, app.InvalidLoginCredentialsError.Error(), response.Msg)
		} else {
			// Check the status code is what we expect.
			assert.Equal(t, http.StatusOK, rr.Code)

			type Response struct {
				AccessToken string `json:"accesstoken"`
			}
			var response Response

			err = json.Unmarshal(responseBody, &response)
			assert.Equal(t, nil, err)
		}
	}
}
