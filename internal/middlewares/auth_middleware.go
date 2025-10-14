package middlewares

import (
	"os"
	"strconv"
	"time"

	"hris_backend/pkg"

	"github.com/gofiber/fiber/v2"
)

func JWTMiddleware(c *fiber.Ctx) error {
	tokenStr := c.Cookies("token")
	if tokenStr == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Missing token"})
	}

	claims, err := pkg.ParseJWT(tokenStr)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid or expired token"})
	}

	roleID, err := strconv.ParseUint(claims.Role, 10, 64)
	if err != nil {
		return pkg.Error(c, fiber.StatusUnauthorized, "Invalid role ID format")
	}

	c.Locals("user_id", claims.UserID)
	c.Locals("user_uuid", claims.UserUUID)
	c.Locals("email", claims.Email)
	c.Locals("role", claims.Role)
	c.Locals("role_id", uint(roleID))
	
	if claims.ExpiresAt != nil {
		timeRemaining := time.Until(claims.ExpiresAt.Time)
		if timeRemaining < 30*time.Minute {
			newToken, err := pkg.GenerateJWT(claims.UserID, claims.UserUUID, claims.Email, claims.Role)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to refresh token"})
			}

			c.Cookie(&fiber.Cookie{
				Name:     "token",
				Value:    newToken,
				HTTPOnly: true,
				Secure:   os.Getenv("APP_ENV") == "production",
				Path:     "/",
				Expires:  time.Now().Add(1 * time.Hour),
			})
		}
	}

	return c.Next()
}
