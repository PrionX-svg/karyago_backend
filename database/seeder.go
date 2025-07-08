package database

import (
	"hris_backend/internal/models"
	"log"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func SeedRoles(db *gorm.DB) {
	roles := []models.Role{
		{UUID: uuid.New().String(), Name: "admin"},
		{UUID: uuid.New().String(), Name: "assistant"},
		{UUID: uuid.New().String(), Name: "owner"},
	}

	for _, role := range roles {
		var existingRole models.Role
		if err := db.Where("name = ?", role.Name).First(&existingRole).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				if err := db.Create(&role).Error; err != nil {
					log.Println("[ERROR]", time.Now().UTC().Format("2006-01-02 15:04:05"), "-", "Failed to create role:", err)
				}
			} else {
				log.Println("[ERROR]", time.Now().UTC().Format("2006-01-02 15:04:05"), "-", "Failed to check existing role:", err)
			}
		}
	}
}
