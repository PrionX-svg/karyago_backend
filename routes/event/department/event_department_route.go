package department

import (
	"hris_backend/internal/handlers"
	"hris_backend/internal/middlewares"
	"hris_backend/internal/repositories"
	"hris_backend/internal/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupEventDepartmentRoutes(router fiber.Router, db *gorm.DB) {
	deptRepo := repositories.NewEventDepartmentRepository(db)
	groupRepo := repositories.NewEventDepartmentGroupRepository(db)
	userRepo := repositories.NewUserRepository(db)

	deptService := services.NewEventDepartmentService(deptRepo, groupRepo, userRepo)
	deptHandler := handlers.NewEventDepartmentHandler(deptService)

	group := router.Group("/event-departments")
	group.Use(middlewares.JWTMiddleware)

	group.Post("/create", middlewares.RequirePermission("event_department.create"), deptHandler.Create)
	group.Get("/get-all", middlewares.RequirePermission("event_department.view-all"), deptHandler.GetAll)
	group.Get("/get/:uuid", middlewares.RequirePermission("event_department.view"), deptHandler.GetByUUID)
	group.Patch("/update/:uuid", middlewares.RequirePermission("event_department.update"), deptHandler.Update)
	group.Delete("/delete/:uuid", middlewares.RequirePermission("event_department.delete"), deptHandler.Delete)
	group.Get("/get-by-group-uuid/:group_uuid", middlewares.RequirePermission("event_department.view-by-group"), deptHandler.GetByGroupUUID)
}

