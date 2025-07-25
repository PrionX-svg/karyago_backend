package models

import (
	"time"
)

type Event struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	UUID      string    `gorm:"type:char(36);uniqueIndex" json:"uuid"`
	CompanyID uint      `gorm:"not null" json:"company_id"`
	Name      string    `gorm:"type:varchar(100);not null" json:"name"`
	StartDate time.Time `gorm:"type:date;not null" json:"start_date"`
	EndDate   time.Time `gorm:"type:date;" json:"end_date"`
	Photo     *string    `gorm:"default:null" json:"photo"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	CreatedBy uint      `gorm:"type:varchar(100)" json:"created_by"`
	ModifyAt  time.Time `gorm:"autoUpdateTime" json:"modify_at"`
	ModifyBy  uint      `gorm:"type:varchar(100)" json:"modify_by"`

	Company Company `gorm:"foreignKey:CompanyID" json:"company"`
}
