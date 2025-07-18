package services

import (
	"bytes"
	"fmt"
	"github.com/xuri/excelize/v2"
	"hris_backend/internal/repositories"
	"hris_backend/internal/request"
	"hris_backend/internal/response"
	"io"
	"log"
	"mime/multipart"
	"strings"
	"time"
)

type UserExcelService interface {
	ExportUsersTemplateToExcel() ([]byte, error)
	ImportUsersFromExcel(
		file multipart.File,
		creatorID uint,
		companyUUID, roleUUID string,
	) ([]response.UserWithEmployeeAndHistoryResponse, error)
}

type userExcelService struct {
	userService              UserService
	companyRepo              repositories.CompanyRepositories
	roleRepo                 repositories.RoleRepositories
	branchRepo               repositories.BranchRepository
	userRepo                 repositories.UserRepository
	employeeRepo             repositories.EmployeeRepository
	employmentHistoryService EmploymentHistoryService
}

func NewUserExcelService(
	userService UserService,
	companyRepo repositories.CompanyRepositories,
	roleRepo repositories.RoleRepositories,
	branchRepo repositories.BranchRepository,
	userRepo repositories.UserRepository,
	employeeRepo repositories.EmployeeRepository,
	employmentHistoryService EmploymentHistoryService,
) UserExcelService {
	return &userExcelService{
		userService:              userService,
		companyRepo:              companyRepo,
		roleRepo:                 roleRepo,
		branchRepo:               branchRepo,
		userRepo:                 userRepo,
		employeeRepo:             employeeRepo,
		employmentHistoryService: employmentHistoryService,
	}
}

func (s *userExcelService) ExportUsersTemplateToExcel() ([]byte, error) {
	f := excelize.NewFile()
	sheet := "Users"
	_, err := f.NewSheet(sheet)
	if err != nil {
		return nil, err
	}

	headers := []string{
		"First Name",
		"Last Name",
		"Email",
		"Password",
		"Phone",
		"Timezone",
		"Is Freelance",

		"Gender",
		"DOB",

		"Position",
		"Is Present",
		"Start Date",
		"End Date",
		"Notes",
	}

	for i, h := range headers {
		cell, _ := excelize.CoordinatesToCellName(i+1, 1)
		if err := f.SetCellValue(sheet, cell, h); err != nil {
			return nil, err
		}
	}

	var buf bytes.Buffer
	if err := f.Write(&buf); err != nil {
		return nil, fmt.Errorf("failed to write excel: %w", err)
	}

	return buf.Bytes(), nil
}

func (s *userExcelService) ImportUsersFromExcel(
	file multipart.File,
	creatorID uint,
	companyUUID, roleUUID string,
) ([]response.UserWithEmployeeAndHistoryResponse, error) {
	var importedUsers []response.UserWithEmployeeAndHistoryResponse

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("failed to read uploaded file: %w", err)
	}

	f, err := excelize.OpenReader(bytes.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("failed to parse excel: %w", err)
	}

	rows, err := f.GetRows("Users")
	if err != nil {
		return nil, fmt.Errorf("failed to get rows: %w", err)
	}

	company, err := s.companyRepo.GetByUUID(companyUUID)
	if err != nil || company.ID == 0 {
		return nil, fmt.Errorf("company not found: %s", companyUUID)
	}

	role, err := s.roleRepo.FindByUUID(roleUUID)
	if err != nil || role.ID == 0 {
		return nil, fmt.Errorf("role not found: %s", roleUUID)
	}

	for i, row := range rows {
		if i == 0 {
			continue // skip header
		}
		if len(row) < 14 {
			log.Printf("Row %d skipped: not enough columns", i+1)
			continue
		}

		email := strings.TrimSpace(row[2])
		if email == "" {
			log.Printf("Row %d skipped: empty email", i+1)
			continue
		}

		existingUser, _ := s.userRepo.FindByEmail(email)
		if existingUser.ID != 0 {
			log.Printf("Row %d skipped: user with email %s already exists", i+1, email)
			continue
		}

		dob, err := time.Parse("2006-01-02", strings.TrimSpace(row[8]))
		if err != nil {
			log.Printf("Row %d skipped: invalid DOB format (%s)", i+1, row[8])
			continue
		}

		isFreelance := strings.EqualFold(strings.TrimSpace(row[6]), "true") || strings.EqualFold(strings.TrimSpace(row[6]), "yes")
		isPresent := strings.EqualFold(strings.TrimSpace(row[10]), "true") || strings.EqualFold(strings.TrimSpace(row[10]), "yes")

		startDate, err := time.Parse("2006-01-02", strings.TrimSpace(row[11]))
		if err != nil {
			log.Printf("Row %d skipped: invalid start date (%s)", i+1, row[11])
			continue
		}

		var endDate *time.Time
		if len(row) > 12 && strings.TrimSpace(row[12]) != "" {
			t, err := time.Parse("2006-01-02", strings.TrimSpace(row[12]))
			if err != nil {
				log.Printf("Row %d skipped: invalid end date (%s)", i+1, row[12])
				continue
			}
			endDate = &t
		}

		var notes *string
		if len(row) > 13 && strings.TrimSpace(row[13]) != "" {
			n := strings.TrimSpace(row[13])
			notes = &n
		}

		userReq := request.UserEmployeeReq{
			FirstName:   strings.TrimSpace(row[0]),
			LastName:    strings.TrimSpace(row[1]),
			Email:       email,
			Password:    strings.TrimSpace(row[3]),
			Phone:       strings.TrimSpace(row[4]),
			Gender:      strings.TrimSpace(row[7]),
			DOB:         dob,
			IsFreelance: isFreelance,
			CompanyUUID: companyUUID,
			RoleUUID:    roleUUID,
		}

		userRes, err := s.userService.CreateUser(userReq, creatorID)
		if err != nil {
			log.Printf("Row %d failed to create user: %v", i+1, err)
			continue
		}

		historyReq := request.EmploymentHistoryRequest{
			EmployeeUUID: userRes.EmployeeUUID,
			CompanyUUID:  companyUUID,
			RoleUUID:     &roleUUID,
			Position:     strings.TrimSpace(row[9]),
			IsPresent:    isPresent,
			StartDate:    startDate,
			EndDate:      endDate,
			Notes:        notes,
		}

		historyRes, err := s.employmentHistoryService.Create(historyReq, creatorID)
		if err != nil {
			log.Printf("Row %d failed to create employment history: %v", i+1, err)
			continue
		}

		importedUsers = append(importedUsers, response.UserWithEmployeeAndHistoryResponse{
			UserUUID:     userRes.UserUUID,
			EmployeeUUID: userRes.EmployeeUUID,
			FirstName:    userRes.FirstName,
			LastName:     userRes.LastName,
			FullName:     fmt.Sprintf("%s %s", userRes.FirstName, userRes.LastName),
			Email:        userRes.Email,
			Phone:        userRes.Phone,
			Gender:       userRes.Gender,
			DOB:          userRes.DOB,
			IsFreelance:  userRes.IsFreelance,
			Role: &response.RoleSimpleResponse{
				UUID: role.UUID,
				Name: role.Name,
			},
			Company: &response.CompanySimpleResponse{
				UUID: company.UUID,
				Name: company.Name,
			},
			Histories: []response.EmploymentHistoryResponse{historyRes},
		})

		log.Printf("Row %d: user %s imported", i+1, email)
	}

	return importedUsers, nil
}
