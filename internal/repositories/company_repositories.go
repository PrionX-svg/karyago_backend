package repositories

import (
	"hris_backend/internal/models"

	"gorm.io/gorm"
)

type CompanyRepositories interface {
	Create(company *models.Company) error

}

type companyRepositories struct {
	db *gorm.DB
}

func NewCompanyRepository(db *gorm.DB) CompanyRepositories {
	return &companyRepositories{db}
}

func (r *companyRepositories) Create(company *models.Company) error {
	if err := r.db.Create(company).Error; err != nil{
		return err
	}
	return nil
}