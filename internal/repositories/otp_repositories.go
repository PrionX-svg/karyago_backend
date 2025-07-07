package repositories

import (
	"hris_backend/internal/models"

	"gorm.io/gorm"
)

type OTPRepositories interface {
}

type otpRepositories struct {
	db *gorm.DB
}

func NewOTPRepositories(db *gorm.DB) OTPRepositories {
	return &otpRepositories{db}
}

func (r *otpRepositories) Create(otp *models.OTP) error {
	if err := r.db.Create(otp).Error; err != nil {
		return err
	}
	return nil
}
