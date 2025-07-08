package handlers

import (
	"github.com/gofiber/fiber/v2"
	"hris_backend/internal/services"
	"hris_backend/pkg"
)

type PermissionHandler struct {
	service services.PermissionService
}

func NewPermissionHandler(service services.PermissionService) *PermissionHandler {
	return &PermissionHandler{service}
}

func (h *PermissionHandler) GetByUUID(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	if uuid == "" {
		return pkg.Error(c, fiber.StatusBadRequest, "UUID is required")
	}

	perm, err := h.service.GetByUUID(uuid)
	if err != nil {
		return pkg.Error(c, fiber.StatusNotFound, "Permission not found")
	}

	return pkg.Success(c, perm)
}

func (h *PermissionHandler) GetAll(c *fiber.Ctx) error {
	list, err := h.service.GetAll()
	if err != nil {
		return pkg.Error(c, fiber.StatusInternalServerError, err.Error())
	}
	return pkg.Success(c, list)
}
