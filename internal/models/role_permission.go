package models

import (
	"time"
)

type RolePermission struct {
	ID           uint      `gorm:"primaryKey;autoIncrement" json:"-"`
	UUID         string    `gorm:"type:char(36);uniqueIndex" json:"uuid"`
	PermissionID uint      `gorm:"not null" json:"permission_id"`
	RoleID       uint      `gorm:"not null" json:"role_id"`
	CreatedAt    time.Time `gorm:"autoCreateTime" json:"-"`
	CreatedBy    uint      `gorm:"not null" json:"-"`
	ModifyAt     time.Time `gorm:"autoUpdateTime" json:"-"`
	ModifyBy     uint      `gorm:"not null" json:"-"`

	Permission Permission `gorm:"foreignKey:PermissionID" json:"-"`
	Role       Role       `gorm:"foreignKey:RoleID" json:"-"`
}
