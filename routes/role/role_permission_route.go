package role

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"hris_backend/internal/handlers"
	"hris_backend/internal/middlewares"
	"hris_backend/internal/repositories"
	"hris_backend/internal/services"
)

func SetupRolePermissionRoutes(router fiber.Router, db *gorm.DB) {
	rpRepo := repositories.NewRolePermissionRepository(db)
	roleRepo := repositories.NewRoleRepositories(db)
	permRepo := repositories.NewPermissionRepository(db)

	rpSvc := services.NewRolePermissionService(rpRepo, roleRepo, permRepo)
	rpHnd := handlers.NewRolePermissionHandler(rpSvc)

	rp := router.Group("/role-permissions")
	rp.Use(middlewares.JWTMiddleware)
	rp.Post("/sync", middlewares.RequirePermission("role.sync"), rpHnd.AssignAndRemovePermissions)
	rp.Get("/:role_uuid", middlewares.RequirePermission("role.view"), rpHnd.ListPermissionsByRole)
}
