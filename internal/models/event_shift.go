package models

import (
	"time"
)

type EventShift struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	UUID      string    `gorm:"type:char(36);uniqueIndex" json:"uuid"`
	EventID   uint      `gorm:"not null" json:"event_id"`
	Name      string    `gorm:"type:varchar(100);not null" json:"name"`
	Date      time.Time `gorm:"type:date;not null" json:"date"`
	HourFrom  time.Time `gorm:"type:time;not null" json:"hour_from"`
	HourUntil time.Time `gorm:"type:time;not null" json:"hour_until"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	CreatedBy uint      `gorm:"type:varchar(100)" json:"created_by"`
	ModifyAt  time.Time `gorm:"autoUpdateTime" json:"modify_at"`
	ModifyBy  uint      `gorm:"type:varchar(100)" json:"modify_by"`

	Event Event `gorm:"foreignKey:EventID" json:"event"`
}
