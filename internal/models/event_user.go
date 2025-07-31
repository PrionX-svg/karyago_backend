package models

import (
	"time"
)

type EventUser struct {
	ID          uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	UUID        string `gorm:"type:char(36);uniqueIndex" json:"uuid"`
	EventID      uint   `gorm:"not null" json:"event_id"`
	UserID      uint   `gorm:"not null" json:"user_id"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	CreatedBy uint      `gorm:"type:varchar(100)" json:"created_by"`
	ModifyAt  time.Time `gorm:"autoUpdateTime" json:"modify_at"`
	ModifyBy  uint      `gorm:"type:varchar(100)" json:"modify_by"`

	User User `gorm:"foreignKey:UserID" json:"user"`
	Event Event `gorm:"foreignKey:EventID" json:"event"`
}
