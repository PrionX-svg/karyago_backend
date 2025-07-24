package department

import (
	"hris_backend/internal/handlers"
	"hris_backend/internal/middlewares"
	"hris_backend/internal/repositories"
	"hris_backend/internal/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupDepartmentRoutes(router fiber.Router, db *gorm.DB) {
	departmentGroupRepo := repositories.NewDepartmentGroupRepositories(db)
	departmentRepo := repositories.NewDepartmentRepositories(db)
	employeeRepo := repositories.NewEmployeeRepository(db)

	departmentService := services.NewDepartmentServices(departmentRepo, departmentGroupRepo, employeeRepo)
	departmentHandler := handlers.NewDepartmentHandler(departmentService)

	department := router.Group("/departments")
	department.Use(middlewares.JWTMiddleware)

	department.Post("/create", middlewares.RequirePermission("department.create"), departmentHandler.Create)
	department.Get("/get-all", middlewares.RequirePermission("department.view-all"), departmentHandler.List)
	department.Get("/get/:uuid", middlewares.RequirePermission("department.view"), departmentHandler.Get)
	department.Get("/get-all/dt", middlewares.RequirePermission("department.view-all-dt"), departmentHandler.GetAllDataTable)
	department.Patch("/update/:uuid", middlewares.RequirePermission("department.update"), departmentHandler.Update)
	department.Delete("/delete/:uuid", middlewares.RequirePermission("department.delete"), departmentHandler.Delete)
}
