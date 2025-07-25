package services

import (
	"hris_backend/internal/models"
	"hris_backend/internal/repositories"
	"hris_backend/internal/request"
	"hris_backend/internal/response"

	"github.com/google/uuid"
)

type EventDepartmentServices interface {
	Create(req request.EventDepartmentRequest) (response.EventDepartmentResponse, error)
	GetAll() ([]response.EventDepartmentResponse, error)
	GetByUUID(uuid string) (response.EventDepartmentResponse, error)
	Update(uuid string, req request.EventDepartmentRequest) (response.EventDepartmentResponse, error)
	Delete(uuid string) (response.EventDepartmentResponse, error)
	GetByGroupID(groupID uint) ([]response.EventDepartmentResponse, error)
}

type eventDepartmentServices struct {
	deptRepo   repositories.EventDepartmentRepository
	groupRepo  repositories.EventDepartmentGroupRepository
	userRepo   repositories.UserRepository
}

func NewEventDepartmentService(
	deptRepo repositories.EventDepartmentRepository,
	groupRepo repositories.EventDepartmentGroupRepository,
	userRepo repositories.UserRepository,
) EventDepartmentServices {
	return &eventDepartmentServices{
		deptRepo:  deptRepo,
		groupRepo: groupRepo,
		userRepo:  userRepo,
	}
}

func (s *eventDepartmentServices) Create(req request.EventDepartmentRequest) (response.EventDepartmentResponse, error) {
	user, err := s.userRepo.GetByUUID(req.UserUUID)
	if err != nil {
		return response.EventDepartmentResponse{}, err
	}

	group, err := s.groupRepo.GetByUUID(req.GroupUUID)
	if err != nil {
		return response.EventDepartmentResponse{}, err
	}

	dept := models.EventDepartment{
		UUID:                  uuid.NewString(),
		Name:                  req.Name,
		Description:           req.Description,
		EventDepartmentGroupID: group.ID,
		CreatedBy:             user.ID,
		ModifyBy:              user.ID,
	}

	created, err := s.deptRepo.Create(&dept)
	if err != nil {
		return response.EventDepartmentResponse{}, err
	}

	return response.EventDepartmentResponse{
		UUID:        created.UUID,
		GroupUUID:   req.GroupUUID,
		Name:        created.Name,
		Description: created.Description,
	}, nil
}

func (s *eventDepartmentServices) GetAll() ([]response.EventDepartmentResponse, error) {
	depts, err := s.deptRepo.GetAll()
	if err != nil {
		return nil, err
	}

	var result []response.EventDepartmentResponse
	for _, d := range depts {
		groupUUID := ""
		if d.EventDepartmentGroup != nil {
			groupUUID = d.EventDepartmentGroup.UUID
		}

		result = append(result, response.EventDepartmentResponse{
			UUID:        d.UUID,
			GroupUUID:   groupUUID,
			Name:        d.Name,
			Description: d.Description,
		})
	}
	return result, nil
}

func (s *eventDepartmentServices) GetByUUID(uuid string) (response.EventDepartmentResponse, error) {
	dept, err := s.deptRepo.GetByUUID(uuid)
	if err != nil {
		return response.EventDepartmentResponse{}, err
	}

	groupUUID := ""
	if dept.EventDepartmentGroup != nil {
		groupUUID = dept.EventDepartmentGroup.UUID
	}

	return response.EventDepartmentResponse{
		UUID:        dept.UUID,
		GroupUUID:   groupUUID,
		Name:        dept.Name,
		Description: dept.Description,
	}, nil
}

func (s *eventDepartmentServices) Update(uuid string, req request.EventDepartmentRequest) (response.EventDepartmentResponse, error) {
	dept, err := s.deptRepo.GetByUUID(uuid)
	if err != nil {
		return response.EventDepartmentResponse{}, err
	}

	user, err := s.userRepo.GetByUUID(req.UserUUID)
	if err != nil {
		return response.EventDepartmentResponse{}, err
	}

	group, err := s.groupRepo.GetByUUID(req.GroupUUID)
	if err != nil {
		return response.EventDepartmentResponse{}, err
	}

	dept.Name = req.Name
	dept.Description = req.Description
	dept.EventDepartmentGroupID = group.ID
	dept.ModifyBy = user.ID

	updated, err := s.deptRepo.Update(dept)
	if err != nil {
		return response.EventDepartmentResponse{}, err
	}

	return response.EventDepartmentResponse{
		UUID:        updated.UUID,
		GroupUUID:   group.UUID,
		Name:        updated.Name,
		Description: updated.Description,
	}, nil
}

func (s *eventDepartmentServices) Delete(uuid string) (response.EventDepartmentResponse, error) {
	dept, err := s.deptRepo.GetByUUID(uuid)
	if err != nil {
		return response.EventDepartmentResponse{}, err
	}

	if err := s.deptRepo.Delete(dept); err != nil {
		return response.EventDepartmentResponse{}, err
	}

	groupUUID := ""
	if dept.EventDepartmentGroup != nil {
		groupUUID = dept.EventDepartmentGroup.UUID
	}

	return response.EventDepartmentResponse{
		UUID:        dept.UUID,
		GroupUUID:   groupUUID,
		Name:        dept.Name,
		Description: dept.Description,
	}, nil
}

func (s *eventDepartmentServices) GetByGroupID(groupID uint) ([]response.EventDepartmentResponse, error) {
	depts, err := s.deptRepo.FindAllByGroupID(groupID)
	if err != nil {
		return nil, err
	}

	var result []response.EventDepartmentResponse
	for _, d := range depts {
		groupUUID := ""
		if d.EventDepartmentGroup != nil {
			groupUUID = d.EventDepartmentGroup.UUID
		}

		result = append(result, response.EventDepartmentResponse{
			UUID:        d.UUID,
			GroupUUID:   groupUUID,
			Name:        d.Name,
			Description: d.Description,
		})
	}
	return result, nil
}

