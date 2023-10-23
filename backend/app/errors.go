package app

import (
	"errors"
)

var UserWithEmailAlreadyExistsError error = errors.New("User with email address already exists")
var WeakPasswordError error = errors.New("Password is weak")
var SomethingWentWrongError error = errors.New("Something went wrong")
var InvalidLoginCredentialsError error = errors.New("Invalid login credentials")

type AuthenticationError struct {
	Message string
}

func (err *AuthenticationError) Error() string {
	return err.Message
}

type EntityNotFoundError struct {
	Message string
}

func (err *EntityNotFoundError) Error() string {
	return err.Message
}

type AuthorizationError struct {
	Message string
}

func (err *AuthorizationError) Error() string {
	return err.Message
}
