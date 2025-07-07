package database

import (
	"gorm.io/gorm"
	"hris_backend/internal/models"
)

func MigrateAll(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.User{},
	)
}
