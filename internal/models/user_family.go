package models

import (
	"time"
)

type UserFamily struct {
	ID       uint       `gorm:"primaryKey;autoIncrement" json:"id"`
	UUID     string    `gorm:"type:char(36);not null;unique" json:"uuid"`
	UserID   uint       `gorm:"not null" json:"user_id"`
	Status   bool      `gorm:"not null" json:"status"`
	Count    uint       `gorm:"not null" json:"count"`
	Member   string  `gorm:"type:varchar(255)" json:"member"`
	CreateAt time.Time `gorm:"autoCreateTime" json:"create_at"`
	CreateBy string    `gorm:"type:varchar(100)" json:"create_by"`
	ModifyAt time.Time `gorm:"autoUpdateTime" json:"modify_at"`
	ModifyBy string    `gorm:"type:varchar(100)" json:"modify_by"`

	User User `gorm:"foreignKey:UserID" json:"user"`
}
