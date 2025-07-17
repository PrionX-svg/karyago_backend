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
	GetAll() ([]response.UserEducationResponse, error)
	GetByID(id uint) (response.UserEducationResponse, error)
	GetByUUID(uuid string) (response.UserEducationResponse, error)
	GetByUserUUID(uuid string) ([]response.UserEducationResponse, error)
	Update(uuid string, req request.UserEducationReq, actorID uint) (response.UserEducationResponse, error)
	Delete(uuid string) (response.UserEducationResponse, error)
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

func (userEducationService *userEducationService) GetAll() ([]response.UserEducationResponse, error) {
	userEducations, err := userEducationService.repo.GetAll()
	if err != nil {
		return nil, err
	}

	var userEducationResponses []response.UserEducationResponse

	for _, edu := range userEducations {
		fullname := edu.User.FirstName + " " + edu.User.LastName
		userEducationResponses = append(userEducationResponses, response.UserEducationResponse{
			UUID:      edu.UUID,
			Name:      edu.Name,
			Location:  edu.Location,
			StartDate: edu.StartDate.Format("2006-01-02"),
			EndDate:   edu.EndDate.Format("2006-01-02"),
			Grade:     edu.Grade,
			User: struct {
				UUID     string `json:"uuid"`
				FullName string `json:"fullname"`
				Email    string `json:"email"`
			}{
				UUID:     edu.User.UUID,
				FullName: fullname,
				Email:    edu.User.Email,
			},
		})
	}

	return userEducationResponses, nil
}

func (s *userEducationService) GetByID(id uint) (response.UserEducationResponse, error) {
	userEducation, err := s.repo.GetByID(id)
	if err != nil {
		return response.UserEducationResponse{}, err
	}

	fullname := userEducation.User.FirstName + " " + userEducation.User.LastName

	userEducationResponse := response.UserEducationResponse{
		UUID:      userEducation.UUID,
		Name:      userEducation.Name,
		Location:  userEducation.Location,
		StartDate: userEducation.StartDate.Format("2006-01-02"),
		EndDate:   userEducation.EndDate.Format("2006-01-02"),
		Grade:     userEducation.Grade,
		User: struct {
			UUID     string `json:"uuid"`
			FullName string `json:"fullname"`
			Email    string `json:"email"`
		}{
			UUID:     userEducation.User.UUID,
			FullName: fullname,
			Email:    userEducation.User.Email,
		},
	}

	return userEducationResponse, nil
}

func (s *userEducationService) GetByUUID(uuid string) (response.UserEducationResponse, error) {
	userEducation, err := s.repo.GetByUUID(uuid)
	if err != nil {
		return response.UserEducationResponse{}, err
	}
	fullname := userEducation.User.FirstName + " " + userEducation.User.LastName
	userEducationResponse := response.UserEducationResponse{
		UUID:      userEducation.UUID,
		Name:      userEducation.Name,
		Location:  userEducation.Location,
		StartDate: userEducation.StartDate.Format("2006-01-02"),
		EndDate:   userEducation.EndDate.Format("2006-01-02"),
		Grade:     userEducation.Grade,
		User: struct {
			UUID     string `json:"uuid"`
			FullName string `json:"fullname"`
			Email    string `json:"email"`
		}{
			UUID:     userEducation.User.UUID,
			FullName: fullname,
			Email:    userEducation.User.Email,
		},
	}
	return userEducationResponse, nil
}

func (s *userEducationService) GetByUserUUID(uuid string) ([]response.UserEducationResponse, error) {
	userEducations, err := s.repo.GetByUserUUID(uuid)
	if err != nil {
		return nil, err
	}

	var userEducationResponses []response.UserEducationResponse

	for _, edu := range userEducations {
		fullname := edu.User.FirstName + " " + edu.User.LastName
		userEducationResponses = append(userEducationResponses, response.UserEducationResponse{
			UUID:      edu.UUID,
			Name:      edu.Name,
			Location:  edu.Location,
			StartDate: edu.StartDate.Format("2006-01-02"),
			EndDate:   edu.EndDate.Format("2006-01-02"),
			Grade:     edu.Grade,
			User: struct {
				UUID     string `json:"uuid"`
				FullName string `json:"fullname"`
				Email    string `json:"email"`
			}{
				UUID:     edu.User.UUID,
				FullName: fullname,
				Email:    edu.User.Email,
			},
		})
	}

	return userEducationResponses, nil
}

func (s *userEducationService) Update(uuid string, req request.UserEducationReq, actorID uint) (response.UserEducationResponse, error) {
	userEducation, err := s.repo.GetByUUID(uuid)
	if err != nil {
		return response.UserEducationResponse{}, err
	}

	userEducation.Name = req.Name
	userEducation.Location = req.Location
	userEducation.StartDate = req.StartDate
	userEducation.EndDate = req.EndDate
	userEducation.Grade = req.Grade
	userEducation.ModifyBy = actorID

	if err := s.repo.Update(userEducation); err != nil {
		return response.UserEducationResponse{}, err
	}

	fullname := userEducation.User.FirstName + " " + userEducation.User.LastName

	userEducationResponse := response.UserEducationResponse{
		UUID:      userEducation.UUID,
		Name:      userEducation.Name,
		Location:  userEducation.Location,
		StartDate: userEducation.StartDate.Format("2006-01-02"),
		EndDate:   userEducation.EndDate.Format("2006-01-02"),
		Grade:     userEducation.Grade,
		User: struct {
			UUID     string `json:"uuid"`
			FullName string `json:"fullname"`
			Email    string `json:"email"`
		}{
			UUID:     userEducation.User.UUID,
			FullName: fullname,
			Email:    userEducation.User.Email,
		},
	}

	return userEducationResponse, nil
}

func (s *userEducationService) Delete(uuid string) (response.UserEducationResponse, error) {
	userEducation, err := s.repo.GetByUUID(uuid)
	if err != nil {
		return response.UserEducationResponse{}, err
	}

	if err := s.repo.Delete(uuid); err != nil {
		return response.UserEducationResponse{}, err
	}

	fullname := userEducation.User.FirstName + " " + userEducation.User.LastName

	userEducationResponse := response.UserEducationResponse{
		UUID:      userEducation.UUID,
		Name:      userEducation.Name,
		Location:  userEducation.Location,
		StartDate: userEducation.StartDate.Format("2006-01-02"),
		EndDate:   userEducation.EndDate.Format("2006-01-02"),
		Grade:     userEducation.Grade,
		User: struct {
			UUID     string `json:"uuid"`
			FullName string `json:"fullname"`
			Email    string `json:"email"`
		}{
			UUID:     userEducation.User.UUID,
			FullName: fullname,
			Email:    userEducation.User.Email,
		},
	}

	return userEducationResponse, nil
} 