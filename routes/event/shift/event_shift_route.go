package eventshift

import (
	"hris_backend/internal/handlers"
	"hris_backend/internal/middlewares"
	"hris_backend/internal/repositories"
	"hris_backend/internal/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupEventShiftRoutes(router fiber.Router, db *gorm.DB) {
	eventRepo := repositories.NewEventRepository(db)
	eventShiftRepo := repositories.NewEventShiftRepository(db)

	eventShiftService := services.NewEventShiftService(eventShiftRepo, eventRepo)
	eventShiftHandler := handlers.NewEventShiftHandler(eventShiftService)

	group := router.Group("/event-shifts")
	group.Use(middlewares.JWTMiddleware)

	group.Post("/create", middlewares.RequirePermission("event_shift.create"), eventShiftHandler.Create)
	group.Get("/get-all", middlewares.RequirePermission("event_shift.view-all"), eventShiftHandler.GetAll)
	group.Get("/get-by-id/:id", middlewares.RequirePermission("event_shift.view"), eventShiftHandler.GetByID)
	group.Get("/get-by-uuid/:uuid", middlewares.RequirePermission("event_shift.view"), eventShiftHandler.GetByUUID)
	group.Get("/get-by-event/:uuid", middlewares.RequirePermission("event_shift.view-by-event"), eventShiftHandler.GetByEventUUID)
	group.Patch("/update/:uuid", middlewares.RequirePermission("event_shift.update"), eventShiftHandler.Update)
	group.Delete("/delete/:uuid", middlewares.RequirePermission("event_shift.delete"), eventShiftHandler.Delete)
}
