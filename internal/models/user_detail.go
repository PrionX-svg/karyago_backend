package models

import (
	"time"
)

type UserDetail struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"-"`
	UUID      string    `gorm:"type:char(36);uniqueIndex" json:"uuid"`
	UserID    uint      `gorm:"not null;index" json:"-"` // foreign key ke User
	UserUUID  string    `gorm:"->" json:"user_uuid"`
	TaxID     string    `gorm:"type:varchar(30)" json:"tax_id"`    // NPWP
	SocialID  string    `gorm:"type:varchar(30)" json:"social_id"` // KTP
	CreatedAt time.Time `gorm:"autoCreateTime" json:"-"`
	CreatedBy uint      `gorm:"not null" json:"-"`
	ModifyAt  time.Time `gorm:"autoUpdateTime" json:"-"`
	ModifyBy  uint      `gorm:"not null" json:"-"`
}
