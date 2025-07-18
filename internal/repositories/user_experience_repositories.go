package repositories

import (
	"hris_backend/internal/models"

	"gorm.io/gorm"
)

type UserExperienceRepository interface {
	Create(userExperience *models.UserExperience) error
	GetAll() ([]models.UserExperience, error)
	GetByID(id uint) (*models.UserExperience, error)
	GetByUUID(uuid string) (*models.UserExperience, error)
	GetByUserUUID(uuid string) ([]models.UserExperience, error)
	Update(userExperience *models.UserExperience) error
	Delete(uuid string) error
}

type userExperienceRepository struct {
	db *gorm.DB
}

func NewUserExperienceRepository(db *gorm.DB) UserExperienceRepository {
	return &userExperienceRepository{db}
}

func (r *userExperienceRepository) Create(userExperience *models.UserExperience) error {
	if err := r.db.Create(userExperience).Error; err != nil {
		return err
	}
	return nil
}

func (r *userExperienceRepository) GetAll() ([]models.UserExperience, error) {
	var userExperiences []models.UserExperience
	err := r.db.
		Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "uuid", "first_name", "last_name", "email")
		}).Find(&userExperiences).Error
	if err != nil {
		return nil, err
	}
	return userExperiences, nil
}

func (r *userExperienceRepository) GetByID(id uint) (*models.UserExperience, error) {
	var userExperience models.UserExperience

	err := r.db.
		Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "uuid", "first_name", "last_name", "email")
		}).
		Where("id = ?", id).
		First(&userExperience).Error

	if err != nil {
		return nil, err
	}
	return &userExperience, nil
}

func (r *userExperienceRepository) GetByUUID(uuid string) (*models.UserExperience, error) {
	var userExperience models.UserExperience

	err := r.db.
		Preload("User").
		Where("uuid = ?", uuid).
		First(&userExperience).Error

	if err != nil {
		return nil, err
	}
	return &userExperience, nil
}

func (r *userExperienceRepository) GetByUserUUID(uuid string) ([]models.UserExperience, error) {
	var experiences []models.UserExperience
	err := r.db.
		Joins("JOIN users ON users.id = user_experiences.user_id").
		Where("users.uuid = ?", uuid).
		Preload("User").
		Find(&experiences).Error

	if err != nil {
		return nil, err
	}
	return experiences, nil
}

func (r *userExperienceRepository) Update(userExperience *models.UserExperience) error {
	if err := r.db.Save(userExperience).Error; err != nil {
		return err
	}
	return nil
}

func (r *userExperienceRepository) Delete(uuid string) error {
	return r.db.Where("uuid = ?", uuid).Delete(&models.UserExperience{}).Error
}