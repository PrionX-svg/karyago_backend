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

type EventUserService interface {
}

type eventUserService struct {
	repo      repositories.EventUserRepository
	userRepo  repositories.UserRepository
	eventRepo repositories.EventRepository
}

func NewEventUserService(repo repositories.EventUserRepository, userRepo repositories.UserRepository, eventRepo repositories.EventRepository) EventUserService {
	return &eventUserService{
		repo:      repo,
		userRepo:  userRepo,
		eventRepo: eventRepo,
	}
}

func (s *eventUserService) Create(request request.EventUserReq, userID uint) (response.EventUserResponse, error) {
	if err := pkg.Validate.Struct(request); err != nil {
		return response.EventUserResponse{}, err
	}

	user, err := s.userRepo.GetByUUID(request.UserUUID)
	if err != nil {
		return response.EventUserResponse{}, fmt.Errorf("Invalid user UUID")
	}

	event, err := s.eventRepo.GetByUUID(request.EventUUID)
	if err != nil {
		return response.EventUserResponse{}, fmt.Errorf("Invalid event UUID")
	}

	eventUser := &models.EventUser{
		UUID:      uuid.NewString(),
		UserID:    user.ID,
		EventID:   event.ID,
		CreatedBy: userID,
		ModifyBy:  userID,
	}

	if err := s.repo.Create(eventUser); err != nil {
		return response.EventUserResponse{}, err
	}

	eventUserModel, err := s.repo.GetByUUID(eventUser.UUID)
	if err != nil {
		return response.EventUserResponse{}, err
	}


	eventUserResponse := response.EventUserResponse{

		UUID: eventUserModel.UUID,
		User: struct {
			UUID     string `json:"uuid"`
			FullName string `json:"fullname"`
			Email    string `json:"email"`
		}{
			UUID:     user.UUID,
			FullName: user.FirstName + " " + user.LastName,
			Email:    user.Email,
		},
		Event: struct {
			UUID      string `json:"uuid"`
			Name      string `json:"name"`
			StartDate string `json:"start_date"`
			EndDate   string `json:"end_date"`
		}{
			UUID:      event.UUID,
			Name:      event.Name,
			StartDate: event.StartDate.Format("2006-01-02 15:04:05"),
			EndDate:   event.EndDate.Format("2006-01-02 15:04:05"),
		},
	}

	return eventUserResponse, nil
}

func (s *eventUserService) GetAll() ([]response.EventUserResponse, error) {
	eventUsers, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	var eventUsersResponses []response.EventUserResponse

	for _, evUs := range eventUsers {
		eventUsersResponses = append(eventUsersResponses, response.EventUserResponse{
			UUID: evUs.UUID,
			User: struct {
				UUID     string `json:"uuid"`
				FullName string `json:"fullname"`
				Email    string `json:"email"`
			}{
				UUID:     evUs.User.UUID,
				FullName: evUs.User.FirstName + " " + evUs.User.LastName,
				Email:    evUs.User.Email,
			},
			Event: struct {
				UUID      string `json:"uuid"`
				Name      string `json:"name"`
				StartDate string `json:"start_date"`
				EndDate   string `json:"end_date"`
			}{
				UUID:      evUs.Event.UUID,
				Name:      evUs.Event.Name,
				StartDate: evUs.Event.StartDate.Format("2006-01-02 15:04:05"),
				EndDate:   evUs.Event.EndDate.Format("2006-01-02 15:04:05"),
			},
		})
	}

	return eventUsersResponses, nil
}

func (s *eventUserService) GetByID(id uint) (response.EventUserResponse, error) {
	eventUser, err := s.repo.GetByID(id)
	if err != nil {
		return response.EventUserResponse{}, err
	}


	eventUserResponse := response.EventUserResponse{

		UUID: eventUser.UUID,
		User: struct {
			UUID     string `json:"uuid"`
			FullName string `json:"fullname"`
			Email    string `json:"email"`
		}{
			UUID:     eventUser.User.UUID,
			FullName: eventUser.User.FirstName + " " + eventUser.User.LastName,
			Email:    eventUser.User.Email,
		},
		Event: struct {
			UUID      string `json:"uuid"`
			Name      string `json:"name"`
			StartDate string `json:"start_date"`
			EndDate   string `json:"end_date"`
		}{
			UUID:      eventUser.Event.UUID,
			Name:      eventUser.Event.Name,
			StartDate: eventUser.Event.StartDate.Format("2006-01-02 15:04:05"),
			EndDate:   eventUser.Event.EndDate.Format("2006-01-02 15:04:05"),
		},
	}

	return eventUserResponse, nil
}

func (s *eventUserService) GetByUUID(uuid string) (response.EventUserResponse, error) {
	eventUser, err := s.repo.GetByUUID(uuid)
	if err != nil {
		return response.EventUserResponse{}, err
	}

	eventUserResponse := response.EventUserResponse{

		UUID: eventUser.UUID,
		User: struct {
			UUID     string `json:"uuid"`
			FullName string `json:"fullname"`
			Email    string `json:"email"`
		}{
			UUID:     eventUser.User.UUID,
			FullName: eventUser.User.FirstName + " " + eventUser.User.LastName,
			Email:    eventUser.User.Email,
		},
		Event: struct {
			UUID      string `json:"uuid"`
			Name      string `json:"name"`
			StartDate string `json:"start_date"`
			EndDate   string `json:"end_date"`
		}{
			UUID:      eventUser.Event.UUID,
			Name:      eventUser.Event.Name,
			StartDate: eventUser.Event.StartDate.Format("2006-01-02 15:04:05"),
			EndDate:   eventUser.Event.EndDate.Format("2006-01-02 15:04:05"),
		},
	}
	return eventUserResponse, nil
}

func (s *eventUserService) GetByUserUUID(uuid string) ([]response.EventUserResponse, error) {
	eventUsers, err := s.repo.GetByUserUUID(uuid)
	if err != nil {
		return nil, err
	}

	var eventUserResponses []response.EventUserResponse

	for _, evUs := range eventUsers {
		eventUserResponses = append(eventUserResponses, response.EventUserResponse{
			UUID:      evUs.UUID,
			User: struct {
				UUID     string `json:"uuid"`
				FullName string `json:"fullname"`
				Email    string `json:"email"`
			}{
				UUID:     evUs.User.UUID,
				FullName: evUs.User.FirstName + " " + evUs.User.LastName,
				Email:    evUs.User.Email,
			},
			Event: struct {
				UUID      string `json:"uuid"`
				Name      string `json:"name"`
				StartDate string `json:"start_date"`
				EndDate   string `json:"end_date"`
			}{
				UUID:      evUs.Event.UUID,
				Name:      evUs.Event.Name,
				StartDate: evUs.Event.StartDate.Format("2006-01-02 15:04:05"),
				EndDate:   evUs.Event.EndDate.Format("2006-01-02 15:04:05"),
			},
		})
	}

	return eventUserResponses, nil
}

func (s *eventUserService) GetByEventUUID(uuid string) ([]response.EventUserResponse, error) {
	eventUsers, err := s.repo.GetByEventUUID(uuid)
	if err != nil {
		return nil, err
	}

	var eventUserResponses []response.EventUserResponse

	for _, evUs := range eventUsers {
		eventUserResponses = append(eventUserResponses, response.EventUserResponse{
			UUID:      evUs.UUID,
			User: struct {
				UUID     string `json:"uuid"`
				FullName string `json:"fullname"`
				Email    string `json:"email"`
			}{
				UUID:     evUs.User.UUID,
				FullName: evUs.User.FirstName + " " + evUs.User.LastName,
				Email:    evUs.User.Email,
			},
			Event: struct {
				UUID      string `json:"uuid"`
				Name      string `json:"name"`
				StartDate string `json:"start_date"`
				EndDate   string `json:"end_date"`
			}{
				UUID:      evUs.Event.UUID,
				Name:      evUs.Event.Name,
				StartDate: evUs.Event.StartDate.Format("2006-01-02 15:04:05"),
				EndDate:   evUs.Event.EndDate.Format("2006-01-02 15:04:05"),
			},
		})
	}

	return eventUserResponses, nil
}

func (s *eventUserService) Update(uuid string, req request.EventUserReq, actorID uint) (response.EventUserResponse, error) {
	eventUser, err := s.repo.GetByUUID(uuid)
	if err != nil {
		return response.EventUserResponse{}, err
	}

	user, err := s.userRepo.GetByUUID(req.UserUUID)
	if err != nil {
		return response.EventUserResponse{}, fmt.Errorf("Invalid user UUID")
	}

	event, err := s.eventRepo.GetByUUID(req.EventUUID)
	if err != nil {
		return response.EventUserResponse{}, fmt.Errorf("Invalid event UUID")
	}

	eventUser.UserID = user.ID
	eventUser.EventID = event.ID
	eventUser.ModifyBy = actorID

	if err := s.repo.Update(eventUser); err != nil {
		return response.EventUserResponse{}, err
	}

	return s.GetByUUID(uuid)
}

func (s *eventUserService) Delete(uuid string) (response.EventUserResponse, error) {
	eventUser, err := s.repo.GetByUUID(uuid)
	if err != nil {
		return response.EventUserResponse{}, err
	}

	if err := s.repo.Delete(uuid); err != nil {
		return response.EventUserResponse{}, err
	}

	fullname := eventUser.User.FirstName + " " + eventUser.User.LastName

	eventUserResponse := response.EventUserResponse{
		UUID: eventUser.UUID,
		User: struct {
			UUID     string `json:"uuid"`
			FullName string `json:"fullname"`
			Email    string `json:"email"`
		}{
			UUID:     eventUser.User.UUID,
			FullName: fullname,
			Email:    eventUser.User.Email,
		},
		Event: struct {
			UUID      string `json:"uuid"`
			Name      string `json:"name"`
			StartDate string `json:"start_date"`
			EndDate   string `json:"end_date"`
		}{
			UUID:      eventUser.Event.UUID,
			Name:      eventUser.Event.Name,
			StartDate: eventUser.Event.StartDate.Format("2006-01-02 15:04:05"),
			EndDate:   eventUser.Event.EndDate.Format("2006-01-02 15:04:05"),
		},
	}

	return eventUserResponse, nil
}

