package handlers

import (
	"fmt"
	"hris_backend/internal/models"
	"hris_backend/internal/services"
	"hris_backend/pkg"
	"time"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler interface {
	Register(c *fiber.Ctx) error
	Verify(c *fiber.Ctx) error
	ResendVerification(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
	Logout(c *fiber.Ctx) error
	ForgotPassword(c *fiber.Ctx) error
	ResetPassword(c *fiber.Ctx) error
	ResendForgotPasswordOTP(c *fiber.Ctx) error
}

type authHandler struct {
	authService services.AuthService
}

func NewAuthHandler(authService services.AuthService) AuthHandler {
	return &authHandler{authService}
}

func (h *authHandler) Register(c *fiber.Ctx) error {
	var newUser models.User
	if err := c.BodyParser(&newUser); err != nil {
		return pkg.Error(c, fiber.StatusBadRequest, "Failed to parse user")
	}

	if err := h.authService.Register(newUser); err != nil {
		return pkg.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return pkg.Created(c, fiber.Map{}, "Registration successful. Please check your email to verify your account.")
}

func (h *authHandler) Verify(c *fiber.Ctx) error {
	type VerifyRequest struct {
		UUID string `json:"uuid"`
	}

	var req VerifyRequest
	if err := c.BodyParser(&req); err != nil || req.UUID == "" {
		return pkg.Error(c, fiber.StatusBadRequest, "Invalid request payload")
	}

	user, err := h.authService.VerifyUser(req.UUID)
	if err != nil {
		return pkg.Error(c, fiber.StatusBadRequest, err.Error())
	}

	return pkg.Success(c, fiber.Map{
		"uuid":  user.UUID,
		"name":  user.FirstName,
		"email": user.Email,
	}, "Email verified successfully")
}

func (h *authHandler) ResendVerification(c *fiber.Ctx) error {
	type Request struct {
		Email string `json:"email"`
	}

	var req Request
	if err := c.BodyParser(&req); err != nil || req.Email == "" {
		return pkg.Error(c, fiber.StatusBadRequest, "Invalid email")
	}

	if err := h.authService.ResendVerificationLink(req.Email); err != nil {
		return pkg.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return pkg.Success(c, fiber.Map{}, "Verification link resent to your email")
}

func (h *authHandler) Login(c *fiber.Ctx) error {
	type LoginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var req LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return pkg.Error(c, fiber.StatusBadRequest, "Invalid login payload")
	}

	user, err := h.authService.Login(req.Email, req.Password)
	if err != nil {
		return pkg.Error(c, fiber.StatusUnauthorized, err.Error())
	}

	role, err := h.authService.GetRoleName(user.RoleID)
	if err != nil {
		return pkg.Error(c, fiber.StatusInternalServerError, "Failed to get role")
	}

	token, err := pkg.GenerateJWT(user.UUID, user.Email, fmt.Sprintf("%d", user.RoleID))
	if err != nil {
		return pkg.Error(c, fiber.StatusInternalServerError, "Failed to generate token")
	}

	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    token,
		Expires:  time.Now().Add(1 * time.Hour),
		HTTPOnly: true,
		Secure:   true,
		SameSite: "Lax",
	})

	return pkg.Success(c, fiber.Map{
		"uuid":      user.UUID,
		"fullname":  user.FirstName + " " + user.LastName,
		"email":     user.Email,
		"role_name": role,
	}, "Login successful")
}

func (h *authHandler) Logout(c *fiber.Ctx) error {
	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    "",
		Expires:  time.Now().Add(-1 * time.Hour),
		HTTPOnly: true,
		Secure:   true,
		SameSite: "Lax",
	})
	return pkg.Success(c, fiber.Map{}, "Logged out successfully")
}

func (h *authHandler) ForgotPassword(c *fiber.Ctx) error {
	type Request struct {
		Email string `json:"email"`
	}

	var req Request
	if err := c.BodyParser(&req); err != nil || req.Email == "" {
		return pkg.Error(c, fiber.StatusBadRequest, "Invalid email")
	}

	if err := h.authService.ForgotPassword(req.Email); err != nil {
		return pkg.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return pkg.Success(c, fiber.Map{}, "Reset link sent to email")
}

func (h *authHandler) ResetPassword(c *fiber.Ctx) error {
	type Request struct {
		Email       string `json:"email"`
		OTPCode     string `json:"otp_code"`
		NewPassword string `json:"new_password"`
	}

	var req Request
	if err := c.BodyParser(&req); err != nil || req.Email == "" || req.OTPCode == "" || req.NewPassword == "" {
		return pkg.Error(c, fiber.StatusBadRequest, "Invalid request data")
	}

	if err := h.authService.ResetPassword(req.Email, req.OTPCode, req.NewPassword); err != nil {
		return pkg.Error(c, fiber.StatusBadRequest, err.Error())
	}

	return pkg.Success(c, fiber.Map{}, "Password reset successfully")
}

func (h *authHandler) ResendForgotPasswordOTP(c *fiber.Ctx) error {
	type Request struct {
		Email string `json:"email"`
	}

	var req Request
	if err := c.BodyParser(&req); err != nil || req.Email == "" {
		return pkg.Error(c, fiber.StatusBadRequest, "Invalid email")
	}

	if err := h.authService.ResendForgotPasswordOTP(req.Email); err != nil {
		return pkg.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return pkg.Success(c, fiber.Map{}, "OTP resent to your email")
}
