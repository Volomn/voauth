package note_test

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
	"github.com/Volomn/voauth/backend/api/router/util"
	"github.com/Volomn/voauth/backend/app"
	"github.com/Volomn/voauth/backend/domain"
	"github.com/Volomn/voauth/backend/mock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestAddNoteHandler(t *testing.T) {
	var mockApplication mock.MockApplication

	secretKey := "test"

	mockApplication.On("GetAuthSecretKey", context.Background()).Return(secretKey)

	newNoteUUID, err := uuid.NewUUID()
	assert.Equal(t, nil, err)

	userUUID, err := uuid.NewUUID()
	assert.Equal(t, nil, err)

	owner := &domain.User{
		Aggregate:      domain.Aggregate{UUID: userUUID},
		FirstName:      "John",
		LastName:       "Doe",
		Email:          "johndoe@test.com",
		HashedPassword: "hashpasswordhash",
	}

	tests := []struct {
		Title                  string
		Content                string
		ApplicationError       error
		ResultingNote          domain.Note
		HandlerStatusCode      int
		HandlerResponseMessage string
	}{
		{
			Title:            "Title one",
			Content:          "Content one",
			ApplicationError: nil,
			ResultingNote: domain.Note{
				Aggregate: domain.Aggregate{UUID: newNoteUUID},
				Title:     "Title one",
				Content:   "Content one",
				OwnerUUID: userUUID,
			},
			HandlerStatusCode:      http.StatusCreated,
			HandlerResponseMessage: "Note added successfully",
		},
		{
			Title:                  "",
			Content:                "Content two",
			ApplicationError:       domain.EmptyNoteTitleError,
			ResultingNote:          domain.Note{},
			HandlerStatusCode:      400,
			HandlerResponseMessage: domain.EmptyNoteTitleError.Error(),
		},
		{
			Title:                  "Title 3",
			Content:                "",
			ApplicationError:       domain.EmptyNoteContentError,
			ResultingNote:          domain.Note{},
			HandlerStatusCode:      400,
			HandlerResponseMessage: domain.EmptyNoteContentError.Error(),
		},
		{
			Title:                  "Title 4",
			Content:                "Content 4",
			ApplicationError:       &app.AuthenticationError{Message: "User not found"},
			ResultingNote:          domain.Note{},
			HandlerStatusCode:      401,
			HandlerResponseMessage: "User not found",
		},
	}
	for _, test := range tests {
		// create context with authentication
		ctx := context.Background()
		ctx = context.WithValue(ctx, "auth", app.Auth{UserUUID: userUUID})

		// mock application AddNote use case
		mockApplication.On("AddNote", ctx, test.Title, test.Content).Return(test.ResultingNote, test.ApplicationError)

		// create access token for the owner user
		accessToken, err := util.CreateUserAccessToken(mockApplication.GetAuthSecretKey(context.Background()), *owner)
		assert.Equal(t, nil, err)

		// construct json request payload
		var jsonData = []byte(fmt.Sprintf(`{
			"title": "%s",
			"content": "%s"
		}`, test.Title, test.Content))

		// construct request and ensure there is no error
		req, err := http.NewRequest("POST", "/notes/", bytes.NewBuffer(jsonData))
		assert.Equal(t, nil, err)

		// set authentication header and content type
		req.Header.Set("Content-Type", "application/json; charset=UTF-8")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))

		// Create a new context.Context and populate it with data.
		requestCtx := context.Background()
		requestCtx = context.WithValue(requestCtx, "app", &mockApplication)

		// handle request with handler
		rr := httptest.NewRecorder()
		handler := api.GetApiRouter(&mockApplication)
		handler.ServeHTTP(rr, req.WithContext(requestCtx))

		// Check the status code is what we expect.
		assert.Equal(t, test.HandlerStatusCode, rr.Code)

		responseBody, err := io.ReadAll(rr.Body)
		assert.Equal(t, nil, err)

		type Response struct {
			Msg string `json:"msg"`
		}
		var response Response

		err = json.Unmarshal(responseBody, &response)
		assert.Equal(t, nil, err)

		assert.Equal(t, test.HandlerResponseMessage, response.Msg)

	}
}
