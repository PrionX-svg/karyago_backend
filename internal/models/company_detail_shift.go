package models

import "time"

type CompanyDetailShift struct {
	ID         uint       `gorm:"primaryKey;autoIncrement" json:"id"`
	UUID       string     `gorm:"type:char(36);uniqueIndex;not null" json:"uuid"`
	ShiftID    uint       `gorm:"not null" json:"shift_id"` // FK ke Shift
	Day1       int        `json:"day_1"`
	Day2       int        `json:"day_2"`
	Day3       int        `json:"day_3"`
	Day4       int        `json:"day_4"`
	Day5       int        `json:"day_5"`
	Day6       int        `json:"day_6"`
	Day7       int        `json:"day_7"`
	CreatedAt  time.Time  `gorm:"autoCreateTime" json:"created_at"`
	CreatedBy  uint       `gorm:"type:varchar(100)" json:"created_by"`
	ModifiedAt *time.Time `gorm:"autoUpdateTime" json:"modified_at"`
	ModifiedBy *uint      `gorm:"type:varchar(100)" json:"modified_by"`

	// Relationship
	Shift Shift `gorm:"foreignKey:ShiftID" json:"shift,omitempty"`
}
