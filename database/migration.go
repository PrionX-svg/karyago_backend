package database

import (
	"hris_backend/internal/models"
	"log"
)

func MigrationAll() {
	err := DB.AutoMigrate(&models.User{}, &models.Role{}, &models.OTP{}, &models.Company{})
	if err != nil {
		log.Fatal("Migration failed:", err)
	}

	SeedRoles(DB)
}
