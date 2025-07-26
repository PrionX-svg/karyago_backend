package eventitem

import (
	"hris_backend/internal/handlers"
	"hris_backend/internal/middlewares"
	"hris_backend/internal/repositories"
	"hris_backend/internal/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupEventItemRoutes(router fiber.Router, db *gorm.DB) {
	itemRepo := repositories.NewEventItemRepository(db)
	eventRepo := repositories.NewEventRepository(db)
	userRepo := repositories.NewUserRepository(db)

	itemService := services.NewEventItemService(itemRepo, eventRepo, userRepo)
	itemHandler := handlers.NewEventItemHandler(itemService)

	group := router.Group("/event-items")
	group.Use(middlewares.JWTMiddleware)

	group.Post("/create", middlewares.RequirePermission("event_item.create"), itemHandler.Create)
	group.Get("/get-all", middlewares.RequirePermission("event_item.view-all"), itemHandler.GetAll)
	group.Get("/get/:uuid", middlewares.RequirePermission("event_item.view"), itemHandler.GetByUUID)
	group.Patch("/update/:uuid", middlewares.RequirePermission("event_item.update"), itemHandler.Update)
	group.Delete("/delete/:uuid", middlewares.RequirePermission("event_item.delete"), itemHandler.Delete)
	group.Get("/get-by-event-uuid/:event_uuid", middlewares.RequirePermission("event_item.view-by-event"), itemHandler.GetByEventUUID)
}
