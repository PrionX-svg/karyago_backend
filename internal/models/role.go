package models

import (
	"time"
)

type Role struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	UUID      string    `gorm:"type:char(36);uniqueIndex" json:"uuid"`
	Name      string    `gorm:"type:varchar(100);not null" json:"name"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"-"`
	CreatedBy uint      `gorm:"not null" json:"-"`
	ModifyAt  time.Time `gorm:"autoUpdateTime" json:"-"`
	ModifyBy  uint      `gorm:"not null" json:"-"`
}
