package mock

import (
	"github.com/Volomn/voauth/backend/domain"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type MockNoteRepository struct{ mock.Mock }

func (repo *MockNoteRepository) Save(db *gorm.DB, note domain.Note) error {
	args := repo.Called(db, note)
	return args.Error(0)
}
