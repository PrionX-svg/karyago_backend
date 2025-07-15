package role

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"hris_backend/internal/handlers"
	"hris_backend/internal/middlewares"
	"hris_backend/internal/repositories"
	"hris_backend/internal/services"
)

func SetupRoleRoutes(router fiber.Router, db *gorm.DB) {
	roleRepo := repositories.NewRoleRepositories(db)
	companyRepo := repositories.NewCompanyRepository(db)
	employeeRepo := repositories.NewEmployeeRepository(db)
	roleService := services.NewRoleService(roleRepo, companyRepo)
	roleHandler := handlers.NewRoleHandler(roleService, companyRepo, employeeRepo, roleRepo)

	role := router.Group("/roles")
	role.Use(middlewares.JWTMiddleware)
	role.Get("/get/:uuid", middlewares.RequirePermission("role.view-uuid"), roleHandler.GetByUUID)
	role.Get("/get-by-name", middlewares.RequirePermission("role.view-name"), roleHandler.GetByName)
	role.Get("/get-all/dt", middlewares.RequirePermission("role.view-all"), roleHandler.GetAll)
	role.Post("/create", middlewares.RequirePermission("role.create"), roleHandler.Create)
	role.Patch("/update/:uuid", middlewares.RequirePermission("role.update"), roleHandler.Update)
	role.Delete("/delete/:uuid", middlewares.RequirePermission("role.delete"), roleHandler.Delete)
}
