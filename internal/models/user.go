package models

import (
	"time"
)

type User struct {
	ID          uint       `gorm:"primaryKey;autoIncrement" json:"-"`
	UUID        string     `gorm:"type:char(36);uniqueIndex" json:"uuid"`
	RoleID      uint       `gorm:"not null" json:"role_id"`
	BranchID    *uint      `gorm:"default:null" json:"-"`
	FirstName   string     `gorm:"type:varchar(100);not null" json:"firstname"`
	LastName    string     `gorm:"type:varchar(100);not null" json:"lastname"`
	Phone       string     `gorm:"type:varchar(20)" json:"phone"`
	DOB         *time.Time `gorm:"type:date" json:"dob"` // nullable
	Gender      *string    `gorm:"type:enum('male','female')" json:"gender"`
	Email       string     `gorm:"type:varchar(100);uniqueIndex;not null" json:"email"`
	Password    string     `gorm:"type:varchar(255);not null" json:"password"`
	Timezone    string     `gorm:"type:varchar(50);default:'UTC'" json:"timezone"`
	IsVerified  bool       `gorm:"default:false" json:"is_verified"`
	IsFreelance bool       `gorm:"default:false" json:"is_freelance"`
	CreatedAt   time.Time  `gorm:"autoCreateTime" json:"-"`
	CreatedBy   uint       `gorm:"not null" json:"-"`
	ModifyAt    time.Time  `gorm:"autoUpdateTime" json:"-"`
	ModifyBy    uint       `gorm:"not null" json:"-"`

	Role Role `json:"role"`
}
