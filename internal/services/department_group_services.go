package services

import (
	"hris_backend/internal/models"
	"hris_backend/internal/repositories"
	"hris_backend/internal/request"
	"hris_backend/pkg"

	"github.com/google/uuid"
)

type DepartmentGroupServices interface {
	Create(request request.DepartmentGroupCreateReq) (models.DepartmentGroup, error)
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

func (s *departmentGroupServices) Create(request request.DepartmentGroupCreateReq) (models.DepartmentGroup, error) {
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
