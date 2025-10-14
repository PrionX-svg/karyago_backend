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
	ListAllEmployee(userID uint, status *models.EditRequestStatus, from, to *time.Time, q *string, limit, offset int) ([]models.AttendanceEditRequest, int64, error)
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

func (s *attendanceEditService) ListAllEmployee(userID uint, status *models.EditRequestStatus, from, to *time.Time, q *string, limit, offset int) ([]models.AttendanceEditRequest, int64, error) {
	reviewer, err := s.getEmployee(userID)
	if err != nil || reviewer.CompanyID == nil {
		return nil, 0, fmt.Errorf("forbidden")
	}
	return s.editRepo.ListAllEmployee(*reviewer.CompanyID, status, from, to, q, limit, offset)
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
		att, err := s.attRepo.EnsureForDay(req.EmployeeID, req.CompanyID, req.ID, req.WorkDate)
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
				subject := "Attendance Edit Approved"
				body := fmt.Sprintf(`
					<html>
					<head>
						<style>
							body {
								font-family: Arial, sans-serif;
								background-color: #f4f4f4;
								color: #333;
								padding: 0;
								margin: 0;
							}
							.container {
								max-width: 600px;
								margin: 40px auto;
								background-color: #ffffff;
								padding: 20px 30px;
								border-radius: 8px;
								box-shadow: 0 0 10px rgba(0,0,0,0.1);
							}
							.header {
								font-size: 20px;
								font-weight: bold;
								color: #1a73e8;
								margin-bottom: 20px;
							}
							.content {
								font-size: 16px;
								line-height: 1.5;
							}
							.footer {
								margin-top: 30px;
								font-size: 14px;
								color: #666;
								text-align: center;
							}
							.highlight {
								font-weight: bold;
								color: #1a73e8;
							}
						</style>
					</head>
					<body>
						<div class="container">
							<div class="header">Attendance Edit Approved</div>
							<div class="content">
								Hello,<br/><br/>
								Your attendance edit request for <span class="highlight">%s</span> has been <span class="highlight">approved</span>.<br/><br/>
								<strong>Details:</strong><br/>
								- Clock In: %v<br/>
								- Clock Out: %v<br/>
								- Home Office: %v<br/><br/>
								You can review your attendance record in the HR system.
							</div>
							<div class="footer">
								This is an automated email from the HR system. Please do not reply.
							</div>
						</div>
					</body>
					</html>
        		`, req.WorkDate.Format("2006-01-02"), att.ClockInAt, att.ClockOutAt, att.Is_homeOffice)

				_ = s.notify.SendEmail(*email, subject, body)
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

	// Kirim email profesional
	if s.notify != nil {
		if email, _ := s.empRepo.GetUserEmailByEmployeeID(req.EmployeeID); email != nil && *email != "" {
			subject := "Attendance Edit Rejected"
			body := fmt.Sprintf(`
			<html>
			<head>
				<style>
					body {
						font-family: Arial, sans-serif;
						background-color: #f4f4f4;
						color: #333;
						margin: 0;
						padding: 0;
					}
					.container {
						max-width: 600px;
						margin: 40px auto;
						background-color: #ffffff;
						padding: 20px 30px;
						border-radius: 8px;
						box-shadow: 0 0 10px rgba(0,0,0,0.1);
					}
					.header {
						font-size: 20px;
						font-weight: bold;
						color: #e53935;
						margin-bottom: 20px;
					}
					.content {
						font-size: 16px;
						line-height: 1.5;
					}
					.footer {
						margin-top: 30px;
						font-size: 14px;
						color: #666;
						text-align: center;
					}
					.highlight {
						font-weight: bold;
						color: #e53935;
					}
				</style>
			</head>
			<body>
				<div class="container">
					<div class="header">Attendance Edit Rejected</div>
					<div class="content">
						Hello,<br/><br/>
						Your attendance edit request for <span class="highlight">%s</span> has been <span class="highlight">rejected</span>.<br/><br/>
						<strong>Reviewer Note:</strong><br/>
						<p style="margin:8px 0; padding: 10px; background-color:#f8d7da; border-left: 4px solid #e53935; border-radius: 4px;">%s</p><br/>
						Please contact your manager or HR for further clarification.
					</div>
					<div class="footer">
						This is an automated email from the HR system. Please do not reply.
					</div>
				</div>
			</body>
			</html>
			`, req.WorkDate.Format("2006-01-02"), note)

			_ = s.notify.SendEmail(*email, subject, body)
		}
	}

	return nil
}
