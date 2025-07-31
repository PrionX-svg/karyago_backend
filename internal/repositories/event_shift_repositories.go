package repositories

import (
	"hris_backend/internal/models"

	"gorm.io/gorm"
)

type EventShiftRepository interface {
	Create(eventShift *models.EventShift) error
	GetAll() ([]models.EventShift, error)
	GetByID(id uint) (*models.EventShift, error)
	GetByUUID(uuid string) (*models.EventShift, error)
	GetByEventUUID(uuid string) ([]models.EventShift, error)
	Update(eventShift *models.EventShift) error
	Delete(uuid string) error
}

type eventShiftRepository struct {
	db *gorm.DB
}

func NewEventShiftRepository(db *gorm.DB) EventShiftRepository {
	return &eventShiftRepository{db}
}

func (r *eventShiftRepository) Create(eventShift *models.EventShift) error {
	if err := r.db.Create(eventShift).Error; err != nil {
		return err
	}
	return nil
}

func (r *eventShiftRepository) GetAll() ([]models.EventShift, error) {
	var shifts []models.EventShift
	err := r.db.
		Preload("Event", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "uuid", "name")
		}).
		Find(&shifts).Error
	if err != nil {
		return nil, err
	}
	return shifts, nil
}

func (r *eventShiftRepository) GetByID(id uint) (*models.EventShift, error) {
	var shift models.EventShift
	err := r.db.
		Preload("Event", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "uuid", "name")
		}).
		Where("id = ?", id).
		First(&shift).Error
	if err != nil {
		return nil, err
	}
	return &shift, nil
}

func (r *eventShiftRepository) GetByUUID(uuid string) (*models.EventShift, error) {
	var shift models.EventShift
	err := r.db.
		Preload("Event", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "uuid", "name")
		}).
		Where("uuid = ?", uuid).
		First(&shift).Error
	if err != nil {
		return nil, err
	}
	return &shift, nil
}

func (r *eventShiftRepository) GetByEventUUID(uuid string) ([]models.EventShift, error) {
	var shifts []models.EventShift
	err := r.db.
		Joins("JOIN events ON events.id = event_shifts.event_id").
		Where("events.uuid = ?", uuid).
		Preload("Event", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "uuid", "name")
		}).
		Find(&shifts).Error
	if err != nil {
		return nil, err
	}
	return shifts, nil
}

func (r *eventShiftRepository) Update(eventShift *models.EventShift) error {
	if err := r.db.
		Omit("Event"). // 👈 prevent GORM from touching the associated Event
		Save(eventShift).Error; err != nil {
		return err
	}
	return nil
}


func (r *eventShiftRepository) Delete(uuid string) error {
	return r.db.Where("uuid = ?", uuid).Delete(&models.EventShift{}).Error
}
