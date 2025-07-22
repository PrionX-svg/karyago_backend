package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"hris_backend/internal/models"
	"hris_backend/internal/request"
	"hris_backend/internal/services"
)

type CompanyDetailShiftHandler interface {
	Create(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
	GetByUUID(c *fiber.Ctx) error
	GetAll(c *fiber.Ctx) error
	GetByShiftUUID(c *fiber.Ctx) error
}

type companyDetailShiftHandler struct {
	service services.CompanyDetailShiftService
}

func NewCompanyDetailShiftHandler(service services.CompanyDetailShiftService) CompanyDetailShiftHandler {
	return &companyDetailShiftHandler{service: service}
}

func (h *companyDetailShiftHandler) Create(c *fiber.Ctx) error {
	var req request.CompanyDetailShiftRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Failed to parse request body",
		})
	}

	shift, err := h.service.FindShiftByUUID(req.ShiftUUID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Shift not found",
		})
	}

	detail := models.CompanyDetailShift{
		UUID:    uuid.NewString(),
		ShiftID: shift.ID,
		Day1:    req.Day1,
		Day2:    req.Day2,
		Day3:    req.Day3,
		Day4:    req.Day4,
		Day5:    req.Day5,
		Day6:    req.Day6,
		Day7:    req.Day7,
	}

	if actorID, ok := c.Locals("user_id").(uint); ok {
		detail.CreatedBy = actorID
	}

	responseData, err := h.service.Create(&detail)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"message": "Company detail shift created successfully",
		"data":    responseData,
	})
}

func (h *companyDetailShiftHandler) Update(c *fiber.Ctx) error {
	uuidParam := c.Params("uuid")
	if uuidParam == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "UUID is required",
		})
	}

	var req request.CompanyDetailShiftRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Failed to parse request body",
		})
	}

	shift, err := h.service.FindShiftByUUID(req.ShiftUUID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Shift not found",
		})
	}

	updated := &models.CompanyDetailShift{
		ShiftID: shift.ID,
		Day1:    req.Day1,
		Day2:    req.Day2,
		Day3:    req.Day3,
		Day4:    req.Day4,
		Day5:    req.Day5,
		Day6:    req.Day6,
		Day7:    req.Day7,
	}

	if actorID, ok := c.Locals("user_id").(uint); ok {
		updated.ModifiedBy = &actorID
	}

	data, err := h.service.Update(uuidParam, updated)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "Company detail shift updated successfully",
		"data":    data,
	})
}

func (h *companyDetailShiftHandler) Delete(c *fiber.Ctx) error {
	uuidParam := c.Params("uuid")
	if uuidParam == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "UUID is required",
		})
	}

	if err := h.service.Delete(uuidParam); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "Company detail shift deleted successfully",
	})
}

func (h *companyDetailShiftHandler) GetByUUID(c *fiber.Ctx) error {
	uuidParam := c.Params("uuid")
	if uuidParam == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "UUID is required",
		})
	}

	result, err := h.service.GetByUUID(uuidParam)
	if err != nil || result == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Company detail shift not found",
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data":    result,
	})
}

func (h *companyDetailShiftHandler) GetAll(c *fiber.Ctx) error {
	result, err := h.service.GetAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"success": true,
		"data":    result,
	})
}

func (h *companyDetailShiftHandler) GetByShiftUUID(c *fiber.Ctx) error {
	shiftUUID := c.Params("shift_uuid")
	if shiftUUID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Shift UUID is required",
		})
	}

	result, err := h.service.GetByShiftUUID(shiftUUID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data":    result,
	})
}
