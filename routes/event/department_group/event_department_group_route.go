package department_group

import (
	"hris_backend/internal/handlers"
	"hris_backend/internal/middlewares"
	"hris_backend/internal/repositories"
	"hris_backend/internal/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupEventDepartmentGroupRoutes(router fiber.Router, db *gorm.DB) {
	groupRepo := repositories.NewEventDepartmentGroupRepository(db)
	userRepo := repositories.NewUserRepository(db)
	groupService := services.NewEventDepartmentGroupService(groupRepo, userRepo)
	groupHandler := handlers.NewEventDepartmentGroupHandler(groupService)

	group := router.Group("/event-department-groups")
	group.Use(middlewares.JWTMiddleware)

	group.Post("/create", middlewares.RequirePermission("event_department_group.create"), groupHandler.Create)
	group.Get("/get-all", middlewares.RequirePermission("event_department_group.view-all"), groupHandler.GetAll)
	group.Get("/get/:uuid", middlewares.RequirePermission("event_department_group.view"), groupHandler.GetByUUID)
	group.Patch("/update/:uuid", middlewares.RequirePermission("event_department_group.update"), groupHandler.Update)
	group.Delete("/delete/:uuid", middlewares.RequirePermission("event_department_group.delete"), groupHandler.Delete)
	group.Get("/get-event-uuid/:event_uuid", middlewares.RequirePermission("event_department_group.view-by-event-uuid"), groupHandler.GetByEventUUID)
}
