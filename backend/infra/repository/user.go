package repository

import (
	"strings"
	"time"

	"github.com/Volomn/voauth/backend/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type DbUser struct {
	UUID           uuid.UUID      `gorm:"type:uuid;primaryKey;column:uuid"`
	FirstName      string         `gorm:"not null;column:first_name"`
	LastName       string         `gorm:"not null;colmn:last_name"`
	Email          string         `gorm:"unique;column:email"`
	HashedPassword string         `gorm:"not null;column:hashed_password"`
	Address        string         `gorm:"not null;column:address"`
	Bio            string         `gorm:"not null;column:bio"`
	PhotoURL       string         `gorm:"not null;column:photo_url"`
	CreatedAt      time.Time      `gorm:"not null;column:created_at"`
	UpdatedAt      time.Time      `gorm:"not null;column:updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"index;column:deleted_at"`
}

type UserRepository struct{}

func (repo *UserRepository) fromOrm(user DbUser) domain.User {
	return domain.User{
		Aggregate:      domain.Aggregate{UUID: user.UUID},
		FirstName:      user.FirstName,
		LastName:       user.LastName,
		Email:          user.Email,
		HashedPassword: user.HashedPassword,
		Bio:            user.Bio,
		Address:        user.Address,
		PhotoURL:       user.PhotoURL,
	}
}

func (repo *UserRepository) toOrm(user domain.User) DbUser {
	return DbUser{
		UUID:           user.UUID,
		FirstName:      user.FirstName,
		LastName:       user.LastName,
		Email:          user.Email,
		Address:        user.Address,
		Bio:            user.Bio,
		PhotoURL:       user.PhotoURL,
		HashedPassword: user.HashedPassword,
	}
}

func (repo *UserRepository) GetUserByEmail(db *gorm.DB, email string) *domain.User {
	var result DbUser
	res := db.Where(&domain.User{Email: strings.ToLower(email)}).First(&result)
	if res.Error == nil {
		user := repo.fromOrm(result)
		return &user
	} else {
		return nil
	}

}

func (repo *UserRepository) GetUserByUUID(db *gorm.DB, userUUID uuid.UUID) *domain.User {
	return nil
}

func (repo *UserRepository) Save(db *gorm.DB, user domain.User) error {
	ormUser := repo.toOrm(user)
	result := db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "uuid"}},
		DoUpdates: clause.AssignmentColumns([]string{"first_name", "last_name", "email", "hashed_password", "address", "bio", "photo_url"}),
	}).Create(&ormUser)
	return result.Error
}
