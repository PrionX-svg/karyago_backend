package services

import (
	"fmt"
	"time"

	"gorm.io/gorm"

	"hris_backend/internal/models"
	"hris_backend/internal/repositories"
)

type NotificationService interface {
	SendEmail(to, subject, htmlBody string) error
}

type CreateEditReq struct {
	WorkDate             time.Time
	RequestType          models.EditRequestType
	ProposedClockInAt    *time.Time
	ProposedClockOutAt   *time.Time
	ProposedIsHomeOffice *bool
	Reason               string
}

type AttendanceEditService interface {
	Create(userID uint, p CreateEditReq) (uint, error)
	ListMine(userID uint, status *models.EditRequestStatus, from, to *time.Time) ([]models.AttendanceEditRequest, error)
	ListForSupervisor(userID uint, status *models.EditRequestStatus, from, to *time.Time, q *string, limit, offset int) ([]models.AttendanceEditRequest, int64, error)
	Approve(userID uint, reqID uint, note *string) error
	Reject(userID uint, reqID uint, note string) error
}

type attendanceEditService struct {
	db          *gorm.DB
	editRepo    repositories.AttendanceEditRepository
	attRepo     repositories.AttendanceRepository
	empRepo     repositories.EmployeeRepository
	userRepo    repositories.UserRepository
	companyRepo repositories.CompanyRepositories
	notify      NotificationService
}

func NewAttendanceEditService(
	db *gorm.DB,
	editRepo repositories.AttendanceEditRepository,
	attRepo repositories.AttendanceRepository,
	empRepo repositories.EmployeeRepository,
	userRepo repositories.UserRepository,
	companyRepo repositories.CompanyRepositories,
	notify NotificationService,
) AttendanceEditService {
	return &attendanceEditService{db, editRepo, attRepo, empRepo, userRepo, companyRepo, notify}
}

/* ---------------- helpers ---------------- */

func (s *attendanceEditService) getEmployee(userID uint) (*models.Employee, error) {
	emp, err := s.empRepo.FindByUserID(userID)
	if err != nil || emp == nil {
		return nil, fmt.Errorf("employee not found")
	}
	return emp, nil
}

func validateCreatePayload(p CreateEditReq) error {
	switch p.RequestType {
	case models.EditReqClockIn:
		if p.ProposedClockInAt == nil {
			return fmt.Errorf("proposed_clock_in_at is required")
		}
	case models.EditReqClockOut:
		if p.ProposedClockOutAt == nil {
			return fmt.Errorf("proposed_clock_out_at is required")
		}
	case models.EditReqBoth:
		if p.ProposedClockInAt == nil && p.ProposedClockOutAt == nil {
			return fmt.Errorf("at least one of proposed_clock_in_at/proposed_clock_out_at is required")
		}
	case models.EditReqHomeFlag:
		if p.ProposedIsHomeOffice == nil {
			return fmt.Errorf("proposed_is_home_office is required")
		}
	case models.EditReqBothPlusFlag:
		if p.ProposedIsHomeOffice == nil && p.ProposedClockInAt == nil && p.ProposedClockOutAt == nil {
			return fmt.Errorf("at least one of time/flag is required")
		}
	default:
		return fmt.Errorf("invalid request_type")
	}
	if p.Reason == "" {
		return fmt.Errorf("reason is required")
	}
	return nil
}

/* ---------------- usecases ---------------- */

func (s *attendanceEditService) Create(userID uint, p CreateEditReq) (uint, error) {
	if err := validateCreatePayload(p); err != nil {
		return 0, err
	}
	emp, err := s.getEmployee(userID)
	if err != nil {
		return 0, err
	}
	req := &models.AttendanceEditRequest{
		EmployeeID:           emp.ID,
		CompanyID:            *emp.CompanyID,
		WorkDate:             p.WorkDate,
		RequestType:          p.RequestType,
		ProposedClockInAt:    p.ProposedClockInAt,
		ProposedClockOutAt:   p.ProposedClockOutAt,
		ProposedIsHomeOffice: p.ProposedIsHomeOffice,
		Reason:               p.Reason,
		Status:               models.EditStatusPending,
	}
	if err := s.editRepo.Create(req); err != nil {
		return 0, err
	}
	return req.ID, nil
}

func (s *attendanceEditService) ListMine(userID uint, status *models.EditRequestStatus, from, to *time.Time) ([]models.AttendanceEditRequest, error) {
	emp, err := s.getEmployee(userID)
	if err != nil {
		return nil, err
	}
	return s.editRepo.ListMine(emp.ID, status, from, to)
}

func (s *attendanceEditService) ListForSupervisor(userID uint, status *models.EditRequestStatus, from, to *time.Time, q *string, limit, offset int) ([]models.AttendanceEditRequest, int64, error) {
	reviewer, err := s.getEmployee(userID)
	if err != nil || reviewer.CompanyID == nil {
		return nil, 0, fmt.Errorf("forbidden")
	}
	return s.editRepo.ListForSupervisor(*reviewer.CompanyID, status, from, to, q, limit, offset)
}

func (s *attendanceEditService) Approve(userID uint, reqID uint, note *string) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		req, err := s.editRepo.FindByID(reqID)
		if err != nil || req == nil || req.Status != models.EditStatusPending {
			return fmt.Errorf("invalid request")
		}
		reviewer, err := s.getEmployee(userID)
		if err != nil || reviewer.CompanyID == nil || *reviewer.CompanyID != req.CompanyID {
			return fmt.Errorf("forbidden")
		}

		// Pastikan ada attendance di tanggal tsb
		att, err := s.attRepo.EnsureForDay(req.EmployeeID, req.CompanyID, req.WorkDate)
		if err != nil {
			return err
		}

		// Terapkan perubahan sesuai tipe
		if req.RequestType == models.EditReqClockIn || req.RequestType == models.EditReqBoth || req.RequestType == models.EditReqBothPlusFlag {
			if req.ProposedClockInAt != nil {
				att.ClockInAt = req.ProposedClockInAt
				if err := tx.Model(att).Update("clock_in_at", att.ClockInAt).Error; err != nil {
					return err
				}
			}
		}
		if req.RequestType == models.EditReqClockOut || req.RequestType == models.EditReqBoth || req.RequestType == models.EditReqBothPlusFlag {
			if req.ProposedClockOutAt != nil {
				att.ClockOutAt = req.ProposedClockOutAt
				if err := tx.Model(att).Update("clock_out_at", att.ClockOutAt).Error; err != nil {
					return err
				}
			}
		}
		if req.RequestType == models.EditReqHomeFlag || req.RequestType == models.EditReqBothPlusFlag {
			if req.ProposedIsHomeOffice != nil {
				att.Is_homeOffice = *req.ProposedIsHomeOffice
				if err := tx.Model(att).Update("is_home_office", att.Is_homeOffice).Error; err != nil {
					return err
				}
			}
		}

		// Set status APPROVED
		if _, err := s.editRepo.Approve(tx, req.ID, userID, note); err != nil {
			return err
		}

		// Notifikasi email
		if s.notify != nil {
			if email, _ := s.empRepo.GetUserEmailByEmployeeID(req.EmployeeID); email != nil && *email != "" {
				_ = s.notify.SendEmail(*email, "Attendance Edit Approved",
					fmt.Sprintf("Permintaan edit %s telah disetujui.", req.WorkDate.Format("2006-01-02")))
			}
		}
		return nil
	})
}

func (s *attendanceEditService) Reject(userID uint, reqID uint, note string) error {
	if note == "" {
		return fmt.Errorf("note is required")
	}
	req, err := s.editRepo.FindByID(reqID)
	if err != nil || req == nil || req.Status != models.EditStatusPending {
		return fmt.Errorf("invalid request")
	}
	reviewer, err := s.getEmployee(userID)
	if err != nil || reviewer.CompanyID == nil || *reviewer.CompanyID != req.CompanyID {
		return fmt.Errorf("forbidden")
	}
	if err := s.editRepo.Reject(req.ID, userID, note); err != nil {
		return err
	}
	if s.notify != nil {
		if email, _ := s.empRepo.GetUserEmailByEmployeeID(req.EmployeeID); email != nil && *email != "" {
			_ = s.notify.SendEmail(*email, "Attendance Edit Rejected",
				fmt.Sprintf("Permintaan edit %s ditolak.", req.WorkDate.Format("2006-01-02")))
		}
	}
	return nil
}
