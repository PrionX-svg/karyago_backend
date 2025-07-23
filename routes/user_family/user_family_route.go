package userfamily

import (
	"hris_backend/internal/handlers"
	"hris_backend/internal/middlewares"
	"hris_backend/internal/repositories"
	"hris_backend/internal/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupUserFamilyRoutes(router fiber.Router, db *gorm.DB) {
	userFamilyRepo := repositories.NewUserFamilyRepository(db)
	userRepo := repositories.NewUserRepository(db)
	userFamilyService := services.NewUserFamilyService(userFamilyRepo, userRepo)
	userFamilyHandler := handlers.NewUserFamilyHandler(userFamilyService)

	userFamily := router.Group("/user-families")
	userFamily.Use(middlewares.JWTMiddleware)

	userFamily.Post("/create", middlewares.RequirePermission("user-family.create"), userFamilyHandler.Create)
	userFamily.Get("/get-all", middlewares.RequirePermission("user-family.view-all"), userFamilyHandler.GetAll)
	userFamily.Get("/get-by-id/:id", middlewares.RequirePermission("user-family.view"), userFamilyHandler.GetByID)
	userFamily.Get("/get-by-uuid/:uuid", middlewares.RequirePermission("user-family.view"), userFamilyHandler.GetByUUID)
	userFamily.Get("/get-by-user/:uuid", middlewares.RequirePermission("user-family.view"), userFamilyHandler.GetByUserUUID)
	userFamily.Patch("/update/:user_uuid", middlewares.RequirePermission("user-family.update"), userFamilyHandler.Update)
	userFamily.Delete("/delete/:uuid", middlewares.RequirePermission("user-family.delete"), userFamilyHandler.Delete)
}
