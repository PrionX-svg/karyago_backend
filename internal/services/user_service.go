package services

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"hris_backend/internal/models"
	"hris_backend/internal/repositories"
	"hris_backend/internal/request"
	"hris_backend/internal/response"
	"hris_backend/pkg"
)

type userService struct {
	db           *gorm.DB
	userRepo     repositories.UserRepository
	otpRepo      repositories.OTPRepositories
	roleRepo     repositories.RoleRepositories
	companyRepo  repositories.CompanyRepositories
	branchRepo   repositories.BranchRepository
	employeeRepo repositories.EmployeeRepository
}

type UserService interface {
	CreateUser(req request.UserEmployeeReq, creatorID uint) error
	GetMe(userID uint) (*response.UserWithEmployeeResponse, error)
	GetAllUsers(companyUUID string) ([]response.UserWithEmployeeResponse, error)
	GetUserByUUID(userUUID string, companyUUID string) (*response.UserWithEmployeeResponse, error)
	GetUsersWithEmployeeDataTable(page, limit int, search, roleUUID, branchUUID, isTerminated string) ([]response.UserWithEmployeeResponse, int64, int64, error)
	UpdateUser(userUUID string, req request.UserEmployeeReq, modifierID uint) error
	DeleteUser(userUUID, companyUUID, reason string) error
	RehireEmployee(userUUID, companyUUID string, req request.RehireEmployeeReq, modifierID uint) error
}

func NewUserService(
	db *gorm.DB,
	userRepo repositories.UserRepository,
	otpRepo repositories.OTPRepositories,
	roleRepo repositories.RoleRepositories,
	branchRepo repositories.BranchRepository,
	employeeRepo repositories.EmployeeRepository,
	companyRepo repositories.CompanyRepositories,
) UserService {
	return &userService{
		db:           db,
		userRepo:     userRepo,
		otpRepo:      otpRepo,
		roleRepo:     roleRepo,
		branchRepo:   branchRepo,
		employeeRepo: employeeRepo,
		companyRepo:  companyRepo,
	}
}

func (s *userService) CreateUser(req request.UserEmployeeReq, creatorID uint) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		var err error
		var role *models.Role

		company, err := s.companyRepo.GetByUUID(req.CompanyUUID)
		if err != nil {
			return fmt.Errorf("company not found: %w", err)
		}

		if req.RoleUUID == "" {
			role, err = s.roleRepo.FindByNameAndCompanyID("employee", req.CompanyUUID)
			if err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					newRole := &models.Role{
						UUID:      uuid.NewString(),
						Name:      "employee",
						CompanyID: &company.ID,
						CreatedBy: creatorID,
						ModifyBy:  creatorID,
					}
					if err := s.roleRepo.Create(newRole); err != nil {
						return fmt.Errorf("failed to create role 'employee': %w", err)
					}
					role = newRole
				} else {
					return fmt.Errorf("failed to check role: %w", err)
				}
			}
		} else {
			role, err = s.roleRepo.FindByUUID(req.RoleUUID)
			if err != nil {
				return fmt.Errorf("role not found: %w", err)
			}
		}

		var branchID *uint
		if req.BranchUUID != "" {
			branch, err := s.branchRepo.FindByUUID(req.BranchUUID)
			if err != nil {
				return fmt.Errorf("branch not found: %w", err)
			}
			branchID = &branch.ID
		}

		hashedPassword, err := pkg.HashPassword(req.Password)
		if err != nil {
			return fmt.Errorf("failed to hash password: %w", err)
		}

		user := &models.User{
			UUID:      uuid.NewString(),
			FirstName: req.FirstName,
			LastName:  req.LastName,
			Phone:     req.Phone,
			Email:     req.Email,
			Password:  hashedPassword,
			DOB:       &req.DOB,
			Gender:    &req.Gender,
			CreatedBy: creatorID,
			ModifyBy:  creatorID,
		}

		if err := s.userRepo.Create(user); err != nil {
			return fmt.Errorf("failed to create user: %w", err)
		}

		employee := &models.Employee{
			UUID:        uuid.NewString(),
			UserID:      user.ID,
			RoleID:      role.ID,
			CompanyID:   &company.ID,
			BranchID:    branchID,
			IsFreelance: req.IsFreelance,
			CreatedBy:   creatorID,
			ModifyBy:    creatorID,
		}

		if err := s.employeeRepo.Create(employee); err != nil {
			return fmt.Errorf("failed to create employee: %w", err)
		}

		otp := &models.OTP{
			UUID:      uuid.NewString(),
			Target:    user.Email,
			Code:      pkg.GenerateOTPCode(6),
			Purpose:   "create user & employee",
			IsUsed:    false,
			ExpiresAt: time.Now().Add(5 * time.Minute),
		}

		if err := s.otpRepo.Create(otp); err != nil {
			return fmt.Errorf("failed to create OTP: %w", err)
		}

		go func(email, name, token string) {
			verificationLink := fmt.Sprintf("%s/en/activation?token=%s", os.Getenv("FRONTEND_URL"), token)
			body := fmt.Sprintf(`
				<html>
					<body>
						<p>Hi %s,</p>
						<p>Welcome! Please verify your email address to activate your account:</p>
						<p><a href="%s">Verify your Email</a></p>
						<p>This link will expire in 5 minutes.</p>
						<p>If you didn't register, you can ignore this email.</p>
						<br/>
						<p>Regards,<br/>The Team</p>
					</body>
				</html>
			`, name, verificationLink)

			if err := pkg.SendEmail(email, "Email Verification", body); err != nil {
				log.Printf("failed to send email: %v", err)
			}
		}(user.Email, user.FirstName, otp.UUID)

		return nil
	})
}

func (s *userService) GetMe(userID uint) (*response.UserWithEmployeeResponse, error) {
	user, err := s.userRepo.GetByID(userID)
	if err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}

	employee, err := s.employeeRepo.FindByUserID(user.ID)
	if err != nil {
		return nil, fmt.Errorf("employee not found: %w", err)
	}

	role, err := s.roleRepo.FindByID(employee.RoleID)
	if err != nil {
		return nil, fmt.Errorf("role not found: %w", err)
	}

	var branchUUID, branchName string
	if employee.BranchID != nil {
		branch, err := s.branchRepo.FindByID(*employee.BranchID)
		if err == nil {
			branchUUID = branch.UUID
			branchName = branch.Name
		}
	}

	return &response.UserWithEmployeeResponse{
		UserUUID:     user.UUID,
		EmployeeUUID: employee.UUID,
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		FullName:     user.FirstName + " " + user.LastName,
		Email:        user.Email,
		Phone:        user.Phone,
		Gender:       user.Gender,
		DOB:          user.DOB,
		IsFreelance:  employee.IsFreelance,
		Role:         role.Name,
		Branch: struct {
			UUID string `json:"uuid"`
			Name string `json:"name"`
		}{UUID: branchUUID, Name: branchName},
	}, nil
}

func (s *userService) GetAllUsers(companyUUID string) ([]response.UserWithEmployeeResponse, error) {
	company, err := s.companyRepo.GetByUUID(companyUUID)
	if err != nil {
		return nil, fmt.Errorf("company not found: %w", err)
	}

	employees, err := s.employeeRepo.FindByCompanyID(company.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to get employees: %w", err)
	}

	var result []response.UserWithEmployeeResponse

	for _, emp := range employees {
		if emp.TerminatedAt != nil {
			continue
		}

		user, err := s.userRepo.GetByID(emp.UserID)
		if err != nil {
			continue
		}

		role, err := s.roleRepo.FindByID(emp.RoleID)
		if err != nil {
			continue
		}

		var branchUUID, branchName string
		if emp.BranchID != nil {
			branch, err := s.branchRepo.FindByID(*emp.BranchID)
			if err == nil {
				branchUUID = branch.UUID
				branchName = branch.Name
			}
		}

		var termination *struct {
			Reason string     `json:"reason"`
			Date   *time.Time `json:"date"`
		}
		if emp.TerminatedAt != nil || emp.TerminationReason != nil {
			termination = &struct {
				Reason string     `json:"reason"`
				Date   *time.Time `json:"date"`
			}{
				Reason: pkg.DerefString(emp.TerminationReason),
				Date:   emp.TerminatedAt,
			}
		}

		result = append(result, response.UserWithEmployeeResponse{
			UserUUID:     user.UUID,
			EmployeeUUID: emp.UUID,
			FirstName:    user.FirstName,
			LastName:     user.LastName,
			FullName:     user.FirstName + " " + user.LastName,
			Email:        user.Email,
			Phone:        user.Phone,
			Gender:       user.Gender,
			DOB:          user.DOB,
			IsFreelance:  emp.IsFreelance,
			Role:         role.Name,
			Branch: struct {
				UUID string `json:"uuid"`
				Name string `json:"name"`
			}{UUID: branchUUID, Name: branchName},
			Termination: termination,
		})
	}

	return result, nil
}

func (s *userService) GetUserByUUID(userUUID string, companyUUID string) (*response.UserWithEmployeeResponse, error) {
	user, err := s.userRepo.GetByUUID(userUUID)
	if err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}

	company, err := s.companyRepo.GetByUUID(companyUUID)
	if err != nil {
		return nil, fmt.Errorf("company not found: %w", err)
	}

	employee, err := s.employeeRepo.FindByUserIDAndCompanyID(user.ID, company.ID)
	if err != nil {
		return nil, fmt.Errorf("employee not found in this company: %w", err)
	}

	role, err := s.roleRepo.FindByID(employee.RoleID)
	if err != nil {
		return nil, fmt.Errorf("role not found: %w", err)
	}

	var branchUUID, branchName string
	if employee.BranchID != nil {
		branch, err := s.branchRepo.FindByID(*employee.BranchID)
		if err == nil {
			branchUUID = branch.UUID
			branchName = branch.Name
		}
	}

	var termination *struct {
		Reason string     `json:"reason"`
		Date   *time.Time `json:"date"`
	}
	if employee.TerminatedAt != nil || employee.TerminationReason != nil {
		termination = &struct {
			Reason string     `json:"reason"`
			Date   *time.Time `json:"date"`
		}{
			Reason: pkg.DerefString(employee.TerminationReason),
			Date:   employee.TerminatedAt,
		}
	}

	return &response.UserWithEmployeeResponse{
		UserUUID:     user.UUID,
		EmployeeUUID: employee.UUID,
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		FullName:     user.FirstName + " " + user.LastName,
		Email:        user.Email,
		Phone:        user.Phone,
		Gender:       user.Gender,
		DOB:          user.DOB,
		IsFreelance:  employee.IsFreelance,
		Role:         role.Name,
		Branch: struct {
			UUID string `json:"uuid"`
			Name string `json:"name"`
		}{UUID: branchUUID, Name: branchName},
		Termination: termination,
	}, nil
}

func (s *userService) GetUsersWithEmployeeDataTable(page, limit int, search, roleUUID, branchUUID, isTerminated string) ([]response.UserWithEmployeeResponse, int64, int64, error) {
	var (
		users  []models.User
		result []response.UserWithEmployeeResponse
		total  int64
		offset = (page - 1) * limit
	)

	query := s.db.Model(&models.User{})
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, 0, fmt.Errorf("failed to count total users: %w", err)
	}

	if search != "" {
		likeQuery := "%" + search + "%"
		query = query.Where("first_name LIKE ? OR last_name LIKE ? OR email LIKE ? OR phone LIKE ?", likeQuery, likeQuery, likeQuery, likeQuery)
	}

	if err := query.Offset(offset).Limit(limit).Order("created_at DESC").Find(&users).Error; err != nil {
		return nil, 0, 0, fmt.Errorf("failed to query users: %w", err)
	}

	for _, user := range users {
		employee, err := s.employeeRepo.FindByUserID(user.ID)
		if err != nil {
			continue
		}

		if isTerminated == "true" && employee.TerminatedAt == nil {
			continue
		}
		if isTerminated == "false" && employee.TerminatedAt != nil {
			continue
		}

		if roleUUID != "" {
			roleFilter, err := s.roleRepo.FindByUUID(roleUUID)
			if err != nil || employee.RoleID != roleFilter.ID {
				continue
			}
		}

		if branchUUID != "" {
			branchFilter, err := s.branchRepo.FindByUUID(branchUUID)
			if err != nil || (employee.BranchID == nil || *employee.BranchID != branchFilter.ID) {
				continue
			}
		}

		role, err := s.roleRepo.FindByID(employee.RoleID)
		if err != nil {
			continue
		}

		var branchUUIDStr, branchName string
		if employee.BranchID != nil {
			branch, err := s.branchRepo.FindByID(*employee.BranchID)
			if err == nil {
				branchUUIDStr = branch.UUID
				branchName = branch.Name
			}
		}

		var termination *struct {
			Reason string     `json:"reason"`
			Date   *time.Time `json:"date"`
		}
		if employee.TerminatedAt != nil || employee.TerminationReason != nil {
			termination = &struct {
				Reason string     `json:"reason"`
				Date   *time.Time `json:"date"`
			}{
				Reason: pkg.DerefString(employee.TerminationReason),
				Date:   employee.TerminatedAt,
			}
		}

		result = append(result, response.UserWithEmployeeResponse{
			UserUUID:     user.UUID,
			EmployeeUUID: employee.UUID,
			FirstName:    user.FirstName,
			LastName:     user.LastName,
			FullName:     user.FirstName + " " + user.LastName,
			Email:        user.Email,
			Phone:        user.Phone,
			Gender:       user.Gender,
			DOB:          user.DOB,
			IsFreelance:  employee.IsFreelance,
			Role:         role.Name,
			Branch: struct {
				UUID string `json:"uuid"`
				Name string `json:"name"`
			}{UUID: branchUUIDStr, Name: branchName},
			Termination: termination,
		})
	}

	return result, total, int64(len(result)), nil
}

func (s *userService) RehireEmployee(userUUID, companyUUID string, req request.RehireEmployeeReq, modifierID uint) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		user, err := s.userRepo.GetByUUID(userUUID)
		if err != nil {
			return fmt.Errorf("user not found: %w", err)
		}

		company, err := s.companyRepo.GetByUUID(companyUUID)
		if err != nil {
			return fmt.Errorf("company not found: %w", err)
		}

		employee, err := s.employeeRepo.FindTerminatedByUserIDAndCompanyID(user.ID, company.ID)
		if err != nil {
			return fmt.Errorf("terminated employee not found: %w", err)
		}

		role, err := s.roleRepo.FindByUUID(req.RoleUUID)
		if err != nil {
			return fmt.Errorf("role not found: %w", err)
		}

		var branchID *uint
		if req.BranchUUID != "" {
			branch, err := s.branchRepo.FindByUUID(req.BranchUUID)
			if err != nil {
				return fmt.Errorf("branch not found: %w", err)
			}
			branchID = &branch.ID
		}

		employee.TerminatedAt = nil
		employee.TerminationReason = nil
		employee.RoleID = role.ID
		employee.BranchID = branchID
		employee.IsFreelance = req.IsFreelance
		employee.ModifyBy = modifierID

		if err := s.employeeRepo.Update(employee); err != nil {
			return fmt.Errorf("failed to rehire employee: %w", err)
		}

		return nil
	})
}

func (s *userService) UpdateUser(userUUID string, req request.UserEmployeeReq, modifierID uint) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		user, err := s.userRepo.GetByUUID(userUUID)
		if err != nil {
			return fmt.Errorf("user not found: %w", err)
		}

		employee, err := s.employeeRepo.FindByUserID(user.ID)
		if err != nil {
			return fmt.Errorf("employee not found: %w", err)
		}

		if employee.TerminatedAt != nil {
			return fmt.Errorf("cannot update employee that has been terminated")
		}

		role, err := s.roleRepo.FindByUUID(req.RoleUUID)
		if err != nil {
			return fmt.Errorf("role not found: %w", err)
		}

		var branchID *uint
		if req.BranchUUID != "" {
			branch, err := s.branchRepo.FindByUUID(req.BranchUUID)
			if err != nil {
				return fmt.Errorf("branch not found: %w", err)
			}
			branchID = &branch.ID
		}

		user.FirstName = req.FirstName
		user.LastName = req.LastName
		user.Phone = req.Phone
		user.Email = req.Email
		user.DOB = &req.DOB
		user.Gender = &req.Gender
		user.ModifyBy = modifierID

		if req.Password != "" {
			hashedPassword, err := pkg.HashPassword(req.Password)
			if err != nil {
				return fmt.Errorf("failed to hash password: %w", err)
			}
			user.Password = hashedPassword
		}

		employee.RoleID = role.ID
		employee.BranchID = branchID
		employee.IsFreelance = req.IsFreelance
		employee.ModifyBy = modifierID

		if err := s.userRepo.Update(&user); err != nil {
			return fmt.Errorf("failed to update user: %w", err)
		}
		if err := s.employeeRepo.Update(employee); err != nil {
			return fmt.Errorf("failed to update employee: %w", err)
		}

		return nil
	})
}

func (s *userService) DeleteUser(userUUID, companyUUID, reason string) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		user, err := s.userRepo.GetByUUID(userUUID)
		if err != nil {
			return fmt.Errorf("user not found: %w", err)
		}

		company, err := s.companyRepo.GetByUUID(companyUUID)
		if err != nil {
			return fmt.Errorf("company not found: %w", err)
		}

		employee, err := s.employeeRepo.FindByUserIDAndCompanyID(user.ID, company.ID)
		if err != nil {
			return fmt.Errorf("employee not found in company: %w", err)
		}

		now := time.Now()
		employee.TerminatedAt = &now
		if reason != "" {
			employee.TerminationReason = &reason
		}
		employee.ModifyBy = user.ID

		if err := s.employeeRepo.Update(employee); err != nil {
			return fmt.Errorf("failed to update employee termination: %w", err)
		}

		return nil
	})
}
