package services

import (
	"hris_backend/internal/models"
	"hris_backend/internal/repositories"
	"hris_backend/internal/request"
	"hris_backend/internal/response"

	"github.com/google/uuid"
)

type EventItemServices interface {
	Create(req request.EventItemRequest) (response.EventItemResponse, error)
	GetAll() ([]response.EventItemResponse, error)
	GetByUUID(uuid string) (response.EventItemResponse, error)
	Update(uuid string, req request.EventItemRequest) (response.EventItemResponse, error)
	Delete(uuid string) (response.EventItemResponse, error)
	GetByEventUUID(eventUUID string) ([]response.EventItemResponse, error)
}

type eventItemServices struct {
	itemRepo  repositories.EventItemRepository
	eventRepo repositories.EventRepository
	userRepo  repositories.UserRepository
}

func NewEventItemService(
	itemRepo repositories.EventItemRepository,
	eventRepo repositories.EventRepository,
	userRepo repositories.UserRepository,
) EventItemServices {
	return &eventItemServices{
		itemRepo:  itemRepo,
		eventRepo: eventRepo,
		userRepo:  userRepo,
	}
}

func (s *eventItemServices) Create(req request.EventItemRequest) (response.EventItemResponse, error) {
	user, err := s.userRepo.GetByUUID(req.UserUUID)
	if err != nil {
		return response.EventItemResponse{}, err
	}

	event, err := s.eventRepo.GetByUUID(req.EventUUID)
	if err != nil {
		return response.EventItemResponse{}, err
	}

	item := models.EventItem{
		UUID:     uuid.NewString(),
		EventID:  event.ID,
		Name:     req.Name,
		Type:     models.EventItemType(req.Type),
		CreatedBy: user.ID,
		ModifyBy:  user.ID,
	}

	created, err := s.itemRepo.Create(&item)
	if err != nil {
		return response.EventItemResponse{}, err
	}

	return response.EventItemResponse{
		UUID:      created.UUID,
		EventUUID: req.EventUUID,
		Name:      created.Name,
		Type:      string(created.Type),
	}, nil
}

func (s *eventItemServices) GetAll() ([]response.EventItemResponse, error) {
	items, err := s.itemRepo.GetAll()
	if err != nil {
		return nil, err
	}

	var result []response.EventItemResponse
	for _, i := range items {
		eventUUID := ""
		if i.Event != nil {
			eventUUID = i.Event.UUID
		}

		result = append(result, response.EventItemResponse{
			UUID:      i.UUID,
			EventUUID: eventUUID,
			Name:      i.Name,
			Type:      string(i.Type),
		})
	}
	return result, nil
}

func (s *eventItemServices) GetByUUID(uuid string) (response.EventItemResponse, error) {
	item, err := s.itemRepo.GetByUUID(uuid)
	if err != nil {
		return response.EventItemResponse{}, err
	}

	eventUUID := ""
	if item.Event != nil {
		eventUUID = item.Event.UUID
	}

	return response.EventItemResponse{
		UUID:      item.UUID,
		EventUUID: eventUUID,
		Name:      item.Name,
		Type:      string(item.Type),
	}, nil
}

func (s *eventItemServices) Update(uuid string, req request.EventItemRequest) (response.EventItemResponse, error) {
	item, err := s.itemRepo.GetByUUID(uuid)
	if err != nil {
		return response.EventItemResponse{}, err
	}

	user, err := s.userRepo.GetByUUID(req.UserUUID)
	if err != nil {
		return response.EventItemResponse{}, err
	}

	event, err := s.eventRepo.GetByUUID(req.EventUUID)
	if err != nil {
		return response.EventItemResponse{}, err
	}

	item.Name = req.Name
	item.Type = models.EventItemType(req.Type)
	item.EventID = event.ID
	item.ModifyBy = user.ID

	updated, err := s.itemRepo.Update(item)
	if err != nil {
		return response.EventItemResponse{}, err
	}

	return response.EventItemResponse{
		UUID:      updated.UUID,
		EventUUID: event.UUID,
		Name:      updated.Name,
		Type:      string(updated.Type),
	}, nil
}

func (s *eventItemServices) Delete(uuid string) (response.EventItemResponse, error) {
	item, err := s.itemRepo.GetByUUID(uuid)
	if err != nil {
		return response.EventItemResponse{}, err
	}

	if err := s.itemRepo.Delete(item); err != nil {
		return response.EventItemResponse{}, err
	}

	eventUUID := ""
	if item.Event != nil {
		eventUUID = item.Event.UUID
	}

	return response.EventItemResponse{
		UUID:      item.UUID,
		EventUUID: eventUUID,
		Name:      item.Name,
		Type:      string(item.Type),
	}, nil
}

func (s *eventItemServices) GetByEventUUID(eventUUID string) ([]response.EventItemResponse, error) {
	event, err := s.eventRepo.GetByUUID(eventUUID)
	if err != nil {
		return nil, err
	}

	items, err := s.itemRepo.FindAllByEventID(event.ID)
	if err != nil {
		return nil, err
	}

	var result []response.EventItemResponse
	for _, i := range items {
		result = append(result, response.EventItemResponse{
			UUID:      i.UUID,
			EventUUID: eventUUID,
			Name:      i.Name,
			Type:      string(i.Type),
		})
	}
	return result, nil
}
