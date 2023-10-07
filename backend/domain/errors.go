package domain

import "errors"

var EmptyNoteTitleError error = errors.New("Note title is empty")
var EmptyNoteContentError error = errors.New("Note content is empty")
