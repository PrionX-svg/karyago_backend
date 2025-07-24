package repositories

import (
	"hris_backend/internal/models"

	"gorm.io/gorm"
)

type EmployeeRepository interface {
	Create(employee *models.Employee) error
	FindByID(id uint) (*models.Employee, error)
	FindByUUID(uuid string) (*models.Employee, error)
	FindByUserID(userID uint) (*models.Employee, error)
	FindByCompanyID(companyID uint) ([]models.Employee, error)
	FindByDepartmentID(departmentID uint) ([]models.Employee, error)
	FindByUserIDAndCompanyID(userID, companyID uint) (*models.Employee, error)
	ClearDepartmentByDepartmentID(departmentID uint) error
	FindTerminatedByUserIDAndCompanyID(userID, companyID uint) (*models.Employee, error)
	Update(employee *models.Employee) error
	DeleteByID(id uint) error
}

type employeeRepository struct {
	db *gorm.DB
}

func NewEmployeeRepository(db *gorm.DB) EmployeeRepository {
	return &employeeRepository{db}
}

func (r *employeeRepository) Create(employee *models.Employee) error {
	return r.db.Create(employee).Error
}

func (r *employeeRepository) FindByID(id uint) (*models.Employee, error) {
	var employee models.Employee
	if err := r.db.Preload("User").Where("id = ?", id).First(&employee).Error; err != nil {
		return nil, err
	}
	return &employee, nil
}

func (r *employeeRepository) FindByUUID(uuid string) (*models.Employee, error) {
	var employee models.Employee
	if err := r.db.Preload("User").Where("uuid = ?", uuid).First(&employee).Error; err != nil {
		return nil, err
	}
	return &employee, nil
}

func (r *employeeRepository) FindByUserID(userID uint) (*models.Employee, error) {
	var employee models.Employee
	if err := r.db.Where("user_id = ?", userID).First(&employee).Error; err != nil {
		return nil, err
	}
	return &employee, nil
}

func (r *employeeRepository) FindByCompanyID(companyID uint) ([]models.Employee, error) {
	var employees []models.Employee
	if err := r.db.Where("company_id = ?", companyID).Find(&employees).Error; err != nil {
		return nil, err
	}
	return employees, nil
}

func (r *employeeRepository) FindByDepartmentID(departmentID uint) ([]models.Employee, error) {
	var employees []models.Employee
	err := r.db.Preload("User").Where("department_id = ?", departmentID).Find(&employees).Error
	if err != nil {
		return nil, err
	}
	return employees, nil
}

func (r *employeeRepository) ClearDepartmentByDepartmentID(departmentID uint) error {
	return r.db.Model(&models.Employee{}).
		Where("department_id = ?", departmentID).
		Update("department_id", nil).Error
}

func (r *employeeRepository) FindByUserIDAndCompanyID(userID, companyID uint) (*models.Employee, error) {
	var employee models.Employee
	if err := r.db.Where("user_id = ? AND company_id = ?", userID, companyID).First(&employee).Error; err != nil {
		return nil, err
	}
	return &employee, nil
}

func (r *employeeRepository) FindTerminatedByUserIDAndCompanyID(userID, companyID uint) (*models.Employee, error) {
	var emp models.Employee
	err := r.db.Where("user_id = ? AND company_id = ? AND terminated_at IS NOT NULL", userID, companyID).
		First(&emp).Error
	return &emp, err
}

func (r *employeeRepository) Update(employee *models.Employee) error {
	return r.db.Save(employee).Error
}

func (r *employeeRepository) DeleteByID(id uint) error {
	return r.db.Delete(&models.Employee{}, id).Error
}
