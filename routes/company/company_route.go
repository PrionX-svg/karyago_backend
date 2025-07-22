package company

import (
	"hris_backend/internal/handlers"
	"hris_backend/internal/middlewares"
	"hris_backend/internal/repositories"
	"hris_backend/internal/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupCompanyRoutes(router fiber.Router, db *gorm.DB) {
	companyRepo := repositories.NewCompanyRepository(db)
	userRepo := repositories.NewUserRepository(db)
	employeeRepo := repositories.NewEmployeeRepository(db)
	roleRepo := repositories.NewRoleRepositories(db)
	companyService := services.NewCompanyService(companyRepo, userRepo, employeeRepo, roleRepo)
	companyHandler := handlers.NewCompanyHandler(companyService)

	company := router.Group("/companies")
	company.Use(middlewares.JWTMiddleware)

	company.Post("/create", middlewares.RequirePermission("company.create"), companyHandler.Create)
	company.Get("/get-all", middlewares.RequirePermission("company.view-all"), companyHandler.GetAll)
	company.Get("/get/:uuid", middlewares.RequirePermission("company.view"), companyHandler.GetByUUID)
	company.Get("/get-by-user/:uuid", middlewares.RequirePermission("company.view"), companyHandler.GetByUserUUID)
	company.Get("/get-companies-by-user/:uuid", middlewares.RequirePermission("company.view-companies"), companyHandler.GetCompaniesByUserUUID)
	company.Patch("/update/:uuid", middlewares.RequirePermission("company.update"), companyHandler.Update)
	company.Delete("/delete/:uuid", middlewares.RequirePermission("company.delete"), companyHandler.Delete)
}
