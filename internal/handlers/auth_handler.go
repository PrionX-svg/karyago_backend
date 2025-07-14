package handlers

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"hris_backend/internal/repositories"
	"hris_backend/internal/request"
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
	VerifyOTP(c *fiber.Ctx) error
	ResetPassword(c *fiber.Ctx) error
	ResendForgotPasswordOTP(c *fiber.Ctx) error
}

type authHandler struct {
	authService services.AuthService
	companyRepo repositories.CompanyRepositories
}

func NewAuthHandler(authService services.AuthService, companyRepo repositories.CompanyRepositories) AuthHandler {
	return &authHandler{authService, companyRepo}
}

func (h *authHandler) Register(c *fiber.Ctx) error {
	var userRequest request.UserRequest
	if err := c.BodyParser(&userRequest); err != nil {
		return pkg.Error(c, fiber.StatusBadRequest, "Failed to parse user")
	}

	if err := h.authService.Register(userRequest); err != nil {
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
		UUID string `json:"uuid"`
	}

	var req Request
	if err := c.BodyParser(&req); err != nil || req.UUID == "" {
		return pkg.Error(c, fiber.StatusBadRequest, "Invalid UUID")
	}

	if err := h.authService.ResendVerificationLink(req.UUID); err != nil {
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

	token, err := pkg.GenerateJWT(user.ID, user.UUID, user.Email, fmt.Sprintf("%d", user.RoleID))
	if err != nil {
		return pkg.Error(c, fiber.StatusInternalServerError, "Failed to generate token")
	}

	isOnboarding := true
	_, err = h.companyRepo.GetByUserID(user.ID)
	if err == nil {
		isOnboarding = false
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return pkg.Error(c, fiber.StatusInternalServerError, "Failed to check company data")
	}

	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    token,
		Expires:  time.Now().Add(1 * time.Hour),
		HTTPOnly: true,
		Secure:   true,
		SameSite: "None",
	})

	return pkg.Success(c, fiber.Map{
		"uuid":          user.UUID,
		"fullname":      user.FirstName + " " + user.LastName,
		"email":         user.Email,
		"role_name":     role,
		"is_onboarding": isOnboarding,
	}, "Login successful")
}

func (h *authHandler) Logout(c *fiber.Ctx) error {
	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    "",
		Expires:  time.Now().Add(-1 * time.Hour),
		HTTPOnly: true,
		Secure:   true,
		SameSite: "None",
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

func (h *authHandler) VerifyOTP(c *fiber.Ctx) error {
	type Request struct {
		Email   string `json:"email"`
		OTPCode string `json:"otp_code"`
	}

	var req Request
	if err := c.BodyParser(&req); err != nil || req.Email == "" || req.OTPCode == "" {
		return pkg.Error(c, fiber.StatusBadRequest, "Invalid request data")
	}

	if err := h.authService.VerifyOTP(req.Email, req.OTPCode); err != nil {
		return pkg.Error(c, fiber.StatusBadRequest, err.Error())
	}

	return pkg.Success(c, fiber.Map{}, "OTP verified successfully")
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
