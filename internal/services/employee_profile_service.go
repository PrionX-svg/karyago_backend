package services

import (
	"fmt"
	"hris_backend/internal/models"
	"hris_backend/internal/repositories"
)

type EmployeeProfileService interface {
	GetMyProfile(userID uint) (*EmployeeProfileResponse, error)
}

type employeeProfileService struct {
	userRepo       repositories.UserRepository
	userDetailRepo repositories.UserDetailRepository
	employeeRepo   repositories.EmployeeRepository
	companyRepo    repositories.CompanyRepositories
	deptRepo       repositories.DepartmentRepositories
	branchRepo     repositories.BranchRepository
}

func NewEmployeeProfileService(
	userRepo repositories.UserRepository,
	userDetailRepo repositories.UserDetailRepository,
	employeeRepo repositories.EmployeeRepository,
	companyRepo    repositories.CompanyRepositories,
	deptRepo       repositories.DepartmentRepositories,
	branchRepo repositories.BranchRepository,
) EmployeeProfileService {
	return &employeeProfileService{
		userRepo, userDetailRepo, employeeRepo, companyRepo, deptRepo, branchRepo,
	}
}

type EmployeeProfileResponse struct {
	Name        string  `json:"name"`
	Email       string  `json:"email"`
	Phone       string  `json:"phone"`
	Gender      *string `json:"gender,omitempty"`
	BirthDate   *string `json:"birth_date,omitempty"`
	Address     *string `json:"address,omitempty"`
	TaxID       *string `json:"tax_id,omitempty"`
	SocialID    *string `json:"social_id,omitempty"`
	Company     *string `json:"company,omitempty"`
	CompanyUUID *string `json:"company_uuid,omitempty"` //new changes
	EmployeeID  *string `json:"employee_id,omitempty"`
	Department  *string `json:"department,omitempty"`
	Branch      *string `json:"branch,omitempty"`
}

func (s *employeeProfileService) GetMyProfile(userID uint) (*EmployeeProfileResponse, error) {
	user, err := s.userRepo.GetByID(userID)
	if err != nil {
		return nil, fmt.Errorf("user not found")
	}

	employee, _ := s.employeeRepo.FindByUserID(userID)
	if employee == nil {
		return nil, fmt.Errorf("employee not found")
	}

	var (
		company    *models.Company
		department *models.Department
		branch     *models.Branch
		userDetail *models.UserDetail
	)

	if employee.CompanyID != nil {
		company, _ = s.companyRepo.GetByUserUUID(user.UUID) //made changes in here
	}
	if employee.DepartmentID != nil {
		department, _ = s.deptRepo.FindByID(*employee.DepartmentID)
	}
	if employee.BranchID != nil {
		branch, _ = s.branchRepo.FindByID(*employee.BranchID)
	}
	userDetail, _ = s.userDetailRepo.FindByUserID(userID)

	res := &EmployeeProfileResponse{
		Name:       fmt.Sprintf("%s %s", user.FirstName, user.LastName),
		Email:      user.Email,
		Phone:      user.Phone,
		Gender:     user.Gender,
		Company:    getString(company, func(c *models.Company) string { return c.Name }),
		CompanyUUID: getString(company, func(c *models.Company) string { return c.UUID }), // 🔥 new changes
		Department: getString(department, func(d *models.Department) string { return d.Name }),
		Branch:     getString(branch, func(b *models.Branch) string { return b.Name }),
	}
	if user.DOB != nil {
		str := user.DOB.Format("2006-01-02")
		res.BirthDate = &str
	}
	if userDetail != nil {
		res.TaxID = &userDetail.TaxID
		res.SocialID = &userDetail.SocialID
	}
	res.EmployeeID = &employee.UUID

	return res, nil
}

// helper buat ambil string aman
func getString[T any](v *T, f func(*T) string) *string {
	if v == nil {
		return nil
	}
	s := f(v)
	return &s
}
