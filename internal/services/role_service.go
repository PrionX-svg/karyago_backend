package services

import (
	"fmt"
	"hris_backend/internal/models"
	"hris_backend/internal/repositories"

	"github.com/google/uuid"
)

type RoleService interface {
	GetByID(id uint) (*models.Role, error)
	GetByUUID(uuid string) (*models.Role, error)
	GetByName(name string) (*models.Role, error)
	GetAll(limit, offset int, search, sort, companyUUID string) ([]models.Role, int64, error)
	Create(role *models.Role, actorID uint) error
	Update(role *models.Role, actorID uint) error
	Delete(uuid string, companyID uint) error
}

type roleService struct {
	repo        repositories.RoleRepositories
	companyRepo repositories.CompanyRepositories
}

func NewRoleService(repo repositories.RoleRepositories, companyRepo repositories.CompanyRepositories) RoleService {
	return &roleService{repo, companyRepo}
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

func (s *roleService) GetAll(limit, offset int, search, sort, companyUUID string) ([]models.Role, int64, error) {
	if companyUUID == "" {
		return nil, 0, fmt.Errorf("company_uuid is required")
	}

	company, err := s.companyRepo.GetByUUID(companyUUID)
	if err != nil {
		return nil, 0, fmt.Errorf("company not found: %w", err)
	}

	return s.repo.FindAll(limit, offset, search, sort, company.ID)
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

func (s *roleService) Delete(uuid string, actorCompanyID uint) error {
	role, err := s.repo.FindByUUID(uuid)
	if err != nil {
		return fmt.Errorf("role not found: %w", err)
	}

	if actorCompanyID != 0 && role.CompanyID != nil && *role.CompanyID != actorCompanyID {
		return fmt.Errorf("you are not allowed to delete this role")
	}

	return s.repo.Delete(uuid)
}
