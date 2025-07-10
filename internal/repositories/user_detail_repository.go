package repositories

import (
	"gorm.io/gorm"
	"hris_backend/internal/models"
)

type UserDetailRepository interface {
	Create(detail *models.UserDetail) error
	Update(detail *models.UserDetail) error
	FindByID(id uint) (*models.UserDetail, error)
	FindByUUID(uuid string) (*models.UserDetail, error)
	FindByUserID(userID uint) (*models.UserDetail, error)
	Delete(id uint) error
}

type userDetailRepository struct {
	db *gorm.DB
}

func NewUserDetailRepository(db *gorm.DB) UserDetailRepository {
	return &userDetailRepository{db: db}
}

func (r *userDetailRepository) Create(detail *models.UserDetail) error {
	return r.db.Create(detail).Error
}

func (r *userDetailRepository) Update(detail *models.UserDetail) error {
	return r.db.Save(detail).Error
}

func (r *userDetailRepository) FindByID(id uint) (*models.UserDetail, error) {
	var detail models.UserDetail
	if err := r.db.First(&detail, id).Error; err != nil {
		return nil, err
	}
	return &detail, nil
}

func (r *userDetailRepository) FindByUUID(uuid string) (*models.UserDetail, error) {
	var detail models.UserDetail
	if err := r.db.Where("uuid = ?", uuid).First(&detail).Error; err != nil {
		return nil, err
	}
	return &detail, nil
}

func (r *userDetailRepository) FindByUserID(userID uint) (*models.UserDetail, error) {
	var detail models.UserDetail

	err := r.db.
		Table("user_details").
		Select("user_details.*, users.uuid as user_uuid").
		Joins("left join users on users.id = user_details.user_id").
		Where("user_details.user_id = ?", userID).
		First(&detail).Error

	if err != nil {
		return nil, err
	}

	return &detail, nil
}

func (r *userDetailRepository) Delete(id uint) error {
	return r.db.Delete(&models.UserDetail{}, id).Error
}
