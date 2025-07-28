package handlers

import (
	"fmt"
	"hris_backend/internal/request"
	"hris_backend/internal/services"
	"hris_backend/pkg"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type EventUserHandler struct {
	service services.EventUserService
}

func NewEventUserHandler(service services.EventUserService) *EventUserHandler {
	return &EventUserHandler{
		service: service,
	}
}

func (h *EventUserHandler) Create(c *fiber.Ctx) error {
	var req request.EventUserReq

	userID, ok := c.Locals("user_id").(uint)
	if !ok {
		fmt.Println("Invalid user_id:", c.Locals("user_id"))
		return pkg.Error(c, fiber.StatusUnauthorized, "Unauthorized")
	}

	if err := c.BodyParser(&req); err != nil {
		return pkg.Error(c, fiber.StatusBadRequest, "Invalid request body")
	}

	eventUser, err := h.service.Create(req, userID)
	if err != nil {
		return pkg.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return pkg.Created(c, eventUser, "Event user created successfully")
}

func (h *EventUserHandler) GetAll(c *fiber.Ctx) error {
	eventUsers, err := h.service.GetAll()
	if err != nil {
		return pkg.Error(c, fiber.StatusInternalServerError, "Failed to retrieve event users")
	}
	return pkg.Success(c, eventUsers, "Successfully retrieved event users")
}

func (h *EventUserHandler) GetByID(c *fiber.Ctx) error {
	idStr := c.Params("id")
	if idStr == "" {
		return pkg.Error(c, fiber.StatusBadRequest, "ID is required")
	}

	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return pkg.Error(c, fiber.StatusBadRequest, "Invalid ID format")
	}

	eventUser, err := h.service.GetByID(uint(id))
	if err != nil {
		return pkg.Error(c, fiber.StatusInternalServerError, "Failed to retrieve event user")
	}

	return pkg.Success(c, eventUser, "Successfully retrieved event user")
}

func (h *EventUserHandler) GetByUUID(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	if uuid == "" {
		return pkg.Error(c, fiber.StatusBadRequest, "UUID is required")
	}

	eventUser, err := h.service.GetByUUID(uuid)
	if err != nil {
		return pkg.Error(c, fiber.StatusInternalServerError, "Failed to retrieve event user")
	}

	return pkg.Success(c, eventUser, "Successfully retrieved event user")
}

func (h *EventUserHandler) GetByUserUUID(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	if uuid == "" {
		return pkg.Error(c, fiber.StatusBadRequest, "User UUID is required")
	}

	eventUsers, err := h.service.GetByUserUUID(uuid)
	if err != nil {
		return pkg.Error(c, fiber.StatusInternalServerError, "Failed to retrieve event users")
	}

	return pkg.Success(c, eventUsers, "Successfully retrieved event users")
}

func (h *EventUserHandler) GetByEventUUID(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	if uuid == "" {
		return pkg.Error(c, fiber.StatusBadRequest, "Event UUID is required")
	}

	eventUsers, err := h.service.GetByEventUUID(uuid)
	if err != nil {
		return pkg.Error(c, fiber.StatusInternalServerError, "Failed to retrieve event users")
	}

	return pkg.Success(c, eventUsers, "Successfully retrieved event users")
}

func (h *EventUserHandler) Update(c *fiber.Ctx) error {
	var req request.EventUserReq

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

	eventUser, err := h.service.Update(uuid, req, userID)
	if err != nil {
		return pkg.Error(c, fiber.StatusInternalServerError, "Failed to update event user")
	}

	return pkg.Success(c, eventUser, "Event user updated successfully")
}

func (h *EventUserHandler) Delete(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	if uuid == "" {
		return pkg.Error(c, fiber.StatusBadRequest, "UUID is required")
	}

	eventUser, err := h.service.Delete(uuid)
	if err != nil {
		return pkg.Error(c, fiber.StatusInternalServerError, "Failed to delete event user")
	}

	return pkg.Success(c, eventUser, "Event user deleted successfully")
}
