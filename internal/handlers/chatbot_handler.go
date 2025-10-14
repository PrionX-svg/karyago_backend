package handlers

import (
	"fmt"
	"hris_backend/internal/request"
	"hris_backend/internal/services"

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
	var req request.ChatbotRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}

	// Ambil dari JWT middleware
	uidVal := c.Locals("user_id")
	roleVal := c.Locals("role_id")

	// Hati-hati: kadang number dari JWT masuk sebagai float64
	var userID uint
	switch v := uidVal.(type) {
	case uint:
		userID = v
	case int:
		userID = uint(v)
	case float64:
		userID = uint(v)
	default:
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "unauthorized"})
	}

	// 🔧 Konversi role ke string (support angka dari JWT)
	var role string
	switch v := roleVal.(type) {
	case string:
		role = v
	case uint:
		m := map[int]string{
			1: "admin",
			2: "assistant",
			3: "owner",
		}
		role = m[int(v)]
	case float64: // JWT decode sebagai float64
		m := map[int]string{
			1: "admin",
			2: "assistant",
			3: "owner",
		}
		role = m[int(v)]
	default:
		role = "employee"
	}

	fmt.Println("🧩 user_id:", userID)
	fmt.Println("🧩 role (mapped):", role)

	// Jalankan service chatbot
	resp, err := h.chatbotServices.Ask(req, userID, role)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(resp)
}
