package repositories

import (
	"gorm.io/gorm"
	"hris_backend/internal/models"
)

type PermissionRepository interface {
	FindByUUID(uuid string) (*models.Permission, error)
	FindByName(name string) (*models.Permission, error)
	GetAll() ([]models.Permission, error)
	RoleHasPermission(roleID uint, permissionName string) (bool, error)
}

type permissionRepository struct {
	db *gorm.DB
}

func NewPermissionRepository(db *gorm.DB) PermissionRepository {
	return &permissionRepository{db}
}

func (r *permissionRepository) FindByUUID(uuid string) (*models.Permission, error) {
	var p models.Permission
	if err := r.db.Where("uuid = ?", uuid).First(&p).Error; err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *permissionRepository) FindByName(name string) (*models.Permission, error) {
	var p models.Permission
	if err := r.db.Where("name = ?", name).First(&p).Error; err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *permissionRepository) GetAll() ([]models.Permission, error) {
	var perms []models.Permission
	err := r.db.Order("label ASC").Find(&perms).Error
	return perms, err
}

func (r *permissionRepository) RoleHasPermission(roleID uint, permissionName string) (bool, error) {
	var count int64
	err := r.db.Table("role_permissions").
		Joins("JOIN permissions ON role_permissions.permission_id = permissions.id").
		Where("role_permissions.role_id = ? AND permissions.label = ?", roleID, permissionName).
		Count(&count).Error
	return count > 0, err
}
