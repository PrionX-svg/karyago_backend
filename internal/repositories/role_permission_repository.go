package repositories

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"hris_backend/internal/models"
	"hris_backend/internal/response"
)

type RolePermissionRepository interface {
	Create(rolePermission *models.RolePermission) error
	DeleteByRoleAndPermission(roleID, permissionID uint) error
	GetByRoleIDDetailed(roleID uint) (response.RolePermissionGrouped, error)
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

func (r *rolePermissionRepository) GetByRoleIDDetailed(roleID uint) (response.RolePermissionGrouped, error) {
	type rawResult struct {
		PermissionName string
		Label          string
		RoleName       string
	}

	var rows []rawResult

	err := r.db.Table("role_permissions").
		Select("permissions.name AS permission_name, permissions.label, roles.name AS role_name").
		Joins("JOIN permissions ON permissions.id = role_permissions.permission_id").
		Joins("JOIN roles ON roles.id = role_permissions.role_id").
		Where("role_permissions.role_id = ?", roleID).
		Scan(&rows).Error
	if err != nil {
		return response.RolePermissionGrouped{}, err
	}

	if len(rows) == 0 {
		return response.RolePermissionGrouped{}, nil
	}

	var permissions []response.PermissionSummary
	for _, row := range rows {
		permissions = append(permissions, response.PermissionSummary{
			PermissionName: row.PermissionName,
			Label:          row.Label,
		})
	}

	return response.RolePermissionGrouped{
		RoleName:    rows[0].RoleName,
		Permissions: permissions,
	}, nil
}

func (r *rolePermissionRepository) GetByRoleID(roleID uint) ([]models.RolePermission, error) {
	var result []models.RolePermission
	err := r.db.Where("role_id = ?", roleID).Find(&result).Error
	return result, err
}
