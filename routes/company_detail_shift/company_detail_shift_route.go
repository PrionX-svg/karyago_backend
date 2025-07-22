package company_detail_shift

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"hris_backend/internal/handlers"
	"hris_backend/internal/middlewares"
	"hris_backend/internal/repositories"
	"hris_backend/internal/services"
)

func SetupCompanyDetailShiftRoutes(router fiber.Router, db *gorm.DB) {
	companyDetailShiftRepo := repositories.NewCompanyDetailShiftRepository(db)
	shiftRepo := repositories.NewShiftRepository(db)
	companyDetailShiftService := services.NewCompanyDetailShiftService(companyDetailShiftRepo, shiftRepo)
	companyDetailShiftHandler := handlers.NewCompanyDetailShiftHandler(companyDetailShiftService)

	r := router.Group("/company-detail-shifts")
	r.Use(middlewares.JWTMiddleware)

	r.Post("/create", middlewares.RequirePermission("company_detail_shift.create"), companyDetailShiftHandler.Create)
	r.Patch("/update/:uuid", middlewares.RequirePermission("company_detail_shift.update"), companyDetailShiftHandler.Update)
	r.Delete("/delete/:uuid", middlewares.RequirePermission("company_detail_shift.delete"), companyDetailShiftHandler.Delete)
	r.Get("/get/:uuid", middlewares.RequirePermission("company_detail_shift.view"), companyDetailShiftHandler.GetByUUID)
	r.Get("/get-all", middlewares.RequirePermission("company_detail_shift.view-all"), companyDetailShiftHandler.GetAll)
	r.Get("/get-by-shift/:shift_uuid", middlewares.RequirePermission("company_detail_shift.view-by-shift-id"), companyDetailShiftHandler.GetByShiftUUID)
}
