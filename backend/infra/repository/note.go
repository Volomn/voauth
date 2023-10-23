package repository

import (
	"github.com/Volomn/voauth/backend/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type NoteRepository struct{}

func (repo *NoteRepository) Save(db *gorm.DB, note domain.Note) error {
	return nil
}

func (repo *NoteRepository) GetNoteByUUID(db *gorm.DB, noteUUID uuid.UUID) *domain.Note {
	return nil
}
