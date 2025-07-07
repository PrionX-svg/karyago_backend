package database

import (
	"hris_backend/internal/models"
	"log"
)

func MigrationAll() {
	err := DB.AutoMigrate(&models.User{}, &models.Role{}, &models.OTP{})
	if err != nil {
		log.Fatal("Migration failed:", err)
	}

	SeedRoles(DB)
}
