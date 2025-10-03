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
	attRepo := repositories.NewAttendanceRepository(db)
	editAttRepo := repositories.NewAttendanceEditRepository(db)
	empRepo := repositories.NewEmployeeRepository(db)
	userRepo := repositories.NewUserRepository(db)
	companyRepo := repositories.NewCompanyRepository(db)

	notify := services.NewSMTPNotifier()
	attService := services.NewAttendanceService(db, attRepo, empRepo, userRepo, companyRepo)
	attHandler := handlers.NewAttendanceHandler(attService)
	editAttService := services.NewAttendanceEditService(db, editAttRepo, attRepo, empRepo, userRepo, companyRepo, notify)
	editAttHandler := handlers.NewAttendanceEditHandler(editAttService)

	att := router.Group("/attendance")
	att.Use(middlewares.JWTMiddleware)

	// Actions
	att.Get("/calendar", attHandler.ListCalendar)
	att.Post("/clock-in", attHandler.ClockIn)
	att.Post("/clock-out", attHandler.ClockOut)
	att.Patch("/toggle-homeOffice", attHandler.ToggleHomeOffice)
	att.Patch("/notes", attHandler.SaveNotes)

	// Employee
	att.Post("/edit-requests",
		middlewares.RequirePermission("attendance.edit.create"),
		editAttHandler.Create,
	)
	att.Get("/edit-requests/my",
		middlewares.RequirePermission("attendance.edit.view_self"),
		editAttHandler.ListMine,
	)

	// Supervisor/HR
	att.Get("/edit-requests",
		middlewares.RequirePermission("attendance.edit.view_all"),
		editAttHandler.ListForSupervisor,
	)
	att.Post("/edit-requests/:id/approve",
		middlewares.RequirePermission("attendance.edit.approve"),
		editAttHandler.Approve,
	)
	att.Post("/edit-requests/:id/reject",
		middlewares.RequirePermission("attendance.edit.reject"),
		editAttHandler.Reject,
	)

	// Queries
	att.Get("/", attHandler.GetByDate)      // ?work_date=YYYY-MM-DD&company_uuid=...
	att.Get("/range", attHandler.ListRange) // ?from=YYYY-MM-DD&to=YYYY-MM-DD&company_uuid=...
}
