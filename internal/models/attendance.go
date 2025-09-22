package models

import "time"

type AttendanceStatus string

const(
	AttendanceStatusOpen AttendanceStatus = "OPEN" //Done clock-in, waiting for clock-out
	AttendanceStatusClosed AttendanceStatus = "COMPLETE" //Done clock-in and clock-out
	AttendanceStatusPendingWFH AttendanceStatus = "PENDING_WFH" //Request work from home, waiting for approval
	AttendanceStatusAbsent AttendanceStatus = "ABSENT" //Did not do clock-in and clock-out
	
	// AttendanceStatusOnLeave AttendanceStatus = "ON_LEAVE" //On leave for the day
)

type Attendance struct {
	ID         uint64           `json:"id" gorm:"primaryKey;autoIncrement"`
	UUID       string           `json:"uuid" gorm:"type:char(36);not null;uniqueIndex"`
	EmployeeID uint64           `json:"employee_id" gorm:"not null;index:idx_emp_workdate,unique"`
	CompanyID  uint64           `json:"company_id" gorm:"not null;index:idx_company_workdate"`
	WorkDate   time.Time        `json:"work_date" gorm:"type:date;not null;index:idx_emp_workdate,unique;index:idx_company_workdate"`

	ClockInAt        *time.Time `json:"clock_in_at" gorm:"type:datetime"`
	ClockInLat       *float64   `json:"clock_in_lat" gorm:"type:decimal(10,7)"`
	ClockInLng       *float64   `json:"clock_in_lng" gorm:"type:decimal(10,7)"`
	ClockInAddr      *string    `json:"clock_in_addr" gorm:"type:text"`
	ClockInAccuracyM *int       `json:"clock_in_accuracy_m"`
	IsInOfficeIn     bool       `json:"is_in_office_in" gorm:"not null;default:false"`
	CompanyLocationIDIn *uint64 `json:"company_location_id_in" gorm:"index"`
	DistanceInM         *float64`json:"distance_in_m"`

	ClockOutAt        *time.Time `json:"clock_out_at" gorm:"type:datetime"`
	ClockOutLat       *float64   `json:"clock_out_lat" gorm:"type:decimal(10,7)"`
	ClockOutLng       *float64   `json:"clock_out_lng" gorm:"type:decimal(10,7)"`
	ClockOutAddr      *string    `json:"clock_out_addr" gorm:"type:text"`
	ClockOutAccuracyM *int       `json:"clock_out_accuracy_m"`
	IsInOfficeOut     bool       `json:"is_in_office_out" gorm:"not null;default:false"`
	CompanyLocationIDOut *uint64 `json:"company_location_id_out" gorm:"index"`
	DistanceOutM         *float64`json:"distance_out_m"`

	IsWFH            bool              `json:"is_wfh" gorm:"not null;default:false"`
	WFHFormRequired  bool              `json:"wfh_form_required" gorm:"not null;default:false"`
	WFHFormSubmitted bool              `json:"wfh_form_submitted" gorm:"not null;default:false"`
	Status           AttendanceStatus  `json:"status" gorm:"type:enum('OPEN','PENDING_WFH_FORM','COMPLETED');not null;default:'OPEN'"`

	Remarks   *string   `json:"remarks" gorm:"type:text"`
	CreatedBy *uint64   `json:"created_by"` // id actor (admin/employee) — bebas tetap pakai user admin id kalau ada
	UpdatedBy *uint64   `json:"updated_by"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}