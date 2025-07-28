package workarea

import (
	"hris_backend/internal/handlers"
	"hris_backend/internal/middlewares"
	"hris_backend/internal/repositories"
	"hris_backend/internal/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupEventWorkAreaRoutes(router fiber.Router, db *gorm.DB) {
	workAreaRepo := repositories.NewEventWorkAreaRepository(db)
	eventRepo := repositories.NewEventRepository(db)
	userRepo := repositories.NewUserRepository(db)

	workAreaService := services.NewEventWorkAreaService(workAreaRepo, eventRepo, userRepo)
	workAreaHandler := handlers.NewEventWorkAreaHandler(workAreaService)

	group := router.Group("/event-work-areas")
	group.Use(middlewares.JWTMiddleware)

	group.Post("/create", middlewares.RequirePermission("event_workarea.create"), workAreaHandler.Create)
	group.Get("/get-all", middlewares.RequirePermission("event_workarea.view-all"), workAreaHandler.GetAll)
	group.Get("/get/:uuid", middlewares.RequirePermission("event_workarea.view"), workAreaHandler.GetByUUID)
	group.Patch("/update/:uuid", middlewares.RequirePermission("event_workarea.update"), workAreaHandler.Update)
	group.Delete("/delete/:uuid", middlewares.RequirePermission("event_workarea.delete"), workAreaHandler.Delete)
	group.Get("/get-by-event-uuid/:event_uuid", middlewares.RequirePermission("event_workarea.view-by-event"), workAreaHandler.GetByEventUUID)
}

