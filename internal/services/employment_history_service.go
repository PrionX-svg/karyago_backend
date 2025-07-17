package services

import (
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"hris_backend/internal/response"

	"hris_backend/internal/models"
	"hris_backend/internal/repositories"
	"hris_backend/internal/request"
)

type EmploymentHistoryService interface {
	Create(req request.EmploymentHistoryRequest, creatorID uint) (response.EmploymentHistoryResponse, error)
	Update(uuid string, req request.EmploymentHistoryRequest, modifierID uint) (response.EmploymentHistoryResponse, error)
	Delete(uuid string) error
	GetByUUID(uuid string) (*response.EmploymentHistoryResponse, error)
	GetByEmployeeUUID(uuid string) ([]response.EmploymentHistoryResponse, error)
}

type employmentHistoryService struct {
	db           *gorm.DB
	historyRepo  repositories.EmploymentHistoryRepository
	employeeRepo repositories.EmployeeRepository
	roleRepo     repositories.RoleRepositories
	companyRepo  repositories.CompanyRepositories
	branchRepo   repositories.BranchRepository
}

func NewEmploymentHistoryService(
	db *gorm.DB,
	historyRepo repositories.EmploymentHistoryRepository,
	employeeRepo repositories.EmployeeRepository,
	roleRepo repositories.RoleRepositories,
	companyRepo repositories.CompanyRepositories,
	branchRepo repositories.BranchRepository,
) EmploymentHistoryService {
	return &employmentHistoryService{
		db:           db,
		historyRepo:  historyRepo,
		employeeRepo: employeeRepo,
		roleRepo:     roleRepo,
		companyRepo:  companyRepo,
		branchRepo:   branchRepo,
	}
}

func (s *employmentHistoryService) Create(
	req request.EmploymentHistoryRequest,
	creatorID uint,
) (response.EmploymentHistoryResponse, error) {
	var returnValue response.EmploymentHistoryResponse

	err := s.db.Transaction(func(tx *gorm.DB) error {
		employee, err := s.employeeRepo.FindByUUID(req.EmployeeUUID)
		if err != nil {
			return fmt.Errorf("employee not found: %w", err)
		}

		var role *models.Role
		if req.RoleUUID == nil || *req.RoleUUID == "" {
			role, err = s.roleRepo.FindByName("Employee")
			if err != nil {
				return fmt.Errorf("default role 'Employee' not found: %w", err)
			}
		} else {
			role, err = s.roleRepo.FindByUUID(*req.RoleUUID)
			if err != nil {
				return fmt.Errorf("role not found: %w", err)
			}
		}

		var companyID *uint
		var company models.Company
		if req.CompanyUUID != "" {
			companyVal, err := s.companyRepo.GetByUUID(req.CompanyUUID)
			if err != nil {
				return fmt.Errorf("company not found: %w", err)
			}
			company = *companyVal
			companyID = &company.ID
		}

		var branchID *uint
		var branch models.Branch
		if req.BranchUUID != "" {
			branchVal, err := s.branchRepo.FindByUUID(req.BranchUUID)
			if err != nil {
				return fmt.Errorf("branch not found: %w", err)
			}
			branch = *branchVal
			branchID = &branch.ID
		}

		history := &models.EmploymentHistory{
			UUID:       uuid.NewString(),
			EmployeeID: employee.ID,
			RoleID:     role.ID,
			CompanyID:  companyID,
			BranchID:   branchID,
			Position:   req.Position,
			IsPresent:  req.IsPresent,
			StartDate:  req.StartDate,
			EndDate:    req.EndDate,
			Notes:      req.Notes,
			CreatedBy:  creatorID,
			ModifyBy:   creatorID,
		}

		if err := s.historyRepo.Create(history); err != nil {
			return fmt.Errorf("failed to create employment history: %w", err)
		}

		returnValue = response.EmploymentHistoryResponse{
			UUID: history.UUID,
			Employee: response.SimpleEmployeeResponse{
				UUID:     employee.UUID,
				FullName: employee.User.FirstName + " " + employee.User.LastName,
				Email:    employee.User.Email,
			},
			Company: func() *response.SimpleCompanyResponse {
				if companyID != nil {
					return &response.SimpleCompanyResponse{
						UUID: company.UUID,
						Name: company.Name,
					}
				}
				return nil
			}(),
			Branch: func() *response.SimpleBranchResponse {
				if branchID != nil {
					return &response.SimpleBranchResponse{
						UUID: branch.UUID,
						Name: branch.Name,
					}
				}
				return nil
			}(),
			Role: response.SimpleRoleResponse{
				UUID: role.UUID,
				Name: role.Name,
			},
			Position:  history.Position,
			IsPresent: history.IsPresent,
			StartDate: history.StartDate,
			EndDate:   history.EndDate,
			Notes:     history.Notes,
		}

		return nil
	})

	return returnValue, err
}

func (s *employmentHistoryService) Update(uuid string, req request.EmploymentHistoryRequest, modifierID uint) (response.EmploymentHistoryResponse, error) {
	var returnValue response.EmploymentHistoryResponse

	err := s.db.Transaction(func(tx *gorm.DB) error {
		history, err := s.historyRepo.FindByUUID(uuid)
		if err != nil {
			return fmt.Errorf("employment history not found: %w", err)
		}

		employee, err := s.employeeRepo.FindByID(history.EmployeeID)
		if err != nil {
			return fmt.Errorf("employee not found: %w", err)
		}

		var role *models.Role
		if req.RoleUUID == nil || *req.RoleUUID == "" {
			role, err = s.roleRepo.FindByName("Employee")
			if err != nil {
				return fmt.Errorf("default role 'Employee' not found: %w", err)
			}
		} else {
			role, err = s.roleRepo.FindByUUID(*req.RoleUUID)
			if err != nil {
				return fmt.Errorf("role not found: %w", err)
			}
		}

		var (
			company   *models.Company
			companyID *uint
		)
		if req.CompanyUUID != "" {
			company, err = s.companyRepo.GetByUUID(req.CompanyUUID)
			if err != nil {
				return fmt.Errorf("company not found: %w", err)
			}
			companyID = &company.ID
		}

		var (
			branch   *models.Branch
			branchID *uint
		)
		if req.BranchUUID != "" {
			branch, err = s.branchRepo.FindByUUID(req.BranchUUID)
			if err != nil {
				return fmt.Errorf("branch not found: %w", err)
			}
			branchID = &branch.ID
		}

		history.RoleID = role.ID
		history.CompanyID = companyID
		history.BranchID = branchID
		history.Position = req.Position
		history.IsPresent = req.IsPresent
		history.StartDate = req.StartDate
		history.EndDate = req.EndDate
		history.Notes = req.Notes
		history.ModifyBy = modifierID

		if err := s.historyRepo.Update(&history); err != nil {
			return err
		}

		returnValue = response.EmploymentHistoryResponse{
			UUID: history.UUID,
			Employee: response.SimpleEmployeeResponse{
				UUID:     employee.UUID,
				FullName: employee.User.FirstName + " " + employee.User.LastName,
				Email:    employee.User.Email,
			},
			Company: func() *response.SimpleCompanyResponse {
				if company != nil {
					return &response.SimpleCompanyResponse{UUID: company.UUID, Name: company.Name}
				}
				return nil
			}(),
			Branch: func() *response.SimpleBranchResponse {
				if branch != nil {
					return &response.SimpleBranchResponse{UUID: branch.UUID, Name: branch.Name}
				}
				return nil
			}(),
			Role: response.SimpleRoleResponse{
				UUID: role.UUID,
				Name: role.Name,
			},
			Position:  history.Position,
			IsPresent: history.IsPresent,
			StartDate: history.StartDate,
			EndDate:   history.EndDate,
			Notes:     history.Notes,
		}

		return nil
	})

	return returnValue, err
}

func (s *employmentHistoryService) Delete(uuid string) error {
	history, err := s.historyRepo.FindByUUID(uuid)
	if err != nil {
		return fmt.Errorf("employment history not found: %w", err)
	}
	return s.historyRepo.Delete(history.ID)
}

func (s *employmentHistoryService) GetByUUID(uuid string) (*response.EmploymentHistoryResponse, error) {
	history, err := s.historyRepo.FindByUUID(uuid)
	if err != nil {
		return nil, fmt.Errorf("employment history not found: %w", err)
	}

	employee, err := s.employeeRepo.FindByID(history.EmployeeID)
	if err != nil {
		return nil, fmt.Errorf("employee not found: %w", err)
	}

	fullName := ""
	email := ""
	if employee.User != nil {
		fullName = employee.User.FirstName + " " + employee.User.LastName
		email = employee.User.Email
	}

	resp := &response.EmploymentHistoryResponse{
		UUID: history.UUID,
		Employee: response.SimpleEmployeeResponse{
			UUID:     employee.UUID,
			FullName: fullName,
			Email:    email,
		},
		Position:  history.Position,
		IsPresent: history.IsPresent,
		StartDate: history.StartDate,
		EndDate:   history.EndDate,
		Notes:     history.Notes,
	}

	if history.Company != nil {
		resp.Company = &response.SimpleCompanyResponse{
			UUID: history.Company.UUID,
			Name: history.Company.Name,
		}
	}

	if history.Branch != nil {
		resp.Branch = &response.SimpleBranchResponse{
			UUID: history.Branch.UUID,
			Name: history.Branch.Name,
		}
	}

	if history.Role != nil {
		resp.Role = response.SimpleRoleResponse{
			UUID: history.Role.UUID,
			Name: history.Role.Name,
		}
	}

	return resp, nil
}

func (s *employmentHistoryService) GetByEmployeeUUID(uuid string) ([]response.EmploymentHistoryResponse, error) {
	employee, err := s.employeeRepo.FindByUUID(uuid)
	if err != nil {
		return nil, fmt.Errorf("employee not found: %w", err)
	}

	fullName := ""
	email := ""
	if employee.User != nil {
		fullName = employee.User.FirstName + " " + employee.User.LastName
		email = employee.User.Email
	}

	histories, err := s.historyRepo.FindByEmployeeID(employee.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch histories: %w", err)
	}

	var result []response.EmploymentHistoryResponse
	for _, h := range histories {
		resp := response.EmploymentHistoryResponse{
			UUID: h.UUID,
			Employee: response.SimpleEmployeeResponse{
				UUID:     employee.UUID,
				FullName: fullName,
				Email:    email,
			},
			Position:  h.Position,
			IsPresent: h.IsPresent,
			StartDate: h.StartDate,
			EndDate:   h.EndDate,
			Notes:     h.Notes,
		}

		if h.Company != nil {
			resp.Company = &response.SimpleCompanyResponse{
				UUID: h.Company.UUID,
				Name: h.Company.Name,
			}
		}

		if h.Branch != nil {
			resp.Branch = &response.SimpleBranchResponse{
				UUID: h.Branch.UUID,
				Name: h.Branch.Name,
			}
		}

		if h.Role != nil {
			resp.Role = response.SimpleRoleResponse{
				UUID: h.Role.UUID,
				Name: h.Role.Name,
			}
		}

		result = append(result, resp)
	}

	return result, nil
}
