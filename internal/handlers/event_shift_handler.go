package handlers

import (
	"fmt"
	"hris_backend/internal/request"
	"hris_backend/internal/response"
	"hris_backend/internal/services"
	"hris_backend/pkg"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type EventShiftHandler struct {
	service services.EventShiftService
}

func NewEventShiftHandler(service services.EventShiftService) *EventShiftHandler {
	return &EventShiftHandler{
		service: service,
	}
}

func (h *EventShiftHandler) Create(c *fiber.Ctx) error {
	var req request.EventShiftReq

	userID, ok := c.Locals("user_id").(uint)
	if !ok {
		fmt.Println("Invalid user_id:", c.Locals("user_id"))
		return pkg.Error(c, fiber.StatusUnauthorized, "Unauthorized")
	}

	if err := c.BodyParser(&req); err != nil {
		return pkg.Error(c, fiber.StatusBadRequest, "Invalid request body")
	}

	eventShift, err := h.service.Create(req, userID)
	if err != nil {
		if err.Error() == "invalid event UUID" {
			return pkg.Error(c, fiber.StatusBadRequest, err.Error())
		}
		return pkg.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return pkg.Created(c, eventShift, "Event shift created successfully")
}

func (h *EventShiftHandler) GetAll(c *fiber.Ctx) error {
	shifts, err := h.service.GetAll()
	if err != nil {
		return pkg.Error(c, fiber.StatusInternalServerError, "Failed to retrieve event shifts")
	}

	return pkg.Success(c, shifts, "Successfully retrieved event shifts")
}

func (h *EventShiftHandler) GetByID(c *fiber.Ctx) error {
	idStr := c.Params("id")
	if idStr == "" {
		return pkg.Error(c, fiber.StatusBadRequest, "ID is required")
	}

	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return pkg.Error(c, fiber.StatusBadRequest, "Invalid ID format")
	}

	shift, err := h.service.GetByID(uint(id))
	if err != nil {
		return pkg.Error(c, fiber.StatusInternalServerError, "Failed to retrieve event shift")
	}

	if (shift == response.EventShiftResponse{}) {
		return pkg.Error(c, fiber.StatusNotFound, "Event shift not found")
	}

	return pkg.Success(c, shift, "Successfully retrieved event shift")
}

func (h *EventShiftHandler) GetByUUID(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	if uuid == "" {
		return pkg.Error(c, fiber.StatusBadRequest, "UUID is required")
	}

	shift, err := h.service.GetByUUID(uuid)
	if err != nil {
		return pkg.Error(c, fiber.StatusInternalServerError, "Failed to retrieve event shift")
	}

	if (shift == response.EventShiftResponse{}) {
		return pkg.Error(c, fiber.StatusNotFound, "Event shift not found")
	}

	return pkg.Success(c, shift, "Successfully retrieved event shift")
}

func (h *EventShiftHandler) GetByEventUUID(c *fiber.Ctx) error {
	eventUUID := c.Params("uuid")
	if eventUUID == "" {
		return pkg.Error(c, fiber.StatusBadRequest, "Event UUID is required")
	}

	shifts, err := h.service.GetByEventUUID(eventUUID)
	if err != nil {
		return pkg.Error(c, fiber.StatusInternalServerError, "Failed to retrieve event shifts")
	}

	if len(shifts) == 0 {
		return pkg.Error(c, fiber.StatusNotFound, "No event shifts found")
	}

	return pkg.Success(c, shifts, "Successfully retrieved event shifts")
}

func (h *EventShiftHandler) Update(c *fiber.Ctx) error {
	var req request.EventShiftReq

	uuid := c.Params("uuid")
	if uuid == "" {
		return pkg.Error(c, fiber.StatusBadRequest, "UUID is required")
	}

	userID, ok := c.Locals("user_id").(uint)
	if !ok {
		fmt.Println("Invalid user_id:", c.Locals("user_id"))
		return pkg.Error(c, fiber.StatusUnauthorized, "Unauthorized")
	}

	if err := c.BodyParser(&req); err != nil {
		return pkg.Error(c, fiber.StatusBadRequest, "Invalid request body")
	}

	shift, err := h.service.Update(uuid, req, userID)
	if err != nil {
		return pkg.Error(c, fiber.StatusInternalServerError, "Failed to update event shift")
	}

	return pkg.Success(c, shift, "Event shift updated successfully")
}

func (h *EventShiftHandler) Delete(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	if uuid == "" {
		return pkg.Error(c, fiber.StatusBadRequest, "UUID is required")
	}

	shift, err := h.service.Delete(uuid)
	if err != nil {
		return pkg.Error(c, fiber.StatusInternalServerError, "Failed to delete event shift")
	}

	return pkg.Success(c, shift, "Event shift deleted successfully")
}
