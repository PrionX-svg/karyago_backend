package database

import (
	"hris_backend/internal/models"
	"log"
)

func MigrationAll() {
	err := DB.AutoMigrate(
		&models.User{},
		&models.Role{},
		&models.OTP{},
		&models.Permission{},
		&models.RolePermission{},
		&models.Company{},
		&models.Branch{},
		&models.UserDetail{},
		&models.DepartmentGroup{},
		&models.Employee{},
		&models.Department{},
		&models.EmploymentHistory{},
		&models.UserEducation{},
		&models.UserBank{},
		&models.UserExperience{},
		&models.Shift{},
		&models.CompanyDetailShift{},
		&models.UserFamily{},
		&models.Event{},
		&models.EventDepartmentGroup{},
		&models.EventDepartment{},
		&models.EventItem{},
		&models.EventWorkArea{},
	)
	if err != nil {
		log.Fatal("Migration failed:", err)
	}

	SeedRoles(DB)
}
