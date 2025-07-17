package handlers

import (
	"github.com/gofiber/fiber/v2"
	"hris_backend/internal/request"
	"hris_backend/internal/services"
)

type EmploymentHistoryHandler struct {
	service services.EmploymentHistoryService
}

func NewEmploymentHistoryHandler(service services.EmploymentHistoryService) *EmploymentHistoryHandler {
	return &EmploymentHistoryHandler{service: service}
}

func (h *EmploymentHistoryHandler) Create(c *fiber.Ctx) error {
	var req request.EmploymentHistoryRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
			"error":   err.Error(),
		})
	}

	userID, ok := c.Locals("user_id").(uint)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	result, err := h.service.Create(req, userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create employment history",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Employment history created successfully",
		"data":    result,
	})
}

func (h *EmploymentHistoryHandler) Update(c *fiber.Ctx) error {
	historyUUID := c.Params("uuid")
	if historyUUID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "UUID is required",
		})
	}

	var req request.EmploymentHistoryRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
			"error":   err.Error(),
		})
	}

	userID, ok := c.Locals("user_id").(uint)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	// Ubah pemanggilan ke service supaya return-nya response
	updatedHistory, err := h.service.Update(historyUUID, req, userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to update employment history",
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Employment history updated successfully",
		"data":    updatedHistory,
	})
}

func (h *EmploymentHistoryHandler) Delete(c *fiber.Ctx) error {
	historyUUID := c.Params("uuid")
	if historyUUID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "UUID is required",
		})
	}

	if err := h.service.Delete(historyUUID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to delete employment history",
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Employment history deleted successfully",
	})
}

func (h *EmploymentHistoryHandler) GetByUUID(c *fiber.Ctx) error {
	historyUUID := c.Params("uuid")
	if historyUUID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "UUID is required",
		})
	}

	data, err := h.service.GetByUUID(historyUUID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Employment history not found",
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"data": data,
	})
}

func (h *EmploymentHistoryHandler) GetByEmployeeUUID(c *fiber.Ctx) error {
	employeeUUID := c.Params("employee_uuid")
	if employeeUUID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "employee_uuid is required",
		})
	}

	data, err := h.service.GetByEmployeeUUID(employeeUUID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to get employment histories",
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"data": data,
	})
}
