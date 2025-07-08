package services

import (
	"hris_backend/internal/models"
	"hris_backend/internal/repositories"
	"hris_backend/internal/request"
	"hris_backend/pkg"

	"github.com/google/uuid"
)


type CompanyServices interface {
	Create(request request.CompanyReq) (models.Company, error)
}

type companyServices struct {
	companyRepo repositories.CompanyRepositories
}

func NewCompanyService(companyRepo repositories.CompanyRepositories) CompanyServices {
	return &companyServices{companyRepo}
}

func (s *companyServices) Create(request request.CompanyReq) (models.Company, error){
	if err := pkg.Validate.Struct(request); err != nil{
		return models.Company{}, err
	}	

	newCompany := models.Company{
		UUID: uuid.NewString(),
		UserId: request.UserId,
		Name: request.Name,
		Logo: request.Logo,
		Address: request.Address,
		Email: request.Email,
		Phone: request.Phone,
	}

	if err := s.companyRepo.Create(&newCompany); err != nil{
		return models.Company{},err
	}
	return newCompany,nil
} 
