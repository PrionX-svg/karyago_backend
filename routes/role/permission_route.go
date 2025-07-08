package role

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"hris_backend/internal/handlers"
	"hris_backend/internal/middlewares"
	"hris_backend/internal/repositories"
	"hris_backend/internal/services"
)

func SetupPermissionRoutes(router fiber.Router, db *gorm.DB) {
	repo := repositories.NewPermissionRepository(db)
	service := services.NewPermissionService(repo)
	handler := handlers.NewPermissionHandler(service)

	p := router.Group("/permissions")
	p.Use(middlewares.JWTMiddleware)
	p.Get("/get-all", middlewares.RequirePermission("permission.view-all"), handler.GetAll)
	p.Get("/get/:uuid", middlewares.RequirePermission("permission.view"), handler.GetByUUID)
}
