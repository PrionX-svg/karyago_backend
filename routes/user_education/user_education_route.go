package usereducation

import (
	"hris_backend/internal/handlers"
	"hris_backend/internal/middlewares"
	"hris_backend/internal/repositories"
	"hris_backend/internal/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupUserEducationRoutes (router fiber.Router, db *gorm.DB) {
	userEducationRepo := repositories.NewUserEducationRepository(db)
	userEducationService := services.NewUserEducationService(userEducationRepo)
	userEducationHandler := handlers.NewUserEducationHandler(userEducationService)

	userEducation := router.Group("/user-educations")
	userEducation.Use(middlewares.JWTMiddleware)

	userEducation.Post("/create", middlewares.RequirePermission("user-education.create"), userEducationHandler.Create)
	// userEducation.Get("/get-all", middlewares.RequirePermission("user-education.view-all"), userEducationHandler.GetAll)
	// userEducation.Get("/get/:uuid", middlewares.RequirePermission("user-education.view"), userEducationHandler.GetByUUID)
	// userEducation.Get("/get-by-user/:uuid", middlewares.RequirePermission("user-education.view"), userEducationHandler.GetByUserUUID)
	// userEducation.Patch("/update/:uuid", middlewares.RequirePermission("user-education.update"), userEducationHandler.Update)
	// userEducation.Delete("/delete/:uuid", middlewares.RequirePermission("user-education.delete"), userEducationHandler.Delete)
}