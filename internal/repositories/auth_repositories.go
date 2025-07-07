package repositories

import (
	"gorm.io/gorm"
	"hris_backend/internal/models"
)

type AuthRepository interface {
	Register(user models.User) error
	FindByEmail(email string) (*models.User, error)
	UpdatePassword(user *models.User) error
	VerifyUser(user *models.User) error
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

func (r *authRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *authRepository) UpdatePassword(user *models.User) error {
	if err := r.db.Model(user).Update("password", user.Password).Error; err != nil {
		return err
	}
	return nil
}

func (r *authRepository) VerifyUser(user *models.User) error {
	if err := r.db.Model(user).Update("isVerified", true).Error; err != nil {
		return err
	}
	return nil
}