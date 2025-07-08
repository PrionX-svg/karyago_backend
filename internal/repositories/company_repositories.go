package repositories

import (
	"hris_backend/internal/models"

	"gorm.io/gorm"
)

type CompanyRepositories interface {
	// Register(user *models.User) error
	// FindByEmail(email string) (*models.User, error)
	// CheckLogin(email string) (*models.User, error)
	// UpdatePasswordByEmail(email, hashedPassword string) error
	// UpdatePassword(user *models.User) error
	// VerifyUser(user *models.User) error
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