package auth

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"hris_backend/internal/handlers"
	"hris_backend/internal/repositories"
	"hris_backend/internal/services"
)

func SetupAuthRoutes(router fiber.Router, db *gorm.DB) {
	otpRepo := repositories.NewOTPRepositories(db)
	roleRepo := repositories.NewRoleRepositories(db)
	authRepo := repositories.NewAuthRepository(db)
	companyRepo := repositories.NewCompanyRepository(db)
	employeeRepo := repositories.NewEmployeeRepository(db)

	authService := services.NewAuthService(authRepo, roleRepo, otpRepo, employeeRepo)
	authHandler := handlers.NewAuthHandler(authService, companyRepo, employeeRepo)

	auth := router.Group("/auth")

	auth.Post("/register", authHandler.Register)
	auth.Post("/verify", authHandler.Verify)
	auth.Post("/resend-verification", authHandler.ResendVerification)
	auth.Post("/login", authHandler.Login)
	auth.Post("/logout", authHandler.Logout)
	auth.Post("/forgot-password", authHandler.ForgotPassword)
	auth.Post("/forgot-password/verify", authHandler.VerifyOTP)
	auth.Post("/forgot-password/reset", authHandler.ResetPassword)
	auth.Post("/forgot-password/resend", authHandler.ResendForgotPasswordOTP)
}
