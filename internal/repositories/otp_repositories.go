package repositories

import (
	"hris_backend/internal/models"
	"time"

	"gorm.io/gorm"
)

type OTPRepositories interface {
	Create(otp *models.OTP) error
	FindByUUID(UUID string) (*models.OTP, error)
	FindValidOTP(uuid string) (*models.OTP, error)
	FindByCodeAndTarget(code, target string) (*models.OTP, error)
	FindValidOTPByTargetAndPurpose(email, purpose string) (*models.OTP, error)
	MarkUsed(UUID string) error
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

func (r *otpRepositories) FindByUUID(UUID string) (*models.OTP, error) {
	otp := &models.OTP{}
	if err := r.db.Where("uuid = ?", UUID).First(otp).Error; err != nil {
		return nil, err
	}
	return otp, nil
}

func (r *otpRepositories) FindValidOTP(uuid string) (*models.OTP, error) {
	var otp models.OTP
	err := r.db.Where("uuid = ? AND is_used = false AND expires_at > ?", uuid, time.Now()).
		First(&otp).Error
	if err != nil {
		return nil, err
	}
	return &otp, nil
}

func (r *otpRepositories) FindByCodeAndTarget(code, target string) (*models.OTP, error) {
	var otp models.OTP
	err := r.db.Where("code = ? AND target = ? AND is_used = false AND expires_at > ?", code, target, time.Now()).
		First(&otp).Error
	if err != nil {
		return nil, err
	}
	return &otp, nil
}

func (r *otpRepositories) FindValidOTPByTargetAndPurpose(email, purpose string) (*models.OTP, error) {
	var otp models.OTP
	err := r.db.
		Where("target = ? AND purpose = ? AND is_used = false AND expires_at > ?", email, purpose, time.Now()).
		Order("created_at DESC").
		First(&otp).Error

	if err != nil {
		return nil, err
	}
	return &otp, nil
}

func (r *otpRepositories) MarkUsed(UUID string) error {
	if err := r.db.Model(&models.OTP{}).Where("UUID = ?", UUID).Update("IsUsed", true).Error; err != nil {
		return err
	}

	return nil
}
