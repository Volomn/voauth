package domain

import (
	"slices"
	"strings"

	valueobject "github.com/Volomn/voauth/backend/domain/valueobjects"
	"github.com/google/uuid"
)

type Note struct {
	Aggregate
	IsPublic    bool
	IsFavorite  bool
	IsArchived  bool
	Title       string
	Content     string
	OwnerUUID   uuid.UUID
	SharedUsers []valueobject.SharedUser
}

func NewNote(noteUUID uuid.UUID, ownerUUID uuid.UUID, isPublic, isFavorite, isArchived bool, title string, content string) (*Note, error) {
	if len(strings.TrimSpace(title)) <= 0 {
		return nil, EmptyNoteTitleError
	}
	if len(strings.TrimSpace(content)) <= 0 {
		return nil, EmptyNoteContentError
	}
	return &Note{
		Aggregate:  Aggregate{UUID: noteUUID},
		IsPublic:   isPublic,
		IsArchived: isArchived,
		IsFavorite: isFavorite,
		Title:      title,
		Content:    content,
		OwnerUUID:  ownerUUID,
	}, nil
}

func (note *Note) MakePublic() {
	note.IsPublic = true
}

func (note *Note) MakePrivate() {
	note.IsPublic = false
}

func (note *Note) Favorite() {
	note.IsFavorite = true
}

func (note *Note) UnFavorite() {
	note.IsFavorite = false
}

func (note *Note) Archive() {
	note.IsArchived = true
}

func (note *Note) UnArchive() {
	note.IsArchived = false
}

func (note *Note) SetTitle(newTitle string) error {
	if len(strings.TrimSpace(newTitle)) <= 0 {
		return EmptyNoteTitleError
	}
	note.Title = newTitle
	return nil
}

func (note *Note) SetContent(newContent string) error {
	if len(strings.TrimSpace(newContent)) <= 0 {
		return EmptyNoteContentError
	}
	note.Content = newContent
	return nil
}

func (note *Note) ShareWithUsers(usersUUIDS []uuid.UUID, permission valueobject.SharedNotePermission) {

	// remove all users from notes SharedUsers if user's uuid is in usersUUIDS parameter
	newSharedUsers := slices.DeleteFunc(note.SharedUsers, func(el valueobject.SharedUser) bool {
		return slices.Contains(usersUUIDS, el.UserUUID)
	})

	// now insert SharedUser for every userUUID in userUUIDS parameter
	for _, userUUID := range usersUUIDS {
		newSharedUsers = append(newSharedUsers, valueobject.SharedUser{UserUUID: userUUID, Permission: permission})
	}

	//set the SharedUsers slice for the note to newSharedUsers
	note.SharedUsers = newSharedUsers
}
