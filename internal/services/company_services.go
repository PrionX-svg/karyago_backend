package services

import (
	"fmt"
	"hris_backend/internal/models"
	"hris_backend/internal/repositories"
	"hris_backend/internal/request"
	"hris_backend/internal/response"
	"hris_backend/pkg"

	"github.com/google/uuid"
)

type CompanyServices interface {
	Create(request request.CompanyReq) (response.CompanyResponse, error)
	GetAll() ([]response.CompanyResponse, error)
	GetByUUID(uuid string) (response.CompanyResponse, error)
	GetByUserUUID(uuid string) (response.CompanyResponse, error)
	Update(uuid string, request request.CompanyReq) (response.CompanyResponse, error)
	Delete(uuid string) (response.CompanyResponse, error)
}

type companyServices struct {
	companyRepo  repositories.CompanyRepositories
	userRepo     repositories.UserRepository
	employeeRepo repositories.EmployeeRepository
}

func NewCompanyService(companyRepo repositories.CompanyRepositories, userRepo repositories.UserRepository, employeeRepo repositories.EmployeeRepository) CompanyServices {
	return &companyServices{
		companyRepo:  companyRepo,
		userRepo:     userRepo,
		employeeRepo: employeeRepo,
	}
}

func (s *companyServices) Create(request request.CompanyReq) (response.CompanyResponse, error) {
	if err := pkg.Validate.Struct(request); err != nil {
		return response.CompanyResponse{}, err
	}

	user, err := s.userRepo.GetByUUID(request.UserUUID)
	if err != nil {
		return response.CompanyResponse{}, err
	}

	newCompany := models.Company{
		UUID:      uuid.NewString(),
		UserId:    user.ID,
		Name:      request.Name,
		Logo:      request.Logo,
		Address:   request.Address,
		Email:     request.Email,
		Phone:     request.Phone,
		CreatedBy: user.ID,
		ModifyBy:  user.ID,
	}

	createdCompany, err := s.companyRepo.CreateWithUser(&newCompany)
	if err != nil {
		return response.CompanyResponse{}, err
	}

	employee, err := s.employeeRepo.FindByUserID(user.ID)
	if err != nil {
		return response.CompanyResponse{}, fmt.Errorf("failed to find employee for user: %w", err)
	}

	employee.CompanyID = &createdCompany.ID
	employee.ModifyBy = user.ID

	if err := s.employeeRepo.Update(employee); err != nil {
		return response.CompanyResponse{}, fmt.Errorf("failed to update employee with company ID: %w", err)
	}

	resp := response.CompanyResponse{
		UUID:    createdCompany.UUID,
		Logo:    createdCompany.Logo,
		Name:    createdCompany.Name,
		Address: createdCompany.Address,
		Email:   createdCompany.Email,
		Phone:   createdCompany.Phone,
		User: struct {
			UUID      string `json:"uuid"`
			FirstName string `json:"firstname"`
			LastName  string `json:"lastname"`
		}{
			UUID:      user.UUID,
			FirstName: user.FirstName,
			LastName:  user.LastName,
		},
	}

	return resp, nil
}

func (s *companyServices) GetAll() ([]response.CompanyResponse, error) {
	companies, err := s.companyRepo.GetAll()
	if err != nil {
		return nil, err
	}

	var result []response.CompanyResponse
	for _, c := range companies {
		res := response.CompanyResponse{
			UUID:    c.UUID,
			Logo:    c.Logo,
			Name:    c.Name,
			Address: c.Address,
			Email:   c.Email,
			Phone:   c.Phone,
		}
		res.User.UUID = c.User.UUID
		res.User.FirstName = c.User.FirstName
		res.User.LastName = c.User.LastName

		result = append(result, res)
	}

	return result, nil
}

func (s *companyServices) GetByUUID(uuid string) (response.CompanyResponse, error) {
	c, err := s.companyRepo.GetByUUID(uuid)
	if err != nil {
		return response.CompanyResponse{}, err
	}

	res := response.CompanyResponse{
		UUID:    c.UUID,
		Logo:    c.Logo,
		Name:    c.Name,
		Address: c.Address,
		Email:   c.Email,
		Phone:   c.Phone,
	}
	res.User.UUID = c.User.UUID
	res.User.FirstName = c.User.FirstName
	res.User.LastName = c.User.LastName

	return res, nil
}

func (s *companyServices) GetByUserUUID(uuid string) (response.CompanyResponse, error) {
	c, err := s.companyRepo.GetByUserUUID(uuid)
	if err != nil {
		return response.CompanyResponse{}, err
	}

	res := response.CompanyResponse{
		UUID:    c.UUID,
		Logo:    c.Logo,
		Name:    c.Name,
		Address: c.Address,
		Email:   c.Email,
		Phone:   c.Phone,
	}
	res.User.UUID = c.User.UUID
	res.User.FirstName = c.User.FirstName
	res.User.LastName = c.User.LastName

	return res, nil
}

func (s *companyServices) Update(uuid string, request request.CompanyReq) (response.CompanyResponse, error) {
	if err := pkg.Validate.Struct(request); err != nil {
		return response.CompanyResponse{}, err
	}

	company, err := s.companyRepo.GetByUUID(uuid)
	if err != nil {
		return response.CompanyResponse{}, err
	}

	company.Name = request.Name
	company.Logo = request.Logo
	company.Address = request.Address
	company.Email = request.Email
	company.Phone = request.Phone
	company.ModifyBy = company.UserId

	updated, err := s.companyRepo.UpdateWithUser(&company)
	if err != nil {
		return response.CompanyResponse{}, err
	}

	return response.CompanyResponse{
		UUID:    updated.UUID,
		Logo:    updated.Logo,
		Name:    updated.Name,
		Address: updated.Address,
		Email:   updated.Email,
		Phone:   updated.Phone,
		User: struct {
			UUID      string `json:"uuid"`
			FirstName string `json:"firstname"`
			LastName  string `json:"lastname"`
		}{
			UUID:      updated.User.UUID,
			FirstName: updated.User.FirstName,
			LastName:  updated.User.LastName,
		},
	}, nil
}

func (s *companyServices) Delete(uuid string) (response.CompanyResponse, error) {
	company, err := s.companyRepo.GetByUUID(uuid)
	if err != nil {
		return response.CompanyResponse{}, err
	}

	hasBranch, err := s.companyRepo.HasAnyBranch(company.ID)
	if err != nil {
		return response.CompanyResponse{}, fmt.Errorf("failed to check company branches: %w", err)
	}
	if hasBranch {
		return response.CompanyResponse{}, fmt.Errorf("cannot delete company because it still has branches")
	}

	if err := s.companyRepo.Delete(&company); err != nil {
		return response.CompanyResponse{}, fmt.Errorf("failed to delete company: %w", err)
	}

	resp := response.CompanyResponse{
		UUID:    company.UUID,
		Logo:    company.Logo,
		Name:    company.Name,
		Address: company.Address,
		Email:   company.Email,
		Phone:   company.Phone,
		User: struct {
			UUID      string `json:"uuid"`
			FirstName string `json:"firstname"`
			LastName  string `json:"lastname"`
		}{
			UUID:      company.User.UUID,
			FirstName: company.User.FirstName,
			LastName:  company.User.LastName,
		},
	}

	return resp, nil
}
