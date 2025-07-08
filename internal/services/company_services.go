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
	GetAll() ([]models.Company, error)
	GetByUUID(uuid string) (models.Company, error)
	Update(uuid string, request request.CompanyReq) (models.Company, error)
	Delete(uuid string) (models.Company, error)
}

type companyServices struct {
	companyRepo repositories.CompanyRepositories
}

func NewCompanyService(companyRepo repositories.CompanyRepositories) CompanyServices {
	return &companyServices{companyRepo}
}

func (s *companyServices) Create(request request.CompanyReq) (models.Company, error) {
	if err := pkg.Validate.Struct(request); err != nil {
		return models.Company{}, err
	}

	newCompany := models.Company{
		UUID:    uuid.NewString(),
		UserId:  request.UserId,
		Name:    request.Name,
		Logo:    request.Logo,
		Address: request.Address,
		Email:   request.Email,
		Phone:   request.Phone,
	}

	if err := s.companyRepo.Create(&newCompany); err != nil {
		return models.Company{}, err
	}
	return newCompany, nil
}

func (s *companyServices) GetAll() ([]models.Company, error) {
	var companies []models.Company
	companies, err := s.companyRepo.GetAll()
	if err != nil {
		return nil, err
	}
	return companies, nil

}

func (s *companyServices) GetByUUID(uuid string) (models.Company, error) {
	company, err := s.companyRepo.GetByUUID(uuid)
	if err != nil {
		return models.Company{}, err
	}
	return company, nil
}

func (s *companyServices) Update(uuid string, request request.CompanyReq) (models.Company, error) {
	if err := pkg.Validate.Struct(request); err != nil{
		return models.Company{}, err
	}

	updatedCompany, err := s.companyRepo.GetByUUID(uuid)
	if err != nil {
		return models.Company{}, err
	}

	updatedCompany.Name = request.Name
	updatedCompany.Logo = request.Logo
	updatedCompany.Address = request.Address
	updatedCompany.Email = request.Email
	updatedCompany.Phone = request.Phone

	if err := s.companyRepo.Update(&updatedCompany); err != nil{
		return models.Company{}, err
	}
	return updatedCompany,nil
}


func (s *companyServices) Delete(uuid string) (models.Company, error) {
	targetCompany, err := s.companyRepo.GetByUUID(uuid)
	if err != nil {
		return models.Company{}, err
	} 

	if err := s.companyRepo.Delete(&targetCompany); err != nil {
		return models.Company{}, err
	}
	return targetCompany, nil

}

