package models

import (
	"time"
)

type EventDepartment struct {
	ID                    uint    `gorm:"primaryKey" json:"id"`
	UUID                  string  `gorm:"type:char(36);uniqueIndex" json:"uuid"`
	Name                  string  `gorm:"type:varchar(255);not null" json:"name"`
	Description           *string `gorm:"type:text;default:null" json:"description,omitempty"`
	EventDepartmentGroupID uint    `gorm:"not null" json:"event_department_group_id"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	CreatedBy uint      `json:"created_by"`
	ModifyAt  time.Time `gorm:"autoUpdateTime" json:"modify_at"`
	ModifyBy  uint      `json:"modify_by"`

	EventDepartmentGroup *EventDepartmentGroup `gorm:"foreignKey:EventDepartmentGroupID" json:"event_department_group,omitempty"`
}

