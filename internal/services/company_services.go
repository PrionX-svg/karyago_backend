package services

import (
	"hris_backend/internal/models"
	"hris_backend/internal/repositories"
)

type CompanyServices interface {
	Create(company *models.Company) error
}

type companyServices struct {
	companyRepo repositories.CompanyRepositories
}

func NewCompanyService(companyRepo repositories.CompanyRepositories) CompanyServices {
	return &companyServices{companyRepo}
}

func (s *companyServices) Create(company *models.Company) (error){
	if err := s.companyRepo.Create(company); err != nil{
		return err
	}
	return nil
} 
