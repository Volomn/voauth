package domain

import (
	"net/mail"
	"strings"

	"github.com/google/uuid"
)

type User struct {
	Aggregate
	FirstName      string
	LastName       string
	Email          string
	HashedPassword string
	Address        string
	Bio            string
	PhotoURL       string
}

func NewUser(userUUID uuid.UUID, firstName, lastName, email, hashedPassword, address, bio, photoURL string) (*User, error) {
	if len(strings.TrimSpace(firstName)) <= 0 {
		return nil, EmptyFirstNameError
	}
	if len(strings.TrimSpace(lastName)) <= 0 {
		return nil, EmptyLastNameError
	}
	if _, err := mail.ParseAddress(email); err != nil {
		return nil, InvalidEmailError
	}
	return &User{
		Aggregate:      Aggregate{UUID: userUUID},
		FirstName:      firstName,
		LastName:       lastName,
		Email:          strings.ToLower(email),
		Address:        address,
		HashedPassword: hashedPassword,
		Bio:            bio,
		PhotoURL:       photoURL,
	}, nil
}

func (user *User) SetAddress(newAddress string) {
	user.Address = newAddress
}

func (user *User) SetBio(newBio string) {
	user.Bio = newBio
}

func (user *User) SetPhotoURL(newPhotoURL string) {
	user.PhotoURL = newPhotoURL
}

func (user *User) SetFirstName(newFirstName string) {
	user.FirstName = newFirstName
}

func (user *User) SetLastName(newLastName string) {
	user.FirstName = newLastName
}
