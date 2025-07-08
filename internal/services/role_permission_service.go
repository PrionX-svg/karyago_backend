package services

import (
	"fmt"

	"github.com/google/uuid"
	"hris_backend/internal/models"
	"hris_backend/internal/repositories"
)

type RolePermissionService interface {
	AssignPermission(roleUUID, permissionUUID string, actorID uint) error
	RemovePermission(roleUUID, permissionUUID string) error
	GetPermissionsByRole(roleUUID string) ([]models.RolePermission, error)
}

type rolePermissionService struct {
	rpRepo   repositories.RolePermissionRepository
	roleRepo repositories.RoleRepositories
	permRepo repositories.PermissionRepository
}

func NewRolePermissionService(
	rpRepo repositories.RolePermissionRepository,
	roleRepo repositories.RoleRepositories,
	permRepo repositories.PermissionRepository,
) RolePermissionService {
	return &rolePermissionService{rpRepo, roleRepo, permRepo}
}

func (s *rolePermissionService) AssignPermission(roleUUID, permissionUUID string, actorID uint) error {
	role, err := s.roleRepo.FindByUUID(roleUUID)
	if err != nil {
		return fmt.Errorf("role not found: %w", err)
	}
	perm, err := s.permRepo.FindByUUID(permissionUUID)
	if err != nil {
		return fmt.Errorf("permission not found: %w", err)
	}
	rp := &models.RolePermission{
		UUID:         uuid.New().String(),
		RoleID:       role.ID,
		PermissionID: perm.ID,
		CreatedBy:    actorID,
		ModifyBy:     actorID,
	}
	return s.rpRepo.Create(rp)
}

func (s *rolePermissionService) RemovePermission(roleUUID, permissionUUID string) error {
	role, err := s.roleRepo.FindByUUID(roleUUID)
	if err != nil {
		return fmt.Errorf("role not found: %w", err)
	}
	perm, err := s.permRepo.FindByUUID(permissionUUID)
	if err != nil {
		return fmt.Errorf("permission not found: %w", err)
	}
	return s.rpRepo.DeleteByRoleAndPermission(role.ID, perm.ID)
}

func (s *rolePermissionService) GetPermissionsByRole(roleUUID string) ([]models.RolePermission, error) {
	role, err := s.roleRepo.FindByUUID(roleUUID)
	if err != nil {
		return nil, fmt.Errorf("role not found: %w", err)
	}
	return s.rpRepo.GetByRoleID(role.ID)
}
