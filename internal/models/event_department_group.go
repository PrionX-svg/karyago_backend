package models

import (
	"time"
)

type EventDepartmentGroup struct {
	ID            uint    `gorm:"primaryKey" json:"id"`
	UUID          string  `gorm:"type:char(36);uniqueIndex" json:"uuid"`
	ResponsibleID *uint   `gorm:"default:null" json:"responsible_id"`
	EventID       uint    `gorm:"not null" json:"event_id"`
	Name          string  `gorm:"type:varchar(255);not null" json:"name"`
	Description   *string `gorm:"type:text;default:null" json:"description,omitempty"`

	CreatedAt time.Time `gorm:"autoUpdateTime" json:"created_at"`
	CreatedBy uint      `json:"created_by"`
	ModifyAt  time.Time `gorm:"autoUpdateTime" json:"modify_at"`
	ModifyBy  uint      `json:"modify_by"`

	Responsible *User `gorm:"foreignKey:ResponsibleID" json:"responsible,omitempty"`
}
