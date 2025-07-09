package branch

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"hris_backend/internal/handlers"
	"hris_backend/internal/middlewares"
	"hris_backend/internal/repositories"
	"hris_backend/internal/services"
)

func SetupBranchRoutes(router fiber.Router, db *gorm.DB) {
	branchRepo := repositories.NewBranchRepository(db)
	companyRepo := repositories.NewCompanyRepository(db)
	branchService := services.NewBranchService(branchRepo, companyRepo)
	branchHandler := handlers.NewBranchHandler(branchService)

	branch := router.Group("/branches")
	branch.Use(middlewares.JWTMiddleware)

	branch.Get("/get/:uuid", middlewares.RequirePermission("branch.view"), branchHandler.Get)
	branch.Get("/get-all", middlewares.RequirePermission("branch.view-all"), branchHandler.List)
	branch.Post("/create", middlewares.RequirePermission("branch.create"), branchHandler.Create)
	branch.Patch("/update/:uuid", middlewares.RequirePermission("branch.update"), branchHandler.Update)
	branch.Delete("/delete/:uuid", middlewares.RequirePermission("branch.delete"), branchHandler.Delete)
}
