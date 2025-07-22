package shift

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"hris_backend/internal/handlers"
	"hris_backend/internal/middlewares"
	"hris_backend/internal/repositories"
	"hris_backend/internal/services"
)

func SetupShiftRoutes(router fiber.Router, db *gorm.DB) {
	shiftRepo := repositories.NewShiftRepository(db)
	companyRepo := repositories.NewCompanyRepository(db)

	shiftService := services.NewShiftService(shiftRepo, companyRepo)
	shiftHandler := handlers.NewShiftHandler(shiftService)

	shift := router.Group("/shifts")
	shift.Use(middlewares.JWTMiddleware)

	shift.Post("/create", middlewares.RequirePermission("shift.create"), shiftHandler.Create)
	shift.Get("/get-all", middlewares.RequirePermission("shift.view-all"), shiftHandler.GetAll)
	shift.Get("/get/:uuid", middlewares.RequirePermission("shift.view"), shiftHandler.GetByUUID)
	shift.Patch("/update/:uuid", middlewares.RequirePermission("shift.update"), shiftHandler.Update)
	shift.Delete("/delete/:uuid", middlewares.RequirePermission("shift.delete"), shiftHandler.Delete)
}
