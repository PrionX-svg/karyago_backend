package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"hris_backend/pkg"
	"os"
	"strconv"
)

func RequirePermission(requiredPermission string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		roleID, ok := c.Locals("role_id").(uint)
		if !ok {
			return pkg.Error(c, fiber.StatusUnauthorized, "Unauthorized")
		}

		ownerRoleIDStr := os.Getenv("OWNER_ROLE_ID")
		ownerRoleIDUint64, err := strconv.ParseUint(ownerRoleIDStr, 10, 64)
		if err != nil {
			return pkg.Error(c, fiber.StatusInternalServerError, "Invalid OWNER_ROLE_ID in environment")
		}
		ownerRoleID := uint(ownerRoleIDUint64)

		if roleID == ownerRoleID {
			return c.Next()
		}

		permRepo := pkg.GetPermissionRepo()
		has, err := permRepo.RoleHasPermission(roleID, requiredPermission)
		if err != nil {
			return pkg.Error(c, fiber.StatusInternalServerError, "Failed to check permission")
		}
		if !has {
			return pkg.Error(c, fiber.StatusForbidden, "Access denied")
		}

		return c.Next()
	}
}
