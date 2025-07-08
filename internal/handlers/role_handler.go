package handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"hris_backend/internal/models"
	"hris_backend/internal/services"
	"hris_backend/pkg"
)

type RoleHandler struct {
	service services.RoleService
}

func NewRoleHandler(service services.RoleService) *RoleHandler {
	return &RoleHandler{service}
}

func (h *RoleHandler) GetByUUID(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	if uuid == "" {
		return pkg.Error(c, fiber.StatusBadRequest, "UUID is required")
	}

	role, err := h.service.GetByUUID(uuid)
	if err != nil {
		return pkg.Error(c, fiber.StatusNotFound, "Role not found")
	}

	return pkg.Success(c, role)
}

func (h *RoleHandler) GetByName(c *fiber.Ctx) error {
	name := c.Query("name")
	if name == "" {
		return pkg.Error(c, fiber.StatusBadRequest, "Name query param is required")
	}

	role, err := h.service.GetByName(name)
	if err != nil {
		return pkg.Error(c, fiber.StatusNotFound, "Role not found")
	}

	return pkg.Success(c, role)
}

func (h *RoleHandler) GetAll(c *fiber.Ctx) error {
	limit := c.QueryInt("limit", 10)
	page := c.QueryInt("page", 1)
	search := c.Query("search", "")
	sort := c.Query("sort", "id desc")

	if limit <= 0 {
		limit = 10
	}
	offset := (page - 1) * limit

	roles, total, err := h.service.GetAll(limit, offset, search, sort)
	if err != nil {
		return pkg.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return pkg.Success(c, fiber.Map{
		"data":       roles,
		"total":      total,
		"page":       page,
		"limit":      limit,
		"totalPages": (total + int64(limit) - 1) / int64(limit),
	})
}

func (h *RoleHandler) Create(c *fiber.Ctx) error {
	var role models.Role
	if err := c.BodyParser(&role); err != nil || role.Name == "" {
		return pkg.Error(c, fiber.StatusBadRequest, "Invalid request body")
	}

	actorID, ok := c.Locals("user_id").(uint)
	if !ok {
		fmt.Println("Invalid user_id:", c.Locals("user_id"))
		return pkg.Error(c, fiber.StatusUnauthorized, "Unauthorized")
	}

	if err := h.service.Create(&role, actorID); err != nil {
		return pkg.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return pkg.Created(c, role, "Role created successfully")
}

func (h *RoleHandler) Update(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	if uuid == "" {
		return pkg.Error(c, fiber.StatusBadRequest, "UUID is required")
	}

	existingRole, err := h.service.GetByUUID(uuid)
	if err != nil {
		return pkg.Error(c, fiber.StatusNotFound, "Role not found")
	}

	var input map[string]interface{}
	if err := c.BodyParser(&input); err != nil {
		return pkg.Error(c, fiber.StatusBadRequest, "Invalid JSON body")
	}

	actorID, ok := c.Locals("user_id").(uint)
	if !ok {
		return pkg.Error(c, fiber.StatusUnauthorized, "Unauthorized")
	}

	if name, ok := input["name"].(string); ok && name != "" {
		existingRole.Name = name
	}
	existingRole.ModifyBy = actorID

	if err := h.service.Update(existingRole, actorID); err != nil {
		return pkg.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return pkg.Success(c, existingRole, "Role updated successfully")
}

func (h *RoleHandler) Delete(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	if uuid == "" {
		return pkg.Error(c, fiber.StatusBadRequest, "UUID is required")
	}

	if err := h.service.Delete(uuid); err != nil {
		return pkg.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return pkg.Success(c, nil, "Role deleted successfully")
}
