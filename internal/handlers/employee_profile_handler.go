package handlers

import (
	"github.com/gofiber/fiber/v2"
	"hris_backend/internal/services"
)

type EmployeeProfileHandler struct {
	service services.EmployeeProfileService
}

func NewEmployeeProfileHandler(service services.EmployeeProfileService) *EmployeeProfileHandler {
	return &EmployeeProfileHandler{service}
}

func (h *EmployeeProfileHandler) GetMyProfile(c *fiber.Ctx) error {
	userID, ok := c.Locals("user_id").(uint)
	if !ok || userID == 0 {
		return fiber.NewError(fiber.StatusUnauthorized, "unauthorized")
	}

	profile, err := h.service.GetMyProfile(userID)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Profile retrieved successfully",
		"data":    profile,
	})
}
