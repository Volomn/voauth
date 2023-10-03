package app

import "gorm.io/gorm"

type Application struct {
	config ApplicationConfig
	db     *gorm.DB
}

func New(config ApplicationConfig, db *gorm.DB) *Application {
	return &Application{
		config: config,
		db:     db,
	}
}
