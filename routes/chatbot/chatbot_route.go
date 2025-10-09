package chatbot

import (
	"hris_backend/internal/handlers"
	"hris_backend/internal/middlewares"
	"hris_backend/internal/services"

	"github.com/gofiber/fiber/v2"
)

func SetupChatbotRoutes(router fiber.Router, apiKey string) {
	chatbotService := services.NewChatbotService(apiKey, "./chatbot-info.txt")
	chatbotHandler := handlers.NewChatbotHandler(chatbotService)

	chat := router.Group("/chatbot")
	chat.Use(middlewares.JWTMiddleware)

	chat.Post("/ask", chatbotHandler.Ask)
}
