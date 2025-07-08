package repositories

import (
	"hris_backend/internal/models"

	"gorm.io/gorm"
)

type CompanyRepositories interface {
	Create(company *models.Company) error
	GetAll() ([]models.Company, error)
	GetByUUID(uuid string) (models.Company, error)
	Update(updatedCompany *models.Company) error
	Delete(targetCompany *models.Company) error 
}

type companyRepositories struct {
	db *gorm.DB
}

func NewCompanyRepository(db *gorm.DB) CompanyRepositories {
	return &companyRepositories{db}
}

func (r *companyRepositories) Create(company *models.Company) error {
	if err := r.db.Create(company).Error; err != nil {
		return err
	}
	return nil
}

func (r *companyRepositories) GetAll() ([]models.Company, error) {
	var companies []models.Company
	if err := r.db.Find(&companies).Error; err != nil {
		return nil, err
	}
	return companies, nil
}

func (r *companyRepositories) GetByUUID(uuid string) (models.Company, error) {
	var company models.Company
	if err := r.db.Where("uuid = ?", uuid).First(&company).Error; err != nil {
		return models.Company{}, err
	}
	return company, nil
}

func (r *companyRepositories) Update(updatedCompany *models.Company) error {
	if err := r.db.Save(updatedCompany).Error; err != nil {
		return err
	}
	return nil
}

func (r *companyRepositories) Delete(targetCompany *models.Company) error {
	if err := r.db.Delete(targetCompany).Error; err != nil {
		return err
	}
	return nil
}
