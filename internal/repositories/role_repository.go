package repositories

import (
	"hris_backend/internal/models"

	"gorm.io/gorm"
)

type RoleRepositories interface {
	FindByName(name string) (*models.Role, error)
	FindByID(id uint) (*models.Role, error)
}

type roleRepositories struct {
	db *gorm.DB
}

func NewRoleRepositories(db *gorm.DB) RoleRepositories {
	return &roleRepositories{db}
}

func (r *roleRepositories) FindByName(name string) (*models.Role, error) {
	var role models.Role
	if err := r.db.Where("name = ?", name).First(&role).Error; err != nil {
		return nil, err
	}
	return &role, nil
}

func (r *roleRepositories) FindByID(id uint) (*models.Role, error) {
	var role models.Role
	if err := r.db.Where("id = ?", id).First(&role).Error; err != nil {
		return nil, err
	}
	return &role, nil
}
