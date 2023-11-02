package note

import "github.com/google/uuid"

type NoteListItem struct {
	UUID       uuid.UUID
	Title      string
	Content    string
	IsPublic   bool
	IsFavorite bool
	IsArchived bool
}

type Note struct {
	UUID       uuid.UUID
	Title      string
	Content    string
	IsPublic   bool
	IsFavorite bool
	IsArchived bool
}
