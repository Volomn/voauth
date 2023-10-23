package mock

import (
	"github.com/Volomn/voauth/backend/domain"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type MockNoteRepository struct{ mock.Mock }

func (repo *MockNoteRepository) Save(db *gorm.DB, note domain.Note) error {
	args := repo.Called(db, note)
	return args.Error(0)
}

func (repo *MockNoteRepository) GetNoteByUUID(db *gorm.DB, noteUUID uuid.UUID) *domain.Note {
	args := repo.Called(db, noteUUID)
	returnValue := args.Get(0)
	if returnValue == nil {
		return nil
	}
	return returnValue.(*domain.Note)
}
