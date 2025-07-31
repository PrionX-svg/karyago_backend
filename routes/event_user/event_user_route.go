package eventuser

import (
	"hris_backend/internal/handlers"
	"hris_backend/internal/middlewares"
	"hris_backend/internal/repositories"
	"hris_backend/internal/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupEventUserRoutes(router fiber.Router, db *gorm.DB) {
	eventUserRepo := repositories.NewEventUserRepository(db)
	userRepo := repositories.NewUserRepository(db)
	eventRepo := repositories.NewEventRepository(db)
	eventUserService := services.NewEventUserService(eventUserRepo, userRepo, eventRepo)
	eventUserHandler := handlers.NewEventUserHandler(eventUserService)

	eventUser := router.Group("/event-users")
	eventUser.Use(middlewares.JWTMiddleware)

	eventUser.Post("/create", middlewares.RequirePermission("event-user.create"), eventUserHandler.Create)
	eventUser.Get("/get-all", middlewares.RequirePermission("event-user.view-all"), eventUserHandler.GetAll)
	eventUser.Get("/get-by-id/:id", middlewares.RequirePermission("event-user.view"), eventUserHandler.GetByID)
	eventUser.Get("/get-by-uuid/:uuid", middlewares.RequirePermission("event-user.view"), eventUserHandler.GetByUUID)
	eventUser.Get("/get-by-user/:uuid", middlewares.RequirePermission("event-user.view"), eventUserHandler.GetByUserUUID)
	eventUser.Get("/get-by-event/:uuid", middlewares.RequirePermission("event-user.view"), eventUserHandler.GetByEventUUID)
	eventUser.Patch("/update/:uuid", middlewares.RequirePermission("event-user.update"), eventUserHandler.Update)
	eventUser.Delete("/delete/:uuid", middlewares.RequirePermission("event-user.delete"), eventUserHandler.Delete)
}
