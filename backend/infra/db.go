package infra

import (
	"fmt"

	"github.com/Volomn/voauth/backend/infra/repository"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDb(dbHost, dbUser, dbPassword, dbName string, dbPort int) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d", dbHost, dbUser, dbPassword, dbName, dbPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

func AutoMigrateDB(db *gorm.DB) {
	db.AutoMigrate(&repository.DbUser{}, &repository.DbNote{})
}

func DropAllTables(db *gorm.DB) {
	db.Migrator().DropTable(&repository.DbUser{})
	db.Migrator().DropTable(&repository.DbUser{})
}
