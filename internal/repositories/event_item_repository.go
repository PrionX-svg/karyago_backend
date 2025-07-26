package repositories

import (
	"hris_backend/internal/models"

	"gorm.io/gorm"
)

type EventItemRepository interface {
	Create(item *models.EventItem) (models.EventItem, error)
	GetAll() ([]models.EventItem, error)
	GetByID(id uint) (models.EventItem, error)
	GetByUUID(uuid string) (*models.EventItem, error)
	Update(item *models.EventItem) (models.EventItem, error)
	Delete(item *models.EventItem) error
	FindAllByEventID(eventID uint) ([]*models.EventItem, error)
}

type eventItemRepo struct {
	db *gorm.DB
}

func NewEventItemRepository(db *gorm.DB) EventItemRepository {
	return &eventItemRepo{db}
}

func (r *eventItemRepo) Create(item *models.EventItem) (models.EventItem, error) {
	if err := r.db.Create(item).Error; err != nil {
		return models.EventItem{}, err
	}

	var created models.EventItem
	err := r.db.Preload("Event").First(&created, item.ID).Error
	return created, err
}

func (r *eventItemRepo) GetAll() ([]models.EventItem, error) {
	var items []models.EventItem
	err := r.db.Preload("Event").Find(&items).Error
	return items, err
}

func (r *eventItemRepo) GetByID(id uint) (models.EventItem, error) {
	var item models.EventItem
	err := r.db.Select("uuid", "name", "type").First(&item, id).Error
	return item, err
}

func (r *eventItemRepo) GetByUUID(uuid string) (*models.EventItem, error) {
	var item models.EventItem
	err := r.db.Preload("Event").Where("uuid = ?", uuid).First(&item).Error
	return &item, err
}

func (r *eventItemRepo) Update(item *models.EventItem) (models.EventItem, error) {
	if err := r.db.Save(item).Error; err != nil {
		return models.EventItem{}, err
	}

	var updated models.EventItem
	err := r.db.Preload("Event").First(&updated, item.ID).Error
	return updated, err
}

func (r *eventItemRepo) Delete(item *models.EventItem) error {
	return r.db.Delete(item).Error
}

func (r *eventItemRepo) FindAllByEventID(eventID uint) ([]*models.EventItem, error) {
	var items []*models.EventItem
	err := r.db.Preload("Event").
		Where("event_id = ?", eventID).
		Find(&items).Error
	return items, err
}
