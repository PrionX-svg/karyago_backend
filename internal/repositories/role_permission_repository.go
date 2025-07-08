package repositories

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"hris_backend/internal/models"
)

type RolePermissionRepository interface {
	Create(rolePermission *models.RolePermission) error
	DeleteByRoleAndPermission(roleID, permissionID uint) error
	GetByRoleID(roleID uint) ([]models.RolePermission, error)
}

type rolePermissionRepository struct {
	db *gorm.DB
}

func NewRolePermissionRepository(db *gorm.DB) RolePermissionRepository {
	return &rolePermissionRepository{db}
}

func (r *rolePermissionRepository) Create(rolePermission *models.RolePermission) error {
	rolePermission.UUID = uuid.New().String()
	return r.db.Create(rolePermission).Error
}

func (r *rolePermissionRepository) DeleteByRoleAndPermission(roleID, permissionID uint) error {
	return r.db.Where("role_id = ? AND permission_id = ?", roleID, permissionID).Delete(&models.RolePermission{}).Error
}

func (r *rolePermissionRepository) GetByRoleID(roleID uint) ([]models.RolePermission, error) {
	var result []models.RolePermission
	err := r.db.Where("role_id = ?", roleID).Find(&result).Error
	return result, err
}
