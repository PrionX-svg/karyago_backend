package employment_history

import (
	"hris_backend/internal/handlers"
	"hris_backend/internal/middlewares"
	"hris_backend/internal/repositories"
	"hris_backend/internal/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupEmploymentHistoryRoutes(router fiber.Router, db *gorm.DB) {
	employmentHistoryRepo := repositories.NewEmploymentHistoryRepository(db)
	employeeRepo := repositories.NewEmployeeRepository(db)
	roleRepo := repositories.NewRoleRepositories(db)
	companyRepo := repositories.NewCompanyRepository(db)
	branchRepo := repositories.NewBranchRepository(db)

	service := services.NewEmploymentHistoryService(
		db,
		employmentHistoryRepo,
		employeeRepo,
		roleRepo,
		companyRepo,
		branchRepo,
	)

	handler := handlers.NewEmploymentHistoryHandler(service)

	r := router.Group("/employment-histories")
	r.Use(middlewares.JWTMiddleware)

	r.Post("/create", middlewares.RequirePermission("employment_history.create"), handler.Create)
	r.Patch("/update/:uuid", middlewares.RequirePermission("employment_history.update"), handler.Update)
	r.Delete("/delete/:uuid", middlewares.RequirePermission("employment_history.delete"), handler.Delete)
	r.Get("/get/:uuid", middlewares.RequirePermission("employment_history.view"), handler.GetByUUID)
	r.Get("/employee/:employee_uuid", middlewares.RequirePermission("employment_history.view"), handler.GetByEmployeeUUID)
}
