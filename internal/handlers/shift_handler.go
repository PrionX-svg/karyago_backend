package handlers

import (
	"fmt"
	"hris_backend/internal/request"
	"hris_backend/internal/services"
	"hris_backend/pkg"

	"github.com/gofiber/fiber/v2"
)

type ShiftHandler interface {
	Create(c *fiber.Ctx) error
	GetAll(c *fiber.Ctx) error
	GetByUUID(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
}

type shiftHandler struct {
	shiftService services.ShiftService
}

func NewShiftHandler(shiftService services.ShiftService) ShiftHandler {
	return &shiftHandler{shiftService}
}

func (h *shiftHandler) Create(c *fiber.Ctx) error {
	var shiftReq request.ShiftRequest
	if err := c.BodyParser(&shiftReq); err != nil {
		return pkg.Error(c, fiber.StatusBadRequest, "Failed to parse shift request")
	}

	if actorID, ok := c.Locals("user_id").(uint); ok {
		shiftReq.CreatedBy = fmt.Sprint(actorID)
		shiftReq.ModifiedBy = fmt.Sprint(actorID)
	} else {
		return pkg.Error(c, fiber.StatusUnauthorized, "Unauthorized: user_id not found")
	}

	shift, err := h.shiftService.Create(shiftReq)
	if err != nil {
		return pkg.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return pkg.Created(c, shift, "Shift created successfully")
}

func (h *shiftHandler) GetAll(c *fiber.Ctx) error {
	shifts, err := h.shiftService.GetAll()
	if err != nil {
		return pkg.Error(c, fiber.StatusInternalServerError, "Failed to get shifts")
	}
	return pkg.Success(c, shifts, "Successfully retrieved all shifts")
}

func (h *shiftHandler) GetByUUID(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	if uuid == "" {
		return pkg.Error(c, fiber.StatusBadRequest, "UUID is required")
	}

	shift, err := h.shiftService.GetByUUID(uuid)
	if err != nil {
		return pkg.Error(c, fiber.StatusNotFound, "Shift not found")
	}

	return pkg.Success(c, shift, "Successfully retrieved shift")
}

func (h *shiftHandler) Update(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	if uuid == "" {
		return pkg.Error(c, fiber.StatusBadRequest, "UUID is required")
	}

	var shiftReq request.ShiftRequest
	if err := c.BodyParser(&shiftReq); err != nil {
		return pkg.Error(c, fiber.StatusBadRequest, "Failed to parse shift request")
	}

	if actorID, ok := c.Locals("user_id").(uint); ok {
		shiftReq.ModifiedBy = fmt.Sprint(actorID)
	} else {
		return pkg.Error(c, fiber.StatusUnauthorized, "Unauthorized: user_id not found")
	}

	updatedShift, err := h.shiftService.Update(uuid, shiftReq)
	if err != nil {
		return pkg.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return pkg.Success(c, updatedShift, "Shift updated successfully")
}

func (h *shiftHandler) Delete(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	if uuid == "" {
		return pkg.Error(c, fiber.StatusBadRequest, "UUID is required")
	}

	err := h.shiftService.Delete(uuid)
	if err != nil {
		return pkg.Error(c, fiber.StatusNotFound, err.Error())
	}

	return pkg.Success(c, nil, "Shift deleted successfully")
}
