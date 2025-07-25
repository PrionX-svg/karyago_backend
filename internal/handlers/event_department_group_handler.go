package handlers

import (
	"hris_backend/internal/request"
	"hris_backend/internal/services"
	"hris_backend/pkg"

	"github.com/gofiber/fiber/v2"
)

type EventDepartmentGroupHandler interface {
	Create(c *fiber.Ctx) error
	GetAll(c *fiber.Ctx) error
	GetByUUID(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
	GetByEventUUID(c *fiber.Ctx) error
}

type eventDepartmentGroupHandler struct {
	service services.EventDepartmentGroupServices
}

func NewEventDepartmentGroupHandler(service services.EventDepartmentGroupServices) EventDepartmentGroupHandler {
	return &eventDepartmentGroupHandler{service}
}

func (h *eventDepartmentGroupHandler) Create(c *fiber.Ctx) error {
	var req request.EventDepartmentGroupRequest
	if err := c.BodyParser(&req); err != nil {
		return pkg.Error(c, fiber.StatusBadRequest, "Failed to parse request")
	}

	actorID, ok := c.Locals("user_id").(uint)
	if !ok {
		return pkg.Error(c, fiber.StatusUnauthorized, "Unauthorized")
	}

	result, err := h.service.Create(actorID, req)
	if err != nil {
		return pkg.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return pkg.Created(c, result, "Event Department Group created successfully")
}

func (h *eventDepartmentGroupHandler) GetAll(c *fiber.Ctx) error {
	result, err := h.service.GetAll()
	if err != nil {
		return pkg.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return pkg.Success(c, result, "Successfully retrieved all Event Department Groups")
}

func (h *eventDepartmentGroupHandler) GetByUUID(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	if uuid == "" {
		return pkg.Error(c, fiber.StatusBadRequest, "UUID is required")
	}

	result, err := h.service.GetByUUID(uuid)
	if err != nil {
		return pkg.Error(c, fiber.StatusNotFound, err.Error())
	}

	return pkg.Success(c, result, "Successfully retrieved Event Department Group")
}

func (h *eventDepartmentGroupHandler) Update(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	if uuid == "" {
		return pkg.Error(c, fiber.StatusBadRequest, "UUID is required")
	}

	var req request.EventDepartmentGroupRequest
	if err := c.BodyParser(&req); err != nil {
		return pkg.Error(c, fiber.StatusBadRequest, "Failed to parse request")
	}

	actorID, ok := c.Locals("user_id").(uint)
	if !ok {
		return pkg.Error(c, fiber.StatusUnauthorized, "Unauthorized")
	}

	result, err := h.service.Update(actorID, uuid, req)
	if err != nil {
		return pkg.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return pkg.Success(c, result, "Event Department Group updated successfully")
}

func (h *eventDepartmentGroupHandler) Delete(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	if uuid == "" {
		return pkg.Error(c, fiber.StatusBadRequest, "UUID is required")
	}

	result, err := h.service.Delete(uuid)
	if err != nil {
		return pkg.Error(c, fiber.StatusNotFound, err.Error())
	}

	return pkg.Success(c, result, "Event Department Group deleted successfully")
}

func (h *eventDepartmentGroupHandler) GetByEventUUID(c *fiber.Ctx) error {
	eventUUID := c.Params("event_uuid")
	if eventUUID == "" {
		return pkg.Error(c, fiber.StatusBadRequest, "Missing event_uuid parameter")
	}

	result, err := h.service.GetByEventUUID(eventUUID)
	if err != nil {
		return pkg.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return pkg.Success(c, result, "Successfully retrieved department groups by event UUID")
}

