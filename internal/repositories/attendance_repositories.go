package repositories

import (
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"hris_backend/internal/models"
)

var (
	ErrAlreadyClockedIn  = errors.New("already clocked in for this date")
	ErrAlreadyClockedOut = errors.New("already clocked out for this date")
)

type AttendanceRepository interface {
	// Pastikan ada 1 record untuk (employee_id, work_date). Dipanggil otomatis oleh ClockIn/SetHomeOffice/UpdateNotes.
	EnsureForDay(employeeID, companyID uint, workDate time.Time) (*models.Attendance, error)

	// Set jam masuk. Idempotent: kalau sudah ada ClockInAt, kembalikan ErrAlreadyClockedIn.
	ClockIn(employeeID, companyID uint, workDate time.Time, at time.Time) (*models.Attendance, error)

	// Set jam keluar. Kalau sudah clock-out, balikan ErrAlreadyClockedOut.
	ClockOut(employeeID uint, workDate time.Time, at time.Time) (*models.Attendance, error)

	// Toggle Home Office / In Office. Boleh diubah selama BELUM clock-out.
	SetHomeOffice(employeeID, companyID uint, workDate time.Time, isHomeOffice bool) (*models.Attendance, error)

	// Update notes bebas (tanpa approval). Boleh diubah selama BELUM clock-out.
	UpdateNotes(employeeID, companyID uint, workDate time.Time, notes *string, updatedBy *uint) (*models.Attendance, error)

	// Query util
	FindByEmployeeAndDate(employeeID uint, workDate time.Time) (*models.Attendance, error)
	ListByEmployeeBetween(employeeID uint, from, to time.Time) ([]models.Attendance, error)
}

type attendanceRepository struct{ db *gorm.DB }

func NewAttendanceRepository(db *gorm.DB) AttendanceRepository {
	return &attendanceRepository{db: db}
}

// EnsureForDay membuat (atau mengambil) record unik per (employee_id, work_date).
func (r *attendanceRepository) EnsureForDay(employeeID, companyID uint, workDate time.Time) (*models.Attendance, error) {
	var att models.Attendance

	if err := r.db.Where("employee_id = ? AND work_date = ?", employeeID, workDate).
		First(&att).Error; err == nil {
		return &att, nil
	}

	att = models.Attendance{
		EmployeeID: employeeID,
		CompanyID:  companyID,
		WorkDate:   workDate,
		UUID: uuid.NewString(),
	}
	if err := r.db.Create(&att).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) || strings.Contains(err.Error(), "Error 400") {
			if err2 := r.db.Where("employee_id = ? AND work_date = ?", employeeID, workDate).
				First(&att).Error; err2 == nil {
				return &att, nil
			}
		}
		return nil, err
	}
	return &att, nil
}

func (r *attendanceRepository) ClockIn(employeeID, companyID uint, workDate time.Time, at time.Time) (*models.Attendance, error) {
	att, err := r.EnsureForDay(employeeID, companyID, workDate)
	if err != nil {
		return nil, err
	}
	if att.ClockInAt != nil {
		return nil, ErrAlreadyClockedIn
	}
	att.ClockInAt = &at
	if err := r.db.Model(att).Update("clock_in_at", att.ClockInAt).Error; err != nil {
		return nil, err
	}
	return att, nil
}

func (r *attendanceRepository) ClockOut(employeeID uint, workDate time.Time, at time.Time) (*models.Attendance, error) {
	var att models.Attendance
	if err := r.db.
		Where("employee_id = ? AND work_date = ?", employeeID, workDate).
		First(&att).Error; err != nil {
		return nil, err
	}
	if att.ClockOutAt != nil {
		return nil, ErrAlreadyClockedOut
	}
	att.ClockOutAt = &at

	// Setelah clock-out, toggle & notes “terkunci” (aturan ada di SetHomeOffice/UpdateNotes).
	if err := r.db.Model(&att).Update("clock_out_at", att.ClockOutAt).Error; err != nil {
		return nil, err
	}
	return &att, nil
}

func (r *attendanceRepository) SetHomeOffice(employeeID, companyID uint, workDate time.Time, isHomeOffice bool) (*models.Attendance, error) {
	att, err := r.EnsureForDay(employeeID, companyID, workDate)
	if err != nil {
		return nil, err
	}
	if att.ClockOutAt != nil {
		return nil, ErrAlreadyClockedOut
	}
	att.Is_homeOffice = isHomeOffice
	if err := r.db.Model(att).Update("is_home_office", att.Is_homeOffice).Error; err != nil {
		return nil, err
	}
	return att, nil
}

func (r *attendanceRepository) UpdateNotes(employeeID, companyID uint, workDate time.Time, notes *string, updatedBy *uint) (*models.Attendance, error) {
	att, err := r.EnsureForDay(employeeID, companyID, workDate)
	if err != nil {
		return nil, err
	}
	if att.ClockOutAt != nil {
		return nil, ErrAlreadyClockedOut
	}
	att.Notes = notes

	updates := map[string]any{"notes": att.Notes}
	if updatedBy != nil {
		att.UpdatedBy = updatedBy
		updates["updated_by"] = *updatedBy
	}
	if err := r.db.Model(att).Updates(updates).Error; err != nil {
		return nil, err
	}
	return att, nil
}

func (r *attendanceRepository) FindByEmployeeAndDate(employeeID uint, workDate time.Time) (*models.Attendance, error) {
	var att models.Attendance
	err := r.db.Where("employee_id = ? AND work_date = ?", employeeID, workDate).First(&att).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &att, err
}

func (r *attendanceRepository) ListByEmployeeBetween(employeeID uint, from, to time.Time) ([]models.Attendance, error) {
	var list []models.Attendance
	err := r.db.
		Where("employee_id = ? AND work_date BETWEEN ? AND ?", employeeID, from, to).
		Order("work_date ASC").
		Find(&list).Error
	return list, err
}
