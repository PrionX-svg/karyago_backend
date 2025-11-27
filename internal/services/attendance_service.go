package services

import (
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"

	"hris_backend/internal/models"
	"hris_backend/internal/repositories"
	"hris_backend/pkg"
)

type AttendanceService interface {
	// Core actions
	ListAllEmployeeAttendance(companyUUID *string, from, to time.Time) ([]models.Attendance, error)
	CalendarDay(userID uint, companyUUIDOpt *string, from, to time.Time) (*models.Attendance, error)
	ToggleHomeOffice(userID uint, companyUUID *string, workDate time.Time, isHome bool) (*models.Attendance, error)
	SaveNotes(userID uint, companyUUID *string, workDate time.Time, notes *string) (*models.Attendance, error)
	ClockIn(userID uint, companyUUID *string, workDate, at time.Time) (*models.Attendance, error)
	ClockOut(userID uint, companyUUID *string, workDate, at time.Time) (*models.Attendance, error)

	// For chatbot
	ListOvertime(companyUUID string) ([]models.Attendance, error)
	ListUnclocked(companyUUID string) ([]models.Attendance, error)

	GetDefaultCompanyUUID() (*string, error)

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

func (s *attendanceService) GetDefaultCompanyUUID() (*string, error) {
	var company models.Company
	if err := s.db.First(&company).Error; err != nil {
		return nil, err
	}
	return &company.UUID, nil
}

type CalendarDay struct {
	WorkDate       time.Time  `json:"work_date"`
	Status         string     `json:"status"` // OPEN | PRESENT | ABSENT
	IsHomeOffice   *bool      `json:"is_home_office,omitempty"`
	ClockInAt      *time.Time `json:"clock_in_at,omitempty"`
	ClockOutAt     *time.Time `json:"clock_out_at,omitempty"`
	Notes          *string    `json:"notes,omitempty"`
	IsPastDue      bool       `json:"is_past_due"`
	CanRequestEdit bool       `json:"can_request_edit"`
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

func (s *attendanceService) CalendarDay(
	userID uint,
	companyUUIDOpt *string,
	from, to time.Time,
) (*models.Attendance, error) {
	empID, _, err := s.resolveActor(userID, companyUUIDOpt)
	if err != nil {
		return nil, err
	}
	return s.attRepo.FindByEmployeeAndDate(empID, from)
}

func (s *attendanceService) ListAllEmployeeAttendance(companyUUID *string, from, to time.Time) ([]models.Attendance, error) {
	if companyUUID == nil {
		return nil, errors.New("company_uuid required")
	}

	loc, _ := time.LoadLocation("Asia/Jakarta")
	fromLocal := from.In(loc)
	toLocal := to.In(loc)

	var list []models.Attendance

	err := s.db.
		Table("employees").
		Select(`
			attendances.*,
			employees.id as employee_id,
			employees.company_id,
			employees.user_id
		`).
		Joins(`
			LEFT JOIN attendances 
			ON attendances.employee_id = employees.id 
			AND attendances.work_date BETWEEN ? AND ?
		`, fromLocal, toLocal).
		Where("employees.company_id = (SELECT id FROM companies WHERE uuid = ?)", *companyUUID).
		Preload("Employee.User").
		Order("attendances.work_date ASC").
		Find(&list).Error

	if err != nil {
		return nil, err
	}

	return list, nil
}

func (s *attendanceService) ListCalendar(
	userID uint,
	companyUUIDOpt *string,
	from, to time.Time,
) ([]CalendarDay, error) {

	empID, userID, err := s.resolveActor(userID, companyUUIDOpt)
	if err != nil {
		return nil, err
	}

	// 1) tentukan timezone company (opsional; ganti sesuai field di company)
	loc := pkg.AppLocation()
	if u, err := s.userRepo.GetByID(userID); err == nil && u.Timezone != "" {
		if l, e := time.LoadLocation(u.Timezone); e == nil {
			loc = l
		}
	}

	// normalize from/to ke awal hari di TZ company
	start := time.Date(from.In(loc).Year(), from.In(loc).Month(), from.In(loc).Day(), 0, 0, 0, 0, loc)
	end := time.Date(to.In(loc).Year(), to.In(loc).Month(), to.In(loc).Day(), 0, 0, 0, 0, loc)

	// 2) ambil attendance yang ada di DB
	rows, err := s.attRepo.ListByEmployeeBetween(empID, start, end)
	if err != nil {
		return nil, err
	}

	byDate := map[string]models.Attendance{}
	for _, r := range rows {
		key := r.WorkDate.In(loc).Format("2006-01-02")
		byDate[key] = r
	}

	// 3) build kalender lengkap
	now := time.Now().In(loc)
	hours := pkg.AttendanceCutoffHours()

	out := make([]CalendarDay, 0, int(end.Sub(start).Hours()/24)+1)

	for d := start; !d.After(end); d = d.AddDate(0, 0, 1) {
		key := d.Format("2006-01-02")
		cutoff := d.Add(time.Duration(hours) * time.Hour)

		if att, ok := byDate[key]; ok {
			status := "OPEN"
			if att.ClockInAt != nil && att.ClockOutAt != nil {
				status = "PRESENT"
			}
			out = append(out, CalendarDay{
				WorkDate:       d,
				Status:         status,
				IsHomeOffice:   &att.Is_homeOffice,
				ClockInAt:      att.ClockInAt,
				ClockOutAt:     att.ClockOutAt,
				Notes:          att.Notes,
				IsPastDue:      now.After(cutoff),
				CanRequestEdit: status != "PRESENT", // boleh edit jika belum lengkap
			})
			continue
		}

		// Tidak ada row → ABSENT
		isPast := now.After(cutoff)
		out = append(out, CalendarDay{
			WorkDate:       d,
			Status:         "ABSENT",
			IsPastDue:      isPast,
			CanRequestEdit: isPast, // kebijakan: boleh request setelah lewat cutoff
		})
	}

	return out, nil
}

// For chatbot
func (s *attendanceService) ListOvertime(companyUUID string) ([]models.Attendance, error) {
	var list []models.Attendance
	err := s.db.
		Joins("JOIN employees ON employees.id = attendances.employee_id").
		Where("attendances.company_id = (SELECT id FROM companies WHERE uuid = ?)", companyUUID).
		Where("attendances.is_overtime = TRUE").
		Preload("Employee.User").
		Order("work_date DESC").
		Find(&list).Error
	return list, err
}

func (s *attendanceService) ListUnclocked(companyUUID string) ([]models.Attendance, error) {
	var list []models.Attendance
	err := s.db.
		Joins("JOIN employees ON employees.id = attendances.employee_id").
		Where("attendances.company_id = (SELECT id FROM companies WHERE uuid = ?)", companyUUID).
		Where("attendances.clock_out_at IS NULL").
		Preload("Employee.User").
		Find(&list).Error
	return list, err
}

func (s *attendanceService) getAdminEmails(companyID uint) ([]string, error) {
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

	return s.attRepo.SetHomeOffice(empID, compID, userID, workDate, isHome)
}

// Notes autosave (bisa diubah sampai clock-out)
func (s *attendanceService) SaveNotes(userID uint, companyUUID *string, workDate time.Time, notes *string) (*models.Attendance, error) {
	empID, compID, err := s.resolveActor(userID, companyUUID)
	if err != nil {
		return nil, err
	}
	return s.attRepo.UpdateNotes(empID, compID, userID, workDate, notes, nil)
}

// Clock-in (idempotent per employee+work_date)
func (s *attendanceService) ClockIn(userID uint, companyUUID *string, workDate time.Time, at time.Time) (*models.Attendance, error) {
	empID, compID, err := s.resolveActor(userID, companyUUID)
	if err != nil {
		return nil, err
	}
	return s.attRepo.ClockIn(empID, compID, userID, workDate, at)
}

// Clock-out (mengunci toggle & notes)
func (s *attendanceService) ClockOut(userID uint, companyUUID *string, workDate time.Time, at time.Time) (*models.Attendance, error) {
	empID, _, err := s.resolveActor(userID, companyUUID)
	if err != nil {
		return nil, err
	}
	// Cross-day cutoff rule (misal: 04:00 pagi WIB)
	const CrossDayCutoffHour = 4

	// Ambil data attendance dari repo (include ClockInAt)
	att, err := s.attRepo.GetByEmployeeAndDate(empID, workDate)
	if err != nil {
		return nil, err
	}

	// 🕒 Kalau clock out lewat tengah malam tapi sebelum cutoff (04:00),
	// tetap dianggap hari sebelumnya
	if att.ClockInAt != nil && at.Day() != att.ClockInAt.Day() && at.Hour() < CrossDayCutoffHour {
		workDate = att.ClockInAt.Truncate(24 * time.Hour)
	}

	// Jalankan clock-out di repository
	att, err = s.attRepo.ClockOut(empID, userID, workDate, at)
	if err != nil {
		return nil, err
	}

	loc, _ := time.LoadLocation("Asia/Jakarta")

	//Kirim email notifikasi kalau lembur
	if att.IsOvertime {
		go func(a *models.Attendance) {
			// pastikan user sudah ke-load
			s.db.Preload("Employee.User").First(&a, a.ID)
			adminEmails, err := s.getAdminEmails(a.CompanyID)
			if err != nil || len(adminEmails) == 0 {
				fmt.Println("⚠️ No admin email found:", err)
				return
			}
			subject := fmt.Sprintf("Overtime Alert: %s", a.WorkDate.Format("02 Jan 2006"))
			body := fmt.Sprintf(`
				<h2 style="color:#ff6600;">KARYAGO Overtime Notification</h2>
				<p><b>%s %s</b> melakukan lembur pada <b>%s</b>.</p>
				<ul>
					<li>Clock In: %s</li>
					<li>Clock Out: %s</li>
					<li>Total Jam: %.2f</li>
					<li>Lembur: %.2f jam</li>
				</ul>
				<p>Silakan dicek lembur ini di attendance list.</p>
			`,
				a.Employee.User.FirstName,
				a.Employee.User.LastName,
				a.WorkDate.Format("02 Jan 2006"),
				a.ClockInAt.In(loc).Format("15:04"),
				a.ClockOutAt.In(loc).Format("15:04"),
				*a.TotalWorkHours,
				*a.OvertimeHours,
			)

			for _, email := range adminEmails {
				if err := pkg.SendEmail(email, subject, body); err != nil {
					fmt.Println("❌ Failed to send overtime email to", email, ":", err)
				}
			}
		}(att)
	}

	return att, nil
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
