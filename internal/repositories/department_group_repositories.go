package repositories

import (
	"gorm.io/gorm"
	"hris_backend/internal/models"
)

type DepartmentGroupRepositories interface {
	Create(departmentGroup *models.DepartmentGroup) error
	GetAll() ([]models.DepartmentGroup, error)
	FindByUUID(UUID string) (*models.DepartmentGroup, error)
	FindByUUIDFromID(id uint) (*models.DepartmentGroup, error)
	GetCompanyByUUID(uuid string) (*models.Company, error)
	GetDataTable(limit, offset int, search string, companyID uint) ([]models.DepartmentGroup, int64, int64, error)
	Update(departmentGroup *models.DepartmentGroup) error
	Delete(UUID string) error
}

type departmentGroupRepositories struct {
	db *gorm.DB
}

func NewDepartmentGroupRepositories(db *gorm.DB) DepartmentGroupRepositories {
	return &departmentGroupRepositories{db}
}

func (r *departmentGroupRepositories) Create(departmentGroup *models.DepartmentGroup) error {
	if err := r.db.Create(departmentGroup).Error; err != nil {
		return err
	}
	return nil
}

func (r *departmentGroupRepositories) GetAll() ([]models.DepartmentGroup, error) {
	var departmentGroups []models.DepartmentGroup
	err := r.db.
		Table("department_groups").
		Select("department_groups.*, companies.uuid AS CompanyUUID, users.uuid AS ResponsibleUUID").
		Joins("JOIN companies ON companies.id = department_groups.company_id").
		Joins("LEFT JOIN users ON users.id = department_groups.responsible_id").
		Find(&departmentGroups).Error
	if err != nil {
		return nil, err
	}
	return departmentGroups, nil
}

func (r *departmentGroupRepositories) FindByUUID(UUID string) (*models.DepartmentGroup, error) {
	var departmentGroup models.DepartmentGroup
	err := r.db.
		Table("department_groups").
		Select("department_groups.*, companies.uuid AS CompanyUUID, users.uuid AS ResponsibleUUID").
		Joins("JOIN companies ON companies.id = department_groups.company_id").
		Joins("LEFT JOIN users ON users.id = department_groups.responsible_id").
		Where("department_groups.uuid = ?", UUID).
		First(&departmentGroup).Error

	if err != nil {
		return nil, err
	}
	return &departmentGroup, nil
}

func (r *departmentGroupRepositories) FindByUUIDFromID(id uint) (*models.DepartmentGroup, error) {
	var group models.DepartmentGroup
	err := r.db.First(&group, id).Error
	if err != nil {
		return nil, err
	}
	return &group, nil
}

func (r *departmentGroupRepositories) GetCompanyByUUID(uuid string) (*models.Company, error) {
	var company models.Company
	err := r.db.
		Table("companies").
		Where("uuid = ?", uuid).
		First(&company).Error
	if err != nil {
		return nil, err
	}
	return &company, nil
}

func (r *departmentGroupRepositories) GetDataTable(limit, offset int, search string, companyID uint) ([]models.DepartmentGroup, int64, int64, error) {
	var results []models.DepartmentGroup
	var total int64
	var filtered int64

	if err := r.db.Model(&models.DepartmentGroup{}).Where("company_id = ?", companyID).Count(&total).Error; err != nil {
		return nil, 0, 0, err
	}

	query := r.db.Model(&models.DepartmentGroup{}).
		Where("company_id = ?", companyID).
		Limit(limit).Offset(offset).
		Order("id DESC")

	if search != "" {
		query = query.Where("name LIKE ? OR `desc` LIKE ?", "%"+search+"%", "%"+search+"%")
	}

	if err := query.Find(&results).Count(&filtered).Error; err != nil {
		return nil, 0, 0, err
	}

	return results, total, filtered, nil
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
