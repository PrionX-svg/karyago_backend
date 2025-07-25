package services

import (
	"fmt"
	"hris_backend/internal/models"
	"hris_backend/internal/repositories"
	"hris_backend/internal/request"
	"hris_backend/internal/response"
	"hris_backend/pkg"

	"github.com/google/uuid"
)

type EventService interface {
	Create(request request.EventReq, userID uint) (response.EventResponse, error)
	GetAll() ([]response.EventResponse, error)
	GetByID(id uint) (response.EventResponse, error)
	GetByUUID(uuid string) (response.EventResponse, error)
	GetByCompanyUUID(uuid string) ([]response.EventResponse, error)
	Update(uuid string, req request.EventReq, actorID uint) (response.EventResponse, error)
	Delete(uuid string) (response.EventResponse, error)
}

type eventService struct {
	repo        repositories.EventRepository
	companyRepo repositories.CompanyRepositories
}

func NewEventService(repo repositories.EventRepository, companyRepo repositories.CompanyRepositories) EventService {
	return &eventService{
		repo:        repo,
		companyRepo: companyRepo,
	}
}

func (s *eventService) Create(request request.EventReq, userID uint) (response.EventResponse, error) {
	if err := pkg.Validate.Struct(request); err != nil {
		return response.EventResponse{}, err
	}

	company, err := s.companyRepo.GetByUUID(request.CompanyUUID)
	if err != nil {
		return response.EventResponse{}, fmt.Errorf("invalid company UUID")
	}

	event := &models.Event{
		UUID:      uuid.NewString(),
		CompanyID: company.ID,
		Name:      request.Name,
		StartDate: request.StartDate,
		EndDate:   request.EndDate,
		Photo:     request.Photo,
		CreatedBy: userID,
		ModifyBy:  userID,
	}

	if err := s.repo.Create(event); err != nil {
		return response.EventResponse{}, err
	}

	eventModel, err := s.repo.GetByUUID(event.UUID)
	if err != nil {
		return response.EventResponse{}, err
	}

	eventResponse := response.EventResponse{

		UUID:      eventModel.UUID,
		Name:      eventModel.Name,
		StartDate: eventModel.StartDate,
		EndDate:   eventModel.EndDate,
		Photo:     eventModel.Photo,
		Company: struct {
			UUID string `json:"uuid"`
			Name string `json:"name"`
		}{
			UUID: eventModel.Company.UUID,
			Name: eventModel.Company.Name,
		},
	}

	return eventResponse, nil
}

func (s *eventService) GetAll() ([]response.EventResponse, error) {
	events, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	var eventsResponses []response.EventResponse

	for _, ev := range events {
		eventsResponses = append(eventsResponses, response.EventResponse{
			UUID:      ev.UUID,
			Name:      ev.Name,
			StartDate: ev.StartDate,
			EndDate:   ev.EndDate,
			Photo:     ev.Photo,
			Company: struct {
				UUID string `json:"uuid"`
				Name string `json:"name"`
			}{
				UUID: ev.Company.UUID,
				Name: ev.Company.Name,
			},
		})
	}

	return eventsResponses, nil
}

func (s *eventService) GetByID(id uint) (response.EventResponse, error) {
	event, err := s.repo.GetByID(id)
	if err != nil {
		return response.EventResponse{}, err
	}

	eventResponse := response.EventResponse{
		UUID:      event.UUID,
		Name:      event.Name,
		StartDate: event.StartDate,
		EndDate:   event.EndDate,
		Photo:     event.Photo,
		Company: struct {
			UUID string `json:"uuid"`
			Name string `json:"name"`
		}{
			UUID: event.Company.UUID,
			Name: event.Company.Name,
		},
	}

	return eventResponse, nil
}

func (s *eventService) GetByUUID(uuid string) (response.EventResponse, error) {
	event, err := s.repo.GetByUUID(uuid)
	if err != nil {
		return response.EventResponse{}, err
	}

	eventResponse := response.EventResponse{
		UUID:      event.UUID,
		Name:      event.Name,
		StartDate: event.StartDate,
		EndDate:   event.EndDate,
		Photo:     event.Photo,
		Company: struct {
			UUID string `json:"uuid"`
			Name string `json:"name"`
		}{
			UUID: event.Company.UUID,
			Name: event.Company.Name,
		},
	}
	return eventResponse, nil
}

func (s *eventService) GetByCompanyUUID(uuid string) ([]response.EventResponse, error) {
	events, err := s.repo.GetByCompanyUUID(uuid)
	if err != nil {
		return nil, err
	}

	var eventResponses []response.EventResponse

	for _, ev := range events {
		eventResponses = append(eventResponses, response.EventResponse{
			UUID:      ev.UUID,
			Name:      ev.Name,
			StartDate: ev.StartDate,
			EndDate:   ev.EndDate,
			Photo:     ev.Photo,
			Company: struct {
				UUID string `json:"uuid"`
				Name string `json:"name"`
			}{
				UUID: ev.Company.UUID,
				Name: ev.Company.Name,
			},
		})
	}

	return eventResponses, nil
}

func (s *eventService) Update(uuid string, req request.EventReq, actorID uint) (response.EventResponse, error) {
	event, err := s.repo.GetByUUID(uuid)
	if err != nil {
		return response.EventResponse{}, err
	}

	event.Name = req.Name
	event.StartDate = req.StartDate
	event.EndDate = req.EndDate
	event.Photo = req.Photo
	event.ModifyBy = actorID

	if err := s.repo.Update(event); err != nil {
		return response.EventResponse{}, err
	}

	eventResponse := response.EventResponse{
		UUID:      event.UUID,
		Name:      event.Name,
		StartDate: event.StartDate,
		EndDate:   event.EndDate,
		Photo:     event.Photo,
		Company: struct {
			UUID string `json:"uuid"`
			Name string `json:"name"`
		}{
			UUID: event.Company.UUID,
			Name: event.Company.Name,
		},
	}

	return eventResponse, nil
}

func (s *eventService) Delete(uuid string) (response.EventResponse, error) {
	event, err := s.repo.GetByUUID(uuid)
	if err != nil {
		return response.EventResponse{}, err
	}

	if err := s.repo.Delete(uuid); err != nil {
		return response.EventResponse{}, err
	}

	eventResponse := response.EventResponse{
		UUID:      event.UUID,
		Name:      event.Name,
		StartDate: event.StartDate,
		EndDate:   event.EndDate,
		Photo:     event.Photo,
		Company: struct {
			UUID string `json:"uuid"`
			Name string `json:"name"`
		}{
			UUID: event.Company.UUID,
			Name: event.Company.Name,
		},
	}

	return eventResponse, nil
}
