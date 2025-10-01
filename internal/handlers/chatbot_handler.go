package handlers

import (
	"hris_backend/internal/request"
	"hris_backend/internal/services"
	"hris_backend/pkg"

	"github.com/gofiber/fiber/v2"
)

type ChatbotHandler interface {
	Ask(c *fiber.Ctx) error
}

type chatbotHandler struct {
	chatbotServices services.ChatbotServices
}

func NewChatbotHandler(chatbotServices services.ChatbotServices) ChatbotHandler {
	return &chatbotHandler{chatbotServices}
}

func (h *chatbotHandler) Ask(c *fiber.Ctx) error {
	var chatbotReq request.ChatbotRequest
	if err := c.BodyParser(&chatbotReq); err != nil {
		return pkg.Error(c, fiber.StatusBadRequest, "Failed to parse chatbot request")
	}

	resp, err := h.chatbotServices.Ask(chatbotReq)
	if err != nil {
		return pkg.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return pkg.Success(c, resp, "Chatbot response retrieved successfully")
}
