package services

import (
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"

	"hris_backend/internal/models"
	"hris_backend/internal/repositories"
)

type AttendanceService interface {
	// Core actions
	ToggleHomeOffice(userID uint, companyUUID *string, workDate time.Time, isHome bool) (*models.Attendance, error)
	SaveNotes(userID uint, companyUUID *string, workDate time.Time, notes *string) (*models.Attendance, error)
	ClockIn(userID uint, companyUUID *string, workDate time.Time, at time.Time) (*models.Attendance, error)
	ClockOut(userID uint, companyUUID *string, workDate time.Time, at time.Time) (*models.Attendance, error)

	// Queries
	GetByDate(userID uint, companyUUID *string, workDate time.Time) (*models.Attendance, error)
	ListRange(userID uint, companyUUID *string, from, to time.Time) ([]models.Attendance, error)
}

type attendanceService struct {
	db          *gorm.DB
	attRepo     repositories.AttendanceRepository
	empRepo     repositories.EmployeeRepository
	userRepo    repositories.UserRepository
	companyRepo repositories.CompanyRepositories
}

func NewAttendanceService(
	db *gorm.DB,
	attRepo repositories.AttendanceRepository,
	empRepo repositories.EmployeeRepository,
	userRepo repositories.UserRepository,
	companyRepo repositories.CompanyRepositories,
) AttendanceService {
	return &attendanceService{db, attRepo, empRepo, userRepo, companyRepo}
}

// resolve employeeID & companyID dari user serta optional companyUUID (kalau multi-company)
func (s *attendanceService) resolveActor(userID uint, companyUUID *string) (employeeID, companyID uint, err error) {
	emp, err := s.empRepo.FindByUserID(userID)
	if err != nil || emp == nil {
		return 0, 0, fmt.Errorf("employee not found for user")
	}

	// default: pakai company dari employee
	if companyUUID == nil || *companyUUID == "" {
		if emp.CompanyID == nil {
			return 0, 0, fmt.Errorf("employee has no company")
		}
		return emp.ID, *emp.CompanyID, nil
	}

	// override: company dari query (mis. employee bisa masuk >1 company)
	company, err := s.companyRepo.GetByUUID(*companyUUID)
	if err != nil {
		return 0, 0, fmt.Errorf("company not found")
	}
	return emp.ID, company.ID, nil
}

// Toggle Home/Office (bisa sebelum/ sesudah clock-in, tapi ditolak setelah clock-out)
func (s *attendanceService) ToggleHomeOffice(userID uint, companyUUID *string, workDate time.Time, isHome bool) (*models.Attendance, error) {
	empID, compID, err := s.resolveActor(userID, companyUUID)
	if err != nil {
		return nil, err
	}
	return s.attRepo.SetHomeOffice(empID, compID, workDate, isHome)
}

// Notes autosave (bisa diubah sampai clock-out)
func (s *attendanceService) SaveNotes(userID uint, companyUUID *string, workDate time.Time, notes *string) (*models.Attendance, error) {
	empID, compID, err := s.resolveActor(userID, companyUUID)
	if err != nil {
		return nil, err
	}
	return s.attRepo.UpdateNotes(empID, compID, workDate, notes, &userID)
}

// Clock-in (idempotent per employee+work_date)
func (s *attendanceService) ClockIn(userID uint, companyUUID *string, workDate time.Time, at time.Time) (*models.Attendance, error) {
	empID, compID, err := s.resolveActor(userID, companyUUID)
	if err != nil {
		return nil, err
	}
	return s.attRepo.ClockIn(empID, compID, workDate, at)
}

// Clock-out (mengunci toggle & notes)
func (s *attendanceService) ClockOut(userID uint, companyUUID *string, workDate time.Time, at time.Time) (*models.Attendance, error) {
	empID, _, err := s.resolveActor(userID, companyUUID)
	if err != nil {
		return nil, err
	}
	return s.attRepo.ClockOut(empID, workDate, at)
}

func (s *attendanceService) GetByDate(userID uint, companyUUID *string, workDate time.Time) (*models.Attendance, error) {
	empID, _, err := s.resolveActor(userID, companyUUID)
	if err != nil {
		return nil, err
	}
	return s.attRepo.FindByEmployeeAndDate(empID, workDate)
}

func (s *attendanceService) ListRange(userID uint, companyUUID *string, from, to time.Time) ([]models.Attendance, error) {
	empID, _, err := s.resolveActor(userID, companyUUID)
	if err != nil {
		return nil, err
	}
	if to.Before(from) {
		return nil, errors.New("invalid range")
	}
	return s.attRepo.ListByEmployeeBetween(empID, from, to)
}
