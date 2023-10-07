package domain

import "errors"

var EmptyNoteTitleError error = errors.New("Note title is empty")
var EmptyNoteContentError error = errors.New("Note content is empty")
var EmptyFirstNameError error = errors.New("User first name is empty")
var EmptyLastNameError error = errors.New("User last name is empty")
var InvalidEmailError error = errors.New("User email is invalid")
