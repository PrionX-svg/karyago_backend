package repositories

import (
	"gorm.io/gorm"
	"hris_backend/internal/models"
)

type AuthRepository interface {
	Register(user models.User) error
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepository{db}
}

func (r *authRepository) Register(user models.User) error {
	if err := r.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}


