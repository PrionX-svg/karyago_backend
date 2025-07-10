package services

import (
	"hris_backend/internal/models"
	"hris_backend/internal/repositories"
	"hris_backend/internal/request"
	"hris_backend/pkg"

	"github.com/google/uuid"
)

type DepartmentGroupServices interface {
	Create(request request.DepartmentGroupReq) (models.DepartmentGroup, error)
	GetAll() ([]models.DepartmentGroup, error)
	FindByUUID(UUID string) (*models.DepartmentGroup, error)
	Update(UUID string, departmentGroupReq request.DepartmentGroupReq) (models.DepartmentGroup, error)
	Delete(UUID string) (models.DepartmentGroup, error)
}

type departmentGroupServices struct {
	departmentGroupRepo repositories.DepartmentGroupRepositories
	companyRepo         repositories.CompanyRepositories
	userRepo            repositories.UserRepository
}

func NewDepartmentGroupService(departmentGroupRepo repositories.DepartmentGroupRepositories, companyRepo repositories.CompanyRepositories, userRepo repositories.UserRepository) DepartmentGroupServices {
	return &departmentGroupServices{
		departmentGroupRepo: departmentGroupRepo,
		companyRepo:         companyRepo,
		userRepo:            userRepo,
	}
}

func (s *departmentGroupServices) Create(request request.DepartmentGroupReq) (models.DepartmentGroup, error) {
	if err := pkg.Validate.Struct(request); err != nil {
		return models.DepartmentGroup{}, err
	}

	company, err := s.companyRepo.GetByUUID(request.CompanyUUID)
	if err != nil {
		return models.DepartmentGroup{}, err
	}

	newDepartmentGroup := models.DepartmentGroup{
		UUID:            uuid.NewString(),
		CompanyID:       company.ID,
		CompanyUUID:     request.CompanyUUID,
		ResponsibleUUID: request.ResponsibleUUID,
		Name:            request.Name,
		Desc:            request.Desc,
	}

	if request.ResponsibleUUID != "" {
		responsible, err := s.userRepo.GetByUUID(request.ResponsibleUUID)
		if err != nil {
			return models.DepartmentGroup{}, err
		}

		newDepartmentGroup.ResponsibleID = responsible.ID
	}

	if err := s.departmentGroupRepo.Create(&newDepartmentGroup); err != nil {
		return models.DepartmentGroup{}, err
	}
	return newDepartmentGroup, nil
}

func (s *departmentGroupServices) GetAll() ([]models.DepartmentGroup, error) {
	var departmentGroups []models.DepartmentGroup
	departmentGroups, err := s.departmentGroupRepo.GetAll()
	if err != nil {
		return nil, err
	}
	return departmentGroups, nil
}

func (s *departmentGroupServices) FindByUUID(UUID string) (*models.DepartmentGroup, error) {
	departmentGroup, err := s.departmentGroupRepo.FindByUUID(UUID)
	if err != nil {
		return nil, err
	}
	return departmentGroup, nil
}

func (s *departmentGroupServices) Update(UUID string, departmentGroupReq request.DepartmentGroupReq) (models.DepartmentGroup, error) {
	if err := pkg.Validate.Struct(departmentGroupReq); err != nil {
		return models.DepartmentGroup{}, err
	}

	departmentGroup, err := s.departmentGroupRepo.FindByUUID(UUID)
	if err != nil {
		return models.DepartmentGroup{}, err
	}

	departmentGroup.Name = departmentGroupReq.Name
	departmentGroup.Desc = departmentGroupReq.Desc
	if departmentGroupReq.ResponsibleUUID != "" {
		responsible, err := s.userRepo.GetByUUID(departmentGroupReq.ResponsibleUUID)
		if err != nil {
			return models.DepartmentGroup{}, err
		}
		departmentGroup.ResponsibleID = responsible.ID
		departmentGroup.ResponsibleUUID = responsible.UUID
	}
	if err := s.departmentGroupRepo.Update(departmentGroup); err != nil {
		return models.DepartmentGroup{}, err
	}
	return *departmentGroup, nil
}

func (s *departmentGroupServices) Delete(UUID string) (models.DepartmentGroup, error) {
	departmentGroup, err := s.departmentGroupRepo.FindByUUID(UUID)
	if err != nil {
		return models.DepartmentGroup{}, err
	}

	if err := s.departmentGroupRepo.Delete(UUID); err != nil {
		return models.DepartmentGroup{}, err
	}
	return *departmentGroup, nil
}
