package repositories

import (
	"hris_backend/internal/models"

	"gorm.io/gorm"
)

type RoleRepositories interface {
	FindByNameAndCompanyID(name string, companyUUID string) (*models.Role, error)
	FindByName(name string) (*models.Role, error)
	FindByID(id uint) (*models.Role, error)
	FindByUUID(uuid string) (*models.Role, error)
	FindAll(limit, offset int, search, sort string, companyID uint) ([]models.Role, int64, error)
	Create(role *models.Role) error
	Update(role *models.Role) error
	Delete(uuid string) error
}

type roleRepositories struct {
	db *gorm.DB
}

func NewRoleRepositories(db *gorm.DB) RoleRepositories {
	return &roleRepositories{db}
}

func (r *roleRepositories) FindByNameAndCompanyID(name string, companyUUID string) (*models.Role, error) {
	var company models.Company
	if err := r.db.Where("uuid = ?", companyUUID).First(&company).Error; err != nil {
		return nil, err
	}

	var role models.Role
	err := r.db.
		Where("name = ? AND company_id = ?", name, company.ID).
		First(&role).Error

	if err != nil {
		return nil, err
	}

	return &role, nil
}

func (r *roleRepositories) FindByName(name string) (*models.Role, error) {
	var role models.Role
	if err := r.db.Where("name = ?", name).First(&role).Error; err != nil {
		return nil, err
	}
	return &role, nil
}

func (r *roleRepositories) FindByID(id uint) (*models.Role, error) {
	var role models.Role
	if err := r.db.Where("id = ?", id).First(&role).Error; err != nil {
		return nil, err
	}
	return &role, nil
}

func (r *roleRepositories) FindByUUID(uuid string) (*models.Role, error) {
	var role models.Role
	if err := r.db.Where("uuid = ?", uuid).First(&role).Error; err != nil {
		return nil, err
	}
	return &role, nil
}

func (r *roleRepositories) FindAll(limit, offset int, search, sort string, companyID uint) ([]models.Role, int64, error) {
	var roles []models.Role
	var total int64

	query := r.db.Model(&models.Role{}).Where("company_id = ?", companyID)

	if search != "" {
		searchPattern := "%" + search + "%"
		query = query.Where("name LIKE ?", searchPattern)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if sort == "" {
		sort = "id desc"
	}

	if err := query.Order(sort).Limit(limit).Offset(offset).Find(&roles).Error; err != nil {
		return nil, 0, err
	}

	return roles, total, nil
}

func (r *roleRepositories) Create(role *models.Role) error {
	return r.db.Create(role).Error
}

func (r *roleRepositories) Update(role *models.Role) error {
	return r.db.Save(role).Error
}

func (r *roleRepositories) Delete(uuid string) error {
	return r.db.Where("uuid = ?", uuid).Delete(&models.Role{}).Error
}
