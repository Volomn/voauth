package domain

import (
	"testing"

	valueobject "github.com/Volomn/voauth/backend/domain/valueobjects"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewNoteWithEmptyTitleFail(t *testing.T) {
	noteUUID, _ := uuid.NewUUID()
	ownerUUID, _ := uuid.NewUUID()
	_, err := NewNote(noteUUID, ownerUUID, false, false, false, "", "Some content")
	assert.EqualError(t, EmptyNoteTitleError, err.Error())
}

func TestNewNoteWithEmptyContentFail(t *testing.T) {
	noteUUID, _ := uuid.NewUUID()
	ownerUUID, _ := uuid.NewUUID()
	_, err := NewNote(noteUUID, ownerUUID, false, false, false, "Title", "")
	assert.EqualError(t, EmptyNoteContentError, err.Error())
}

func TestMakeNotePublic(t *testing.T) {
	noteUUID, _ := uuid.NewUUID()
	ownerUUID, _ := uuid.NewUUID()
	note, _ := NewNote(noteUUID, ownerUUID, false, false, false, "Title", "Content")
	assert.Equal(t, false, note.IsPublic)
	note.MakePublic()
	assert.Equal(t, true, note.IsPublic)
}

func TestMakeNotePrivate(t *testing.T) {
	noteUUID, _ := uuid.NewUUID()
	ownerUUID, _ := uuid.NewUUID()
	note, _ := NewNote(noteUUID, ownerUUID, true, false, false, "Title", "Content")
	assert.Equal(t, true, note.IsPublic)
	note.MakePrivate()
	assert.Equal(t, false, note.IsPublic)
}

func TestArchiveNote(t *testing.T) {
	noteUUID, _ := uuid.NewUUID()
	ownerUUID, _ := uuid.NewUUID()
	note, _ := NewNote(noteUUID, ownerUUID, true, false, false, "Title", "Content")
	assert.Equal(t, false, note.IsArchived)
	note.Archive()
	assert.Equal(t, true, note.IsArchived)
}

func TestUnarchiveNote(t *testing.T) {
	noteUUID, _ := uuid.NewUUID()
	ownerUUID, _ := uuid.NewUUID()
	note, _ := NewNote(noteUUID, ownerUUID, true, false, true, "Title", "Content")
	assert.Equal(t, true, note.IsArchived)
	note.UnArchive()
	assert.Equal(t, false, note.IsArchived)
}

func TestShareNoteWithNewUser(t *testing.T) {
	noteUUID, _ := uuid.NewUUID()
	ownerUUID, _ := uuid.NewUUID()
	note, _ := NewNote(noteUUID, ownerUUID, false, false, false, "Title", "Content")
	assert.Equal(t, 0, len(note.SharedUsers))
	newUserUUID, _ := uuid.NewUUID()
	note.ShareWithUsers([]uuid.UUID{newUserUUID}, valueobject.SharedNotePermission(valueobject.WRITE))
	assert.Equal(t, 1, len(note.SharedUsers))
	assert.Equal(t, newUserUUID, note.SharedUsers[0].UserUUID)
	assert.Equal(t, valueobject.SharedNotePermission(valueobject.WRITE), note.SharedUsers[0].Permission)
}

func TestShareNoteWithMultipleNewUsers(t *testing.T) {
	noteUUID, _ := uuid.NewUUID()
	ownerUUID, _ := uuid.NewUUID()
	note, _ := NewNote(noteUUID, ownerUUID, false, false, false, "Title", "Content")
	assert.Equal(t, 0, len(note.SharedUsers))
	newUserUUID, _ := uuid.NewUUID()
	newUserUUID2, _ := uuid.NewUUID()
	note.ShareWithUsers([]uuid.UUID{newUserUUID, newUserUUID2}, valueobject.SharedNotePermission(valueobject.WRITE))
	assert.Equal(t, 2, len(note.SharedUsers))
	assert.Equal(t, newUserUUID, note.SharedUsers[0].UserUUID)
	assert.Equal(t, newUserUUID2, note.SharedUsers[1].UserUUID)
	assert.Equal(t, valueobject.SharedNotePermission(valueobject.WRITE), note.SharedUsers[0].Permission)
	assert.Equal(t, valueobject.SharedNotePermission(valueobject.WRITE), note.SharedUsers[1].Permission)
}

func TestShareNoteWithExistingSharedUserButWithDifferentPermission(t *testing.T) {
	noteUUID, _ := uuid.NewUUID()
	ownerUUID, _ := uuid.NewUUID()
	exisingUserUUID1, _ := uuid.NewUUID()
	exisingUserUUID2, _ := uuid.NewUUID()
	exisingUserUUID3, _ := uuid.NewUUID()
	exisingUserUUID4, _ := uuid.NewUUID()

	// create new note
	note, _ := NewNote(noteUUID, ownerUUID, false, false, false, "Title", "Content")

	// share note with 4 users all with read permissions
	note.ShareWithUsers([]uuid.UUID{exisingUserUUID1, exisingUserUUID2, exisingUserUUID3, exisingUserUUID4}, valueobject.SharedNotePermission(valueobject.READ))

	// share note with two of the existing users again but with write permission
	note.ShareWithUsers([]uuid.UUID{exisingUserUUID2, exisingUserUUID3}, valueobject.SharedNotePermission(valueobject.WRITE))

	//ensure we still have only 4 entries
	assert.Equal(t, 4, len(note.SharedUsers))

	// ensure the permission of the 2nd and 3rd user is now write
	for _, sharedUser := range note.SharedUsers {
		if sharedUser.UserUUID == exisingUserUUID2 || sharedUser.UserUUID == exisingUserUUID3 {
			assert.Equal(t, valueobject.SharedNotePermission(valueobject.WRITE), sharedUser.Permission)
		}
	}
}
