package repositories

import (
	"hris_backend/internal/models"

	"gorm.io/gorm"
)

type EventRepository interface {
	Create(event *models.Event) error
	GetAll() ([]models.Event, error)
	GetByID(id uint) (*models.Event, error)
	GetByUUID(uuid string) (*models.Event, error)
	GetByCompanyUUID(uuid string) ([]models.Event, error)
	Update(event *models.Event) error
	Delete(uuid string) error
}

type eventRepository struct {
	db *gorm.DB
}

func NewEventRepository(db *gorm.DB) EventRepository {
	return &eventRepository{db}
}

func (r *eventRepository) Create(event *models.Event) error {
	return r.db.Create(event).Error
}

func (r *eventRepository) GetAll() ([]models.Event, error) {
	var events []models.Event
	err := r.db.
		Preload("Company", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "uuid", "name")
		}).
		Find(&events).Error

	if err != nil {
		return nil, err
	}
	return events, nil
}

func (r *eventRepository) GetByID(id uint) (*models.Event, error) {
	var event models.Event
	err := r.db.
		Preload("Company", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "uuid", "name")
		}).
		First(&event, id).Error

	if err != nil {
		return nil, err
	}
	return &event, nil
}

func (r *eventRepository) GetByUUID(uuid string) (*models.Event, error) {
	var event models.Event
	err := r.db.
		Preload("Company", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "uuid", "name")
		}).
		Where("uuid = ?", uuid).
		First(&event).Error

	if err != nil {
		return nil, err
	}
	return &event, nil
}

func (r *eventRepository) GetByCompanyUUID(uuid string) ([]models.Event, error) {
	var events []models.Event
	err := r.db.
		Joins("JOIN companies ON companies.id = events.company_id").
		Where("companies.uuid = ?", uuid).
		Preload("Company", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "uuid", "name")
		}).
		Find(&events).Error

	if err != nil {
		return nil, err
	}
	return events, nil
}

func (r *eventRepository) Update(event *models.Event) error {
	return r.db.Save(event).Error
}

func (r *eventRepository) Delete(uuid string) error {
	return r.db.Where("uuid = ?", uuid).Delete(&models.Event{}).Error
}
