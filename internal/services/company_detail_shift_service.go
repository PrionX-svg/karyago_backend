package services

import (
	"errors"
	"hris_backend/internal/models"
	"hris_backend/internal/repositories"
	"hris_backend/internal/response"
)

type CompanyDetailShiftService interface {
	Create(detail *models.CompanyDetailShift) (*response.CompanyDetailShiftWithShiftResponse, error)
	Update(uuid string, updated *models.CompanyDetailShift) (*response.CompanyDetailShiftWithShiftResponse, error)
	Delete(uuid string) error
	GetByUUID(uuid string) (*response.CompanyDetailShiftWithShiftResponse, error)
	GetAll() ([]response.CompanyDetailShiftWithShiftResponse, error)
	GetByShiftUUID(shiftUUID string) ([]response.CompanyDetailShiftWithShiftResponse, error)
	FindShiftByUUID(uuid string) (*models.Shift, error)
}

type companyDetailShiftService struct {
	repo      repositories.CompanyDetailShiftRepository
	shiftRepo repositories.ShiftRepository
}

func NewCompanyDetailShiftService(
	repo repositories.CompanyDetailShiftRepository,
	shiftRepo repositories.ShiftRepository,
) CompanyDetailShiftService {
	return &companyDetailShiftService{
		repo:      repo,
		shiftRepo: shiftRepo,
	}
}

func (s *companyDetailShiftService) Create(detail *models.CompanyDetailShift) (*response.CompanyDetailShiftWithShiftResponse, error) {
	if err := s.repo.Create(detail); err != nil {
		return nil, err
	}

	shift, err := s.shiftRepo.FindByID(detail.ShiftID)
	if err != nil {
		return nil, err
	}

	return &response.CompanyDetailShiftWithShiftResponse{
		UUID:      detail.UUID,
		ShiftUUID: shift.UUID,
		ShiftName: shift.Name,
		Day1:      detail.Day1,
		Day2:      detail.Day2,
		Day3:      detail.Day3,
		Day4:      detail.Day4,
		Day5:      detail.Day5,
		Day6:      detail.Day6,
		Day7:      detail.Day7,
	}, nil
}

func (s *companyDetailShiftService) Update(uuid string, updated *models.CompanyDetailShift) (*response.CompanyDetailShiftWithShiftResponse, error) {
	existing, err := s.repo.FindByUUID(uuid)
	if err != nil {
		return nil, err
	}
	if existing == nil {
		return nil, errors.New("detail shift not found")
	}

	updated.ID = existing.ID
	updated.UUID = existing.UUID
	updated.CreatedAt = existing.CreatedAt
	updated.CreatedBy = existing.CreatedBy

	if err := s.repo.Update(updated); err != nil {
		return nil, err
	}

	shift, err := s.shiftRepo.FindByID(updated.ShiftID)
	if err != nil {
		return nil, err
	}

	return &response.CompanyDetailShiftWithShiftResponse{
		UUID:      updated.UUID,
		ShiftUUID: shift.UUID,
		ShiftName: shift.Name,
		Day1:      updated.Day1,
		Day2:      updated.Day2,
		Day3:      updated.Day3,
		Day4:      updated.Day4,
		Day5:      updated.Day5,
		Day6:      updated.Day6,
		Day7:      updated.Day7,
	}, nil
}

func (s *companyDetailShiftService) Delete(uuid string) error {
	return s.repo.DeleteByUUID(uuid)
}

func (s *companyDetailShiftService) GetByUUID(uuid string) (*response.CompanyDetailShiftWithShiftResponse, error) {
	detail, err := s.repo.FindByUUID(uuid)
	if err != nil {
		return nil, err
	}
	if detail == nil {
		return nil, errors.New("detail shift not found")
	}

	shift, err := s.shiftRepo.FindByID(detail.ShiftID)
	if err != nil {
		return nil, err
	}

	return &response.CompanyDetailShiftWithShiftResponse{
		UUID:      detail.UUID,
		ShiftUUID: shift.UUID,
		ShiftName: shift.Name,
		Day1:      detail.Day1,
		Day2:      detail.Day2,
		Day3:      detail.Day3,
		Day4:      detail.Day4,
		Day5:      detail.Day5,
		Day6:      detail.Day6,
		Day7:      detail.Day7,
	}, nil
}

func (s *companyDetailShiftService) GetAll() ([]response.CompanyDetailShiftWithShiftResponse, error) {
	details, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}

	var responses []response.CompanyDetailShiftWithShiftResponse
	for _, d := range details {
		shift, err := s.shiftRepo.FindByID(d.ShiftID)
		if err != nil {
			continue
		}

		responses = append(responses, response.CompanyDetailShiftWithShiftResponse{
			UUID:      d.UUID,
			ShiftUUID: shift.UUID,
			ShiftName: shift.Name,
			Day1:      d.Day1,
			Day2:      d.Day2,
			Day3:      d.Day3,
			Day4:      d.Day4,
			Day5:      d.Day5,
			Day6:      d.Day6,
			Day7:      d.Day7,
		})
	}

	return responses, nil
}

func (s *companyDetailShiftService) GetByShiftUUID(shiftUUID string) ([]response.CompanyDetailShiftWithShiftResponse, error) {
	shift, err := s.shiftRepo.FindByUUID(shiftUUID)
	if err != nil {
		return nil, err
	}
	if shift == nil {
		return nil, errors.New("shift not found")
	}

	details, err := s.repo.FindByShiftID(shift.ID)
	if err != nil {
		return nil, err
	}

	var responses []response.CompanyDetailShiftWithShiftResponse
	for _, d := range details {
		responses = append(responses, response.CompanyDetailShiftWithShiftResponse{
			UUID:      d.UUID,
			ShiftUUID: shift.UUID,
			ShiftName: shift.Name,
			Day1:      d.Day1,
			Day2:      d.Day2,
			Day3:      d.Day3,
			Day4:      d.Day4,
			Day5:      d.Day5,
			Day6:      d.Day6,
			Day7:      d.Day7,
		})
	}

	return responses, nil
}

func (s *companyDetailShiftService) FindShiftByUUID(uuid string) (*models.Shift, error) {
	return s.shiftRepo.FindByUUID(uuid)
}
