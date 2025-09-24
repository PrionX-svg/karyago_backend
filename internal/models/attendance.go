package models

import "time"

type Attendance struct {
	ID         uint64    `json:"id" gorm:"primaryKey;autoIncrement"`
	UUID       string    `json:"uuid" gorm:"type:char(36);not null;uniqueIndex"`
	EmployeeID uint64    `json:"employee_id" gorm:"not null;index:idx_emp_workdate,unique"`
	CompanyID  uint64    `json:"company_id" gorm:"not null;index:idx_company_workdate"`
	WorkDate   time.Time `json:"work_date" gorm:"type:date;not null;index:idx_emp_workdate,unique;index:idx_company_workdate"`

	ClockInAt  *time.Time `json:"clock_in_at" gorm:"type:datetime"`
	ClockInLat *float64   `json:"clock_in_lat" gorm:"type:decimal(10,7)"`
	ClockInLng *float64   `json:"clock_in_lng" gorm:"type:decimal(10,7)"`

	ClockOutAt  *time.Time `json:"clock_out_at" gorm:"type:datetime"`
	ClockOutLat *float64   `json:"clock_out_lat" gorm:"type:decimal(10,7)"`
	ClockOutLng *float64   `json:"clock_out_lng" gorm:"type:decimal(10,7)"`

	Is_homeOfiice    bool             `json:"Is_homeOfiice" gorm:"not null;default:false"`
	
	Notes   *string   `json:"Notes" gorm:"type:varchar(400)"`
	CreatedBy *uint64   `json:"created_by"`
	UpdatedBy *uint64   `json:"updated_by"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
