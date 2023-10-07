package app

import "errors"

var UserWithEmailAlreadyExistsError error = errors.New("User with email address already exists")
var WeakPasswordError error = errors.New("Password is weak")
var SomethingWentWrongError error = errors.New("Something went wrong")
