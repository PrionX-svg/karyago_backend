package models

import "time"

type EditRequestType string

const (
	EditReqClockIn      EditRequestType = "CLOCK_IN"
	EditReqClockOut     EditRequestType = "CLOCK_OUT"
	EditReqBoth         EditRequestType = "BOTH"
	EditReqHomeFlag     EditRequestType = "HOME_FLAG"
	EditReqBothPlusFlag EditRequestType = "BOTH_PLUS_FLAG"
)

type EditRequestStatus string

const (
	EditStatusPending  EditRequestStatus = "PENDING"
	EditStatusApproved EditRequestStatus = "APPROVED"
	EditStatusRejected EditRequestStatus = "REJECTED"
	EditStatusCanceled EditRequestStatus = "CANCELED"
)

type AttendanceEditRequest struct {
	ID         uint               `json:"id" gorm:"primaryKey;autoIncrement"`
	EmployeeID uint               `json:"employee_id" gorm:"not null;index:idx_employee"`
	CompanyID  uint               `json:"company_id" gorm:"not null;index:idx_company_status"`
	AttendanceID *uint            `json:"attendance_id" gorm:"index"` // Boleh null jika record attendance harinya belum ada (lupa absen sama sekali)
	WorkDate     time.Time        `json:"work_date" gorm:"type:date;not null;index:idx_work_date"`
	RequestType  EditRequestType  `json:"request_type" gorm:"type:enum('CLOCK_IN','CLOCK_OUT','BOTH','HOME_FLAG','BOTH_PLUS_FLAG');not null;index:idx_company_status"`
	ProposedClockInAt   *time.Time `json:"proposed_clock_in_at" gorm:"type:datetime"`
	ProposedClockOutAt  *time.Time `json:"proposed_clock_out_at" gorm:"type:datetime"`
	ProposedIsHomeOffice *bool     `json:"proposed_is_home_office" gorm:"column:proposed_is_home_office"`
	Reason string `json:"reason" gorm:"type:text;not null"`
	Status      EditRequestStatus `json:"status" gorm:"type:enum('PENDING','APPROVED','REJECTED','CANCELED');not null;default:'PENDING';index:idx_company_status"`
	ReviewedBy  *uint             `json:"reviewed_by"`
	ReviewedAt  *time.Time        `json:"reviewed_at" gorm:"type:datetime"`
	ReviewNote  *string           `json:"review_note" gorm:"type:text"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
