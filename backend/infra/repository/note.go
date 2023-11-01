package repository

import (
	"time"

	"github.com/Volomn/voauth/backend/domain"
	valueobject "github.com/Volomn/voauth/backend/domain/valueobjects"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type DbNote struct {
	UUID       uuid.UUID      `gorm:"type:uuid;primaryKey;column:uuid"`
	Title      string         `gorm:"not null;column:title"`
	Content    string         `gorm:"not null;colmn:content"`
	IsPublic   bool           `gorm:"not null;column:is_public"`
	IsFavorite bool           `gorm:"not null;column:is_favorite"`
	IsArchived bool           `gorm:"not null;column:is_archived"`
	OwnerUUID  uuid.UUID      `gorm:"not null;column:owner_uuid;type:uuid"`
	CreatedAt  time.Time      `gorm:"not null;column:created_at"`
	UpdatedAt  time.Time      `gorm:"not null;column:updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index;column:deleted_at"`
}

type NoteRepository struct{}

func (repo *NoteRepository) fromOrm(note DbNote) domain.Note {
	return domain.Note{
		Aggregate:   domain.Aggregate{UUID: note.UUID},
		Title:       note.Title,
		Content:     note.Content,
		IsPublic:    note.IsPublic,
		IsFavorite:  note.IsFavorite,
		IsArchived:  note.IsArchived,
		OwnerUUID:   note.OwnerUUID,
		SharedUsers: []valueobject.SharedUser{},
	}
}

func (repo *NoteRepository) toOrm(note domain.Note) DbNote {
	return DbNote{
		UUID:       note.UUID,
		Title:      note.Title,
		Content:    note.Content,
		IsPublic:   note.IsPublic,
		IsArchived: note.IsArchived,
		IsFavorite: note.IsFavorite,
		OwnerUUID:  note.OwnerUUID,
	}
}

func (repo *NoteRepository) Save(db *gorm.DB, note domain.Note) error {
	ormNote := repo.toOrm(note)
	result := db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "uuid"}},
		DoUpdates: clause.AssignmentColumns([]string{"title", "content", "is_public", "is_favorite", "is_archived", "owner_uuid"}),
	}).Create(&ormNote)
	return result.Error
}

func (repo *NoteRepository) Delete(db *gorm.DB, note domain.Note) error {
	result := db.Delete(&DbNote{}, note.UUID)
	return result.Error
}

func (repo *NoteRepository) GetNoteByUUID(db *gorm.DB, noteUUID uuid.UUID) *domain.Note {
	dbNote := DbNote{UUID: noteUUID}
	result := db.First(&dbNote)
	if result.Error != nil {
		return nil
	}
	note := repo.fromOrm(dbNote)
	return &note
}
