package repositories

import (
	"gorm.io/gorm"
	"hris_backend/internal/models"
)

type UserRepository interface {
	GetByUUID(uuid string) (models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) GetByUUID(uuid string) (models.User, error) {
	var user models.User
	if err := r.db.Where("uuid = ?", uuid).First(&user).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}
