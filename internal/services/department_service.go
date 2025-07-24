package services

import (
	"hris_backend/internal/models"
	"hris_backend/internal/repositories"
	"hris_backend/internal/request"
	"hris_backend/internal/response"
	"hris_backend/pkg"

	"github.com/google/uuid"
)

type DepartmentServices interface {
	Create(req request.DepartmentReq) (response.DepartmentResponse, error)
	GetAll() ([]response.DepartmentResponse, error)
	FindByUUID(UUID string) (*response.DepartmentResponse, error)
	GetDataTable(page, limit int, search, companyUUID, departmentGroupUUID string) ([]response.DepartmentResponse, int64, int64, error)
	Update(UUID string, req request.DepartmentReq) (response.DepartmentResponse, error)
	Delete(UUID string) (response.DepartmentResponse, error)
}

type departmentServices struct {
	departmentRepo      repositories.DepartmentRepositories
	departmentGroupRepo repositories.DepartmentGroupRepositories
	employeeRepo        repositories.EmployeeRepository
}

func NewDepartmentServices(departmentRepo repositories.DepartmentRepositories, departmentGroupRepo repositories.DepartmentGroupRepositories, employeeRepo repositories.EmployeeRepository) DepartmentServices {
	return &departmentServices{
		departmentRepo:      departmentRepo,
		departmentGroupRepo: departmentGroupRepo,
		employeeRepo:        employeeRepo,
	}
}

func (s *departmentServices) Create(req request.DepartmentReq) (response.DepartmentResponse, error) {
	if err := pkg.Validate.Struct(req); err != nil {
		return response.DepartmentResponse{}, err
	}

	group, err := s.departmentGroupRepo.FindByUUID(req.DepartmentGroupUUID)
	if err != nil {
		return response.DepartmentResponse{}, err
	}

	newDepartment := models.Department{
		UUID:              uuid.NewString(),
		DepartmentGroupID: group.ID,
		Name:              req.Name,
		Description:       req.Description,
		CreatedBy:         req.CreatedBy,
		ModifyBy:          req.CreatedBy,
	}

	if err := s.departmentRepo.Create(&newDepartment); err != nil {
		return response.DepartmentResponse{}, err
	}

	return response.DepartmentResponse{
		UUID:        newDepartment.UUID,
		Name:        newDepartment.Name,
		Description: newDepartment.Description,
		Group: response.DepartmentGroupMiniResponse{
			UUID: group.UUID,
			Name: group.Name,
		},
		Employees: []response.UserMiniResponse{},
	}, nil
}

func (s *departmentServices) GetAll() ([]response.DepartmentResponse, error) {
	departments, err := s.departmentRepo.GetAll()
	if err != nil {
		return nil, err
	}

	var responses []response.DepartmentResponse
	for _, dept := range departments {
		group, err := s.departmentGroupRepo.FindByUUIDFromID(dept.DepartmentGroupID)
		if err != nil {
			return nil, err
		}

		employees, err := s.employeeRepo.FindByDepartmentID(dept.ID)
		if err != nil {
			return nil, err
		}

		var employeeResponses []response.UserMiniResponse
		for _, emp := range employees {
			employeeResponses = append(employeeResponses, response.UserMiniResponse{
				UUID:  emp.User.UUID,
				Name:  emp.User.FirstName + " " + emp.User.LastName,
				Email: emp.User.Email,
			})
		}

		responses = append(responses, response.DepartmentResponse{
			UUID:        dept.UUID,
			Name:        dept.Name,
			Description: dept.Description,
			Group: response.DepartmentGroupMiniResponse{
				UUID: group.UUID,
				Name: group.Name,
			},
			Employees: employeeResponses,
		})
	}

	return responses, nil
}

func (s *departmentServices) FindByUUID(UUID string) (*response.DepartmentResponse, error) {
	dept, err := s.departmentRepo.FindByUUID(UUID)
	if err != nil {
		return nil, err
	}

	group, err := s.departmentGroupRepo.FindByUUIDFromID(dept.DepartmentGroupID)
	if err != nil {
		return nil, err
	}

	employees, err := s.employeeRepo.FindByDepartmentID(dept.ID)
	if err != nil {
		return nil, err
	}

	var employeeResponses []response.UserMiniResponse
	for _, emp := range employees {
		employeeResponses = append(employeeResponses, response.UserMiniResponse{
			UUID:  emp.User.UUID,
			Name:  emp.User.FirstName + " " + emp.User.LastName,
			Email: emp.User.Email,
		})
	}

	res := &response.DepartmentResponse{
		UUID:        dept.UUID,
		Name:        dept.Name,
		Description: dept.Description,
		Group: response.DepartmentGroupMiniResponse{
			UUID: group.UUID,
			Name: group.Name,
		},
		Employees: employeeResponses,
	}

	return res, nil
}

func (s *departmentServices) GetDataTable(page, limit int, search, companyUUID, departmentGroupUUID string) ([]response.DepartmentResponse, int64, int64, error) {
	offset := (page - 1) * limit

	company, err := s.departmentGroupRepo.GetCompanyByUUID(companyUUID)
	if err != nil {
		return nil, 0, 0, err
	}

	var departmentGroupID *uint
	if departmentGroupUUID != "" {
		group, err := s.departmentGroupRepo.FindByUUID(departmentGroupUUID)
		if err != nil {
			return nil, 0, 0, err
		}
		departmentGroupID = &group.ID
	}

	departments, total, filtered, err := s.departmentRepo.GetDataTable(limit, offset, search, company.ID, departmentGroupID)
	if err != nil {
		return nil, 0, 0, err
	}

	var responses []response.DepartmentResponse
	for _, dept := range departments {
		group, err := s.departmentGroupRepo.FindByUUIDFromID(dept.DepartmentGroupID)
		if err != nil {
			continue
		}

		employees, err := s.employeeRepo.FindByDepartmentID(dept.ID)
		if err != nil {
			continue
		}

		var employeeResponses []response.UserMiniResponse
		for _, emp := range employees {
			employeeResponses = append(employeeResponses, response.UserMiniResponse{
				UUID:  emp.User.UUID,
				Name:  emp.User.FirstName + " " + emp.User.LastName,
				Email: emp.User.Email,
			})
		}

		res := response.DepartmentResponse{
			UUID:        dept.UUID,
			Name:        dept.Name,
			Description: dept.Description,
			Group: response.DepartmentGroupMiniResponse{
				UUID: group.UUID,
				Name: group.Name,
			},
			Employees: employeeResponses,
		}
		responses = append(responses, res)
	}

	return responses, total, filtered, nil
}

func (s *departmentServices) Update(UUID string, req request.DepartmentReq) (response.DepartmentResponse, error) {
	if err := pkg.Validate.Struct(req); err != nil {
		return response.DepartmentResponse{}, err
	}

	dept, err := s.departmentRepo.FindByUUID(UUID)
	if err != nil {
		return response.DepartmentResponse{}, err
	}

	group, err := s.departmentGroupRepo.FindByUUID(req.DepartmentGroupUUID)
	if err != nil {
		return response.DepartmentResponse{}, err
	}

	dept.Name = req.Name
	dept.Description = req.Description
	dept.DepartmentGroupID = group.ID
	dept.ModifyBy = req.ModifyBy

	if err := s.departmentRepo.Update(dept); err != nil {
		return response.DepartmentResponse{}, err
	}

	return response.DepartmentResponse{
		UUID:        dept.UUID,
		Name:        dept.Name,
		Description: dept.Description,
		Group: response.DepartmentGroupMiniResponse{
			UUID: group.UUID,
			Name: group.Name,
		},
		Employees: []response.UserMiniResponse{},
	}, nil
}

func (s *departmentServices) Delete(UUID string) (response.DepartmentResponse, error) {
	dept, err := s.departmentRepo.FindByUUID(UUID)
	if err != nil {
		return response.DepartmentResponse{}, err
	}

	group, err := s.departmentGroupRepo.FindByUUIDFromID(dept.DepartmentGroupID)
	if err != nil {
		return response.DepartmentResponse{}, err
	}

	if err := s.departmentRepo.Delete(UUID); err != nil {
		return response.DepartmentResponse{}, err
	}

	return response.DepartmentResponse{
		UUID:        dept.UUID,
		Name:        dept.Name,
		Description: dept.Description,
		Group: response.DepartmentGroupMiniResponse{
			UUID: group.UUID,
			Name: group.Name,
		},
		Employees: []response.UserMiniResponse{},
	}, nil
}
