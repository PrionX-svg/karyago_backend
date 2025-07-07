package services

import "hris_backend/internal/repositories"

type AuthService interface {

}

type authService struct {
	authRepo repositories.AuthRepository
}

func NewAuthService(authRepo repositories.AuthRepository) AuthService {
	return &authService{authRepo}
}

func (s *authService) RegisterWithOTP() error {

}