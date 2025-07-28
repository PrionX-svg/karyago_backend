package models

import (
	"time"
)

type EventWorkArea struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	UUID     string `gorm:"type:char(36);uniqueIndex" json:"uuid"`
	EventID  uint   `gorm:"not null" json:"event_id"`
	Name     string `gorm:"type:varchar(255);not null" json:"name"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	CreatedBy uint      `json:"created_by"`
	ModifyAt  time.Time `gorm:"autoUpdateTime" json:"modify_at"`
	ModifyBy  uint      `json:"modify_by"`

	Event *Event `gorm:"foreignKey:EventID" json:"event,omitempty"`
}
