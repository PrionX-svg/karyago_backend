package repositories

import (
	"hris_backend/internal/models"

	"gorm.io/gorm"
)

type EventDepartmentGroupRepository interface {
	Create(group *models.EventDepartmentGroup) (models.EventDepartmentGroup, error)
	GetAll() ([]models.EventDepartmentGroup, error)
	GetByID(id uint) (models.EventDepartmentGroup, error)
	GetByUUID(uuid string) (*models.EventDepartmentGroup, error)
	FindEventIDByUUID(eventUUID string) (uint, error)
	FindEventUUIDByID(eventID uint) (string, error)
	Update(group *models.EventDepartmentGroup) (models.EventDepartmentGroup, error)
	Delete(group *models.EventDepartmentGroup) error
	FindAllByEventID(eventID uint) ([]*models.EventDepartmentGroup, error)
}

type eventDepartmentGroupRepo struct {
	db *gorm.DB
}

func NewEventDepartmentGroupRepository(db *gorm.DB) EventDepartmentGroupRepository {
	return &eventDepartmentGroupRepo{db}
}

func (r *eventDepartmentGroupRepo) Create(group *models.EventDepartmentGroup) (models.EventDepartmentGroup, error) {
	if err := r.db.Create(group).Error; err != nil {
		return models.EventDepartmentGroup{}, err
	}

	var created models.EventDepartmentGroup
	err := r.db.Preload("Responsible").First(&created, group.ID).Error
	return created, err
}

func (r *eventDepartmentGroupRepo) GetAll() ([]models.EventDepartmentGroup, error) {
	var groups []models.EventDepartmentGroup
	err := r.db.Preload("Responsible", func(db *gorm.DB) *gorm.DB {
		return db.Select("id", "uuid", "first_name", "last_name")
	}).Find(&groups).Error
	return groups, err
}

func (r *eventDepartmentGroupRepo) GetByID(id uint) (models.EventDepartmentGroup, error) {
	var group models.EventDepartmentGroup
	err := r.db.Select("uuid", "name").First(&group, id).Error
	return group, err
}

func (r *eventDepartmentGroupRepo) GetByUUID(uuid string) (*models.EventDepartmentGroup, error) {
	var group models.EventDepartmentGroup
	err := r.db.Preload("Responsible", func(db *gorm.DB) *gorm.DB {
		return db.Select("id", "uuid", "first_name", "last_name")
	}).Where("uuid = ?", uuid).First(&group).Error
	return &group, err
}

func (r *eventDepartmentGroupRepo) FindEventIDByUUID(eventUUID string) (uint, error) {
	var event models.Event
	if err := r.db.Select("id").Where("uuid = ?", eventUUID).First(&event).Error; err != nil {
		return 0, err
	}
	return event.ID, nil
}

func (r *eventDepartmentGroupRepo) FindEventUUIDByID(eventID uint) (string, error) {
	var event models.Event
	if err := r.db.Select("uuid").Where("id = ?", eventID).First(&event).Error; err != nil {
		return "", err
	}
	return event.UUID, nil
}

func (r *eventDepartmentGroupRepo) Update(group *models.EventDepartmentGroup) (models.EventDepartmentGroup, error) {
	if err := r.db.Save(group).Error; err != nil {
		return models.EventDepartmentGroup{}, err
	}

	var updated models.EventDepartmentGroup
	err := r.db.Preload("Responsible").First(&updated, group.ID).Error
	return updated, err
}

func (r *eventDepartmentGroupRepo) Delete(group *models.EventDepartmentGroup) error {
	if err := r.db.Delete(group).Error; err != nil {
		return err
	}
	return nil
}

func (r *eventDepartmentGroupRepo) FindAllByEventID(eventID uint) ([]*models.EventDepartmentGroup, error) {
	var groups []*models.EventDepartmentGroup
	err := r.db.Preload("Responsible", func(db *gorm.DB) *gorm.DB {
		return db.Select("id", "uuid", "first_name", "last_name")
	}).Where("event_id = ?", eventID).Find(&groups).Error
	return groups, err
}
