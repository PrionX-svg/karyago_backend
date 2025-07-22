package repositories

import (
	"hris_backend/internal/models"

	"gorm.io/gorm"
)

type UserFamilyRepository interface {
	Create(userFamily *models.UserFamily) error
	GetAll() ([]models.UserFamily, error)
	GetByID(id uint) (*models.UserFamily, error)
	GetByUUID(uuid string) (*models.UserFamily, error)
	GetByUserUUID(uuid string) ([]models.UserFamily, error)
	Update(userFamily *models.UserFamily) error
	Delete(uuid string) error
	DeleteByUserID(userID uint) error
}

type userFamilyRepository struct {
	db *gorm.DB
}

func NewUserFamilyRepository(db *gorm.DB) UserFamilyRepository {
	return &userFamilyRepository{db}
}

func (r *userFamilyRepository) Create(userFamily *models.UserFamily) error {
	if err := r.db.Create(userFamily).Error; err != nil {
		return err
	}
	return nil
}

func (r *userFamilyRepository) GetAll() ([]models.UserFamily, error) {
	var userFamilies []models.UserFamily
	err := r.db.
		Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "uuid", "first_name", "last_name", "email")
		}).Find(&userFamilies).Error
	if err != nil {
		return nil, err
	}
	return userFamilies, nil
}

func (r *userFamilyRepository) GetByID(id uint) (*models.UserFamily, error) {
	var userFamily models.UserFamily

	err := r.db.
		Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "uuid", "first_name", "last_name", "email")
		}).
		Where("id = ?", id).
		First(&userFamily).Error

	if err != nil {
		return nil, err
	}
	return &userFamily, nil
}

func (r *userFamilyRepository) GetByUUID(uuid string) (*models.UserFamily, error) {
	var userFamily models.UserFamily

	err := r.db.
		Preload("User").
		Where("uuid = ?", uuid).
		First(&userFamily).Error

	if err != nil {
		return nil, err
	}
	return &userFamily, nil
}

func (r *userFamilyRepository) GetByUserUUID(uuid string) ([]models.UserFamily, error) {
	var userFamilies []models.UserFamily
	err := r.db.
		Joins("JOIN users ON users.id = user_families.user_id").
		Where("users.uuid = ?", uuid).
		Preload("User").
		Find(&userFamilies).Error

	if err != nil {
		return nil, err
	}
	return userFamilies, nil
}

func (r *userFamilyRepository) Update(userFamily *models.UserFamily) error {
	if err := r.db.Save(userFamily).Error; err != nil {
		return err
	}
	return nil
}

func (r *userFamilyRepository) Delete(uuid string) error {
	return r.db.Where("uuid = ?", uuid).Delete(&models.UserFamily{}).Error
}

func (r *userFamilyRepository) DeleteByUserID(userID uint) error {
	return r.db.Where("user_id = ?", userID).Delete(&models.UserFamily{}).Error
}
