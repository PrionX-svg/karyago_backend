package models

import (
	"time"
)

type UserExperience struct {
	ID          uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	UUID        string `gorm:"type:char(36);uniqueIndex" json:"uuid"`
	UserID      uint   `gorm:"not null" json:"user_id"`
	Name        string `gorm:"type:varchar(100);not null" json:"name"`
	Description string `gorm:"type:text" json:"description,omitempty"`
	StartDate time.Time `gorm:"type:date;not null" json:"start_date"`
	EndDate   time.Time `gorm:"type:date;" json:"end_date"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	CreatedBy uint      `gorm:"type:varchar(100)" json:"created_by"`
	ModifyAt  time.Time `gorm:"autoUpdateTime" json:"modify_at"`
	ModifyBy  uint      `gorm:"type:varchar(100)" json:"modify_by"`

	User User `gorm:"foreignKey:UserID" json:"user"`
}
