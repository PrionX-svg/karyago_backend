package services

import (
	"fmt"
	"hris_backend/internal/models"
	"hris_backend/internal/repositories"
	"hris_backend/pkg"
	"os"
	"time"

	"github.com/google/uuid"
)

type AuthService interface {
}

type authService struct {
	authRepo repositories.AuthRepository
}

func NewAuthService(authRepo repositories.AuthRepository) AuthService {
	return &authService{authRepo}
}

func (s *authService) Register(newUser models.User) error {
	if err := s.authRepo.Register(newUser); err != nil {
		return err
	}

	code := pkg.GenerateOTPCode(6)

	email := newUser.Email

	otp := &models.OTP{
		UUID:      uuid.NewString(),
		Target:    email,
		Code:      code,
		Purpose:   "register",
		IsUsed:    false,
		ExpiresAt: time.Now().Add(5 * time.Minute),
	}

	subject := "Email Verification"

	frontendURL := os.Getenv("FRONTEND_URL")
	verificationLink := fmt.Sprintf("%s/en/activation?token=%s", frontendURL,otp.UUID)

	body := fmt.Sprintf(`
		<html>
			<body>
				<p>Hi %s,</p>
				<p>Thank you for registering with us!</p>
				<p>Please verify your email through the link below</p>
				<a href="%s">Verify your Email</a>
				<p>This link will expire in 5 minutes.</p>
				<p>If you didn't request this, please ignore this email.</p>
				<br/>
				<p>Regards,<br/>The Team</p>
			</body>
		</html>
	`, newUser.FirstName, verificationLink)

	if err := pkg.SendEmail(email, subject, body); err != nil {
		return err
	}

	return nil

}
