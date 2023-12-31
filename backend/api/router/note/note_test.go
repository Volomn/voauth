package note_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Volomn/voauth/backend/api"
	noterouter "github.com/Volomn/voauth/backend/api/router/note"
	"github.com/Volomn/voauth/backend/api/router/util"
	"github.com/Volomn/voauth/backend/app"
	"github.com/Volomn/voauth/backend/domain"
	"github.com/Volomn/voauth/backend/mock"
	qs "github.com/Volomn/voauth/backend/queryservice"
	"github.com/Volomn/voauth/backend/queryservice/note"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestAddNoteHandler(t *testing.T) {
	var mockApplication mock.MockApplication
	var mockNotequeryService mock.MockNoteQueryService

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

		// handle request with handler
		rr := httptest.NewRecorder()
		handler := api.GetApiRouter(&mockApplication, &mockNotequeryService)
		handler.ServeHTTP(rr, req)

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

func TestUpdateNoteHandler(t *testing.T) {
	var mockApplication mock.MockApplication
	var mockNotequeryService mock.MockNoteQueryService

	secretKey := "test"

	mockApplication.On("GetAuthSecretKey", context.Background()).Return(secretKey)

	noteUUID, err := uuid.NewUUID()
	assert.Equal(t, nil, err)

	notFoundNoteUUID, err := uuid.NewUUID()
	assert.Equal(t, nil, err)

	ownerUUID, err := uuid.NewUUID()
	assert.Equal(t, nil, err)

	notFoundUserUUID, err := uuid.NewUUID()
	assert.Equal(t, nil, err)

	notOwnerUUID, err := uuid.NewUUID()
	assert.Equal(t, nil, err)

	owner := &domain.User{
		Aggregate:      domain.Aggregate{UUID: ownerUUID},
		FirstName:      "John",
		LastName:       "Doe",
		Email:          "johndoe@test.com",
		HashedPassword: "hashpasswordhash",
	}

	notOwner := &domain.User{
		Aggregate:      domain.Aggregate{UUID: notOwnerUUID},
		FirstName:      "Jane",
		LastName:       "Doe",
		Email:          "janedoe@test.com",
		HashedPassword: "hashpasswordhash",
	}

	note := &domain.Note{
		Aggregate: domain.Aggregate{UUID: noteUUID},
		Title:     "tone",
		Content:   "cone",
		OwnerUUID: owner.UUID,
	}

	tests := []struct {
		AuthUserUUID           uuid.UUID
		NoteUUID               uuid.UUID
		Title                  string
		Content                string
		ApplicationError       error
		ResultingNote          domain.Note
		HandlerStatusCode      int
		HandlerResponseMessage string
	}{
		{
			AuthUserUUID:     owner.UUID,
			NoteUUID:         note.UUID,
			Title:            "Title one",
			Content:          "Content one",
			ApplicationError: nil,
			ResultingNote: domain.Note{
				Aggregate: domain.Aggregate{UUID: note.UUID},
				Title:     "Title one",
				Content:   "Content one",
				OwnerUUID: note.OwnerUUID,
			},
			HandlerStatusCode:      200,
			HandlerResponseMessage: "Note updated successfully",
		},
		{
			AuthUserUUID:           owner.UUID,
			NoteUUID:               note.UUID,
			Title:                  "",
			Content:                "Content two",
			ApplicationError:       domain.EmptyNoteTitleError,
			ResultingNote:          domain.Note{},
			HandlerStatusCode:      400,
			HandlerResponseMessage: domain.EmptyNoteTitleError.Error(),
		},
		{
			AuthUserUUID:           owner.UUID,
			NoteUUID:               note.UUID,
			Title:                  "Title 3",
			Content:                "",
			ApplicationError:       domain.EmptyNoteContentError,
			ResultingNote:          domain.Note{},
			HandlerStatusCode:      400,
			HandlerResponseMessage: domain.EmptyNoteContentError.Error(),
		},
		{
			AuthUserUUID:           notFoundUserUUID,
			NoteUUID:               note.UUID,
			Title:                  "Title 4",
			Content:                "Content 4",
			ApplicationError:       &app.AuthenticationError{Message: "User not found"},
			ResultingNote:          domain.Note{},
			HandlerStatusCode:      401,
			HandlerResponseMessage: "User not found",
		},
		{
			AuthUserUUID:           notOwner.UUID,
			NoteUUID:               note.UUID,
			Title:                  "Title 5",
			Content:                "Content 5",
			ApplicationError:       &app.AuthorizationError{Message: "User is not permitted to update this note"},
			ResultingNote:          domain.Note{},
			HandlerStatusCode:      403,
			HandlerResponseMessage: "Not allowed",
		},
		{
			AuthUserUUID:           owner.UUID,
			NoteUUID:               notFoundNoteUUID,
			Title:                  "Title 6",
			Content:                "Content 6",
			ApplicationError:       &app.EntityNotFoundError{Message: "Note not found"},
			ResultingNote:          domain.Note{},
			HandlerStatusCode:      404,
			HandlerResponseMessage: "Note not found",
		},
	}
	for _, test := range tests {
		// create context with authentication
		ctx := context.Background()
		ctx = context.WithValue(ctx, "auth", app.Auth{UserUUID: test.AuthUserUUID})

		// mock application AddNote use case
		mockApplication.On("UpdateNote", ctx, test.NoteUUID, test.Title, test.Content).Return(test.ResultingNote, test.ApplicationError)

		// create access token for the owner user
		accessToken, err := util.CreateUserAccessToken(mockApplication.GetAuthSecretKey(context.Background()), domain.User{Aggregate: domain.Aggregate{UUID: test.AuthUserUUID}})
		assert.Equal(t, nil, err)

		// construct json request payload
		var jsonData = []byte(fmt.Sprintf(`{
			"title": "%s",
			"content": "%s"
		}`, test.Title, test.Content))

		// construct request and ensure there is no error
		requestURL := fmt.Sprintf("/notes/%s", test.NoteUUID.String())
		slog.Info("Request url", "requestURL", requestURL)
		req, err := http.NewRequest("PUT", requestURL, bytes.NewBuffer(jsonData))
		assert.Equal(t, nil, err)

		// set authentication header and content type
		req.Header.Set("Content-Type", "application/json; charset=UTF-8")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))

		// handle request with handler
		rr := httptest.NewRecorder()
		handler := api.GetApiRouter(&mockApplication, &mockNotequeryService)
		handler.ServeHTTP(rr, req)

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

func TestDeleteNoteHandler(t *testing.T) {
	var mockApplication mock.MockApplication
	var mockNotequeryService mock.MockNoteQueryService

	secretKey := "test"

	mockApplication.On("GetAuthSecretKey", context.Background()).Return(secretKey)

	noteUUID, err := uuid.NewUUID()
	assert.Equal(t, nil, err)

	notFoundNoteUUID, err := uuid.NewUUID()
	assert.Equal(t, nil, err)

	ownerUUID, err := uuid.NewUUID()
	assert.Equal(t, nil, err)

	notFoundUserUUID, err := uuid.NewUUID()
	assert.Equal(t, nil, err)

	notOwnerUUID, err := uuid.NewUUID()
	assert.Equal(t, nil, err)

	owner := &domain.User{
		Aggregate:      domain.Aggregate{UUID: ownerUUID},
		FirstName:      "John",
		LastName:       "Doe",
		Email:          "johndoe@test.com",
		HashedPassword: "hashpasswordhash",
	}

	notOwner := &domain.User{
		Aggregate:      domain.Aggregate{UUID: notOwnerUUID},
		FirstName:      "Jane",
		LastName:       "Doe",
		Email:          "janedoe@test.com",
		HashedPassword: "hashpasswordhash",
	}

	note := &domain.Note{
		Aggregate: domain.Aggregate{UUID: noteUUID},
		Title:     "tone",
		Content:   "cone",
		OwnerUUID: owner.UUID,
	}

	tests := []struct {
		AuthUserUUID           uuid.UUID
		NoteUUID               uuid.UUID
		ApplicationError       error
		HandlerStatusCode      int
		HandlerResponseMessage string
	}{
		{
			AuthUserUUID:           owner.UUID,
			NoteUUID:               note.UUID,
			ApplicationError:       nil,
			HandlerStatusCode:      200,
			HandlerResponseMessage: "Note deleted successfully",
		},
		{
			AuthUserUUID:           notFoundUserUUID,
			NoteUUID:               note.UUID,
			ApplicationError:       &app.AuthenticationError{Message: "User not found"},
			HandlerStatusCode:      401,
			HandlerResponseMessage: "User not found",
		},
		{
			AuthUserUUID:           notOwner.UUID,
			NoteUUID:               note.UUID,
			ApplicationError:       &app.AuthorizationError{Message: "User is not permitted to delete this note"},
			HandlerStatusCode:      403,
			HandlerResponseMessage: "Not allowed",
		},
		{
			AuthUserUUID:           owner.UUID,
			NoteUUID:               notFoundNoteUUID,
			ApplicationError:       &app.EntityNotFoundError{Message: "Note not found"},
			HandlerStatusCode:      404,
			HandlerResponseMessage: "Note not found",
		},
	}
	for _, test := range tests {
		// create context with authentication
		ctx := context.Background()
		ctx = context.WithValue(ctx, "auth", app.Auth{UserUUID: test.AuthUserUUID})

		// mock application AddNote use case
		mockApplication.On("DeleteNote", ctx, test.NoteUUID).Return(test.ApplicationError)

		// create access token for the owner user
		accessToken, err := util.CreateUserAccessToken(mockApplication.GetAuthSecretKey(context.Background()), domain.User{Aggregate: domain.Aggregate{UUID: test.AuthUserUUID}})
		assert.Equal(t, nil, err)

		// construct request and ensure there is no error
		requestURL := fmt.Sprintf("/notes/%s", test.NoteUUID.String())
		slog.Info("Request url", "requestURL", requestURL)
		req, err := http.NewRequest("DELETE", requestURL, nil)
		assert.Equal(t, nil, err)

		// set authentication header and content type
		req.Header.Set("Content-Type", "application/json; charset=UTF-8")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))

		// handle request with handler
		rr := httptest.NewRecorder()
		handler := api.GetApiRouter(&mockApplication, &mockNotequeryService)
		handler.ServeHTTP(rr, req)

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

func TestFetchUserNotes(t *testing.T) {
	userUUID, err := uuid.NewUUID()
	assert.Equal(t, nil, err)

	anotherUserUUID, err := uuid.NewUUID()
	assert.Equal(t, nil, err)

	noteUUID, err := uuid.NewUUID()
	assert.Equal(t, nil, err)

	contextWithoutAuth := context.Background()
	contextWithUserAuth := context.WithValue(contextWithoutAuth, "auth", qs.Auth{UserUUID: userUUID})
	contextWithAnotherUserAuth := context.WithValue(contextWithoutAuth, "auth", qs.Auth{UserUUID: anotherUserUUID})

	var mockNotequeryService mock.MockNoteQueryService
	mockNotequeryService.On("FetchUserNotes", contextWithoutAuth).Return([]note.NoteListItem{}, &qs.AuthenticationError{Message: "Not Authorized"})
	mockNotequeryService.On("FetchUserNotes", contextWithUserAuth).Return([]note.NoteListItem{
		{UUID: noteUUID, Title: "Lorem", Content: "Ipsum", IsPublic: true, IsFavorite: false, IsArchived: false},
	}, nil)
	mockNotequeryService.On("FetchUserNotes", contextWithAnotherUserAuth).Return([]note.NoteListItem{}, nil)

	var mockApplication mock.MockApplication
	mockApplication.On("GetAuthSecretKey", context.Background()).Return("secret")

	userAccessToken, err := util.CreateUserAccessToken(mockApplication.GetAuthSecretKey(context.Background()), domain.User{Aggregate: domain.Aggregate{UUID: userUUID}})
	assert.Equal(t, nil, err)

	anotherUserAccessToken, err := util.CreateUserAccessToken(mockApplication.GetAuthSecretKey(context.Background()), domain.User{Aggregate: domain.Aggregate{UUID: anotherUserUUID}})
	assert.Equal(t, nil, err)

	tests := []struct {
		AccessToken     string
		StatusCode      int
		ResponseMessage string
		ResponseBody    []noterouter.NoteListItemResponse
	}{
		{AccessToken: "", StatusCode: 401, ResponseMessage: "Unauthorized", ResponseBody: nil},
		{AccessToken: userAccessToken, StatusCode: 200, ResponseBody: []noterouter.NoteListItemResponse{
			{
				UUID:       noteUUID,
				Title:      "Lorem",
				Content:    "Ipsum",
				IsPublic:   true,
				IsFavorite: false,
				IsArchived: false,
			},
		}},
		{AccessToken: anotherUserAccessToken, StatusCode: 200, ResponseBody: []noterouter.NoteListItemResponse{}},
	}
	for _, test := range tests {
		// construct request and ensure there is no error
		requestURL := fmt.Sprintf("/notes/")
		req, err := http.NewRequest("GET", requestURL, nil)
		assert.Equal(t, nil, err)

		// set authentication header and content type
		req.Header.Set("Content-Type", "application/json; charset=UTF-8")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", test.AccessToken))

		// Create a new context.Context and populate it with data.
		// requestCtx := context.Background()
		// requestCtx = context.WithValue(requestCtx, "app", &mockApplication)
		// requestCtx = context.WithValue(requestCtx, "noteQueryService", &mockNotequeryService)

		// handle request with handler
		rr := httptest.NewRecorder()
		handler := api.GetApiRouter(&mockApplication, &mockNotequeryService)
		handler.ServeHTTP(rr, req)

		// Check the status code is what we expect.
		assert.Equal(t, test.StatusCode, rr.Code)

		responseBody, err := io.ReadAll(rr.Body)
		assert.Equal(t, nil, err)

		slog.Info("Response body is", "responseBody", responseBody)

		type ErrorResponse struct {
			Msg string `json:"msg"`
		}

		var errorResponse ErrorResponse
		var successResponse []noterouter.NoteListItemResponse

		if test.StatusCode == 200 {
			err = json.Unmarshal(responseBody, &successResponse)
			assert.Equal(t, nil, err)
			assert.Equal(t, test.ResponseBody, successResponse)
		} else {
			err = json.Unmarshal(responseBody, &errorResponse)
			assert.Equal(t, nil, err)
			assert.Equal(t, test.ResponseMessage, errorResponse.Msg)
		}

	}
}

func TestGetUserNote(t *testing.T) {
	userUUID, err := uuid.NewUUID()
	assert.Equal(t, nil, err)

	anotherUserUUID, err := uuid.NewUUID()
	assert.Equal(t, nil, err)

	noteUUID, err := uuid.NewUUID()
	assert.Equal(t, nil, err)

	contextWithoutAuth := context.Background()
	contextWithUserAuth := context.WithValue(contextWithoutAuth, "auth", qs.Auth{UserUUID: userUUID})
	contextWithAnotherUserAuth := context.WithValue(contextWithoutAuth, "auth", qs.Auth{UserUUID: anotherUserUUID})

	var mockNotequeryService mock.MockNoteQueryService
	mockNotequeryService.On("GetUserNote", contextWithoutAuth, noteUUID).Return(note.Note{}, &qs.AuthenticationError{Message: "Not Authorized"})
	mockNotequeryService.On("GetUserNote", contextWithUserAuth, noteUUID).Return(note.Note{UUID: noteUUID, Title: "Lorem", Content: "Ipsum", IsPublic: true, IsFavorite: false, IsArchived: false}, nil)
	mockNotequeryService.On("GetUserNote", contextWithAnotherUserAuth, noteUUID).Return(note.Note{}, &qs.EntityNotFoundError{Message: "Note not found"})

	var mockApplication mock.MockApplication
	mockApplication.On("GetAuthSecretKey", context.Background()).Return("secret")

	userAccessToken, err := util.CreateUserAccessToken(mockApplication.GetAuthSecretKey(context.Background()), domain.User{Aggregate: domain.Aggregate{UUID: userUUID}})
	assert.Equal(t, nil, err)

	anotherUserAccessToken, err := util.CreateUserAccessToken(mockApplication.GetAuthSecretKey(context.Background()), domain.User{Aggregate: domain.Aggregate{UUID: anotherUserUUID}})
	assert.Equal(t, nil, err)

	tests := []struct {
		AccessToken     string
		StatusCode      int
		ResponseMessage string
		ResponseBody    noterouter.NoteResponse
	}{
		{AccessToken: "", StatusCode: 401, ResponseMessage: "Unauthorized", ResponseBody: noterouter.NoteResponse{}},
		{AccessToken: userAccessToken, StatusCode: 200, ResponseBody: noterouter.NoteResponse{
			UUID:       noteUUID,
			Title:      "Lorem",
			Content:    "Ipsum",
			IsPublic:   true,
			IsFavorite: false,
			IsArchived: false,
		}},
		{AccessToken: anotherUserAccessToken, StatusCode: 404, ResponseBody: noterouter.NoteResponse{}, ResponseMessage: "Note not found"},
	}
	for _, test := range tests {
		// construct request and ensure there is no error
		requestURL := fmt.Sprintf("/notes/%s", noteUUID.String())
		req, err := http.NewRequest("GET", requestURL, nil)
		assert.Equal(t, nil, err)

		// set authentication header and content type
		req.Header.Set("Content-Type", "application/json; charset=UTF-8")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", test.AccessToken))

		// Create a new context.Context and populate it with data.
		// requestCtx := context.Background()
		// requestCtx = context.WithValue(requestCtx, "app", &mockApplication)
		// requestCtx = context.WithValue(requestCtx, "noteQueryService", &mockNotequeryService)

		// handle request with handler
		rr := httptest.NewRecorder()
		handler := api.GetApiRouter(&mockApplication, &mockNotequeryService)
		handler.ServeHTTP(rr, req)

		// Check the status code is what we expect.
		assert.Equal(t, test.StatusCode, rr.Code)

		responseBody, err := io.ReadAll(rr.Body)
		assert.Equal(t, nil, err)

		slog.Info("Response body is", "responseBody", responseBody)

		type ErrorResponse struct {
			Msg string `json:"msg"`
		}

		var errorResponse ErrorResponse
		var successResponse noterouter.NoteResponse

		if test.StatusCode == 200 {
			err = json.Unmarshal(responseBody, &successResponse)
			assert.Equal(t, nil, err)
			assert.Equal(t, test.ResponseBody, successResponse)
		} else {
			err = json.Unmarshal(responseBody, &errorResponse)
			assert.Equal(t, nil, err)
			assert.Equal(t, test.ResponseMessage, errorResponse.Msg)
		}

	}
}
