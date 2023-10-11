package domain

import (
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	tests := []struct {
		FirstName      string
		LastName       string
		Email          string
		HashedPassword string
		Address        string
		Bio            string
		PhotoURL       string
		Err            error
	}{
		{FirstName: "  ", LastName: "Doe", Email: "johndoe@test.com", HashedPassword: "hashedpassword", Address: "Home", Bio: "My bio", PhotoURL: "", Err: EmptyFirstNameError},
		{FirstName: "John", LastName: "  ", Email: "johndoe@test.com", HashedPassword: "hashedpassword", Address: "Home", Bio: "My bio", PhotoURL: "", Err: EmptyLastNameError},
		{FirstName: "John", LastName: "Doe", Email: "invalidemail", HashedPassword: "hashedpassword", Address: "Home", Bio: "My bio", PhotoURL: "", Err: InvalidEmailError},
		{FirstName: "John", LastName: "Doe", Email: "johndoe@test.com", HashedPassword: "hashedpassword", Address: "Home", Bio: "My bio", PhotoURL: "", Err: nil},
		{FirstName: "John", LastName: "Doe", Email: "JohnDoe@test.com", HashedPassword: "hashedpassword", Address: "Home", Bio: "My bio", PhotoURL: "", Err: nil},
	}

	for _, test := range tests {
		userUUID, _ := uuid.NewUUID()
		user, err := NewUser(userUUID, test.FirstName, test.LastName, test.Email, test.HashedPassword, test.Address, test.Bio, test.PhotoURL)
		if test.Err != nil {
			assert.EqualError(t, err, test.Err.Error())
		} else {
			assert.Equal(t, nil, err)
			assert.Equal(t, test.FirstName, user.FirstName)
			assert.Equal(t, test.LastName, user.LastName)
			assert.Equal(t, strings.ToLower(test.Email), user.Email)
			assert.Equal(t, test.HashedPassword, user.HashedPassword)
			assert.Equal(t, test.Bio, user.Bio)
			assert.Equal(t, test.Address, user.Address)
			assert.Equal(t, test.PhotoURL, user.PhotoURL)

		}
	}
}
