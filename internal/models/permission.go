package models

import (
	"time"
)

type Permission struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"-"`
	UUID      string    `gorm:"type:char(36);uniqueIndex" json:"uuid"`
	Label     string    `gorm:"type:varchar(100);not null" json:"label"`
	Name      string    `gorm:"type:varchar(100);not null" json:"name"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"-"`
	CreatedBy uint      `gorm:"not null" json:"-"`
	ModifyAt  time.Time `gorm:"autoUpdateTime" json:"-"`
	ModifyBy  uint      `gorm:"not null" json:"-"`
}
