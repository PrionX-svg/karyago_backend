package repositories

import (
	"gorm.io/gorm"
	"hris_backend/internal/models"
	"time"
)

type AuthRepository interface {
	Register(user *models.User) error
	FindByEmail(email string) (*models.User, error)
	CheckLogin(email string) (*models.User, error)
	UpdateLastLogin(userID uint, loginTime time.Time) error
	UpdatePasswordByEmail(email, hashedPassword string) error
	UpdatePassword(user *models.User) error
	VerifyUser(user *models.User) error
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepository{db}
}

func (r *authRepository) Register(user *models.User) error {
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

func (r *authRepository) CheckLogin(email string) (*models.User, error) {
	var user models.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *authRepository) UpdateLastLogin(userID uint, loginTime time.Time) error {
	return r.db.Model(&models.User{}).
		Where("id = ?", userID).
		Update("last_login_at", loginTime).Error
}

func (r *authRepository) UpdatePasswordByEmail(email, hashedPassword string) error {
	return r.db.Model(&models.User{}).Where("email = ?", email).Update("password", hashedPassword).Error
}

func (r *authRepository) UpdatePassword(user *models.User) error {
	if err := r.db.Model(user).Update("password", user.Password).Error; err != nil {
		return err
	}
	return nil
}

func (r *authRepository) VerifyUser(user *models.User) error {
	if err := r.db.Model(user).Update("is_verified", true).Error; err != nil {
		return err
	}
	return nil
}
