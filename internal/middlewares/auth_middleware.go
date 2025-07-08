package middlewares

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"hris_backend/pkg"
)

func JWTMiddleware(c *fiber.Ctx) error {
	tokenStr := c.Cookies("token")
	if tokenStr == "" {
		fmt.Println("JWTMiddleware: Token missing")
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Missing token"})
	}

	claims, err := pkg.ParseJWT(tokenStr)
	if err != nil {
		fmt.Println("JWTMiddleware: Failed to parse token:", err)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid or expired token"})
	}

	// Debug print claims
	fmt.Printf("JWTMiddleware: Claims => user_id: %v, user_uuid: %v, email: %v, role: %v\n",
		claims.UserID, claims.UserUUID, claims.Email, claims.Role)

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
		fmt.Println("JWTMiddleware: Time remaining:", timeRemaining)

		if timeRemaining < 15*time.Minute {
			fmt.Println("JWTMiddleware: Token expiring soon, refreshing...")

			newToken, err := pkg.GenerateJWT(claims.UserID, claims.UserUUID, claims.Email, claims.Role)
			if err != nil {
				fmt.Println("JWTMiddleware: Failed to generate new token:", err)
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to refresh token"})
			}

			c.Cookie(&fiber.Cookie{
				Name:     "token",
				Value:    newToken,
				Expires:  time.Now().Add(1 * time.Hour),
				HTTPOnly: true,
				Secure:   os.Getenv("ENV") == "production",
				SameSite: fiber.CookieSameSiteLaxMode,
			})
			fmt.Println("JWTMiddleware: Token refreshed")
		}
	} else {
		fmt.Println("JWTMiddleware: claims.ExpiresAt is nil")
	}

	return c.Next()
}
