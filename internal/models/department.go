package models

import (
	"time"
)

type Department struct {
	ID                uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	UUID              string    `gorm:"type:char(36);uniqueIndex" json:"uuid"`
	DepartmentGroupID uint      `gorm:"not null" json:"department_group_id"` // foreign key ke table department group
	Name              string    `gorm:"type:varchar(100);not null" json:"name"`
	Description       string    `gorm:"type:text" json:"description,omitempty"`
	CreatedAt         time.Time `gorm:"autoCreateTime" json:"created_at"`
	CreatedBy         uint      `gorm:"type:varchar(100)" json:"created_by"`
	ModifyAt          time.Time `gorm:"autoUpdateTime" json:"modify_at"`
	ModifyBy          uint      `gorm:"type:varchar(100)" json:"modify_by"`
}
