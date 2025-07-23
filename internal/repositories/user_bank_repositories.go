package repositories

import (
	"hris_backend/internal/models"

	"gorm.io/gorm"
)

type UserBankRepository interface {
	Create(userBank *models.UserBank) error
	GetAll() ([]models.UserBank, error)
	GetByID(id uint) (*models.UserBank, error)
	GetByUUID(uuid string) (*models.UserBank, error)
	GetByUserUUID(uuid string) ([]models.UserBank, error)
	Update(userBank *models.UserBank) error
	Delete(uuid string) error
}

type userBankRepository struct {
	db *gorm.DB
}

func NewUserBankRepository(db *gorm.DB) UserBankRepository {
	return &userBankRepository{db}
}

func (r *userBankRepository) Create(userBank *models.UserBank) error {
	if err := r.db.Create(userBank).Error; err != nil {
		return err
	}
	return nil
}

func (r *userBankRepository) GetAll() ([]models.UserBank, error) {
	var userBanks []models.UserBank
	err := r.db.
		Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "uuid", "first_name", "last_name", "email")
		}).Find(&userBanks).Error
	if err != nil {
		return nil, err
	}
	return userBanks, nil
}

func (r *userBankRepository) GetByID(id uint) (*models.UserBank, error) {
	var userBank models.UserBank

	err := r.db.
		Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "uuid", "first_name", "last_name", "email")
		}).
		Where("id = ?", id).
		First(&userBank).Error

	if err != nil {
		return nil, err
	}
	return &userBank, nil
}

func (r *userBankRepository) GetByUUID(uuid string) (*models.UserBank, error) {
	var userBank models.UserBank

	err := r.db.
		Preload("User").
		Where("uuid = ?", uuid).
		First(&userBank).Error

	if err != nil {
		return nil, err
	}
	return &userBank, nil
}

func (r *userBankRepository) GetByUserUUID(uuid string) ([]models.UserBank, error) {
	var userBanks []models.UserBank
	err := r.db.
		Joins("JOIN users ON users.id = user_banks.user_id").
		Where("users.uuid = ?", uuid).
		Preload("User").
		Find(&userBanks).Error

	if err != nil {
		return nil, err
	}
	return userBanks, nil
}

func (r *userBankRepository) Update(userBank *models.UserBank) error {
	if err := r.db.Save(userBank).Error; err != nil {
		return err
	}
	return nil
}

func (r *userBankRepository) Delete(uuid string) error {
	return r.db.Where("uuid = ?", uuid).Delete(&models.UserBank{}).Error
}