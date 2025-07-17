package models

import (
	"time"
)

type UserEducation struct {
	ID          uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	UUID        string `gorm:"type:char(36);uniqueIndex" json:"uuid"`
	UserID      uint   `gorm:"not null" json:"user_id"`
	UserUUID    string `gorm:"->" json:"user_uuid"`
	Name        string `gorm:"type:varchar(100);not null" json:"name"`
	Description string `gorm:"type:text" json:"description,omitempty"`
	Location    string `gorm:"type:varchar(100);not null" json:"location"`
	StartDate time.Time `gorm:"type:date;not null" json:"start_date"`
	EndDate   time.Time `gorm:"type:date;" json:"end_date"`
	Grade float64 `gorm:"type:decimal(4,2);not null" json:"grade"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	CreatedBy uint      `gorm:"type:varchar(100)" json:"created_by"`
	ModifyAt  time.Time `gorm:"autoUpdateTime" json:"modify_at"`
	ModifyBy  uint      `gorm:"type:varchar(100)" json:"modify_by"`

	User User `gorm:"foreignKey:UserID" json:"user"`
}
