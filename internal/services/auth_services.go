package services

import (
	"fmt"
	"hris_backend/internal/models"
	"hris_backend/internal/repositories"
	"hris_backend/internal/request"
	"hris_backend/pkg"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
)

type AuthService interface {
	Register(request request.UserRequest) error
	VerifyUser(otpUUID string) (*models.User, error)
	ResendVerificationLink(email string) error
	Login(email, password string) (*models.User, error)
	ForgotPassword(email string) error
	VerifyOTP(email, code string) error
	ResetPassword(email, code, newPassword string) error
	ResendForgotPasswordOTP(email string) error
	GetRoleName(roleID uint) (string, error)
}

type authService struct {
	authRepo repositories.AuthRepository
	roleRepo repositories.RoleRepositories
	otpRepo  repositories.OTPRepositories
}

func NewAuthService(authRepo repositories.AuthRepository, roleRepo repositories.RoleRepositories, otpRepo repositories.OTPRepositories) AuthService {
	return &authService{authRepo, roleRepo, otpRepo}
}

func (s *authService) Register(request request.UserRequest) error {
	if err := pkg.Validate.Struct(request); err != nil {
		return err
	}

	role, err := s.roleRepo.FindByName("owner")
	if err != nil {
		return fmt.Errorf("failed to find role: %w", err)
	}

	password, err := pkg.HashPassword(request.Password)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}

	newUser := models.User{
		UUID:      uuid.NewString(),
		RoleID:    role.ID,
		CompanyID: request.CompanyID,
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Phone:     request.Phone,
		Email:     request.Email,
		Password:  password,
		Timezone:  request.Timezone,
	}

	if err := s.authRepo.Register(&newUser); err != nil {
		return fmt.Errorf("failed to register user: %w", err)
	}

	otp := &models.OTP{
		UUID:      uuid.NewString(),
		Target:    newUser.Email,
		Code:      pkg.GenerateOTPCode(6),
		Purpose:   "register",
		IsUsed:    false,
		ExpiresAt: time.Now().Add(5 * time.Minute),
	}

	if err := s.otpRepo.Create(otp); err != nil {
		return fmt.Errorf("failed to create OTP: %w", err)
	}

	go func(email, name, token string) {
		verificationLink := fmt.Sprintf("%s/en/activation?token=%s", os.Getenv("FRONTEND_URL"), token)
		body := fmt.Sprintf(`
			<html>
				<body>
					<p>Hi %s,</p>
					<p>Thank you for registering with us!</p>
					<p>Please verify your email through the link below:</p>
					<a href="%s">Verify your Email</a>
					<p>This link will expire in 5 minutes.</p>
					<p>If you didn't request this, please ignore this email.</p>
					<br/>
					<p>Regards,<br/>The Team</p>
				</body>
			</html>
		`, name, verificationLink)

		err = pkg.SendEmail(email, "Email Verification", body)
		if err != nil {
			log.Printf("failed to send email: %v", err)
		}
	}(newUser.Email, newUser.FirstName, otp.UUID)

	return nil
}

func (s *authService) VerifyUser(otpUUID string) (*models.User, error) {
	otp, err := s.otpRepo.FindByUUID(otpUUID)
	if err != nil {
		return nil, fmt.Errorf("OTP not found: %w", err)
	}

	if otp.IsUsed {
		return nil, fmt.Errorf("OTP has already been used")
	}
	if time.Now().After(otp.ExpiresAt) {
		return nil, fmt.Errorf("OTP has expired")
	}

	user, err := s.authRepo.FindByEmail(otp.Target)
	if err != nil {
		return nil, fmt.Errorf("user not found for OTP target: %w", err)
	}

	if err := s.authRepo.VerifyUser(user); err != nil {
		return nil, fmt.Errorf("failed to verify user: %w", err)
	}
	if err := s.otpRepo.MarkUsed(otpUUID); err != nil {
		return nil, fmt.Errorf("failed to mark OTP as used: %w", err)
	}

	return user, nil
}

func (s *authService) ResendVerificationLink(oldUUID string) error {
	otp, err := s.otpRepo.FindByUUID(oldUUID)
	if err != nil {
		return fmt.Errorf("OTP not found")
	}

	user, err := s.authRepo.FindByEmail(otp.Target)
	if err != nil {
		return fmt.Errorf("user not found for this OTP")
	}

	if user.IsVerified {
		return fmt.Errorf("user already verified")
	}

	newOTP := &models.OTP{
		UUID:      uuid.NewString(),
		Target:    user.Email,
		Code:      pkg.GenerateOTPCode(6),
		Purpose:   "register",
		IsUsed:    false,
		ExpiresAt: time.Now().Add(5 * time.Minute),
	}

	if err := s.otpRepo.Create(newOTP); err != nil {
		return fmt.Errorf("failed to generate new OTP: %w", err)
	}

	link := fmt.Sprintf("%s/en/activation?token=%s", os.Getenv("FRONTEND_URL"), newOTP.UUID)
	body := fmt.Sprintf(`
		<p>Hi %s,</p>
		<p>You requested a new verification link.</p>
		<p>Please verify your email using the link below:</p>
		<a href="%s">Verify your Email</a>
		<p>This link will expire in 5 minutes.</p>`,
		user.FirstName, link)

	go func() {
		if err := pkg.SendEmail(user.Email, "Email Verification", body); err != nil {
			log.Printf("failed to send email: %v", err)
		}
	}()

	return nil
}

func (s *authService) Login(email, password string) (*models.User, error) {
	user, err := s.authRepo.CheckLogin(email)
	if err != nil {
		return nil, fmt.Errorf("email or password is incorrect")
	}

	if !user.IsVerified {
		return nil, fmt.Errorf("account is not verified")
	}

	if !pkg.CheckPasswordHash(password, user.Password) {
		return nil, fmt.Errorf("email or password is incorrect")
	}

	return user, nil
}

func (s *authService) ForgotPassword(email string) error {
	user, err := s.authRepo.FindByEmail(email)
	if err != nil {
		return fmt.Errorf("email not found")
	}

	otp := &models.OTP{
		UUID:      uuid.NewString(),
		Target:    user.Email,
		Code:      pkg.GenerateOTPCode(6),
		Purpose:   "forgot_password",
		IsUsed:    false,
		ExpiresAt: time.Now().Add(10 * time.Minute),
	}
	if err := s.otpRepo.Create(otp); err != nil {
		return err
	}

	body := fmt.Sprintf(`
		<p>Hi %s,</p>
		<p>Your OTP to reset your password is:</p>
		<h2>%s</h2>
		<p>This code will expire in 10 minutes.</p>
		<p>If you didn’t request a password reset, you can ignore this email.</p>
	`, user.FirstName, otp.Code)

	go func() {
		err := pkg.SendEmail(user.Email, "Reset Password OTP", body)
		if err != nil {
			log.Printf("failed to send email: %v", err)
		}
	}()

	return nil
}

func (s *authService) VerifyOTP(email, code string) error {
	otp, err := s.otpRepo.FindByCodeAndTarget(code, email)
	if err != nil || otp.IsUsed || otp.ExpiresAt.Before(time.Now()) {
		return fmt.Errorf("invalid or expired OTP")
	}

	return nil
}

func (s *authService) ResetPassword(email, code, newPassword string) error {
	otp, err := s.otpRepo.FindByCodeAndTarget(code, email)
	if err != nil || otp.IsUsed || otp.ExpiresAt.Before(time.Now()) {
		return fmt.Errorf("invalid or expired OTP")
	}

	user, err := s.authRepo.FindByEmail(email)
	if err != nil {
		return err
	}

	hashed, err := pkg.HashPassword(newPassword)
	if err != nil {
		return err
	}

	if err := s.authRepo.UpdatePasswordByEmail(user.Email, hashed); err != nil {
		return err
	}

	return s.otpRepo.MarkUsed(otp.UUID)
}

func (s *authService) ResendForgotPasswordOTP(email string) error {
	user, err := s.authRepo.FindByEmail(email)
	if err != nil {
		return fmt.Errorf("email not found")
	}

	otp, err := s.otpRepo.FindValidOTPByTargetAndPurpose(user.Email, "forgot_password")
	if err != nil {
		return fmt.Errorf("no valid OTP found")
	}

	body := fmt.Sprintf(`
		<p>Hi %s,</p>
		<p>Your OTP code to reset your password is:</p>
		<h3>%s</h3>
		<p>This code will expire at %s.</p>`,
		user.FirstName, otp.Code, otp.ExpiresAt.Format("15:04:05 MST"))

	go func() {
		err := pkg.SendEmail(user.Email, "Your Password Reset OTP", body)
		if err != nil {
			log.Printf("failed to send email: %v", err)
		}
	}()

	return nil
}

func (s *authService) GetRoleName(roleID uint) (string, error) {
	role, err := s.roleRepo.FindByID(roleID)
	if err != nil {
		return "", err
	}
	return role.Name, nil
}
