package repository

import (
	"github.com/Volomn/voauth/backend/domain"
	"gorm.io/gorm"
)

type UserRepository struct{}

func (repo *UserRepository) GetUserByEmail(db *gorm.DB, email string) *domain.User {
	return nil
}

func (repo *UserRepository) Save(db *gorm.DB, user domain.User) error {
	return nil
}
