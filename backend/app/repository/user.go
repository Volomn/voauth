package repository

import (
	"github.com/Volomn/voauth/backend/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetUserByEmail(db *gorm.DB, email string) *domain.User
	GetUserByUUID(db *gorm.DB, userUUID uuid.UUID) *domain.User
	Save(db *gorm.DB, user domain.User) error
}

type NoteRepository interface {
	Save(db *gorm.DB, note domain.Note) error
}
