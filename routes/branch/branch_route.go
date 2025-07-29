package branch

import (
	"hris_backend/internal/handlers"
	"hris_backend/internal/middlewares"
	"hris_backend/internal/repositories"
	"hris_backend/internal/services"
	"hris_backend/pkg/r2"

	"github.com/gofiber/fiber/v2"
	"github.com/minio/minio-go/v7"
	"gorm.io/gorm"
)

func SetupBranchRoutes(router fiber.Router, db *gorm.DB, r2Client *minio.Client, bucketName string) {
	uploader := r2.NewUploader(r2Client, bucketName)
	branchRepo := repositories.NewBranchRepository(db)
	companyRepo := repositories.NewCompanyRepository(db)
	branchService := services.NewBranchService(branchRepo, companyRepo, *uploader)
	branchHandler := handlers.NewBranchHandler(branchService)

	branch := router.Group("/branches")
	branch.Use(middlewares.JWTMiddleware)

	branch.Get("/get/:uuid", middlewares.RequirePermission("branch.view"), branchHandler.Get)
	branch.Get("/get-by-company-uuid/:uuid", middlewares.RequirePermission("branch.view"), branchHandler.GetByCompanyUUID)
	branch.Get("/get-all", middlewares.RequirePermission("branch.view-all"), branchHandler.List)
	branch.Post("/create", middlewares.RequirePermission("branch.create"), branchHandler.Create)
	branch.Patch("/update/:uuid", middlewares.RequirePermission("branch.update"), branchHandler.Update)
	branch.Delete("/delete/:uuid", middlewares.RequirePermission("branch.delete"), branchHandler.Delete)
}
