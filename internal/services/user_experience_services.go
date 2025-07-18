package services

import (
	"fmt"
	"hris_backend/internal/models"
	"hris_backend/internal/repositories"
	"hris_backend/internal/request"
	"hris_backend/internal/response"
	"hris_backend/pkg"

	"github.com/google/uuid"
)

type UserExperienceService interface {
	Create(request request.UserExperienceReq, userID uint) (response.UserExperienceResponse, error)
	GetAll() ([]response.UserExperienceResponse, error)
	GetByID(id uint) (response.UserExperienceResponse, error)
	GetByUUID(uuid string) (response.UserExperienceResponse, error)
	GetByUserUUID(uuid string) ([]response.UserExperienceResponse, error)
	Update(uuid string, req request.UserExperienceReq, actorID uint) (response.UserExperienceResponse, error)
	Delete(uuid string) (response.UserExperienceResponse, error)
}

type userExperienceService struct {
	repo repositories.UserExperienceRepository
	userRepo repositories.UserRepository
}

func NewUserExperienceService(repo repositories.UserExperienceRepository, userRepo repositories.UserRepository) userExperienceService {
	return userExperienceService{
		repo: repo,
		userRepo: userRepo,
	}
}

func (s *userExperienceService) Create(request request.UserExperienceReq, userID uint) (response.UserExperienceResponse, error) {
	if err := pkg.Validate.Struct(request); err != nil {
		return response.UserExperienceResponse{}, err
	}

	user, err := s.userRepo.GetByUUID(request.UserUUID)
	if err != nil {
		return response.UserExperienceResponse{}, fmt.Errorf("invalid user UUID")
	}

	userExperience := &models.UserExperience{
		UUID:      uuid.NewString(),
		UserID:    user.ID,
		Name:      request.Name,
		Description:  request.Description,
		StartDate: request.StartDate,
		EndDate:   request.EndDate,
		CreatedBy: userID,
		ModifyBy:  userID,
	}

	if err := s.repo.Create(userExperience); err != nil {
		return response.UserExperienceResponse{}, err
	}

	userExperienceModel, err := s.repo.GetByUUID(userExperience.UUID)
	if err != nil {
		return response.UserExperienceResponse{}, err
	}

	fullname := userExperienceModel.User.FirstName + " " + userExperienceModel.User.LastName

	userExperienceResponse := response.UserExperienceResponse{

		UUID:      userExperienceModel.UUID,
		Name:      userExperienceModel.Name,
		Description:  userExperienceModel.Description,
		StartDate: userExperienceModel.StartDate.Format("2006-01-02"),
		EndDate:   userExperienceModel.EndDate.Format("2006-01-02"),
		User: struct {
			UUID     string `json:"uuid"`
			FullName string `json:"fullname"`
			Email    string `json:"email"`
		}{
			UUID:     userExperienceModel.User.UUID,
			FullName: fullname,
			Email:    userExperienceModel.User.Email,
		},
	}

	return userExperienceResponse, nil
}

func (userExperienceService *userExperienceService) GetAll() ([]response.UserExperienceResponse, error) {
	userExperiences, err := userExperienceService.repo.GetAll()
	if err != nil {
		return nil, err
	}

	var userExperienceResponses []response.UserExperienceResponse

	for _, exp := range userExperiences {
		fullname := exp.User.FirstName + " " + exp.User.LastName
		userExperienceResponses = append(userExperienceResponses, response.UserExperienceResponse{
			UUID:      exp.UUID,
			Name:      exp.Name,
			Description:  exp.Description,
			StartDate: exp.StartDate.Format("2006-01-02"),
			EndDate:   exp.EndDate.Format("2006-01-02"),
			User: struct {
				UUID     string `json:"uuid"`
				FullName string `json:"fullname"`
				Email    string `json:"email"`
			}{
				UUID:     exp.User.UUID,
				FullName: fullname,
				Email:    exp.User.Email,
			},
		})
	}

	return userExperienceResponses, nil
}

func (s *userExperienceService) GetByID(id uint) (response.UserExperienceResponse, error) {
	userExperience, err := s.repo.GetByID(id)
	if err != nil {
		return response.UserExperienceResponse{}, err
	}

	fullname := userExperience.User.FirstName + " " + userExperience.User.LastName

	userExperienceResponse := response.UserExperienceResponse{
		UUID:      userExperience.UUID,
		Name:      userExperience.Name,
		Description:  userExperience.Description,
		StartDate: userExperience.StartDate.Format("2006-01-02"),
		EndDate:   userExperience.EndDate.Format("2006-01-02"),
		User: struct {
			UUID     string `json:"uuid"`
			FullName string `json:"fullname"`
			Email    string `json:"email"`
		}{
			UUID:     userExperience.User.UUID,
			FullName: fullname,
			Email:    userExperience.User.Email,
		},
	}

	return userExperienceResponse, nil
}

func (s *userExperienceService) GetByUUID(uuid string) (response.UserExperienceResponse, error) {
	userExperience, err := s.repo.GetByUUID(uuid)
	if err != nil {
		return response.UserExperienceResponse{}, err
	}
	fullname := userExperience.User.FirstName + " " + userExperience.User.LastName
	userExperienceResponse := response.UserExperienceResponse{
		UUID:      userExperience.UUID,
		Name:      userExperience.Name,
		Description:  userExperience.Description,
		StartDate: userExperience.StartDate.Format("2006-01-02"),
		EndDate:   userExperience.EndDate.Format("2006-01-02"),
		User: struct {
			UUID     string `json:"uuid"`
			FullName string `json:"fullname"`
			Email    string `json:"email"`
		}{
			UUID:     userExperience.User.UUID,
			FullName: fullname,
			Email:    userExperience.User.Email,
		},
	}
	return userExperienceResponse, nil
}

func (s *userExperienceService) GetByUserUUID(uuid string) ([]response.UserExperienceResponse, error) {
	userEducations, err := s.repo.GetByUserUUID(uuid)
	if err != nil {
		return nil, err
	}

	var userExperienceResponse []response.UserExperienceResponse

	for _, exp := range userEducations {
		fullname := exp.User.FirstName + " " + exp.User.LastName
		userExperienceResponse = append(userExperienceResponse, response.UserExperienceResponse{
			UUID:      exp.UUID,
			Name:      exp.Name,
			Description:  exp.Description,
			StartDate: exp.StartDate.Format("2006-01-02"),
			EndDate:   exp.EndDate.Format("2006-01-02"),
			User: struct {
				UUID     string `json:"uuid"`
				FullName string `json:"fullname"`
				Email    string `json:"email"`
			}{
				UUID:     exp.User.UUID,
				FullName: fullname,
				Email:    exp.User.Email,
			},
		})
	}

	return userExperienceResponse, nil
}

func (s *userExperienceService) Update(uuid string, req request.UserExperienceReq, actorID uint) (response.UserExperienceResponse, error) {
	userExperience, err := s.repo.GetByUUID(uuid)
	if err != nil {
		return response.UserExperienceResponse{}, err
	}

	userExperience.Name = req.Name
	userExperience.Description = req.Description
	userExperience.StartDate = req.StartDate
	userExperience.EndDate = req.EndDate
	userExperience.ModifyBy = actorID

	if err := s.repo.Update(userExperience); err != nil {
		return response.UserExperienceResponse{}, err
	}

	fullname := userExperience.User.FirstName + " " + userExperience.User.LastName

	userExperienceResponse := response.UserExperienceResponse{
		UUID:      userExperience.UUID,
		Name:      userExperience.Name,
		Description:  userExperience.Description,
		StartDate: userExperience.StartDate.Format("2006-01-02"),
		EndDate:   userExperience.EndDate.Format("2006-01-02"),

		User: struct {
			UUID     string `json:"uuid"`
			FullName string `json:"fullname"`
			Email    string `json:"email"`
		}{
			UUID:     userExperience.User.UUID,
			FullName: fullname,
			Email:    userExperience.User.Email,
		},
	}

	return userExperienceResponse, nil
}

func (s *userExperienceService) Delete(uuid string) (response.UserExperienceResponse, error) {
	userExperience, err := s.repo.GetByUUID(uuid)
	if err != nil {
		return response.UserExperienceResponse{}, err
	}

	if err := s.repo.Delete(uuid); err != nil {
		return response.UserExperienceResponse{}, err
	}

	fullname := userExperience.User.FirstName + " " + userExperience.User.LastName

	userExperienceResponse := response.UserExperienceResponse{
		UUID:      userExperience.UUID,
		Name:      userExperience.Name,
		Description:  userExperience.Description,
		StartDate: userExperience.StartDate.Format("2006-01-02"),
		EndDate:   userExperience.EndDate.Format("2006-01-02"),

		User: struct {
			UUID     string `json:"uuid"`
			FullName string `json:"fullname"`
			Email    string `json:"email"`
		}{
			UUID:     userExperience.User.UUID,
			FullName: fullname,
			Email:    userExperience.User.Email,
		},
	}

	return userExperienceResponse, nil
} 