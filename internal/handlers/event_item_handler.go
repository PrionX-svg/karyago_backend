package handlers

import (
	"hris_backend/internal/request"
	"hris_backend/internal/services"
	"hris_backend/pkg"

	"github.com/gofiber/fiber/v2"
)

type EventItemHandler interface {
	Create(c *fiber.Ctx) error
	GetAll(c *fiber.Ctx) error
	GetByUUID(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
	GetByEventUUID(c *fiber.Ctx) error
}

type eventItemHandler struct {
	service services.EventItemServices
}

func NewEventItemHandler(service services.EventItemServices) EventItemHandler {
	return &eventItemHandler{service}
}

func (h *eventItemHandler) Create(c *fiber.Ctx) error {
	var req request.EventItemRequest
	if err := c.BodyParser(&req); err != nil {
		return pkg.Error(c, fiber.StatusBadRequest, "Failed to parse request")
	}

	result, err := h.service.Create(req)
	if err != nil {
		return pkg.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return pkg.Created(c, result, "Event Item created successfully")
}

func (h *eventItemHandler) GetAll(c *fiber.Ctx) error {
	result, err := h.service.GetAll()
	if err != nil {
		return pkg.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return pkg.Success(c, result, "Successfully retrieved all Event Items")
}

func (h *eventItemHandler) GetByUUID(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	if uuid == "" {
		return pkg.Error(c, fiber.StatusBadRequest, "UUID is required")
	}

	result, err := h.service.GetByUUID(uuid)
	if err != nil {
		return pkg.Error(c, fiber.StatusNotFound, err.Error())
	}

	return pkg.Success(c, result, "Successfully retrieved Event Item")
}

func (h *eventItemHandler) Update(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	if uuid == "" {
		return pkg.Error(c, fiber.StatusBadRequest, "UUID is required")
	}

	var req request.EventItemRequest
	if err := c.BodyParser(&req); err != nil {
		return pkg.Error(c, fiber.StatusBadRequest, "Failed to parse request")
	}

	result, err := h.service.Update(uuid, req)
	if err != nil {
		return pkg.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return pkg.Success(c, result, "Event Item updated successfully")
}

func (h *eventItemHandler) Delete(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	if uuid == "" {
		return pkg.Error(c, fiber.StatusBadRequest, "UUID is required")
	}

	result, err := h.service.Delete(uuid)
	if err != nil {
		return pkg.Error(c, fiber.StatusNotFound, err.Error())
	}

	return pkg.Success(c, result, "Event Item deleted successfully")
}

func (h *eventItemHandler) GetByEventUUID(c *fiber.Ctx) error {
	eventUUID := c.Params("event_uuid")
	if eventUUID == "" {
		return pkg.Error(c, fiber.StatusBadRequest, "Missing event_uuid")
	}

	result, err := h.service.GetByEventUUID(eventUUID)
	if err != nil {
		return pkg.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return pkg.Success(c, result, "Successfully retrieved items by event UUID")
}
