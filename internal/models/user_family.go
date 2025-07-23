package models

import (
	"time"
)

type UserFamily struct {
	ID       uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	UUID     string    `gorm:"type:char(36);not null;unique" json:"uuid"`
	UserID   uint      `gorm:"not null" json:"user_id"`
	Status   bool      `gorm:"not null" json:"status"`
	Name     string    `gorm:"type:varchar(255)" json:"name"`
	Relation string    `gorm:"type:varchar(100)" json:"relation"`
	Phone    string    `gorm:"type:varchar(20)" json:"phone"`
	CreateAt time.Time `gorm:"autoCreateTime" json:"create_at"`
	CreateBy uint      `gorm:"type:varchar(100)" json:"create_by"`
	ModifyAt time.Time `gorm:"autoUpdateTime" json:"modify_at"`
	ModifyBy uint      `gorm:"type:varchar(100)" json:"modify_by"`

	User User `gorm:"foreignKey:UserID" json:"user"`
}
