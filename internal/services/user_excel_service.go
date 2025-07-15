package services

import (
	"bytes"
	"fmt"
	"github.com/xuri/excelize/v2"
	"hris_backend/internal/repositories"
	"hris_backend/internal/request"
	"hris_backend/pkg"
	"io"
	"mime/multipart"
	"time"
)

type UserExcelService interface {
	ExportUsersToExcel(companyUUID string) ([]byte, error)
	ImportUsersFromExcel(file multipart.File, creatorID uint, companyUUID string) error
}

type userExcelService struct {
	userService  UserService
	companyRepo  repositories.CompanyRepositories
	roleRepo     repositories.RoleRepositories
	branchRepo   repositories.BranchRepository
	userRepo     repositories.UserRepository
	employeeRepo repositories.EmployeeRepository
}

func NewUserExcelService(
	userService UserService,
	companyRepo repositories.CompanyRepositories,
	roleRepo repositories.RoleRepositories,
	branchRepo repositories.BranchRepository,
	userRepo repositories.UserRepository,
	employeeRepo repositories.EmployeeRepository,
) UserExcelService {
	return &userExcelService{
		userService:  userService,
		companyRepo:  companyRepo,
		roleRepo:     roleRepo,
		branchRepo:   branchRepo,
		userRepo:     userRepo,
		employeeRepo: employeeRepo,
	}
}

func (s *userExcelService) ExportUsersToExcel(companyUUID string) ([]byte, error) {
	users, err := s.userService.GetAllUsers(companyUUID)
	if err != nil {
		return nil, fmt.Errorf("failed to get users: %w", err)
	}

	f := excelize.NewFile()
	sheet := "Users"
	_, err = f.NewSheet(sheet)
	if err != nil {
		return nil, err
	}

	headers := []string{
		"First Name", "Last Name", "Email", "Phone", "Gender", "DOB", "Is Freelance", "Role", "Branch",
	}

	for i, h := range headers {
		cell, _ := excelize.CoordinatesToCellName(i+1, 1)
		err := f.SetCellValue(sheet, cell, h)
		if err != nil {
			return nil, err
		}
	}

	for i, user := range users {
		row := i + 2
		values := []interface{}{
			user.FirstName, user.LastName, user.Email, user.Phone,
			pkg.DerefString(user.Gender),
			func() string {
				if user.DOB != nil {
					return user.DOB.Format("2006-01-02")
				}
				return ""
			}(),
			user.IsFreelance,
			user.Role,
			user.Branch.Name,
		}

		for j, val := range values {
			cell, _ := excelize.CoordinatesToCellName(j+1, row)
			err := f.SetCellValue(sheet, cell, val)
			if err != nil {
				return nil, err
			}
		}
	}

	var buf bytes.Buffer
	if err := f.Write(&buf); err != nil {
		return nil, fmt.Errorf("failed to write excel: %w", err)
	}

	return buf.Bytes(), nil
}

func (s *userExcelService) ImportUsersFromExcel(file multipart.File, creatorID uint, companyUUID string) error {
	data, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("failed to read uploaded file: %w", err)
	}

	f, err := excelize.OpenReader(bytes.NewReader(data))
	if err != nil {
		return fmt.Errorf("failed to parse excel: %w", err)
	}

	rows, err := f.GetRows("Users")
	if err != nil {
		return fmt.Errorf("failed to get rows: %w", err)
	}

	for i, row := range rows {
		if i == 0 {
			continue
		}

		if len(row) < 9 {
			continue
		}

		existingUser, err := s.userRepo.FindByEmail(row[2])
		if err == nil && existingUser.ID != 0 {
			continue
		}

		dob, _ := time.Parse("2006-01-02", row[5])
		isFreelance := row[6] == "true"

		req := request.UserEmployeeReq{
			FirstName:   row[0],
			LastName:    row[1],
			Email:       row[2],
			Phone:       row[3],
			Gender:      row[4],
			DOB:         dob,
			IsFreelance: isFreelance,
			Password:    "default123",
			CompanyUUID: companyUUID,
		}

		role, _ := s.roleRepo.FindByName(row[7])
		branch, _ := s.branchRepo.FindByName(row[8])

		if role.ID == 0 {
			continue
		}

		req.RoleUUID = role.UUID
		if branch.ID != 0 {
			req.BranchUUID = branch.UUID
		}

		_ = s.userService.CreateUser(req, creatorID)
	}

	return nil
}
