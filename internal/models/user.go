package models

import (
	"time"
)

type User struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	UUID      string    `gorm:"type:char(36);uniqueIndex" json:"uuid"`
	RoleID    uint      `gorm:"not null" json:"role_id"`
	CompanyID *uint     `gorm:"default:null" json:"company_id"` // nullable
	FirstName string    `gorm:"type:varchar(100);not null" json:"firstname"`
	LastName  string    `gorm:"type:varchar(100);not null" json:"lastname"`
	Phone     string    `gorm:"type:varchar(20)" json:"phone"`
	Email     string    `gorm:"type:varchar(100);uniqueIndex;not null" json:"email"`
	Password  string    `gorm:"type:varchar(255);not null" json:"password"`
	Timezone  string    `gorm:"type:varchar(50);default:'UTC'" json:"timezone"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"-"`
	CreatedBy uint      `gorm:"not null" json:"-"`
	ModifyAt  time.Time `gorm:"autoUpdateTime" json:"-"`
	ModifyBy  uint      `gorm:"not null" json:"-"`
}
