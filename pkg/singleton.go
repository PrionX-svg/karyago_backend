package pkg

import (
	"gorm.io/gorm"
	"hris_backend/internal/repositories"
	"sync"
)

var (
	oncePermRepo           sync.Once
	permissionRepoInstance repositories.PermissionRepository
	GlobalDB               *gorm.DB
)

func InitPermissionRepo(db *gorm.DB) {
	GlobalDB = db
}

func GetPermissionRepo() repositories.PermissionRepository {
	oncePermRepo.Do(func() {
		permissionRepoInstance = repositories.NewPermissionRepository(GlobalDB)
	})
	return permissionRepoInstance
}
