package services

import (
	"hris_backend/internal/models"
	"hris_backend/internal/repositories"
)

type PermissionService interface {
	GetByUUID(uuid string) (*models.Permission, error)
	GetByName(name string) (*models.Permission, error)
	GetAll() ([]models.Permission, error)
}

type permissionService struct {
	repo repositories.PermissionRepository
}

func NewPermissionService(repo repositories.PermissionRepository) PermissionService {
	return &permissionService{repo}
}

func (s *permissionService) GetByUUID(uuid string) (*models.Permission, error) {
	return s.repo.FindByUUID(uuid)
}

func (s *permissionService) GetByName(name string) (*models.Permission, error) {
	return s.repo.FindByName(name)
}

func (s *permissionService) GetAll() ([]models.Permission, error) {
	return s.repo.GetAll()
}
