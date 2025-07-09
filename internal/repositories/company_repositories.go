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

	err := r.db.
		Table("companies").
		Select("companies.*, users.uuid AS UserUUID").
		Joins("JOIN users ON users.id = companies.user_id").
		Scan(&companies).Error

	return companies, err
}

func (r *companyRepositories) GetByUUID(uuid string) (models.Company, error) {
	var company models.Company

	err := r.db.
		Table("companies").
		Select("companies.*, users.uuid AS UserUUID").
		Joins("JOIN users ON users.id = companies.user_id").
		Where("companies.uuid = ?", uuid).
		Scan(&company).Error

	if err != nil {
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
