package repositories

import (
	"hris_backend/internal/models"

	"gorm.io/gorm"
)

type CompanyDetailShiftRepository interface {
	Create(detail *models.CompanyDetailShift) error
	Update(detail *models.CompanyDetailShift) error
	DeleteByUUID(uuid string) error
	FindByUUID(uuid string) (*models.CompanyDetailShift, error)
	FindAll() ([]models.CompanyDetailShift, error)
	FindByShiftID(shiftID uint) ([]models.CompanyDetailShift, error)
}

type companyDetailShiftRepository struct {
	db *gorm.DB
}

func NewCompanyDetailShiftRepository(db *gorm.DB) CompanyDetailShiftRepository {
	return &companyDetailShiftRepository{db}
}

func (r *companyDetailShiftRepository) Create(detail *models.CompanyDetailShift) error {
	return r.db.Create(detail).Error
}

func (r *companyDetailShiftRepository) Update(detail *models.CompanyDetailShift) error {
	return r.db.Save(detail).Error
}

func (r *companyDetailShiftRepository) DeleteByUUID(uuid string) error {
	return r.db.Where("uuid = ?", uuid).Delete(&models.CompanyDetailShift{}).Error
}

func (r *companyDetailShiftRepository) FindByUUID(uuid string) (*models.CompanyDetailShift, error) {
	var detail models.CompanyDetailShift
	err := r.db.Preload("Shift").Where("uuid = ?", uuid).First(&detail).Error
	if err != nil {
		return nil, err
	}
	return &detail, nil
}

func (r *companyDetailShiftRepository) FindAll() ([]models.CompanyDetailShift, error) {
	var details []models.CompanyDetailShift
	err := r.db.Preload("Shift").Find(&details).Error
	return details, err
}

func (r *companyDetailShiftRepository) FindByShiftID(shiftID uint) ([]models.CompanyDetailShift, error) {
	var details []models.CompanyDetailShift
	err := r.db.Where("shift_id = ?", shiftID).Find(&details).Error
	return details, err
}
