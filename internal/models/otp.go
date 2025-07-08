package models

import (
	"time"
)

type OTP struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"-"`
	UUID      string    `gorm:"type:char(36);uniqueIndex" json:"uuid"`
	UserID    *uint     `gorm:"index" json:"user_id,omitempty"`
	Target    string    `gorm:"type:varchar(100);index;not null" json:"target"`
	Code      string    `gorm:"type:varchar(10);not null" json:"code"`
	Purpose   string    `gorm:"type:varchar(50);index;not null" json:"purpose"`
	IsUsed    bool      `gorm:"default:false" json:"is_used"`
	ExpiresAt time.Time `json:"-"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"-"`
	CreatedBy uint      `gorm:"not null" json:"-"`
	ModifyAt  time.Time `gorm:"autoUpdateTime" json:"-"`
	ModifyBy  uint      `gorm:"not null" json:"-"`
}
