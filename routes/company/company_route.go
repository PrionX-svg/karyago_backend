package company

import (
	"hris_backend/internal/handlers"
	"hris_backend/internal/repositories"
	"hris_backend/internal/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupCompanyRoutes(router fiber.Router, db *gorm.DB){
	companyRepo := repositories.NewCompanyRepository(db)

	companyService := services.NewCompanyService(companyRepo)
	companyHandler := handlers.NewCompanyHandler(companyService)

	company := router.Group("/company")

	company.Post("/create", companyHandler.Create)
	company.Get("/get-all", companyHandler.GetAll)
	company.Get("/get/:uuid", companyHandler.GetByUUID)
	company.Patch("/update/:uuid", companyHandler.Update)
	company.Delete("/delete/:uuid", companyHandler.Delete)
}