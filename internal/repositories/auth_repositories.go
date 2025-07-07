package repositories

import (
	"hris_backend/internal/models"
	"gorm.io/gorm"
)

type AuthRepository interface {
	Register(user models.User) error
	FindByEmail(email string) (*models.User, error)
	UpdatePassword(user *models.User) error
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

func (r* authRepository) FindByEmail(email string) (*models.User, error)  {
	var user models.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil{
		return nil,err
	}
	return &user, nil
}

func (r* authRepository) UpdatePassword(user *models.User) error {
	if err := r.db.Model(user).Update("password", user.Password).Error; err != nil{
		return err
	}
	return nil
}

