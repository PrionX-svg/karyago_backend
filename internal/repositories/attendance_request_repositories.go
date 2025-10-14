package repositories

import (
	"errors"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"hris_backend/internal/models"
)

type AttendanceEditRepository interface {
	Create(req *models.AttendanceEditRequest) error
	FindByID(id uint) (*models.AttendanceEditRequest, error)

	// Employee view
	ListMine(employeeID uint, status *models.EditRequestStatus, from, to *time.Time) ([]models.AttendanceEditRequest, error)

	// Supervisor view (with pagination)
	ListAllEmployee(companyID uint, status *models.EditRequestStatus, from, to *time.Time, q *string, limit, offset int) ([]models.AttendanceEditRequest, int64, error)

	// Approval updates (Approve MUST be called within a transaction)
	Approve(tx *gorm.DB, reqID uint, reviewerID uint, note *string) (*models.AttendanceEditRequest, error)
	Reject(reqID uint, reviewerID uint, note string) error
}

type attendanceEditRepository struct{ db *gorm.DB }

func NewAttendanceEditRepository(db *gorm.DB) AttendanceEditRepository {
	return &attendanceEditRepository{db: db}
}

func (r *attendanceEditRepository) Create(req *models.AttendanceEditRequest) error {
	if req.Status == "" {
		req.Status = models.EditStatusPending
	}
	return r.db.Create(req).Error
}

func (r *attendanceEditRepository) FindByID(id uint) (*models.AttendanceEditRequest, error) {
	var m models.AttendanceEditRequest
	err := r.db.First(&m, "id = ?", id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &m, err
}

func (r *attendanceEditRepository) ListMine(employeeID uint, status *models.EditRequestStatus, from, to *time.Time) ([]models.AttendanceEditRequest, error) {
	q := r.db.Model(&models.AttendanceEditRequest{}).Where("employee_id = ?", employeeID)

	if status != nil && *status != "" {
		q = q.Where("status = ?", *status)
	}
	if from != nil {
		q = q.Where("work_date >= ?", from.Format("2006-01-02"))
	}
	if to != nil {
		q = q.Where("work_date <= ?", to.Format("2006-01-02"))
	}

	var rows []models.AttendanceEditRequest
	err := q.Order("created_at DESC").Find(&rows).Error
	return rows, err
}

func (r *attendanceEditRepository) ListAllEmployee(companyID uint, status *models.EditRequestStatus, from, to *time.Time, qtext *string, limit, offset int) ([]models.AttendanceEditRequest, int64, error) {
	q := r.db.Model(&models.AttendanceEditRequest{}).
		Preload("Employee.User").       
		Preload("Employee.Department").
		Where("company_id = ?", companyID)

	if status != nil && *status != "" {
		q = q.Where("status = ?", *status)
	}
	if from != nil {
		q = q.Where("work_date >= ?", from.Format("2006-01-02"))
	}
	if to != nil {
		q = q.Where("work_date <= ?", to.Format("2006-01-02"))
	}
	if qtext != nil && *qtext != "" {
		q = q.Where("reason LIKE ?", "%"+*qtext+"%")
	}

	var total int64
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if limit <= 0 || limit > 200 {
		limit = 50
	}
	if offset < 0 {
		offset = 0
	}

	var rows []models.AttendanceEditRequest
	err := q.Order("created_at DESC").Limit(limit).Offset(offset).Find(&rows).Error
	return rows, total, err
}

func (r *attendanceEditRepository) Approve(tx *gorm.DB, reqID uint, reviewerID uint, note *string) (*models.AttendanceEditRequest, error) {
	now := time.Now().UTC()
	upd := map[string]any{
		"status":      models.EditStatusApproved,
		"reviewed_by": reviewerID,
		"reviewed_at": now,
	}
	if note != nil {
		upd["review_note"] = *note
	}

	if err := tx.Model(&models.AttendanceEditRequest{}).
		Where("id = ? AND status = ?", reqID, models.EditStatusPending).
		Updates(upd).Error; err != nil {
		return nil, err
	}

	var out models.AttendanceEditRequest
	if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
		First(&out, "id = ?", reqID).Error; err != nil {
		return nil, err
	}
	return &out, nil
}

func (r *attendanceEditRepository) Reject(reqID uint, reviewerID uint, note string) error {
	now := time.Now().UTC()
	return r.db.Model(&models.AttendanceEditRequest{}).
		Where("id = ? AND status = ?", reqID, models.EditStatusPending).
		Updates(map[string]any{
			"status":      models.EditStatusRejected,
			"reviewed_by": reviewerID,
			"reviewed_at": now,
			"review_note": note,
		}).Error
}
