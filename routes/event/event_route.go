package event

import (
	"hris_backend/internal/handlers"
	"hris_backend/internal/middlewares"
	"hris_backend/internal/repositories"
	"hris_backend/internal/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupEventRoutes(router fiber.Router, db *gorm.DB) {
	eventRepo := repositories.NewEventRepository(db)
	companyRepo := repositories.NewCompanyRepository(db)

	eventService := services.NewEventService(eventRepo, companyRepo)
	eventHandler := handlers.NewEventHandler(eventService)

	event := router.Group("/events")
	event.Use(middlewares.JWTMiddleware)

	event.Post("/create", middlewares.RequirePermission("event.create"), eventHandler.Create)
	event.Get("/get-all", middlewares.RequirePermission("event.view-all"), eventHandler.GetAll)
	event.Get("/get-by-id/:id", middlewares.RequirePermission("event.view"), eventHandler.GetByID)
	event.Get("/get-by-uuid/:uuid", middlewares.RequirePermission("event.view"), eventHandler.GetByUUID)
	event.Get("/get-by-company/:company_uuid", middlewares.RequirePermission("event.view"), eventHandler.GetByCompanyUUID)
	event.Patch("/update/:uuid", middlewares.RequirePermission("event.update"), eventHandler.Update)
	event.Delete("/delete/:uuid", middlewares.RequirePermission("event.delete"), eventHandler.Delete)
}
