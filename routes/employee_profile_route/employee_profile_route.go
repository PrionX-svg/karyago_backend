package employee

import (
	"hris_backend/internal/handlers"
	"hris_backend/internal/middlewares"
	"hris_backend/internal/repositories"
	"hris_backend/internal/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupEmployeeProfileRoutes(router fiber.Router, db *gorm.DB) {
	userRepo := repositories.NewUserRepository(db)
	userDetailRepo := repositories.NewUserDetailRepository(db)
	employeeRepo := repositories.NewEmployeeRepository(db)
	companyRepo := repositories.NewCompanyRepository(db)
	deptRepo := repositories.NewDepartmentRepositories(db)
	branchRepo := repositories.NewBranchRepository(db)

	service := services.NewEmployeeProfileService(
		userRepo, userDetailRepo, employeeRepo, companyRepo, deptRepo, branchRepo,
	)
	handler := handlers.NewEmployeeProfileHandler(service)

	r := router.Group("/employee")
	r.Use(middlewares.JWTMiddleware)
	r.Get("/me", handler.GetMyProfile)
}
