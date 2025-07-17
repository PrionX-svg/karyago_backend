package repositories

import (
	"hris_backend/internal/models"

	"gorm.io/gorm"
)

type UserEducationRepository interface {
	Create(userEducation *models.UserEducation) error
	GetAll() ([]models.UserEducation, error)
	GetByID(id uint) (*models.UserEducation, error)
	GetByUUID(uuid string) (*models.UserEducation, error)
	GetByUserUUID(uuid string) ([]models.UserEducation, error)
	Update(userEducation *models.UserEducation) error
	Delete(uuid string) error
}

type userEducationRepository struct {
	db *gorm.DB
}

func NewUserEducationRepository(db *gorm.DB) UserEducationRepository {
	return &userEducationRepository{db}
}

func (r *userEducationRepository) Create(userEducation *models.UserEducation) error {
	if err := r.db.Create(userEducation).Error; err != nil {
		return err
	}
	return nil
}

func (r *userEducationRepository) GetAll() ([]models.UserEducation, error) {
	var userEducations []models.UserEducation
	err := r.db.
		Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "uuid", "first_name", "last_name", "email")
		}).Find(&userEducations).Error
	if err != nil {
		return nil, err
	}
	return userEducations, nil
}

func (r *userEducationRepository) GetByID(id uint) (*models.UserEducation, error) {
	var userEducation models.UserEducation

	err := r.db.
		Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "uuid", "first_name", "last_name", "email")
		}).
		Where("id = ?", id).
		First(&userEducation).Error

	if err != nil {
		return nil, err
	}
	return &userEducation, nil
}

func (r *userEducationRepository) GetByUUID(uuid string) (*models.UserEducation, error) {
	var userEducation models.UserEducation

	err := r.db.
		Preload("User").
		Where("uuid = ?", uuid).
		First(&userEducation).Error

	if err != nil {
		return nil, err
	}
	return &userEducation, nil
}

func (r *userEducationRepository) GetByUserUUID(uuid string) ([]models.UserEducation, error) {
	var userEducations []models.UserEducation

	err := r.db.
		Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "uuid", "first_name", "last_name", "email")
		}).
		Where("user_uuid = ?", uuid).
		Find(&userEducations).Error

	if err != nil {
		return nil, err
	}
	return userEducations, nil
}

func (r *userEducationRepository) Update(userEducation *models.UserEducation) error {
	if err := r.db.Save(userEducation).Error; err != nil {
		return err
	}
	return nil
}

func (r *userEducationRepository) Delete(uuid string) error {
	return r.db.Where("uuid = ?", uuid).Delete(&models.UserEducation{}).Error
}
