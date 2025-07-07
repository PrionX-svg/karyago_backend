package handlers

import (
	"hris_backend/internal/models"
	"hris_backend/internal/services"
	"hris_backend/pkg"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler interface {

}

type authHandler struct {
	authService services.AuthService
}

func NewAuthHandler(authService services.AuthService) AuthHandler{
	return &authHandler{authService}
}

func (h *authHandler) Register(c *fiber.Ctx) error {
	var newUser models.User
	if err := c.BodyParser(&newUser); err != nil{
		return pkg.Error(c, 400, "Failed to parse user")
	}

	if err := h.authService.Register(newUser); err != nil {
		return err
	}

	return pkg.Created(c, newUser)

}