package models

import "time"

type Attendance struct {
	ID         uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	UUID       string    `gorm:"type:char(36);uniqueIndex" json:"uuid"`
	EmployeeID uint      `json:"employee_id" gorm:"not null;index:idx_emp_workdate,unique"`
	UserID     uint      `json:"user_id" gorm:"not null"` // denormalized for easier join with users table
	CompanyID  uint      `json:"company_id" gorm:"not null;index:idx_company_workdate"`
	WorkDate   time.Time `json:"work_date" gorm:"type:date;not null;index:idx_emp_workdate,unique;index:idx_company_workdate"`

	ClockInAt  *time.Time `json:"clock_in_at" gorm:"type:datetime"`
	ClockInLat *float64   `json:"clock_in_lat" gorm:"type:decimal(10,7)"`
	ClockInLng *float64   `json:"clock_in_lng" gorm:"type:decimal(10,7)"`

	ClockOutAt  *time.Time `json:"clock_out_at" gorm:"type:datetime"`
	ClockOutLat *float64   `json:"clock_out_lat" gorm:"type:decimal(10,7)"`
	ClockOutLng *float64   `json:"clock_out_lng" gorm:"type:decimal(10,7)"`

	IsOvertime     bool     `json:"is_overtime" gorm:"not null;default:false"`
	OvertimeHours  *float64 `json:"overtime_hours" gorm:"column:overtime_hours;type:decimal(5,2)"`
	OvertimeReason *string  `json:"overtime_reason" gorm:"column:overtime_reason;type:varchar(255)"`

	TotalWorkHours *float64 `json:"total_work_hours" gorm:"type:decimal(5,2)"`

	Is_homeOffice bool `json:"is_home_office" gorm:"not null;default:false"`

	Notes     *string   `json:"notes" gorm:"type:varchar(400)"`
	CreatedBy *uint     `json:"created_by"`
	UpdatedBy *uint     `json:"updated_by"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	Employee Employee `json:"employee" gorm:"foreignKey:EmployeeID;references:ID"`
	User     User     `json:"user" gorm:"foreignKey:UserID;references:ID"`
}
