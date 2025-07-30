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

type EventShiftService interface {
	Create(req request.EventShiftReq, actorID uint) (response.EventShiftResponse, error)
	GetAll() ([]response.EventShiftResponse, error)
	GetByID(id uint) (response.EventShiftResponse, error)
	GetByUUID(uuid string) (response.EventShiftResponse, error)
	GetByEventUUID(uuid string) ([]response.EventShiftResponse, error)
	Update(uuid string, req request.EventShiftReq, actorID uint) (response.EventShiftResponse, error)
	Delete(uuid string) (response.EventShiftResponse, error)
}

type eventShiftService struct {
	repo      repositories.EventShiftRepository
	eventRepo repositories.EventRepository
}

func NewEventShiftService(repo repositories.EventShiftRepository, eventRepo repositories.EventRepository) EventShiftService {
	return &eventShiftService{
		repo:      repo,
		eventRepo: eventRepo,
	}
}

func (s *eventShiftService) Create(req request.EventShiftReq, actorID uint) (response.EventShiftResponse, error) {
	if err := pkg.Validate.Struct(req); err != nil {
		return response.EventShiftResponse{}, err
	}

	event, err := s.eventRepo.GetByUUID(req.EventUUID)
	if err != nil {
		return response.EventShiftResponse{}, fmt.Errorf("invalid event UUID")
	}

	shift := &models.EventShift{
		UUID:      uuid.NewString(),
		EventID:   event.ID,
		Name:      req.Name,
		Date:      req.Date,
		HourFrom:  req.HourFrom,
		HourUntil: req.HourUntil,
		CreatedBy: actorID,
		ModifyBy:  actorID,
	}

	if err := s.repo.Create(shift); err != nil {
		return response.EventShiftResponse{}, err
	}

	created, err := s.repo.GetByUUID(shift.UUID)
	if err != nil {
		return response.EventShiftResponse{}, err
	}

	return response.EventShiftResponse{
		UUID:      created.UUID,
		Name:      created.Name,
		Date:      req.Date,
		HourFrom:  req.HourFrom,
		HourUntil: req.HourUntil,
		Event: struct {
			UUID string `json:"uuid"`
			Name string `json:"name"`
		}{
			UUID: created.Event.UUID,
			Name: created.Event.Name,
		},
	}, nil
}

func (s *eventShiftService) GetAll() ([]response.EventShiftResponse, error) {
	shifts, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	var responses []response.EventShiftResponse
	for _, shift := range shifts {
		responses = append(responses, response.EventShiftResponse{
			UUID:      shift.UUID,
			Name:      shift.Name,
			Date :      shift.Date,
			HourFrom:  shift.HourFrom,
			HourUntil: shift.HourUntil,
			Event: struct {
				UUID string `json:"uuid"`
				Name string `json:"name"`
			}{
				UUID: shift.Event.UUID,
				Name: shift.Event.Name,
			},
		})
	}

	return responses, nil
}

func (s *eventShiftService) GetByID(id uint) (response.EventShiftResponse, error) {
	shift, err := s.repo.GetByID(id)
	if err != nil {
		return response.EventShiftResponse{}, err
	}

	return response.EventShiftResponse{
		UUID:      shift.UUID,
		Name:      shift.Name,
		Date: 	shift.Date,
		HourFrom:  shift.HourFrom,
		HourUntil: shift.HourUntil,
		Event: struct {
			UUID string `json:"uuid"`
			Name string `json:"name"`
		}{
			UUID: shift.Event.UUID,
			Name: shift.Event.Name,
		},
	}, nil
}

func (s *eventShiftService) GetByUUID(uuid string) (response.EventShiftResponse, error) {
	shift, err := s.repo.GetByUUID(uuid)
	if err != nil {
		return response.EventShiftResponse{}, err
	}

	return response.EventShiftResponse{
		UUID:      shift.UUID,
		Name:      shift.Name,
		Date: shift.Date,
		HourFrom:  shift.HourFrom,
		HourUntil: shift.HourUntil,
		Event: struct {
			UUID string `json:"uuid"`
			Name string `json:"name"`
		}{
			UUID: shift.Event.UUID,
			Name: shift.Event.Name,
		},
	}, nil
}

func (s *eventShiftService) GetByEventUUID(uuid string) ([]response.EventShiftResponse, error) {
	shifts, err := s.repo.GetByEventUUID(uuid)
	if err != nil {
		return nil, err
	}

	var responses []response.EventShiftResponse
	for _, shift := range shifts {
		responses = append(responses, response.EventShiftResponse{
			UUID:      shift.UUID,
			Name:      shift.Name,
			Date:      shift.Date,
			HourFrom:  shift.HourFrom,
			HourUntil: shift.HourUntil,
			Event: struct {
				UUID string `json:"uuid"`
				Name string `json:"name"`
			}{
				UUID: shift.Event.UUID,
				Name: shift.Event.Name,
			},
		})
	}

	return responses, nil
}

func (s *eventShiftService) Update(uuid string, req request.EventShiftReq, actorID uint) (response.EventShiftResponse, error) {
	fmt.Print("Update Event Shift Service dengan uuid : ", uuid)
	shift, err := s.repo.GetByUUID(uuid)
	if err != nil {
		return response.EventShiftResponse{}, err
	}
	fmt.Print("\n\nberhasil get : ", uuid)
	
	shift.Name = req.Name
	shift.Date = req.Date
	shift.HourFrom = req.HourFrom
	shift.HourUntil = req.HourUntil
	shift.ModifyBy = actorID
	
	if err := s.repo.Update(shift); err != nil {
		fmt.Print("\nwah error!!!!")
		return response.EventShiftResponse{}, err
	}
	fmt.Print("\n\nberhasil update : ", uuid)

	return response.EventShiftResponse{
		UUID:      shift.UUID,
		Name:      shift.Name,
		Date:      shift.Date,
		HourFrom:  shift.HourFrom,
		HourUntil: shift.HourUntil,
		Event: struct {
			UUID string `json:"uuid"`
			Name string `json:"name"`
		}{
			UUID: shift.Event.UUID,
			Name: shift.Event.Name,
		},
	}, nil
}

func (s *eventShiftService) Delete(uuid string) (response.EventShiftResponse, error) {
	shift, err := s.repo.GetByUUID(uuid)
	if err != nil {
		return response.EventShiftResponse{}, err
	}

	if err := s.repo.Delete(uuid); err != nil {
		return response.EventShiftResponse{}, err
	}

	return response.EventShiftResponse{
		UUID:      shift.UUID,
		Name:      shift.Name,
		Date:      shift.Date,
		HourFrom:  shift.HourFrom,
		HourUntil: shift.HourUntil,
		Event: struct {
			UUID string `json:"uuid"`
			Name string `json:"name"`
		}{
			UUID: shift.Event.UUID,
			Name: shift.Event.Name,
		},
	}, nil
}
