package repositories

import (
	"hris_backend/internal/models"
	"gorm.io/gorm"
)

type DepartmentGroupRepositories interface {
	Create(departmentGroup *models.DepartmentGroup) error
	GetAll() ([]models.DepartmentGroup, error)
	FindByUUID(UUID string) (*models.DepartmentGroup, error)
	Update(departmentGroup *models.DepartmentGroup) error
	Delete(UUID string) error
}

type departmentGroupRepositories struct {
	db *gorm.DB
}

func NewDepartmentGroupRepositories(db *gorm.DB) DepartmentGroupRepositories {
	return &departmentGroupRepositories{db}
}

func (r *departmentGroupRepositories) Create(departmentGroup *models.DepartmentGroup) error{
	if err := r.db.Create(departmentGroup).Error; err != nil {
		return err
	}
	return nil
}

func (r *departmentGroupRepositories) GetAll() ([]models.DepartmentGroup, error) {
	var departmentGroups []models.DepartmentGroup
	if err := r.db.Find(&departmentGroups).Error; err != nil {
		return nil, err
	}
	return departmentGroups, nil
}

func (r *departmentGroupRepositories) FindByUUID(UUID string) (*models.DepartmentGroup, error) {
	departmentGroup := &models.DepartmentGroup{}
	if err := r.db.Where("uuid = ?", UUID).First(departmentGroup).Error; err != nil {
		return nil, err
	}
	return departmentGroup, nil
}

func (r *departmentGroupRepositories) Update(departmentGroup *models.DepartmentGroup) error {
	if err := r.db.Save(departmentGroup).Error; err != nil {
		return err
	}
	return nil
}

func (r *departmentGroupRepositories) Delete(UUID string) error {
	if err := r.db.Where("uuid = ?", UUID).Delete(&models.DepartmentGroup{}).Error; err != nil {
		return err
	}
	return nil
}