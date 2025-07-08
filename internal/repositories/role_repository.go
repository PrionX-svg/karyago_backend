package repositories

import (
	"hris_backend/internal/models"

	"gorm.io/gorm"
)

type RoleRepositories interface {
	FindByName(name string) (*models.Role, error)
	FindByID(id uint) (*models.Role, error)
	FindByUUID(uuid string) (*models.Role, error)
	FindAll(limit, offset int, search, sort string) ([]models.Role, int64, error)
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

func (r *roleRepositories) FindAll(limit, offset int, search, sort string) ([]models.Role, int64, error) {
	var roles []models.Role
	var total int64

	query := r.db.Model(&models.Role{})

	if search != "" {
		query = query.Where("name LIKE ?", "%"+search+"%")
	}

	query.Count(&total)

	if sort != "" {
		query = query.Order(sort)
	} else {
		query = query.Order("created_at DESC")
	}

	err := query.Limit(limit).Offset(offset).Find(&roles).Error
	if err != nil {
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
