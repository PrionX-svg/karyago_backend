package models

import (
	"time"
)

type EventItemType string

const (
	EventItemTypeEvent EventItemType = "event"
	EventItemTypeDate  EventItemType = "date"
	EventItemTypeShift EventItemType = "shift"
)

type EventItem struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	UUID      string         `gorm:"type:char(36);uniqueIndex" json:"uuid"`
	EventID   uint           `gorm:"not null" json:"event_id"`
	Name      string         `gorm:"type:varchar(255);not null" json:"name"`
	Type      EventItemType  `gorm:"type:enum('event','date','shift');not null" json:"type"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	CreatedBy uint      `json:"created_by"`
	ModifyAt  time.Time `gorm:"autoUpdateTime" json:"modify_at"`
	ModifyBy  uint      `json:"modify_by"`

	Event *Event `gorm:"foreignKey:EventID" json:"event,omitempty"`
}

