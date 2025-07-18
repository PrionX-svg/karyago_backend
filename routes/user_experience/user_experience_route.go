package userexperience

import (
	"hris_backend/internal/handlers"
	"hris_backend/internal/middlewares"
	"hris_backend/internal/repositories"
	"hris_backend/internal/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupUserExperienceRoutes (router fiber.Router, db *gorm.DB) {
	userExperienceRepo := repositories.NewUserExperienceRepository(db)
	userRepo := repositories.NewUserRepository(db)
	userExperienceService := services.NewUserExperienceService(userExperienceRepo, userRepo)
	userExperienceHandler := handlers.NewUserExperienceHandler(&userExperienceService)

	userEducation := router.Group("/user-experiences")
	userEducation.Use(middlewares.JWTMiddleware)

	userEducation.Post("/create", middlewares.RequirePermission("user-experience.create"), userExperienceHandler.Create)
	userEducation.Get("/get-all", middlewares.RequirePermission("user-experience.view-all"), userExperienceHandler.GetAll)
	userEducation.Get("/get-by-id/:id", middlewares.RequirePermission("user-experience.view"), userExperienceHandler.GetByID)
	userEducation.Get("/get-by-uuid/:uuid", middlewares.RequirePermission("user-experience.view"), userExperienceHandler.GetByUUID)
	userEducation.Get("/get-by-user/:uuid", middlewares.RequirePermission("user-experience.view"), userExperienceHandler.GetByUserUUID)
	userEducation.Patch("/update/:uuid", middlewares.RequirePermission("user-experience.update"), userExperienceHandler.Update)
	userEducation.Delete("/delete/:uuid", middlewares.RequirePermission("user-experience.delete"), userExperienceHandler.Delete)
}