package mock

import (
	"github.com/Volomn/voauth/backend/domain"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type MockUserRepository struct{ mock.Mock }

func (repo *MockUserRepository) GetUserByEmail(db *gorm.DB, email string) *domain.User {
	args := repo.Called(db, email)
	returnValue := args.Get(0)
	if returnValue == nil {
		return nil
	}
	return returnValue.(*domain.User)
}

func (repo *MockUserRepository) GetUserByUUID(db *gorm.DB, userUUID uuid.UUID) *domain.User {
	args := repo.Called(db, userUUID)
	returnValue := args.Get(0)
	if returnValue == nil {
		return nil
	}
	return returnValue.(*domain.User)
}

func (repo *MockUserRepository) Save(db *gorm.DB, user domain.User) error {
	args := repo.Called(db, user)
	return args.Error(0)
}
