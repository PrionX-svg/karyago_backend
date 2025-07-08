package services

import (
	"hris_backend/internal/models"
	"hris_backend/internal/repositories"

	"github.com/google/uuid"
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
	company.UUID = uuid.NewString()
	if err := s.companyRepo.Create(company); err != nil{
		return err
	}
	return nil
} 
