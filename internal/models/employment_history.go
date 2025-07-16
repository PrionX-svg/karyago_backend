package models

import "time"

type EmploymentHistory struct {
	ID         uint       `gorm:"primaryKey;autoIncrement" json:"-"`
	UUID       string     `gorm:"type:char(36);uniqueIndex" json:"uuid"`
	EmployeeID uint       `gorm:"not null" json:"employee_id"`
	CompanyID  *uint      `gorm:"default:null" json:"company_id,omitempty"`
	BranchID   *uint      `gorm:"default:null" json:"branch_id,omitempty"`
	RoleID     uint       `gorm:"not null" json:"role_id"`
	Position   string     `gorm:"type:varchar(100);not null" json:"position"`
	IsPresent  bool       `gorm:"default:false" json:"is_present"`
	StartDate  time.Time  `gorm:"not null" json:"start_date"`
	EndDate    *time.Time `json:"end_date,omitempty"`
	Notes      *string    `gorm:"type:text" json:"notes,omitempty"`
	CreatedAt  time.Time  `gorm:"autoCreateTime" json:"-"`
	CreatedBy  uint       `gorm:"not null" json:"-"`
	ModifyAt   time.Time  `gorm:"autoUpdateTime" json:"-"`
	ModifyBy   uint       `gorm:"not null" json:"-"`

	Employee *Employee `gorm:"foreignKey:EmployeeID" json:"employee,omitempty"`
	Company  *Company  `gorm:"foreignKey:CompanyID" json:"company,omitempty"`
	Branch   *Branch   `gorm:"foreignKey:BranchID" json:"branch,omitempty"`
	Role     *Role     `gorm:"foreignKey:RoleID" json:"role,omitempty"`
}
