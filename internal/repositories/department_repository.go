package repositories

import (
	"gorm.io/gorm"
	"hris_backend/internal/models"
)

type DepartmentRepositories interface {
	Create(department *models.Department) error
	GetAll() ([]models.Department, error)
	FindByUUID(UUID string) (*models.Department, error)
	Update(department *models.Department) error
	Delete(UUID string) error
}

type departmentRepositories struct {
	db *gorm.DB
}

func NewDepartmentRepositories(db *gorm.DB) DepartmentRepositories {
	return &departmentRepositories{db}
}

func (r *departmentRepositories) Create(department *models.Department) error {
	return r.db.Create(department).Error
}

func (r *departmentRepositories) GetAll() ([]models.Department, error) {
	var departments []models.Department
	err := r.db.
		Table("departments").
		Select("departments.*").
		Find(&departments).Error
	if err != nil {
		return nil, err
	}
	return departments, nil
}

func (r *departmentRepositories) FindByUUID(UUID string) (*models.Department, error) {
	var department models.Department
	err := r.db.
		Table("departments").
		Select("departments.*").
		Where("departments.uuid = ?", UUID).
		First(&department).Error

	if err != nil {
		return nil, err
	}
	return &department, nil
}

func (r *departmentRepositories) Update(department *models.Department) error {
	return r.db.Save(department).Error
}

func (r *departmentRepositories) Delete(UUID string) error {
	return r.db.Where("uuid = ?", UUID).Delete(&models.Department{}).Error
}
