package handlers

import (
	"github.com/gofiber/fiber/v2"
	"hris_backend/internal/services"
	"hris_backend/pkg"
)

type RolePermissionHandler struct {
	service services.RolePermissionService
}

func NewRolePermissionHandler(svc services.RolePermissionService) *RolePermissionHandler {
	return &RolePermissionHandler{svc}
}

func (h *RolePermissionHandler) AssignAndRemovePermissions(c *fiber.Ctx) error {
	type Item struct {
		RoleUUID       string `json:"role_uuid"`
		PermissionUUID string `json:"permission_uuid"`
	}
	type Request struct {
		Assign []Item `json:"assign"`
		Remove []Item `json:"remove"`
	}

	var body Request
	if err := c.BodyParser(&body); err != nil {
		return pkg.Error(c, fiber.StatusBadRequest, "Invalid request body")
	}

	actorID, ok := c.Locals("user_id").(uint)
	if !ok {
		return pkg.Error(c, fiber.StatusUnauthorized, "Unauthorized")
	}

	var assignErrors, removeErrors []string

	for _, item := range body.Assign {
		if item.RoleUUID == "" || item.PermissionUUID == "" {
			assignErrors = append(assignErrors, "Invalid assign item with empty UUIDs")
			continue
		}
		if err := h.service.AssignPermission(item.RoleUUID, item.PermissionUUID, actorID); err != nil {
			assignErrors = append(assignErrors, err.Error())
		}
	}

	for _, item := range body.Remove {
		if item.RoleUUID == "" || item.PermissionUUID == "" {
			removeErrors = append(removeErrors, "Invalid remove item with empty UUIDs")
			continue
		}
		if err := h.service.RemovePermission(item.RoleUUID, item.PermissionUUID); err != nil {
			removeErrors = append(removeErrors, err.Error())
		}
	}

	result := fiber.Map{
		"assigned": len(body.Assign) - len(assignErrors),
		"removed":  len(body.Remove) - len(removeErrors),
	}

	if len(assignErrors) > 0 || len(removeErrors) > 0 {
		result["assign_errors"] = assignErrors
		result["remove_errors"] = removeErrors
		return pkg.Success(c, result, "Processed with some errors")
	}

	return pkg.Success(c, result, "Permissions processed successfully")
}

func (h *RolePermissionHandler) ListPermissionsByRole(c *fiber.Ctx) error {
	roleUUID := c.Params("role_uuid")
	if roleUUID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "role_uuid is required"})
	}
	list, err := h.service.GetPermissionsByRole(roleUUID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(list)
}
