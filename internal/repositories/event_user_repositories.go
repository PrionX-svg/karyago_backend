package repositories

import (
	"hris_backend/internal/models"

	"gorm.io/gorm"
)

type EventUserRepository interface {
	Create(eventUser *models.EventUser) error
	GetAll() ([]models.EventUser, error)
	GetByID(id uint) (*models.EventUser, error)
	GetByUUID(uuid string) (*models.EventUser, error)
	GetByUserUUID(uuid string) ([]models.EventUser, error)
	GetByEventUUID(uuid string) ([]models.EventUser, error)
	Update(eventUser *models.EventUser) error
	Delete(uuid string) error
}

type eventUserRepository struct {
	db *gorm.DB
}

func NewEventUserRepository(db *gorm.DB) EventUserRepository {
	return &eventUserRepository{db}
}

func (r *eventUserRepository) Create(eventUser *models.EventUser) error {
	return r.db.Create(eventUser).Error
}

func (r *eventUserRepository) GetAll() ([]models.EventUser, error) {
	var eventUsers []models.EventUser

	err := r.db.
		Preload("User").
		Preload("Event").
		Find(&eventUsers).Error

	if err != nil {
		return nil, err
	}
	return eventUsers, nil
}

func (r *eventUserRepository) GetByID(id uint) (*models.EventUser, error) {
	var eventUser models.EventUser

	err := r.db.
		Preload("User").
		Preload("Event").
		First(&eventUser, id).Error

	if err != nil {
		return nil, err
	}
	return &eventUser, nil
}

func (r *eventUserRepository) GetByUUID(uuid string) (*models.EventUser, error) {
var eventUser models.EventUser
	err := r.db.
		Preload("User").
		Preload("Event").
		Where("uuid = ?", uuid).
		First(&eventUser).Error

	if err != nil {
		return nil, err
	}
	return &eventUser, nil
}

func (r *eventUserRepository) GetByUserUUID(uuid string) ([]models.EventUser, error) {
	var eventUsers []models.EventUser

	err := r.db.
		Joins("JOIN users ON users.id = event_users.user_id").
		Joins("JOIN events ON events.id = event_users.event_id").
		Where("users.uuid = ?", uuid).
		Preload("User").
		Preload("Event", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "uuid", "name", "start_date", "end_date") // include other fields if needed
		}).
		Find(&eventUsers).Error

	if err != nil {
		return nil, err
	}
	return eventUsers, nil
}



func (r *eventUserRepository) GetByEventUUID(uuid string) ([]models.EventUser, error) {
	var eventUsers []models.EventUser

	err := r.db.
		Joins("JOIN events ON events.id = event_users.event_id").
		Joins("JOIN users ON users.id = event_users.user_id").
		Where("events.uuid = ?", uuid).
		Preload("User").
		Preload("Event", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "uuid", "name", "start_date", "end_date")
		}).
		Find(&eventUsers).Error

	if err != nil {
		return nil, err
	}
	return eventUsers, nil
}


func (r *eventUserRepository) Update(eventUser *models.EventUser) error {
	return r.db.Save(eventUser).Error
}

func (r *eventUserRepository) Delete(uuid string) error {
	return r.db.Where("uuid = ?", uuid).Delete(&models.EventUser{}).Error
}
