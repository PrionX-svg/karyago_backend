package handlers

import (
	"fmt"
	"hris_backend/internal/request"
	"hris_backend/internal/services"
	"hris_backend/pkg"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type EventHandler interface {
	Create(c *fiber.Ctx) error
	GetAll(c *fiber.Ctx) error
	GetByID(c *fiber.Ctx) error
	GetByUUID(c *fiber.Ctx) error
	GetByCompanyUUID(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
}

type eventHandler struct {
	eventService services.EventService
}

func NewEventHandler(service services.EventService) EventHandler {
	return &eventHandler{
		eventService: service,
	}
}

func (h *eventHandler) Create(c *fiber.Ctx) error {
	fmt.Print("Creating event... di handler nih coy\n")
	var eventReq request.EventReq
	if err := c.BodyParser(&eventReq); err != nil {
		return pkg.Error(c, fiber.StatusBadRequest, "Failed to parse event")
	}

	actorID, ok := c.Locals("user_id").(uint)
	if !ok {
		return pkg.Error(c, fiber.StatusUnauthorized, "User ID not found or invalid")
	}

	event, err := h.eventService.Create(eventReq, actorID)
	if err != nil {
		return pkg.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return pkg.Created(c, event, "Event created successfully")
}

func (h *eventHandler) GetAll(c *fiber.Ctx) error {
	events, err := h.eventService.GetAll()
	if err != nil {
		return pkg.Error(c, fiber.StatusInternalServerError, "Failed to retrieve events")
	}
	return pkg.Success(c, events, "Successfully retrieved all events")
}

func (h *eventHandler) GetByID(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil || id <= 0 {
		return pkg.Error(c, fiber.StatusBadRequest, "Invalid event ID")
	}

	event, err := h.eventService.GetByID(uint(id))
	if err != nil {
		return pkg.Error(c, fiber.StatusNotFound, "Event not found")
	}

	return pkg.Success(c, event, "Successfully retrieved event by ID")
}

func (h *eventHandler) GetByUUID(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	if uuid == "" {
		return pkg.Error(c, fiber.StatusBadRequest, "UUID is required")
	}

	event, err := h.eventService.GetByUUID(uuid)
	if err != nil {
		return pkg.Error(c, fiber.StatusNotFound, "Event not found")
	}

	return pkg.Success(c, event, "Successfully retrieved event")
}

func (h *eventHandler) GetByCompanyUUID(c *fiber.Ctx) error {
	companyUUID := c.Params("company_uuid")
	if companyUUID == "" {
		return pkg.Error(c, fiber.StatusBadRequest, "Company UUID is required")
	}

	events, err := h.eventService.GetByCompanyUUID(companyUUID)
	if err != nil {
		return pkg.Error(c, fiber.StatusInternalServerError, "Failed to retrieve events for company")
	}

	return pkg.Success(c, events, "Successfully retrieved company events")
}

func (h *eventHandler) Update(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	if uuid == "" {
		return pkg.Error(c, fiber.StatusBadRequest, "UUID is required")
	}

	var eventReq request.EventReq
	if err := c.BodyParser(&eventReq); err != nil {
		return pkg.Error(c, fiber.StatusBadRequest, "Failed to parse event")
	}

	actorID, ok := c.Locals("user_id").(uint)
	if !ok {
		return pkg.Error(c, fiber.StatusUnauthorized, "User ID not found or invalid")
	}

	updatedEvent, err := h.eventService.Update(uuid, eventReq, actorID)
	if err != nil {
		return pkg.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return pkg.Success(c, updatedEvent, "Event updated successfully")
}

func (h *eventHandler) Delete(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	if uuid == "" {
		return pkg.Error(c, fiber.StatusBadRequest, "UUID is required")
	}

	deletedEvent, err := h.eventService.Delete(uuid)
	if err != nil {
		return pkg.Error(c, fiber.StatusInternalServerError, "Failed to delete event")
	}

	return pkg.Success(c, deletedEvent, "Event deleted successfully")
}
