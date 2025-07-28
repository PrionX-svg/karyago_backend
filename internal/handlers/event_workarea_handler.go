package handlers

import (
	"hris_backend/internal/request"
	"hris_backend/internal/services"
	"hris_backend/pkg"

	"github.com/gofiber/fiber/v2"
)

type EventWorkAreaHandler interface {
	Create(c *fiber.Ctx) error
	GetAll(c *fiber.Ctx) error
	GetByUUID(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
	GetByEventUUID(c *fiber.Ctx) error
}

type eventWorkAreaHandler struct {
	service services.EventWorkAreaServices
}

func NewEventWorkAreaHandler(service services.EventWorkAreaServices) EventWorkAreaHandler {
	return &eventWorkAreaHandler{service}
}

func (h *eventWorkAreaHandler) Create(c *fiber.Ctx) error {
	var req request.EventWorkAreaRequest
	if err := c.BodyParser(&req); err != nil {
		return pkg.Error(c, fiber.StatusBadRequest, "Failed to parse request")
	}

	result, err := h.service.Create(req)
	if err != nil {
		return pkg.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return pkg.Created(c, result, "Event Work Area created successfully")
}

func (h *eventWorkAreaHandler) GetAll(c *fiber.Ctx) error {
	result, err := h.service.GetAll()
	if err != nil {
		return pkg.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return pkg.Success(c, result, "Successfully retrieved all Event Work Areas")
}

func (h *eventWorkAreaHandler) GetByUUID(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	if uuid == "" {
		return pkg.Error(c, fiber.StatusBadRequest, "UUID is required")
	}

	result, err := h.service.GetByUUID(uuid)
	if err != nil {
		return pkg.Error(c, fiber.StatusNotFound, err.Error())
	}

	return pkg.Success(c, result, "Successfully retrieved Event Work Area")
}

func (h *eventWorkAreaHandler) Update(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	if uuid == "" {
		return pkg.Error(c, fiber.StatusBadRequest, "UUID is required")
	}

	var req request.EventWorkAreaRequest
	if err := c.BodyParser(&req); err != nil {
		return pkg.Error(c, fiber.StatusBadRequest, "Failed to parse request")
	}

	result, err := h.service.Update(uuid, req)
	if err != nil {
		return pkg.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return pkg.Success(c, result, "Event Work Area updated successfully")
}

func (h *eventWorkAreaHandler) Delete(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	if uuid == "" {
		return pkg.Error(c, fiber.StatusBadRequest, "UUID is required")
	}

	result, err := h.service.Delete(uuid)
	if err != nil {
		return pkg.Error(c, fiber.StatusNotFound, err.Error())
	}

	return pkg.Success(c, result, "Event Work Area deleted successfully")
}

func (h *eventWorkAreaHandler) GetByEventUUID(c *fiber.Ctx) error {
	eventUUID := c.Params("event_uuid")
	if eventUUID == "" {
		return pkg.Error(c, fiber.StatusBadRequest, "Missing event_uuid")
	}

	result, err := h.service.GetByEventUUID(eventUUID)
	if err != nil {
		return pkg.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return pkg.Success(c, result, "Successfully retrieved work areas by event UUID")
}
