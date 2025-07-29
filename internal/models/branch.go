package models

import (
	"time"
)

type Branch struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"-"`
	UUID      string    `gorm:"type:char(36);uniqueIndex" json:"uuid"`
	CompanyID uint      `gorm:"not null" json:"-"`
	Image     *string   `gorm:"type:varchar(255)" json:"image"`
	Name      string    `gorm:"type:varchar(100);not null" json:"name"`
	Address   string    `gorm:"type:varchar(255);not null" json:"address"`
	Email     string    `gorm:"type:varchar(100);uniqueIndex;not null" json:"email"`
	Phone     string    `gorm:"type:varchar(20)" json:"phone"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"-"`
	CreatedBy uint      `gorm:"not null" json:"-"`
	ModifyAt  time.Time `gorm:"autoUpdateTime" json:"-"`
	ModifyBy  uint      `gorm:"not null" json:"-"`

	Company Company `gorm:"foreignKey:CompanyID" json:"company"`
}
