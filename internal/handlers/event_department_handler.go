package handlers

import (
	"hris_backend/internal/request"
	"hris_backend/internal/services"
	"hris_backend/pkg"

	"github.com/gofiber/fiber/v2"
)

type EventDepartmentHandler interface {
	Create(c *fiber.Ctx) error
	GetAll(c *fiber.Ctx) error
	GetByUUID(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
	GetByGroupID(c *fiber.Ctx) error
}

type eventDepartmentHandler struct {
	service services.EventDepartmentServices
}

func NewEventDepartmentHandler(service services.EventDepartmentServices) EventDepartmentHandler {
	return &eventDepartmentHandler{service}
}

func (h *eventDepartmentHandler) Create(c *fiber.Ctx) error {
	var req request.EventDepartmentRequest
	if err := c.BodyParser(&req); err != nil {
		return pkg.Error(c, fiber.StatusBadRequest, "Failed to parse request")
	}

	result, err := h.service.Create(req)
	if err != nil {
		return pkg.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return pkg.Created(c, result, "Event Department created successfully")
}

func (h *eventDepartmentHandler) GetAll(c *fiber.Ctx) error {
	result, err := h.service.GetAll()
	if err != nil {
		return pkg.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return pkg.Success(c, result, "Successfully retrieved all Event Departments")
}

func (h *eventDepartmentHandler) GetByUUID(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	if uuid == "" {
		return pkg.Error(c, fiber.StatusBadRequest, "UUID is required")
	}

	result, err := h.service.GetByUUID(uuid)
	if err != nil {
		return pkg.Error(c, fiber.StatusNotFound, err.Error())
	}

	return pkg.Success(c, result, "Successfully retrieved Event Department")
}

func (h *eventDepartmentHandler) Update(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	if uuid == "" {
		return pkg.Error(c, fiber.StatusBadRequest, "UUID is required")
	}

	var req request.EventDepartmentRequest
	if err := c.BodyParser(&req); err != nil {
		return pkg.Error(c, fiber.StatusBadRequest, "Failed to parse request")
	}

	result, err := h.service.Update(uuid, req)
	if err != nil {
		return pkg.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return pkg.Success(c, result, "Event Department updated successfully")
}

func (h *eventDepartmentHandler) Delete(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	if uuid == "" {
		return pkg.Error(c, fiber.StatusBadRequest, "UUID is required")
	}

	result, err := h.service.Delete(uuid)
	if err != nil {
		return pkg.Error(c, fiber.StatusNotFound, err.Error())
	}

	return pkg.Success(c, result, "Event Department deleted successfully")
}

func (h *eventDepartmentHandler) GetByGroupID(c *fiber.Ctx) error {
	groupID, err := c.ParamsInt("group_id")
	if err != nil || groupID <= 0 {
		return pkg.Error(c, fiber.StatusBadRequest, "Invalid group_id")
	}

	result, err := h.service.GetByGroupID(uint(groupID))
	if err != nil {
		return pkg.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return pkg.Success(c, result, "Successfully retrieved departments by group ID")
}

