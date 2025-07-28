package repositories

import (
	"hris_backend/internal/models"

	"gorm.io/gorm"
)

type EventWorkAreaRepository interface {
	Create(area *models.EventWorkArea) (models.EventWorkArea, error)
	GetAll() ([]models.EventWorkArea, error)
	GetByID(id uint) (models.EventWorkArea, error)
	GetByUUID(uuid string) (*models.EventWorkArea, error)
	Update(area *models.EventWorkArea) (models.EventWorkArea, error)
	Delete(area *models.EventWorkArea) error
	FindAllByEventID(eventID uint) ([]*models.EventWorkArea, error)
}

type eventWorkAreaRepo struct {
	db *gorm.DB
}

func NewEventWorkAreaRepository(db *gorm.DB) EventWorkAreaRepository {
	return &eventWorkAreaRepo{db}
}

func (r *eventWorkAreaRepo) Create(area *models.EventWorkArea) (models.EventWorkArea, error) {
	if err := r.db.Create(area).Error; err != nil {
		return models.EventWorkArea{}, err
	}

	var created models.EventWorkArea
	err := r.db.Preload("Event").First(&created, area.ID).Error
	return created, err
}

func (r *eventWorkAreaRepo) GetAll() ([]models.EventWorkArea, error) {
	var areas []models.EventWorkArea
	err := r.db.Preload("Event").Find(&areas).Error
	return areas, err
}

func (r *eventWorkAreaRepo) GetByID(id uint) (models.EventWorkArea, error) {
	var area models.EventWorkArea
	err := r.db.Select("uuid", "name").First(&area, id).Error
	return area, err
}

func (r *eventWorkAreaRepo) GetByUUID(uuid string) (*models.EventWorkArea, error) {
	var area models.EventWorkArea
	err := r.db.Preload("Event").Where("uuid = ?", uuid).First(&area).Error
	return &area, err
}

func (r *eventWorkAreaRepo) Update(area *models.EventWorkArea) (models.EventWorkArea, error) {
	if err := r.db.Save(area).Error; err != nil {
		return models.EventWorkArea{}, err
	}

	var updated models.EventWorkArea
	err := r.db.Preload("Event").First(&updated, area.ID).Error
	return updated, err
}

func (r *eventWorkAreaRepo) Delete(area *models.EventWorkArea) error {
	if err := r.db.Delete(area).Error; err != nil {
		return err
	}
	return nil
}

func (r *eventWorkAreaRepo) FindAllByEventID(eventID uint) ([]*models.EventWorkArea, error) {
	var areas []*models.EventWorkArea
	err := r.db.Preload("Event").
		Where("event_id = ?", eventID).
		Find(&areas).Error
	return areas, err
}
