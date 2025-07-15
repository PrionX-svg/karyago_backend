package repositories

import (
	"hris_backend/internal/models"

	"gorm.io/gorm"
)

type CompanyRepositories interface {
	CreateWithUser(company *models.Company) (models.Company, error)
	GetAll() ([]models.Company, error)
	GetByID(id uint) (models.Company, error)
	GetByUUID(uuid string) (models.Company, error)
	GetByUserUUID(uuid string) (models.Company, error)
	UpdateWithUser(company *models.Company) (models.Company, error)
	HasAnyBranch(companyID uint) (bool, error)
	GetByUserID(userID uint) (models.Company, error)
	Delete(targetCompany *models.Company) error
}

type companyRepositories struct {
	db *gorm.DB
}

func NewCompanyRepository(db *gorm.DB) CompanyRepositories {
	return &companyRepositories{db}
}

func (r *companyRepositories) CreateWithUser(company *models.Company) (models.Company, error) {
	if err := r.db.Create(company).Error; err != nil {
		return models.Company{}, err
	}

	var created models.Company
	err := r.db.Preload("User").First(&created, company.ID).Error
	if err != nil {
		return models.Company{}, err
	}

	return created, nil
}

func (r *companyRepositories) GetAll() ([]models.Company, error) {
	var companies []models.Company
	err := r.db.Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("id", "uuid", "first_name", "last_name")
	}).Find(&companies).Error
	return companies, err
}

func (r *companyRepositories) GetByID(id uint) (models.Company, error) {
	var company models.Company
	err := r.db.Select("uuid").First(&company, id).Error
	return company, err
}

func (r *companyRepositories) GetByUUID(uuid string) (models.Company, error) {
	var company models.Company
	err := r.db.Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("id", "uuid", "first_name", "last_name")
	}).Where("uuid = ?", uuid).First(&company).Error
	return company, err
}

func (r *companyRepositories) GetByUserUUID(uuid string) (models.Company, error) {
	var company models.Company
	err := r.db.Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("id", "uuid", "first_name", "last_name")
	}).Where("user_uuid = ?", uuid).First(&company).Error
	return company, err
}

func (r *companyRepositories) UpdateWithUser(company *models.Company) (models.Company, error) {
	if err := r.db.Save(company).Error; err != nil {
		return models.Company{}, err
	}

	var updated models.Company
	if err := r.db.Preload("User").First(&updated, company.ID).Error; err != nil {
		return models.Company{}, err
	}

	return updated, nil
}

func (r *companyRepositories) HasAnyBranch(companyID uint) (bool, error) {
	var count int64
	err := r.db.Table("branches").
		Where("company_id = ?", companyID).
		Count(&count).Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (r *companyRepositories) GetByUserID(userID uint) (models.Company, error) {
	var company models.Company
	err := r.db.Where("user_id = ?", userID).First(&company).Error
	return company, err
}

func (r *companyRepositories) Delete(targetCompany *models.Company) error {
	if err := r.db.Delete(targetCompany).Error; err != nil {
		return err
	}
	return nil
}
