package models

import "time"

type DepartmentGroup struct {
	ID              uint      `gorm:"primaryKey;autoIncrement" json:"-"`
	UUID            string    `gorm:"type:char(36);uniqueIndex" json:"uuid"`
	CompanyID       uint      `gorm:"not null" json:"-"`
	CompanyUUID     string    `gorm:"->" json:"company_uuid"`
	ResponsibleID   uint      `gorm:"not null" json:"-"`
	ResponsibleUUID string    `gorm:"->" json:"responsible_uuid"`
	Name            string    `gorm:"type:varchar(100);not null" json:"name"`
	Desc            string    `gorm:"type:varchar(255);not null" json:"desc"`
	CreatedAt       time.Time `gorm:"autoCreateTime" json:"-"`
	CreatedBy       uint      `gorm:"not null" json:"-"`
	ModifyAt        time.Time `gorm:"autoUpdateTime" json:"-"`
	ModifyBy        uint      `gorm:"not null" json:"-"`
}
