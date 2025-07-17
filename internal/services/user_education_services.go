package services

import (
	"hris_backend/internal/models"
	"hris_backend/internal/repositories"
	"hris_backend/internal/request"
	"hris_backend/internal/response"
	"hris_backend/pkg"

	"github.com/google/uuid"
)

type UserEducationService interface {
	Create(request request.UserEducationReq, userID uint) (response.UserEducationResponse, error)
}

type userEducationService struct {
	repo repositories.UserEducationRepository
}

func NewUserEducationService(repo repositories.UserEducationRepository) UserEducationService {
	return &userEducationService{
		repo: repo,
	}
}

func (s *userEducationService) Create(request request.UserEducationReq, userID uint) (response.UserEducationResponse, error) {
	if err := pkg.Validate.Struct(request); err != nil {
		return response.UserEducationResponse{}, err
	}

	userEducation := &models.UserEducation{
		UUID:      uuid.NewString(),
		UserUUID:  request.UserUUID,
		UserID:    userID,
		Name:      request.Name,
		Location:  request.Location,
		StartDate: request.StartDate,
		EndDate:   request.EndDate,
		Grade:     request.Grade,
		CreatedBy: userID,
		ModifyBy:  userID,
	}

	if err := s.repo.Create(userEducation); err != nil {
		return response.UserEducationResponse{}, err
	}

	userEducationModel, err := s.repo.GetByUUID(userEducation.UUID)
	if err != nil {
		return response.UserEducationResponse{}, err
	}

	fullname := userEducationModel.User.FirstName + " " + userEducationModel.User.LastName

	userEducationResponse := response.UserEducationResponse{

		UUID:      userEducationModel.UUID,
		Name:      userEducationModel.Name,
		Location:  userEducationModel.Location,
		StartDate: userEducationModel.StartDate.Format("2006-01-02"),
		EndDate:   userEducationModel.EndDate.Format("2006-01-02"),
		Grade:     userEducationModel.Grade,
		User: struct {
			UUID     string `json:"uuid"`
			FullName string `json:"fullname"`
			Email    string `json:"email"`
		}{
			UUID:     userEducationModel.User.UUID,
			FullName: fullname,
			Email:    userEducationModel.User.Email,
		},
	}

	return userEducationResponse, nil
}
