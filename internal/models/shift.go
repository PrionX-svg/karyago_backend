package models

import (
	"time"
)

type Shift struct {
	ID         uint       `gorm:"primaryKey;autoIncrement" json:"id"`
	UUID       string     `gorm:"type:char(36);uniqueIndex;not null" json:"uuid"`
	CompanyID  uint       `gorm:"not null" json:"company_id"`
	Name       string     `gorm:"type:varchar(100);not null" json:"name"`
	Date       time.Time  `gorm:"type:date;not null" json:"date"`
	HourFrom   time.Time  `gorm:"type:time;not null" json:"hour_from"`
	HourUntil  time.Time  `gorm:"type:time;not null" json:"hour_until"`
	CreatedAt  time.Time  `gorm:"autoCreateTime" json:"created_at"`
	CreatedBy  string     `gorm:"type:varchar(100)" json:"created_by"`
	ModifiedAt *time.Time `gorm:"autoUpdateTime" json:"modified_at"`
	ModifiedBy *string    `gorm:"type:varchar(100)" json:"modified_by"`

	// Relationship
	Details []CompanyDetailShift `gorm:"foreignKey:ShiftID" json:"details,omitempty"`
}
