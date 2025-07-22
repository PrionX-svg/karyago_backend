package repositories

import (
	"gorm.io/gorm"
	"hris_backend/internal/models"
)

type shiftRepository struct {
	db *gorm.DB
}

type ShiftRepository interface {
	Create(shift *models.Shift) error
	FindByID(id uint) (*models.Shift, error)
	FindByUUID(uuid string) (*models.Shift, error)
	FindAll() ([]models.Shift, error)
	Update(shift *models.Shift) error
	Delete(id uint) error
}

func NewShiftRepository(db *gorm.DB) ShiftRepository {
	return &shiftRepository{db}
}

func (r *shiftRepository) Create(shift *models.Shift) error {
	return r.db.Create(shift).Error
}

func (r *shiftRepository) FindByID(id uint) (*models.Shift, error) {
	var shift models.Shift
	err := r.db.Preload("Details").First(&shift, id).Error
	if err != nil {
		return nil, err
	}
	return &shift, nil
}

func (r *shiftRepository) FindByUUID(uuid string) (*models.Shift, error) {
	var shift models.Shift
	err := r.db.
		Preload("Details").
		Where("uuid = ?", uuid).
		First(&shift).Error

	if err != nil {
		return nil, err
	}
	return &shift, nil
}

func (r *shiftRepository) FindAll() ([]models.Shift, error) {
	var shifts []models.Shift
	err := r.db.Preload("Details").Find(&shifts).Error
	if err != nil {
		return nil, err
	}
	return shifts, nil
}

func (r *shiftRepository) Update(shift *models.Shift) error {
	return r.db.Save(shift).Error
}

func (r *shiftRepository) Delete(id uint) error {
	return r.db.Delete(&models.Shift{}, id).Error
}
