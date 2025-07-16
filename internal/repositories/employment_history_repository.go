package repositories

import (
	"gorm.io/gorm"
	"hris_backend/internal/models"
)

type EmploymentHistoryRepository interface {
	Create(data *models.EmploymentHistory) error
	Update(data *models.EmploymentHistory) error
	Delete(id uint) error
	FindByID(id uint) (models.EmploymentHistory, error)
	FindByUUID(uuid string) (models.EmploymentHistory, error)
	FindByEmployeeID(employeeID uint) ([]models.EmploymentHistory, error)
	ListAll() ([]models.EmploymentHistory, error)
}

type employmentHistoryRepository struct {
	db *gorm.DB
}

func NewEmploymentHistoryRepository(db *gorm.DB) EmploymentHistoryRepository {
	return &employmentHistoryRepository{db}
}

func (r *employmentHistoryRepository) Create(data *models.EmploymentHistory) error {
	return r.db.Create(data).Error
}

func (r *employmentHistoryRepository) Update(data *models.EmploymentHistory) error {
	return r.db.Save(data).Error
}

func (r *employmentHistoryRepository) Delete(id uint) error {
	return r.db.Delete(&models.EmploymentHistory{}, id).Error
}

func (r *employmentHistoryRepository) FindByID(id uint) (models.EmploymentHistory, error) {
	var history models.EmploymentHistory
	err := r.db.Preload("Employee").
		Preload("Company").
		Preload("Branch").
		Preload("Role").
		First(&history, id).Error
	return history, err
}

func (r *employmentHistoryRepository) FindByUUID(uuid string) (models.EmploymentHistory, error) {
	var history models.EmploymentHistory
	err := r.db.Preload("Employee.User").
		Preload("Company").
		Preload("Branch").
		Preload("Role").
		Where("uuid = ?", uuid).
		First(&history).Error
	return history, err
}

func (r *employmentHistoryRepository) FindByEmployeeID(employeeID uint) ([]models.EmploymentHistory, error) {
	var histories []models.EmploymentHistory
	err := r.db.Preload("Employee.User").
		Preload("Company").
		Preload("Branch").
		Preload("Role").
		Where("employee_id = ?", employeeID).
		Order("start_date desc").
		Find(&histories).Error
	return histories, err
}

func (r *employmentHistoryRepository) ListAll() ([]models.EmploymentHistory, error) {
	var histories []models.EmploymentHistory
	err := r.db.Order("created_at desc").Find(&histories).Error
	return histories, err
}
