package attendance

import (
	"hris_backend/internal/handlers"
	"hris_backend/internal/middlewares"
	"hris_backend/internal/repositories"
	"hris_backend/internal/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupAttendanceRoutes(router fiber.Router, db *gorm.DB) {
	// repos
	attRepo := repositories.NewAttendanceRepository(db)
	empRepo := repositories.NewEmployeeRepository(db)
	userRepo := repositories.NewUserRepository(db)
	companyRepo := repositories.NewCompanyRepository(db)

	// service & handler
	attService := services.NewAttendanceService(db, attRepo, empRepo, userRepo, companyRepo)
	attHandler := handlers.NewAttendanceHandler(attService)

	att := router.Group("/attendance")
	att.Use(middlewares.JWTMiddleware)

	// Actions
	att.Post("/clock-in", attHandler.ClockIn)
	att.Post("/clock-out", attHandler.ClockOut)
	att.Patch("/toggle-homeOffice", attHandler.ToggleHomeOffice)
	att.Patch("/notes", attHandler.SaveNotes)

	// Queries
	att.Get("/", attHandler.GetByDate)        // ?work_date=YYYY-MM-DD&company_uuid=...
	att.Get("/range", attHandler.ListRange)   // ?from=YYYY-MM-DD&to=YYYY-MM-DD&company_uuid=...
}
