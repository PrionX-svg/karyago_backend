package services

import (
	"errors"
	"fmt"
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
	db             *gorm.DB
	userRepo       repositories.UserRepository
	otpRepo        repositories.OTPRepositories
	roleRepo       repositories.RoleRepositories
	companyRepo    repositories.CompanyRepositories
	branchRepo     repositories.BranchRepository
	employeeRepo   repositories.EmployeeRepository
	departmentRepo repositories.DepartmentRepositories
}

type UserService interface {
	CreateUser(req request.UserEmployeeReq, creatorID uint) (response.UserWithEmployeeResponse, error)
	GetMe(userID uint) (*response.UserWithEmployeeResponse, error)
	GetAllUsers(companyUUID string) ([]response.UserWithEmployeeResponse, error)
	GetUserByUUID(userUUID string, companyUUID string) (*response.UserWithEmployeeResponse, error)
	GetUsersWithEmployeeDataTable(
		page, limit int,
		search, roleUUID, branchUUID, isTerminated, companyUUID string,
	) ([]response.UserWithEmployeeResponse, int64, int64, error)
	UpdateUser(userUUID string, req request.UserEmployeeReq, modifierID uint) (response.UserWithEmployeeResponse, error)
	UpdateEmployeeDepartment(userUUID string, departmentUUID string, modifierID uint) error
	RemoveEmployeeFromDepartment(userUUID string, departmentUUID string, modifierID uint) error
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
	departmentRepo repositories.DepartmentRepositories,
) UserService {
	return &userService{
		db:             db,
		userRepo:       userRepo,
		otpRepo:        otpRepo,
		roleRepo:       roleRepo,
		branchRepo:     branchRepo,
		employeeRepo:   employeeRepo,
		companyRepo:    companyRepo,
		departmentRepo: departmentRepo,
	}
}

func (s *userService) CreateUser(req request.UserEmployeeReq, creatorID uint) (response.UserWithEmployeeResponse, error) {
	var result response.UserWithEmployeeResponse

	err := s.db.Transaction(func(tx *gorm.DB) error {
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
		var branchData *models.Branch
		if req.BranchUUID != "" {
			branch, err := s.branchRepo.FindByUUID(req.BranchUUID)
			if err != nil {
				return fmt.Errorf("branch not found: %w", err)
			}
			branchID = &branch.ID
			branchData = branch
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

		// Applies FIFO
		otp := &models.OTP{
			UUID:      uuid.NewString(),
			Target:    user.Email,
			Code:      pkg.GenerateOTPCode(6),
			Purpose:   "create user & employee",
			IsUsed:    false,
			ExpiresAt: time.Now().Add(5 * time.Minute),
			CreatedAt: time.Now(),
		}

		if err := s.otpRepo.Create(otp); err != nil {
			return fmt.Errorf("failed to create OTP: %w", err)
		}

		fmt.Printf("[INFO] OTP created: UUID=%s, email=%s\n", otp.UUID, user.Email)

		// Kirim email verifikasi (FIFO queue + eksplisit verificationLink)
		verificationLink := fmt.Sprintf("%s/en/activation?token=%s", os.Getenv("FRONTEND_URL"), otp.UUID)

		body := fmt.Sprintf(`
			<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8" />
			<meta name="viewport" content="width=device-width, initial-scale=1.0" />
			<title>Email Verification</title>
			<style>
				body {
					font-family: 'Segoe UI', Roboto, Helvetica, Arial, sans-serif;
					background-color: #f9fafb;
					margin: 0;
					padding: 0;
				}
				.container {
					max-width: 600px;
					margin: 40px auto;
					background: #ffffff;
					border-radius: 8px;
					box-shadow: 0 2px 6px rgba(0,0,0,0.08);
					padding: 40px;
				}
				h2 {
					color: #111827;
					font-size: 22px;
					margin-bottom: 12px;
				}
				p {
					color: #374151;
					font-size: 15px;
					line-height: 1.6;
				}
				.button {
					display: inline-block;
					background-color: #2563eb;
					color: #ffffff !important;
					padding: 12px 24px;
					margin-top: 24px;
					border-radius: 6px;
					text-decoration: none;
					font-weight: 500;
				}
				.footer {
					margin-top: 32px;
					border-top: 1px solid #e5e7eb;
					padding-top: 16px;
					font-size: 13px;
					color: #6b7280;
				}
				.link {
					word-break: break-all;
					color: #2563eb;
					text-decoration: none;
				}
			</style>
		</head>
		<body>
			<div class="container">
				<h2>Hi %s 👋,</h2>
				<p>Welcome to <strong>KARYAGO</strong>! Before you can start exploring, we just need to verify your email address.</p>
				<p>Click the button below to confirm your account:</p>

				<a href="%s" class="button">Verify Your Email</a>

				<p style="margin-top: 24px;">If the button above doesn't work, you can also copy and paste this link into your browser:</p>
				<p><a href="%s" class="link">%s</a></p>

				<p>This verification link will expire in <strong>5 minutes</strong>.</p>

				<p>If you didn't sign up for this account, you can safely ignore this email.</p>

				<div class="footer">
					<p>Best regards,<br/><strong>The KARYAGO Team</strong></p>
					<p>&copy; %d KARYAGO. All rights reserved.</p>
				</div>
			</div>
		</body>
		</html>
		`, user.FirstName, verificationLink, verificationLink, verificationLink, time.Now().Year())

		// Baru di sini email dikirim, setelah link & body siap
		pkg.EnqueueEmail(pkg.EmailJob{
			To:      user.Email,
			Subject: "Email Verification | KARYAGO",
			Body:    body,
		})

		fmt.Printf("[INFO] Enqueued email job for user: %s\n", user.Email)

		result = response.UserWithEmployeeResponse{
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
			Role: struct {
				UUID string `json:"uuid"`
				Name string `json:"name"`
			}{
				UUID: role.UUID,
				Name: role.Name,
			},
			Branch: struct {
				UUID string `json:"uuid"`
				Name string `json:"name"`
			}{
				UUID: req.BranchUUID,
				Name: func() string {
					if branchData != nil {
						return branchData.Name
					}
					return ""
				}(),
			},
			Company: &struct {
				UUID string `json:"uuid"`
				Name string `json:"name"`
			}{
				UUID: req.CompanyUUID,
				Name: company.Name,
			},
		}

		return nil
	})

	return result, err
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
		if err == nil && branch != nil {
			branchUUID = branch.UUID
			branchName = branch.Name
		}
	}

	var companyResp *struct {
		UUID string `json:"uuid"`
		Name string `json:"name"`
	}

	company, err := s.companyRepo.GetByUserID(user.ID)
	if err == nil && company != nil {
		companyResp = &struct {
			UUID string `json:"uuid"`
			Name string `json:"name"`
		}{
			UUID: company.UUID,
			Name: company.Name,
		}
	}

	var department *struct {
		UUID string `json:"uuid"`
		Name string `json:"name"`
	}
	if employee.DepartmentID != nil {
		dept, err := s.departmentRepo.FindByID(*employee.DepartmentID)
		if err == nil && dept != nil {
			department = &struct {
				UUID string `json:"uuid"`
				Name string `json:"name"`
			}{
				UUID: dept.UUID,
				Name: dept.Name,
			}
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
		Role: struct {
			UUID string `json:"uuid"`
			Name string `json:"name"`
		}{
			UUID: role.UUID,
			Name: role.Name,
		},
		Branch: struct {
			UUID string `json:"uuid"`
			Name string `json:"name"`
		}{
			UUID: branchUUID,
			Name: branchName,
		},
		Company:    companyResp,
		Department: department,
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

		var department *struct {
			UUID string `json:"uuid"`
			Name string `json:"name"`
		}
		if emp.DepartmentID != nil {
			dept, err := s.departmentRepo.FindByID(*emp.DepartmentID)
			if err == nil && dept != nil {
				department = &struct {
					UUID string `json:"uuid"`
					Name string `json:"name"`
				}{
					UUID: dept.UUID,
					Name: dept.Name,
				}
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
			Role: struct {
				UUID string `json:"uuid"`
				Name string `json:"name"`
			}{
				UUID: role.UUID,
				Name: role.Name,
			},
			Branch: struct {
				UUID string `json:"uuid"`
				Name string `json:"name"`
			}{UUID: branchUUID, Name: branchName},
			Department:  department,
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

	var department *struct {
		UUID string `json:"uuid"`
		Name string `json:"name"`
	}
	if employee.DepartmentID != nil {
		dept, err := s.departmentRepo.FindByID(*employee.DepartmentID)
		if err == nil && dept != nil {
			department = &struct {
				UUID string `json:"uuid"`
				Name string `json:"name"`
			}{
				UUID: dept.UUID,
				Name: dept.Name,
			}
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
		Role: struct {
			UUID string `json:"uuid"`
			Name string `json:"name"`
		}{
			UUID: role.UUID,
			Name: role.Name,
		},
		Branch: struct {
			UUID string `json:"uuid"`
			Name string `json:"name"`
		}{UUID: branchUUID, Name: branchName},
		Department:  department,
		Termination: termination,
	}, nil
}

func (s *userService) GetUsersWithEmployeeDataTable(
	page, limit int,
	search, roleUUID, branchUUID, isTerminated, companyUUID string,
) ([]response.UserWithEmployeeResponse, int64, int64, error) {
	var (
		users  []models.User
		result []response.UserWithEmployeeResponse
		total  int64
		offset = (page - 1) * limit
	)

	query := s.db.Model(&models.User{}).
		Joins("JOIN employees ON employees.user_id = users.id")

	if companyUUID != "" {
		company, err := s.companyRepo.GetByUUID(companyUUID)
		if err != nil {
			return nil, 0, 0, fmt.Errorf("company not found: %w", err)
		}
		query = query.Where("employees.company_id = ?", company.ID)
	}

	if search != "" {
		likeQuery := "%" + search + "%"
		query = query.Where("users.first_name LIKE ? OR users.last_name LIKE ? OR users.email LIKE ? OR users.phone LIKE ?", likeQuery, likeQuery, likeQuery, likeQuery)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, 0, fmt.Errorf("failed to count total users: %w", err)
	}

	if err := query.Offset(offset).Limit(limit).Order("users.created_at DESC").Find(&users).Error; err != nil {
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

		var companyUUIDStr string
		var companyName string
		if employee.CompanyID != nil {
			company, err := s.companyRepo.GetByID(*employee.CompanyID)
			if err == nil {
				companyUUIDStr = company.UUID
				companyName = company.Name
			}
		}

		var department *struct {
			UUID string `json:"uuid"`
			Name string `json:"name"`
		}
		if employee.DepartmentID != nil {
			dept, err := s.departmentRepo.FindByID(*employee.DepartmentID)
			if err == nil && dept != nil {
				department = &struct {
					UUID string `json:"uuid"`
					Name string `json:"name"`
				}{
					UUID: dept.UUID,
					Name: dept.Name,
				}
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
			Role: struct {
				UUID string `json:"uuid"`
				Name string `json:"name"`
			}{
				UUID: role.UUID,
				Name: role.Name,
			},
			Branch: struct {
				UUID string `json:"uuid"`
				Name string `json:"name"`
			}{UUID: branchUUIDStr, Name: branchName},
			Company: &struct {
				UUID string `json:"uuid"`
				Name string `json:"name"`
			}{UUID: companyUUIDStr, Name: companyName},
			Department:  department,
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

		employee.TerminatedAt = nil
		employee.TerminationReason = nil
		employee.IsFreelance = req.IsFreelance
		employee.ModifyBy = modifierID

		if err := s.employeeRepo.Update(employee); err != nil {
			return fmt.Errorf("failed to rehire employee: %w", err)
		}

		return nil
	})
}

func (s *userService) UpdateUser(userUUID string, req request.UserEmployeeReq, modifierID uint) (response.UserWithEmployeeResponse, error) {
	var result response.UserWithEmployeeResponse

	err := s.db.Transaction(func(tx *gorm.DB) error {
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

		company, err := s.companyRepo.GetByUUID(req.CompanyUUID)
		if err != nil {
			return fmt.Errorf("company not found: %w", err)
		}

		var role *models.Role
		if req.RoleUUID == "" {
			role, err = s.roleRepo.FindByNameAndCompanyID("employee", req.CompanyUUID)
			if err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					newRole := &models.Role{
						UUID:      uuid.NewString(),
						Name:      "employee",
						CompanyID: &company.ID,
						CreatedBy: modifierID,
						ModifyBy:  modifierID,
					}
					if err := s.roleRepo.Create(newRole); err != nil {
						return fmt.Errorf("failed to create role 'employee': %w", err)
					}
					role = newRole
				} else {
					return fmt.Errorf("failed to fetch fallback role: %w", err)
				}
			}
		} else {
			role, err = s.roleRepo.FindByUUID(req.RoleUUID)
			if err != nil {
				return fmt.Errorf("role not found: %w", err)
			}
		}

		var branchID *uint
		var branchData *models.Branch
		if req.BranchUUID != "" {
			branch, err := s.branchRepo.FindByUUID(req.BranchUUID)
			if err != nil {
				return fmt.Errorf("branch not found: %w", err)
			}
			branchID = &branch.ID
			branchData = branch
		}

		var department *struct {
			UUID string `json:"uuid"`
			Name string `json:"name"`
		}
		if employee.DepartmentID != nil {
			dept, err := s.departmentRepo.FindByID(*employee.DepartmentID)
			if err == nil && dept != nil {
				department = &struct {
					UUID string `json:"uuid"`
					Name string `json:"name"`
				}{
					UUID: dept.UUID,
					Name: dept.Name,
				}
			}
		}

		//Simpan email lama untuk deteksi perubahan
		oldEmail := user.Email

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

		// ✅ Jika email berubah → kirim ulang link verifikasi (styled HTML)
		if oldEmail != req.Email {
			otp := &models.OTP{
				UUID:      uuid.NewString(),
				Target:    req.Email,
				Code:      pkg.GenerateOTPCode(6),
				Purpose:   "update email verification",
				IsUsed:    false,
				ExpiresAt: time.Now().Add(5 * time.Minute),
				CreatedAt: time.Now(),
			}

			fmt.Printf("[INFO] OTP created (email update): UUID=%s, new_email=%s\n", otp.UUID, req.Email)

			if err := s.otpRepo.Create(otp); err != nil {
				return fmt.Errorf("failed to create OTP for new email: %w", err)
			}

			verificationLink := fmt.Sprintf("%s/en/activation?token=%s", os.Getenv("FRONTEND_URL"), otp.UUID)

			body := fmt.Sprintf(`
			<!DOCTYPE html>
			<html lang="en">
			<head>
				<meta charset="UTF-8" />
				<meta name="viewport" content="width=device-width, initial-scale=1.0" />
				<title>Verify Your New Email</title>
				<style>
					body {
						font-family: 'Segoe UI', Roboto, Helvetica, Arial, sans-serif;
						background-color: #f9fafb;
						margin: 0;
						padding: 0;
					}
					.container {
						max-width: 600px;
						margin: 40px auto;
						background: #ffffff;
						border-radius: 8px;
						box-shadow: 0 2px 6px rgba(0,0,0,0.08);
						padding: 40px;
					}
					h2 {
						color: #111827;
						font-size: 22px;
						margin-bottom: 12px;
					}
					p {
						color: #374151;
						font-size: 15px;
						line-height: 1.6;
					}
					.button {
						display: inline-block;
						background-color: #2563eb;
						color: #ffffff !important;
						padding: 12px 24px;
						margin-top: 24px;
						border-radius: 6px;
						text-decoration: none;
						font-weight: 500;
					}
					.footer {
						margin-top: 32px;
						border-top: 1px solid #e5e7eb;
						padding-top: 16px;
						font-size: 13px;
						color: #6b7280;
					}
					.link {
						word-break: break-all;
						color: #2563eb;
						text-decoration: none;
					}
				</style>
			</head>
			<body>
				<div class="container">
					<h2>Hi %s 👋,</h2>
					<p>We noticed you've updated your email address for <strong>KARYAGO</strong>.</p>
					<p>To complete this change, please verify your new email by clicking the button below:</p>

					<a href="%s" class="button">Verify New Email</a>

					<p style="margin-top: 24px;">If the button above doesn't work, you can also copy and paste this link into your browser:</p>
					<p><a href="%s" class="link">%s</a></p>

					<p>This verification link will expire in <strong>5 minutes</strong>.</p>

					<p>If you didn’t request this change, please contact your administrator.</p>

					<div class="footer">
						<p>Best regards,<br/><strong>The KARYAGO Team</strong></p>
						<p>&copy; %d KARYAGO. All rights reserved.</p>
					</div>
				</div>
			</body>
			</html>
			`, user.FirstName, verificationLink, verificationLink, verificationLink, time.Now().Year())

			pkg.EnqueueEmail(pkg.EmailJob{
				To:      req.Email,
				Subject: "Verify Your New Email | KARYAGO",
				Body:    body,
			})

			fmt.Printf("[INFO] Enqueued email verification for updated email: %s\n", req.Email)
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

		result = response.UserWithEmployeeResponse{
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
			Role: struct {
				UUID string `json:"uuid"`
				Name string `json:"name"`
			}{
				UUID: role.UUID,
				Name: role.Name,
			},
			Branch: struct {
				UUID string `json:"uuid"`
				Name string `json:"name"`
			}{
				UUID: req.BranchUUID,
				Name: func() string {
					if branchData != nil {
						return branchData.Name
					}
					return ""
				}(),
			},
			Company: &struct {
				UUID string `json:"uuid"`
				Name string `json:"name"`
			}{
				UUID: req.CompanyUUID,
				Name: company.Name,
			},
			Department: department,
		}

		return nil
	})

	return result, err
}

func (s *userService) UpdateEmployeeDepartment(userUUID string, departmentUUID string, modifierID uint) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		employee, err := s.employeeRepo.FindByUUID(userUUID)
		if err != nil {
			return fmt.Errorf("employee not found: %w", err)
		}

		if employee.TerminatedAt != nil {
			return fmt.Errorf("cannot update department for terminated employee")
		}

		var departmentID *uint
		if departmentUUID != "" {
			department, err := s.departmentRepo.FindByUUID(departmentUUID)
			if err != nil {
				return fmt.Errorf("department not found: %w", err)
			}
			departmentID = &department.ID
		}

		employee.DepartmentID = departmentID
		employee.ModifyBy = modifierID

		if err := s.employeeRepo.Update(employee); err != nil {
			return fmt.Errorf("failed to update employee department: %w", err)
		}

		return nil
	})
}

func (s *userService) RemoveEmployeeFromDepartment(userUUID string, departmentUUID string, modifierID uint) error {
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
			return fmt.Errorf("cannot update department for terminated employee")
		}

		if departmentUUID != "" {
			_, err := s.departmentRepo.FindByUUID(departmentUUID)
			if err != nil {
				return fmt.Errorf("department not found: %w", err)
			}
		}

		employee.DepartmentID = nil
		employee.ModifyBy = modifierID

		if err := s.employeeRepo.Update(employee); err != nil {
			return fmt.Errorf("failed to remove employee from department: %w", err)
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
