package services

import (
	"hris_backend/internal/models"
	"hris_backend/internal/repositories"
	"hris_backend/internal/request"
	"hris_backend/internal/response"
	"hris_backend/pkg"

	"github.com/google/uuid"
)

type DepartmentGroupServices interface {
	Create(request request.DepartmentGroupReq) (response.DepartmentGroupResponse, error)
	GetAll() ([]response.DepartmentGroupResponse, error)
	FindByUUID(UUID string) (*response.DepartmentGroupResponse, error)
	GetDataTable(page, limit int, search, companyUUID string) ([]response.DepartmentGroupResponse, int64, int64, error)
	Update(UUID string, departmentGroupReq request.DepartmentGroupReq) (response.DepartmentGroupResponse, error)
	Delete(UUID string) (response.DepartmentGroupResponse, error)
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

func (s *departmentGroupServices) Create(request request.DepartmentGroupReq) (response.DepartmentGroupResponse, error) {
	if err := pkg.Validate.Struct(request); err != nil {
		return response.DepartmentGroupResponse{}, err
	}

	company, err := s.companyRepo.GetByUUID(request.CompanyUUID)
	if err != nil {
		return response.DepartmentGroupResponse{}, err
	}

	newDepartmentGroup := models.DepartmentGroup{
		UUID:            uuid.NewString(),
		CompanyID:       company.ID,
		ResponsibleUUID: request.ResponsibleUUID,
		Name:            request.Name,
		Desc:            request.Desc,
	}

	if request.ResponsibleUUID != "" {
		responsible, err := s.userRepo.GetByUUID(request.ResponsibleUUID)
		if err != nil {
			return response.DepartmentGroupResponse{}, err
		}
		newDepartmentGroup.ResponsibleID = responsible.ID
	}

	if err := s.departmentGroupRepo.Create(&newDepartmentGroup); err != nil {
		return response.DepartmentGroupResponse{}, err
	}

	return response.DepartmentGroupResponse{
		UUID:            newDepartmentGroup.UUID,
		CompanyUUID:     request.CompanyUUID,
		ResponsibleUUID: request.ResponsibleUUID,
		Name:            newDepartmentGroup.Name,
		Desc:            newDepartmentGroup.Desc,
	}, nil
}

func (s *departmentGroupServices) GetAll() ([]response.DepartmentGroupResponse, error) {
	departmentGroups, err := s.departmentGroupRepo.GetAll()
	if err != nil {
		return nil, err
	}

	var responses []response.DepartmentGroupResponse
	for _, group := range departmentGroups {
		company, err := s.companyRepo.GetByID(group.CompanyID)
		if err != nil {
			return nil, err
		}

		res := response.DepartmentGroupResponse{
			UUID:            group.UUID,
			CompanyUUID:     company.UUID,
			ResponsibleUUID: group.ResponsibleUUID,
			Name:            group.Name,
			Desc:            group.Desc,
		}
		responses = append(responses, res)
	}

	return responses, nil
}

func (s *departmentGroupServices) FindByUUID(UUID string) (*response.DepartmentGroupResponse, error) {
	group, err := s.departmentGroupRepo.FindByUUID(UUID)
	if err != nil {
		return nil, err
	}

	company, err := s.companyRepo.GetByID(group.CompanyID)
	if err != nil {
		return nil, err
	}

	res := &response.DepartmentGroupResponse{
		UUID:            group.UUID,
		CompanyUUID:     company.UUID,
		ResponsibleUUID: group.ResponsibleUUID,
		Name:            group.Name,
		Desc:            group.Desc,
	}

	return res, nil
}

func (s *departmentGroupServices) GetDataTable(page, limit int, search, companyUUID string) ([]response.DepartmentGroupResponse, int64, int64, error) {
	offset := (page - 1) * limit

	company, err := s.companyRepo.GetByUUID(companyUUID)
	if err != nil {
		return nil, 0, 0, err
	}

	groups, total, filtered, err := s.departmentGroupRepo.GetDataTable(limit, offset, search, company.ID)
	if err != nil {
		return nil, 0, 0, err
	}

	var responses []response.DepartmentGroupResponse
	for _, group := range groups {
		var responsibleUUID string
		if group.ResponsibleID != 0 {
			responsible, err := s.userRepo.GetByID(group.ResponsibleID)
			if err != nil {
				continue
			}
			responsibleUUID = responsible.UUID
		}

		res := response.DepartmentGroupResponse{
			UUID:            group.UUID,
			CompanyUUID:     company.UUID,
			ResponsibleUUID: responsibleUUID,
			Name:            group.Name,
			Desc:            group.Desc,
		}
		responses = append(responses, res)
	}

	return responses, total, filtered, nil
}

func (s *departmentGroupServices) Update(UUID string, departmentGroupReq request.DepartmentGroupReq) (response.DepartmentGroupResponse, error) {
	if err := pkg.Validate.Struct(departmentGroupReq); err != nil {
		return response.DepartmentGroupResponse{}, err
	}

	departmentGroup, err := s.departmentGroupRepo.FindByUUID(UUID)
	if err != nil {
		return response.DepartmentGroupResponse{}, err
	}

	departmentGroup.Name = departmentGroupReq.Name
	departmentGroup.Desc = departmentGroupReq.Desc

	if departmentGroupReq.ResponsibleUUID != "" {
		responsible, err := s.userRepo.GetByUUID(departmentGroupReq.ResponsibleUUID)
		if err != nil {
			return response.DepartmentGroupResponse{}, err
		}
		departmentGroup.ResponsibleID = responsible.ID
		departmentGroup.ResponsibleUUID = responsible.UUID
	}

	if err := s.departmentGroupRepo.Update(departmentGroup); err != nil {
		return response.DepartmentGroupResponse{}, err
	}

	company, err := s.companyRepo.GetByID(departmentGroup.CompanyID)
	if err != nil {
		return response.DepartmentGroupResponse{}, err
	}

	return response.DepartmentGroupResponse{
		UUID:            departmentGroup.UUID,
		CompanyUUID:     company.UUID,
		ResponsibleUUID: departmentGroup.ResponsibleUUID,
		Name:            departmentGroup.Name,
		Desc:            departmentGroup.Desc,
	}, nil
}

func (s *departmentGroupServices) Delete(UUID string) (response.DepartmentGroupResponse, error) {
	departmentGroup, err := s.departmentGroupRepo.FindByUUID(UUID)
	if err != nil {
		return response.DepartmentGroupResponse{}, err
	}

	if err := s.departmentGroupRepo.Delete(UUID); err != nil {
		return response.DepartmentGroupResponse{}, err
	}

	company, err := s.companyRepo.GetByID(departmentGroup.CompanyID)
	if err != nil {
		return response.DepartmentGroupResponse{}, err
	}

	return response.DepartmentGroupResponse{
		UUID:            departmentGroup.UUID,
		CompanyUUID:     company.UUID,
		ResponsibleUUID: departmentGroup.ResponsibleUUID,
		Name:            departmentGroup.Name,
		Desc:            departmentGroup.Desc,
	}, nil
}
