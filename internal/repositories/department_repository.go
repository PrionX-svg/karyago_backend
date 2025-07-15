package repositories

import (
	"gorm.io/gorm"
	"hris_backend/internal/models"
)

type DepartmentRepositories interface {
	Create(department *models.Department) error
	GetAll() ([]models.Department, error)
	FindByUUID(UUID string) (*models.Department, error)
	GetDataTable(limit, offset int, search string, companyID uint) ([]models.Department, int64, int64, error)
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

func (r *departmentRepositories) GetDataTable(limit, offset int, search string, companyID uint) ([]models.Department, int64, int64, error) {
	var departments []models.Department
	var total int64
	var filtered int64

	if err := r.db.Model(&models.Department{}).
		Joins("JOIN department_groups ON departments.department_group_id = department_groups.id").
		Where("department_groups.company_id = ?", companyID).
		Count(&total).Error; err != nil {
		return nil, 0, 0, err
	}

	query := r.db.Model(&models.Department{}).
		Joins("JOIN department_groups ON departments.department_group_id = department_groups.id").
		Where("department_groups.company_id = ?", companyID)

	if search != "" {
		query = query.Where("departments.name LIKE ?", "%"+search+"%")
	}

	if err := query.Count(&filtered).Error; err != nil {
		return nil, 0, 0, err
	}

	err := query.
		Select("departments.*").
		Limit(limit).
		Offset(offset).
		Order("departments.id desc").
		Find(&departments).Error
	if err != nil {
		return nil, 0, 0, err
	}

	return departments, total, filtered, nil
}

func (r *departmentRepositories) Update(department *models.Department) error {
	return r.db.Save(department).Error
}

func (r *departmentRepositories) Delete(UUID string) error {
	return r.db.Where("uuid = ?", UUID).Delete(&models.Department{}).Error
}
