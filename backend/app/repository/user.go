package repository

import (
	"github.com/Volomn/voauth/backend/domain"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetUserByEmail(db *gorm.DB, email string) *domain.User
	Save(db *gorm.DB, user domain.User) error
}
