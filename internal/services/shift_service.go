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

type ShiftService interface {
	Create(req request.ShiftRequest) (response.ShiftResponse, error)
	GetAll() ([]response.ShiftResponse, error)
	GetByUUID(uuid string) (response.ShiftResponse, error)
	Update(uuid string, req request.ShiftRequest) (response.ShiftResponse, error)
	Delete(uuid string) error
}

type shiftService struct {
	shiftRepo   repositories.ShiftRepository
	companyRepo repositories.CompanyRepositories
}

func NewShiftService(shiftRepo repositories.ShiftRepository, companyRepo repositories.CompanyRepositories) ShiftService {
	return &shiftService{
		shiftRepo:   shiftRepo,
		companyRepo: companyRepo,
	}
}

func (s *shiftService) Create(req request.ShiftRequest) (response.ShiftResponse, error) {
	if err := pkg.Validate.Struct(req); err != nil {
		return response.ShiftResponse{}, err
	}

	company, err := s.companyRepo.GetByUUID(req.CompanyUUID)
	if err != nil {
		return response.ShiftResponse{}, fmt.Errorf("company not found: %w", err)
	}

	shift := models.Shift{
		UUID:       uuid.NewString(),
		CompanyID:  company.ID,
		Name:       req.Name,
		Date:       req.Date,
		HourFrom:   req.HourFrom,
		HourUntil:  req.HourUntil,
		CreatedBy:  req.CreatedBy,
		ModifiedBy: &req.CreatedBy,
	}

	if err := s.shiftRepo.Create(&shift); err != nil {
		return response.ShiftResponse{}, err
	}

	return response.ShiftResponse{
		UUID:      shift.UUID,
		Name:      shift.Name,
		Date:      shift.Date,
		HourFrom:  shift.HourFrom,
		HourUntil: shift.HourUntil,
	}, nil
}

func (s *shiftService) GetAll() ([]response.ShiftResponse, error) {
	shifts, err := s.shiftRepo.FindAll()
	if err != nil {
		return nil, err
	}

	var result []response.ShiftResponse
	for _, shift := range shifts {
		result = append(result, response.ShiftResponse{
			UUID:      shift.UUID,
			Name:      shift.Name,
			Date:      shift.Date,
			HourFrom:  shift.HourFrom,
			HourUntil: shift.HourUntil,
		})
	}
	return result, nil
}

func (s *shiftService) GetByUUID(uuid string) (response.ShiftResponse, error) {
	shift, err := s.shiftRepo.FindByUUID(uuid)
	if err != nil {
		return response.ShiftResponse{}, err
	}

	return response.ShiftResponse{
		UUID:      shift.UUID,
		Name:      shift.Name,
		Date:      shift.Date,
		HourFrom:  shift.HourFrom,
		HourUntil: shift.HourUntil,
	}, nil
}

func (s *shiftService) Update(uuid string, req request.ShiftRequest) (response.ShiftResponse, error) {
	if err := pkg.Validate.Struct(req); err != nil {
		return response.ShiftResponse{}, err
	}

	shift, err := s.shiftRepo.FindByUUID(uuid)
	if err != nil {
		return response.ShiftResponse{}, err
	}

	shift.Name = req.Name
	shift.Date = req.Date
	shift.HourFrom = req.HourFrom
	shift.HourUntil = req.HourUntil
	shift.ModifiedBy = &req.ModifiedBy

	if err := s.shiftRepo.Update(shift); err != nil {
		return response.ShiftResponse{}, err
	}

	return response.ShiftResponse{
		UUID:      shift.UUID,
		Name:      shift.Name,
		Date:      shift.Date,
		HourFrom:  shift.HourFrom,
		HourUntil: shift.HourUntil,
	}, nil
}

func (s *shiftService) Delete(uuid string) error {
	shift, err := s.shiftRepo.FindByUUID(uuid)
	if err != nil {
		return err
	}

	return s.shiftRepo.Delete(shift.ID)
}
