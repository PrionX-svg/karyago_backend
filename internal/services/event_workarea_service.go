package services

import (
	"hris_backend/internal/models"
	"hris_backend/internal/repositories"
	"hris_backend/internal/request"
	"hris_backend/internal/response"

	"github.com/google/uuid"
)

type EventWorkAreaServices interface {
	Create(req request.EventWorkAreaRequest) (response.EventWorkAreaResponse, error)
	GetAll() ([]response.EventWorkAreaResponse, error)
	GetByUUID(uuid string) (response.EventWorkAreaResponse, error)
	Update(uuid string, req request.EventWorkAreaRequest) (response.EventWorkAreaResponse, error)
	Delete(uuid string) (response.EventWorkAreaResponse, error)
	GetByEventUUID(eventUUID string) ([]response.EventWorkAreaResponse, error)
}

type eventWorkAreaServices struct {
	workAreaRepo repositories.EventWorkAreaRepository
	eventRepo    repositories.EventRepository
	userRepo     repositories.UserRepository
}

func NewEventWorkAreaService(
	workAreaRepo repositories.EventWorkAreaRepository,
	eventRepo repositories.EventRepository,
	userRepo repositories.UserRepository,
) EventWorkAreaServices {
	return &eventWorkAreaServices{
		workAreaRepo: workAreaRepo,
		eventRepo:    eventRepo,
		userRepo:     userRepo,
	}
}

func (s *eventWorkAreaServices) Create(req request.EventWorkAreaRequest) (response.EventWorkAreaResponse, error) {
	user, err := s.userRepo.GetByUUID(req.UserUUID)
	if err != nil {
		return response.EventWorkAreaResponse{}, err
	}

	event, err := s.eventRepo.GetByUUID(req.EventUUID)
	if err != nil {
		return response.EventWorkAreaResponse{}, err
	}

	area := models.EventWorkArea{
		UUID:     uuid.NewString(),
		Name:     req.Name,
		EventID:  event.ID,
		CreatedBy: user.ID,
		ModifyBy:  user.ID,
	}

	created, err := s.workAreaRepo.Create(&area)
	if err != nil {
		return response.EventWorkAreaResponse{}, err
	}

	return response.EventWorkAreaResponse{
		UUID:      created.UUID,
		EventUUID: req.EventUUID,
		Name:      created.Name,
	}, nil
}

func (s *eventWorkAreaServices) GetAll() ([]response.EventWorkAreaResponse, error) {
	areas, err := s.workAreaRepo.GetAll()
	if err != nil {
		return nil, err
	}

	var result []response.EventWorkAreaResponse
	for _, a := range areas {
		eventUUID := ""
		if a.Event != nil {
			eventUUID = a.Event.UUID
		}

		result = append(result, response.EventWorkAreaResponse{
			UUID:      a.UUID,
			EventUUID: eventUUID,
			Name:      a.Name,
		})
	}
	return result, nil
}

func (s *eventWorkAreaServices) GetByUUID(uuid string) (response.EventWorkAreaResponse, error) {
	area, err := s.workAreaRepo.GetByUUID(uuid)
	if err != nil {
		return response.EventWorkAreaResponse{}, err
	}

	eventUUID := ""
	if area.Event != nil {
		eventUUID = area.Event.UUID
	}

	return response.EventWorkAreaResponse{
		UUID:      area.UUID,
		EventUUID: eventUUID,
		Name:      area.Name,
	}, nil
}

func (s *eventWorkAreaServices) Update(uuid string, req request.EventWorkAreaRequest) (response.EventWorkAreaResponse, error) {
	area, err := s.workAreaRepo.GetByUUID(uuid)
	if err != nil {
		return response.EventWorkAreaResponse{}, err
	}

	user, err := s.userRepo.GetByUUID(req.UserUUID)
	if err != nil {
		return response.EventWorkAreaResponse{}, err
	}

	event, err := s.eventRepo.GetByUUID(req.EventUUID)
	if err != nil {
		return response.EventWorkAreaResponse{}, err
	}

	area.Name = req.Name
	area.EventID = event.ID
	area.ModifyBy = user.ID

	updated, err := s.workAreaRepo.Update(area)
	if err != nil {
		return response.EventWorkAreaResponse{}, err
	}

	return response.EventWorkAreaResponse{
		UUID:      updated.UUID,
		EventUUID: req.EventUUID,
		Name:      updated.Name,
	}, nil
}

func (s *eventWorkAreaServices) Delete(uuid string) (response.EventWorkAreaResponse, error) {
	area, err := s.workAreaRepo.GetByUUID(uuid)
	if err != nil {
		return response.EventWorkAreaResponse{}, err
	}

	if err := s.workAreaRepo.Delete(area); err != nil {
		return response.EventWorkAreaResponse{}, err
	}

	eventUUID := ""
	if area.Event != nil {
		eventUUID = area.Event.UUID
	}

	return response.EventWorkAreaResponse{
		UUID:      area.UUID,
		EventUUID: eventUUID,
		Name:      area.Name,
	}, nil
}

func (s *eventWorkAreaServices) GetByEventUUID(eventUUID string) ([]response.EventWorkAreaResponse, error) {
	event, err := s.eventRepo.GetByUUID(eventUUID)
	if err != nil {
		return nil, err
	}

	areas, err := s.workAreaRepo.FindAllByEventID(event.ID)
	if err != nil {
		return nil, err
	}

	var result []response.EventWorkAreaResponse
	for _, a := range areas {
		result = append(result, response.EventWorkAreaResponse{
			UUID:      a.UUID,
			EventUUID: eventUUID,
			Name:      a.Name,
		})
	}
	return result, nil
}
