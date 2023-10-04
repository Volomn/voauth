package valueobject

import (
	"github.com/google/uuid"
)

var (
	WRITE = "WRITE"
	READ  = "READ"
)

type SharedNotePermission string

type SharedUser struct {
	UserUUID   uuid.UUID
	Permission SharedNotePermission
}
