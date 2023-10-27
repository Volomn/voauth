package repository

import (
	"github.com/Volomn/voauth/backend/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type NoteRepository interface {
	GetNoteByUUID(db *gorm.DB, noteUUID uuid.UUID) *domain.Note
	Delete(db *gorm.DB, note domain.Note) error
	Save(db *gorm.DB, note domain.Note) error
}
