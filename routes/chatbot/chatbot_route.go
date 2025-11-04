package chatbot

import (
	"hris_backend/internal/handlers"
	"hris_backend/internal/middlewares"
	"hris_backend/internal/repositories"
	"hris_backend/internal/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupChatbotRoutes(router fiber.Router, apiKey string, db *gorm.DB) {
	attRepo := repositories.NewAttendanceRepository(db)
	empRepo := repositories.NewEmployeeRepository(db)
	userRepo := repositories.NewUserRepository(db)
	companyRepo := repositories.NewCompanyRepository(db)
	attendanceService := services.NewAttendanceService(db, attRepo, empRepo, userRepo, companyRepo)
	chatbotService := services.NewChatbotService(apiKey, "chatbot-info.txt", attendanceService)
	chatbotHandler := handlers.NewChatbotHandler(chatbotService)

	chat := router.Group("/chatbot")
	chat.Use(middlewares.JWTMiddleware)

	chat.Post("/ask", chatbotHandler.Ask)
}
