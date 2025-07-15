package repositories

import (
	"gorm.io/gorm"
	"hris_backend/internal/models"
)

type BranchRepository interface {
	Create(branch *models.Branch) error
	FindByUUID(uuid string) (*models.Branch, error)
	FindByID(id uint) (*models.Branch, error)
	FindByName(name string) (*models.Branch, error)
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
		Preload("Company").
		Where("uuid = ?", uuid).
		First(&branch).Error

	if err != nil {
		return nil, err
	}
	return &branch, nil
}

func (r *branchRepository) FindByID(id uint) (*models.Branch, error) {
	var branch models.Branch

	err := r.db.
		Preload("Company").
		Where("id = ?", id).
		First(&branch).Error

	if err != nil {
		return nil, err
	}
	return &branch, nil
}

func (r *branchRepository) FindByName(name string) (*models.Branch, error) {
	var branch models.Branch

	err := r.db.
		Preload("Company").
		Where("name = ?", name).
		First(&branch).Error

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
		Preload("Company").
		Find(&branches).Error

	if err != nil {
		return nil, err
	}
	return branches, nil
}
