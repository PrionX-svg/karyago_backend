package handlers

import (
	"fmt"
	"hris_backend/internal/request"
	"hris_backend/internal/services"
	"hris_backend/pkg"

	"github.com/gofiber/fiber/v2"
)

type UserEducationHandler struct {
	service services.UserEducationService
}

func NewUserEducationHandler(service services.UserEducationService) *UserEducationHandler {
	return &UserEducationHandler{
		service: service,
	}
}

func (h *UserEducationHandler) Create(c *fiber.Ctx) error {
	var req request.UserEducationReq

	userID, ok := c.Locals("user_id").(uint)
	if !ok {
		fmt.Println("Invalid user_id:", c.Locals("user_id"))
		return pkg.Error(c, fiber.StatusUnauthorized, "Unauthorized")
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
			"error":   err.Error(),
		})
	}

	userEducation, err := h.service.Create(req, userID); 
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create user education",
			"error":   err.Error(),
		})
	}

	return pkg.Created(c, userEducation, "User education created successfully")
}

// func (h *UserEducationHandler) GetAll() ([]models.UserEducation, error) {
// 	userEducations, err := h.service.GetAll()
// 	if err != nil {
// 		return nil, err
// 	}
// 	return userEducations, nil
// }
