package models

import "time"

type Employee struct {
	ID                uint       `gorm:"primaryKey;autoIncrement" json:"-"`
	UUID              string     `gorm:"type:char(36);uniqueIndex" json:"uuid"`
	CompanyID         *uint      `gorm:"default:null" json:"company_id"`    // Nullable
	BranchID          *uint      `gorm:"default:null" json:"branch_id"`     // Nullable
	RoleID            uint       `gorm:"not null" json:"role_id"`           // Required
	UserID            uint       `gorm:"not null" json:"user_id"`           // Required
	DepartmentID      *uint      `gorm:"default:null" json:"department_id"` // Nullable
	IsFreelance       bool       `gorm:"default:false" json:"is_freelance"`
	TerminationReason *string    `gorm:"type:text" json:"termination_reason"`
	TerminatedAt      *time.Time `json:"terminated_at"`
	CreatedAt         time.Time  `gorm:"autoCreateTime" json:"-"`
	CreatedBy         uint       `gorm:"not null" json:"-"`
	ModifyAt          time.Time  `gorm:"autoUpdateTime" json:"-"`
	ModifyBy          uint       `gorm:"not null" json:"-"`

	Role    *Role    `gorm:"foreignKey:RoleID" json:"role"` // optional eager loading
	User    *User    `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Company *Company `gorm:"foreignKey:CompanyID" json:"company,omitempty"`
	Branch  *Branch  `gorm:"foreignKey:BranchID" json:"branch,omitempty"`
}
