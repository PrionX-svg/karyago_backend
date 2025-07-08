package services

import (
	"hris_backend/internal/models"
	"hris_backend/internal/repositories"

	"github.com/google/uuid"
)

type RoleService interface {
	GetByID(id uint) (*models.Role, error)
	GetByUUID(uuid string) (*models.Role, error)
	GetByName(name string) (*models.Role, error)
	GetAll(limit, offset int, search, sort string) ([]models.Role, int64, error)
	Create(role *models.Role, actorID uint) error
	Update(role *models.Role, actorID uint) error
	Delete(uuid string) error
}

type roleService struct {
	repo repositories.RoleRepositories
}

func NewRoleService(repo repositories.RoleRepositories) RoleService {
	return &roleService{repo}
}

func (s *roleService) GetByID(id uint) (*models.Role, error) {
	return s.repo.FindByID(id)
}

func (s *roleService) GetByUUID(uuid string) (*models.Role, error) {
	return s.repo.FindByUUID(uuid)
}

func (s *roleService) GetByName(name string) (*models.Role, error) {
	return s.repo.FindByName(name)
}

func (s *roleService) GetAll(limit, offset int, search, sort string) ([]models.Role, int64, error) {
	return s.repo.FindAll(limit, offset, search, sort)
}

func (s *roleService) Create(role *models.Role, actorID uint) error {
	role.UUID = uuid.New().String()
	role.CreatedBy = actorID
	role.ModifyBy = actorID
	return s.repo.Create(role)
}

func (s *roleService) Update(role *models.Role, actorID uint) error {
	role.ModifyBy = actorID
	return s.repo.Update(role)
}

func (s *roleService) Delete(uuid string) error {
	return s.repo.Delete(uuid)
}
