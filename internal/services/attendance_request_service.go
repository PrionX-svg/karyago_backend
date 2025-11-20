package services

import (
	"fmt"
	"strings"
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
	var overtimeAtt *models.Attendance

	txErr := s.db.Transaction(func(tx *gorm.DB) error {
		req, err := s.editRepo.FindByID(reqID)
		if err != nil || req == nil || req.Status != models.EditStatusPending {
			return fmt.Errorf("invalid request")
		}

		reviewer, err := s.getEmployee(userID)
		if err != nil || reviewer.CompanyID == nil || *reviewer.CompanyID != req.CompanyID {
			return fmt.Errorf("forbidden")
		}

		// Pastikan attendance untuk hari tersebut sudah ada / dibuat
		att, err := s.attRepo.EnsureForDay(req.EmployeeID, req.CompanyID, req.ID, req.WorkDate)
		if err != nil {
			return err
		}

		loc, _ := time.LoadLocation("Asia/Jakarta")

		// Terapkan perubahan berdasarkan tipe request
		updateFields := map[string]any{}
		if req.ProposedClockInAt != nil {
			updateFields["clock_in_at"] = req.ProposedClockInAt
			att.ClockInAt = req.ProposedClockInAt
		}
		if req.ProposedClockOutAt != nil {
			updateFields["clock_out_at"] = req.ProposedClockOutAt
			att.ClockOutAt = req.ProposedClockOutAt
		}
		if req.ProposedIsHomeOffice != nil {
			updateFields["is_home_office"] = *req.ProposedIsHomeOffice
			att.Is_homeOffice = *req.ProposedIsHomeOffice
		}

		if len(updateFields) > 0 {
			if err := tx.Model(att).Updates(updateFields).Error; err != nil {
				return err
			}
		}

		// Recalculate overtime
		if err := s.recalculateOvertime(tx, att); err != nil {
			return err
		}

		if att.IsOvertime {
			overtimeAtt = att
		}

		// Update status request menjadi APPROVED
		if _, err := s.editRepo.Approve(tx, req.ID, userID, note); err != nil {
			return err
		}

		// Kirim notifikasi ke karyawan
		if s.notify != nil {
			if email, _ := s.empRepo.GetUserEmailByEmployeeID(req.EmployeeID); email != nil && *email != "" {
				clockIn := "-"
				if att.ClockInAt != nil {
					clockIn = att.ClockInAt.In(loc).Format("15:04")
				}

				clockOut := "-"
				if att.ClockOutAt != nil {
					clockOut = att.ClockOutAt.In(loc).Format("15:04")
				}
				subject := "Attendance Edit Approved"
				body := fmt.Sprintf(`
					<html>
					<body style="font-family:Arial,sans-serif">
						<h3 style="color:#1a73e8;">Attendance Edit Approved</h3>
						<p>Your attendance edit request for <b>%s</b> has been approved.</p>
						<ul>
							<li>Clock In: %v</li>
							<li>Clock Out: %v</li>
							<li>Home Office: %v</li>
						</ul>
						<p>Please review your updated record in KARYAGO.</p>
					</body>
					</html>`,
					req.WorkDate.Format("2006-01-02"),
					clockIn, 
					clockOut, 
					att.Is_homeOffice,
				)
				_ = s.notify.SendEmail(*email, subject, body)
			}
		}

		return nil
	})

	if txErr == nil && overtimeAtt != nil {
		go func() {
			time.Sleep(2 * time.Second)
			s.notifyAdminsAboutOvertime(overtimeAtt)
		}()
	}

	return txErr
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
						Please contact your Supervisor or HR for further clarification.
					</div>
					<div class="footer">
						This is an automated email from the KARYAGO system. Please do not reply.
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

func (s *attendanceEditService) recalculateOvertime(tx *gorm.DB, att *models.Attendance) error {
	var (
		totalHours   *float64
		overtimePart *float64
		isOvertime   bool
	)

	if att.ClockInAt != nil && att.ClockOutAt != nil {
		duration := att.ClockOutAt.Sub(*att.ClockInAt).Hours()
		totalHours = &duration
		if duration > 8 {
			isOvertime = true
			ot := duration - 8
			overtimePart = &ot
		}
	}

	return tx.Model(att).Updates(map[string]any{
		"total_work_hours": totalHours,
		"is_overtime":      isOvertime,
		"overtime_hours":   overtimePart,
	}).Error
}

func (s *attendanceEditService) notifyAdminsAboutOvertime(att *models.Attendance) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("⚠️ recovered from panic in notifyAdminsAboutOvertime:", r)
		}
	}()

	var detailed models.Attendance
	if err := s.db.Preload("Employee.User").First(&detailed, att.ID).Error; err != nil {
		fmt.Println("⚠️ failed to load attendance for overtime notification:", err)
		return
	}

	loc, _ := time.LoadLocation("Asia/Jakarta")

	emails, err := s.getAdminEmails(detailed.CompanyID)
	if err != nil || len(emails) == 0 {
		fmt.Println("⚠️ no admin email available:", err)
		return
	}

	// Safe handling untuk waktu & jam kerja
	clockIn := "-"
	if detailed.ClockInAt != nil {
		clockIn = detailed.ClockInAt.In(loc).Format("15:04")
	}
	clockOut := "-"
	if detailed.ClockOutAt != nil {
		clockOut = detailed.ClockOutAt.In(loc).Format("15:04")
	}

	totalHours := 0.0
	if detailed.TotalWorkHours != nil {
		totalHours = *detailed.TotalWorkHours
	}
	overtimeHours := 0.0
	if detailed.OvertimeHours != nil {
		overtimeHours = *detailed.OvertimeHours
	}

	subject := fmt.Sprintf("Overtime Alert: %s", detailed.WorkDate.Format("02 Jan 2006"))
	body := fmt.Sprintf(`
		<h2 style="color:#ff6600;">KARYAGO Overtime Notification</h2>
		<p><b>%s %s</b> melakukan lembur pada <b>%s</b>.</p>
		<ul>
			<li>Clock In: %s</li>
			<li>Clock Out: %s</li>
			<li>Total Jam: %.2f</li>
			<li>Lembur: %.2f jam</li>
		</ul>
	`,
		detailed.Employee.User.FirstName,
		detailed.Employee.User.LastName,
		detailed.WorkDate.Format("02 Jan 2006"),
		clockIn, clockOut, totalHours, overtimeHours,
	)

	for _, email := range emails {
		fmt.Printf("[DEBUG] Trying to send overtime email to: '%s'\n", email)
		if strings.TrimSpace(email) == "" {
			fmt.Println("⚠️ skipped empty email in admin list")
			continue
		}
		if err := s.notify.SendEmail(email, subject, body); err != nil {
			fmt.Println("❌ failed to send overtime email to", email, ":", err)
		} else {
			fmt.Println("✅ overtime email sent to", email)
		}
	}

}

func (s *attendanceEditService) getAdminEmails(companyID uint) ([]string, error) {
	var emails []string

	const roleID = 1

	err := s.db.
		Table("employees AS e").
		Joins("JOIN users u ON u.id = e.user_id").
		Where("e.company_id = ? AND e.role_id = ? AND e.terminated_at IS NULL", companyID, roleID).
		Pluck("u.email", &emails).Error
	fmt.Println("[DEBUG] Admin emails:", emails)

	return emails, err
}
