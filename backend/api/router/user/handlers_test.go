package user

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

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

func TestSignupUserHandler(t *testing.T) {
	var mockApplication MockApplication

	newUserUUID, err := uuid.NewUUID()
	assert.Equal(t, nil, err)

	newUser := &domain.User{
		Aggregate:      domain.Aggregate{UUID: newUserUUID},
		FirstName:      "John",
		LastName:       "Doe",
		Email:          "johndoe@test.com",
		HashedPassword: "hashpasswordhash",
		Address:        "",
		Bio:            "",
	}

	mockApplication.On("SignupUser", newUser.FirstName, newUser.LastName, newUser.Email, "password").Return(*newUser, nil)

	var jsonData = []byte(fmt.Sprintf(`{
		"firstName": "%s",
		"lastName": "%s",
		"email": "%s",
		"password": "%s"
	}`, newUser.FirstName, newUser.LastName, newUser.Email, "password"))

	req, err := http.NewRequest("POST", "/api/users/", bytes.NewBuffer(jsonData))
	assert.Equal(t, nil, err)

	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(SignupUserHandler)

	// Create a new context.Context and populate it with data.
	ctx := context.Background()
	ctx = context.WithValue(ctx, "app", &mockApplication)
	handler.ServeHTTP(rr, req.WithContext(ctx))

	// Check the status code is what we expect.
	assert.Equal(t, http.StatusCreated, rr.Code)

	responseBody, err := io.ReadAll(rr.Body)
	assert.Equal(t, nil, err)

	type Response struct {
		Msg string `json:"msg"`
	}
	var response Response

	err = json.Unmarshal(responseBody, &response)
	assert.Equal(t, nil, err)

	assert.Equal(t, "User sign up successful", response.Msg)
}
