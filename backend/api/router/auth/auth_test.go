package auth_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Volomn/voauth/backend/api"
	"github.com/Volomn/voauth/backend/app"
	"github.com/Volomn/voauth/backend/domain"
	"github.com/Volomn/voauth/backend/mock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestEmailAndPasswordAuthenticationHandler(t *testing.T) {
	var mockApplication mock.MockApplication
	var mockNotequeryService mock.MockNoteQueryService

	mockApplication.On("GetAuthSecretKey", context.Background()).Return("secret")

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
		mockApplication.On("AuthenticateWithEmailAndPassword", context.Background(), test.Email, test.Password).Return(test.ExpectedUser, test.ExpectedError)
		var jsonData = []byte(fmt.Sprintf(`{
			"email": "%s",
			"password": "%s"
		}`, test.Email, test.Password))

		req, err := http.NewRequest("POST", "/auth/", bytes.NewBuffer(jsonData))
		assert.Equal(t, nil, err)

		req.Header.Set("Content-Type", "application/json; charset=UTF-8")

		// Create a new context.Context and populate it with data.
		ctx := context.Background()
		ctx = context.WithValue(ctx, "app", &mockApplication)

		rr := httptest.NewRecorder()
		handler := api.GetApiRouter(&mockApplication, &mockNotequeryService)
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
