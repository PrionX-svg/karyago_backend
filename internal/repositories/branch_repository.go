package repositories

import (
	"gorm.io/gorm"
	"hris_backend/internal/models"
)

type BranchRepository interface {
	Create(branch *models.Branch) error
	FindByUUID(uuid string) (*models.Branch, error)
	Update(branch *models.Branch) error
	Delete(uuid string) error
	FindAll() ([]models.Branch, error)
}

type branchRepository struct {
	db *gorm.DB
}

func NewBranchRepository(db *gorm.DB) BranchRepository {
	return &branchRepository{db}
}

func (r *branchRepository) Create(branch *models.Branch) error {
	return r.db.Create(branch).Error
}

func (r *branchRepository) FindByUUID(uuid string) (*models.Branch, error) {
	var branch models.Branch

	err := r.db.
		Table("branches").
		Select("branches.*, companies.uuid AS CompanyUUID").
		Joins("JOIN companies ON companies.id = branches.company_id").
		Where("branches.uuid = ?", uuid).
		Scan(&branch).Error

	if err != nil {
		return nil, err
	}
	return &branch, nil
}

func (r *branchRepository) Update(branch *models.Branch) error {
	return r.db.Save(branch).Error
}

func (r *branchRepository) Delete(uuid string) error {
	return r.db.Where("uuid = ?", uuid).Delete(&models.Branch{}).Error
}

func (r *branchRepository) FindAll() ([]models.Branch, error) {
	var branches []models.Branch

	err := r.db.
		Table("branches").
		Select("branches.*, companies.uuid AS CompanyUUID").
		Joins("JOIN companies ON companies.id = branches.company_id").
		Scan(&branches).Error

	if err != nil {
		return nil, err
	}
	return branches, nil
}
