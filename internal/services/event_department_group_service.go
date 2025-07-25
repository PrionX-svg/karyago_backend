package services

import (
	"fmt"
	"hris_backend/internal/models"
	"hris_backend/internal/repositories"
	"hris_backend/internal/request"
	"hris_backend/internal/response"

	"github.com/google/uuid"
)

type EventDepartmentGroupServices interface {
	Create(userID uint, req request.EventDepartmentGroupRequest) (response.EventDepartmentGroupResponse, error)
	GetAll() ([]response.EventDepartmentGroupResponse, error)
	GetByUUID(uuid string) (response.EventDepartmentGroupResponse, error)
	Update(userID uint, uuid string, req request.EventDepartmentGroupRequest) (response.EventDepartmentGroupResponse, error)	
	Delete(uuid string) (response.EventDepartmentGroupResponse, error)
	GetByEventUUID(eventUUID string) ([]response.EventDepartmentGroupResponse, error)
}

type eventDepartmentGroupServices struct {
	groupRepo repositories.EventDepartmentGroupRepository
	userRepo  repositories.UserRepository
}

func NewEventDepartmentGroupService(
	groupRepo repositories.EventDepartmentGroupRepository,
	userRepo repositories.UserRepository,
) EventDepartmentGroupServices {
	return &eventDepartmentGroupServices{
		groupRepo: groupRepo,
		userRepo:  userRepo,
	}
}

func (s *eventDepartmentGroupServices) Create(userID uint, req request.EventDepartmentGroupRequest) (response.EventDepartmentGroupResponse, error) {
	eventID, err := s.groupRepo.FindEventIDByUUID(req.EventUUID)
	if err != nil {
		return response.EventDepartmentGroupResponse{}, err
	}

	var responsibleID *uint
	if req.ResponsibleUUID != nil {
		responsibleUser, err := s.userRepo.GetByUUID(*req.ResponsibleUUID)
		if err != nil {
			return response.EventDepartmentGroupResponse{}, err
		}
		responsibleID = &responsibleUser.ID
	}

	group := models.EventDepartmentGroup{
		UUID:          uuid.NewString(),
		EventID:       eventID,
		Name:          req.Name,
		Description:   req.Description,
		ResponsibleID: responsibleID,
		CreatedBy:     userID,
		ModifyBy:      userID,
	}

	created, err := s.groupRepo.Create(&group)
	if err != nil {
		return response.EventDepartmentGroupResponse{}, err
	}

	var responsible *response.UserBasicResponse
	if created.Responsible != nil {
		responsible = &response.UserBasicResponse{
			UUID: created.Responsible.UUID,
			Name: fmt.Sprintf("%s %s", created.Responsible.FirstName, created.Responsible.LastName),	
		}
	}

	return response.EventDepartmentGroupResponse{
		UUID:        created.UUID,
		EventUUID:   req.EventUUID,
		Name:        created.Name,
		Description: created.Description,
		Responsible: responsible,
	}, nil
}

func (s *eventDepartmentGroupServices) GetAll() ([]response.EventDepartmentGroupResponse, error) {
	groups, err := s.groupRepo.GetAll()
	if err != nil {
		return nil, err
	}

	var result []response.EventDepartmentGroupResponse
	for _, g := range groups {
		eventUUID, _ := s.groupRepo.FindEventUUIDByID(g.EventID) 

		var responsible *response.UserBasicResponse
		if g.Responsible != nil {
			responsible = &response.UserBasicResponse{
				UUID:      g.Responsible.UUID,
				Name: fmt.Sprintf("%s %s", g.Responsible.FirstName, g.Responsible.LastName),
			}
		}

		result = append(result, response.EventDepartmentGroupResponse{
			UUID:        g.UUID,
			EventUUID:   eventUUID,
			Name:        g.Name,
			Description: g.Description,
			Responsible: responsible,
		})
	}

	return result, nil
}

func (s *eventDepartmentGroupServices) GetByUUID(uuid string) (response.EventDepartmentGroupResponse, error) {
	group, err := s.groupRepo.GetByUUID(uuid)
	if err != nil {
		return response.EventDepartmentGroupResponse{}, err
	}

	eventUUID, _ := s.groupRepo.FindEventUUIDByID(group.EventID)

	var responsible *response.UserBasicResponse
	if group.Responsible != nil {
		responsible = &response.UserBasicResponse{
			UUID:      group.Responsible.UUID,
			Name: fmt.Sprintf("%s %s", group.Responsible.FirstName, group.Responsible.LastName),
		}
	}

	return response.EventDepartmentGroupResponse{
		UUID:        group.UUID,
		EventUUID:   eventUUID,
		Name:        group.Name,
		Description: group.Description,
		Responsible: responsible,
	}, nil
}

func (s *eventDepartmentGroupServices) Update(userID uint, uuid string, req request.EventDepartmentGroupRequest) (response.EventDepartmentGroupResponse, error) {
	group, err := s.groupRepo.GetByUUID(uuid)
	if err != nil {
		return response.EventDepartmentGroupResponse{}, err
	}

	group.Name = req.Name
	group.Description = req.Description
	group.ModifyBy = userID

	if req.ResponsibleUUID != nil {
		responsibleUser, err := s.userRepo.GetByUUID(*req.ResponsibleUUID)
		if err != nil {
			return response.EventDepartmentGroupResponse{}, err
		}
		group.ResponsibleID = &responsibleUser.ID
	} else {
		group.ResponsibleID = nil
	}

	updated, err := s.groupRepo.Update(group)
	if err != nil {
		return response.EventDepartmentGroupResponse{}, err
	}

	eventUUID, _ := s.groupRepo.FindEventUUIDByID(updated.EventID)

	var responsible *response.UserBasicResponse
	if updated.Responsible != nil {
		responsible = &response.UserBasicResponse{
			UUID:      updated.Responsible.UUID,
			Name: fmt.Sprintf("%s %s", updated.Responsible.FirstName, updated.Responsible.LastName),
		}
	}

	return response.EventDepartmentGroupResponse{
		UUID:        updated.UUID,
		EventUUID:   eventUUID,
		Name:        updated.Name,
		Description: updated.Description,
		Responsible: responsible,
	}, nil
}

func (s *eventDepartmentGroupServices) Delete(uuid string) (response.EventDepartmentGroupResponse, error) {
	group, err := s.groupRepo.GetByUUID(uuid)
	if err != nil {
		return response.EventDepartmentGroupResponse{}, err
	}

	if err := s.groupRepo.Delete(group); err != nil {
		return response.EventDepartmentGroupResponse{}, err
	}

	eventUUID, _ := s.groupRepo.FindEventUUIDByID(group.EventID)

	var responsible *response.UserBasicResponse
	if group.Responsible != nil {
		responsible = &response.UserBasicResponse{
			UUID:      group.Responsible.UUID,
			Name: fmt.Sprintf("%s %s", group.Responsible.FirstName, group.Responsible.LastName),
		}
	}

	return response.EventDepartmentGroupResponse{
		UUID:        group.UUID,
		EventUUID:   eventUUID,
		Name:        group.Name,
		Description: group.Description,
		Responsible: responsible,
	}, nil
}

func (s *eventDepartmentGroupServices) GetByEventUUID(eventUUID string) ([]response.EventDepartmentGroupResponse, error) {
	eventID, err := s.groupRepo.FindEventIDByUUID(eventUUID)
	if err != nil {
		return nil, err
	}

	groups, err := s.groupRepo.FindAllByEventID(eventID)
	if err != nil {
		return nil, err
	}

	var result []response.EventDepartmentGroupResponse
	for _, group := range groups {
		var responsible *response.UserBasicResponse
		if group.Responsible != nil {
			responsible = &response.UserBasicResponse{
				UUID: group.Responsible.UUID,
				Name: fmt.Sprintf("%s %s", group.Responsible.FirstName, group.Responsible.LastName),
			}
		}

		result = append(result, response.EventDepartmentGroupResponse{
			UUID:        group.UUID,
			EventUUID:   eventUUID,
			Name:        group.Name,
			Description: group.Description,
			Responsible: responsible,
		})
	}

	return result, nil
}
