package repositories

import (
	"hris_backend/internal/models"

	"gorm.io/gorm"
)

type EventDepartmentRepository interface {
	Create(dept *models.EventDepartment) (models.EventDepartment, error)
	GetAll() ([]models.EventDepartment, error)
	GetByID(id uint) (models.EventDepartment, error)
	GetByUUID(uuid string) (*models.EventDepartment, error)
	Update(dept *models.EventDepartment) (models.EventDepartment, error)
	Delete(dept *models.EventDepartment) error
	FindAllByGroupID(groupID uint) ([]*models.EventDepartment, error)
}

type eventDepartmentRepo struct {
	db *gorm.DB
}

func NewEventDepartmentRepository(db *gorm.DB) EventDepartmentRepository {
	return &eventDepartmentRepo{db}
}

func (r *eventDepartmentRepo) Create(dept *models.EventDepartment) (models.EventDepartment, error) {
	if err := r.db.Create(dept).Error; err != nil {
		return models.EventDepartment{}, err
	}

	var created models.EventDepartment
	err := r.db.Preload("EventDepartmentGroup").First(&created, dept.ID).Error
	return created, err
}

func (r *eventDepartmentRepo) GetAll() ([]models.EventDepartment, error) {
	var depts []models.EventDepartment
	err := r.db.Preload("EventDepartmentGroup").Find(&depts).Error
	return depts, err
}

func (r *eventDepartmentRepo) GetByID(id uint) (models.EventDepartment, error) {
	var dept models.EventDepartment
	err := r.db.Select("uuid", "name").First(&dept, id).Error
	return dept, err
}

func (r *eventDepartmentRepo) GetByUUID(uuid string) (*models.EventDepartment, error) {
	var dept models.EventDepartment
	err := r.db.Preload("EventDepartmentGroup").Where("uuid = ?", uuid).First(&dept).Error
	return &dept, err
}

func (r *eventDepartmentRepo) Update(dept *models.EventDepartment) (models.EventDepartment, error) {
	if err := r.db.Save(dept).Error; err != nil {
		return models.EventDepartment{}, err
	}

	var updated models.EventDepartment
	err := r.db.Preload("EventDepartmentGroup").First(&updated, dept.ID).Error
	return updated, err
}

func (r *eventDepartmentRepo) Delete(dept *models.EventDepartment) error {
	if err := r.db.Delete(dept).Error; err != nil {
		return err
	}
	return nil
}

func (r *eventDepartmentRepo) FindAllByGroupID(groupID uint) ([]*models.EventDepartment, error) {
	var depts []*models.EventDepartment
	err := r.db.Preload("EventDepartmentGroup").
		Where("event_department_group_id = ?", groupID).
		Find(&depts).Error
	return depts, err
}

