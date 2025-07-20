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

type UserBankService interface {
	Create(request request.UserBankReq, userID uint) (response.UserBankResponse, error)
	GetAll() ([]response.UserBankResponse, error)
	GetByID(id uint) (response.UserBankResponse, error)
	GetByUUID(uuid string) (response.UserBankResponse, error)
	GetByUserUUID(uuid string) ([]response.UserBankResponse, error)
	Update(uuid string, req request.UserBankReq, actorID uint) (response.UserBankResponse, error)
	Delete(uuid string) (response.UserBankResponse, error)
}

type userBankService struct {
	repo     repositories.UserBankRepository
	userRepo repositories.UserRepository
}

func NewUserBankService(repo repositories.UserBankRepository, userRepo repositories.UserRepository) userBankService {
	return userBankService{
		repo:     repo,
		userRepo: userRepo,
	}
}

func (s *userBankService) Create(request request.UserBankReq, userID uint) (response.UserBankResponse, error) {
	if err := pkg.Validate.Struct(request); err != nil {
		return response.UserBankResponse{}, err
	}

	user, err := s.userRepo.GetByUUID(request.UserUUID)
	if err != nil {
		return response.UserBankResponse{}, fmt.Errorf("invalid user UUID")
	}

	userBank := &models.UserBank{
		UUID:      uuid.NewString(),
		UserID:    user.ID,
		Name:      request.Name,
		Number:    request.Number,
		ExpDate:   request.ExpDate,
		CreatedBy: userID,
		ModifyBy:  userID,
	}

	if err := s.repo.Create(userBank); err != nil {
		return response.UserBankResponse{}, err
	}

	userBankModel, err := s.repo.GetByUUID(userBank.UUID)
	if err != nil {
		return response.UserBankResponse{}, err
	}

	fullname := userBankModel.User.FirstName + " " + userBankModel.User.LastName

	userBankResponse := response.UserBankResponse{
		UUID:    userBankModel.UUID,
		Name:    userBankModel.Name,
		Number:  userBankModel.Number,
		ExpDate: userBankModel.ExpDate.Format("2006-01-02"),
		User: struct {
			UUID     string `json:"uuid"`
			FullName string `json:"fullname"`
			Email    string `json:"email"`
		}{
			UUID:     userBankModel.User.UUID,
			FullName: fullname,
			Email:    userBankModel.User.Email,
		},
	}

	return userBankResponse, nil
}

func (s *userBankService) GetAll() ([]response.UserBankResponse, error) {
	userBanks, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	var userBankResponses []response.UserBankResponse

	for _, bank := range userBanks {
		fullname := bank.User.FirstName + " " + bank.User.LastName
		userBankResponses = append(userBankResponses, response.UserBankResponse{
			UUID:    bank.UUID,
			Name:    bank.Name,
			Number:  bank.Number,
			ExpDate: bank.ExpDate.Format("2006-01-02"),
			User: struct {
				UUID     string `json:"uuid"`
				FullName string `json:"fullname"`
				Email    string `json:"email"`
			}{
				UUID:     bank.User.UUID,
				FullName: fullname,
				Email:    bank.User.Email,
			},
		})
	}

	return userBankResponses, nil
}

func (s *userBankService) GetByID(id uint) (response.UserBankResponse, error) {
	userBank, err := s.repo.GetByID(id)
	if err != nil {
		return response.UserBankResponse{}, err
	}

	fullname := userBank.User.FirstName + " " + userBank.User.LastName

	userBankResponse := response.UserBankResponse{
		UUID:    userBank.UUID,
		Name:    userBank.Name,
		Number:  userBank.Number,
		ExpDate: userBank.ExpDate.Format("2006-01-02"),
		User: struct {
			UUID     string `json:"uuid"`
			FullName string `json:"fullname"`
			Email    string `json:"email"`
		}{
			UUID:     userBank.User.UUID,
			FullName: fullname,
			Email:    userBank.User.Email,
		},
	}

	return userBankResponse, nil
}

func (s *userBankService) GetByUUID(uuid string) (response.UserBankResponse, error) {
	userBank, err := s.repo.GetByUUID(uuid)
	if err != nil {
		return response.UserBankResponse{}, err
	}
	fullname := userBank.User.FirstName + " " + userBank.User.LastName
	userBankResponse := response.UserBankResponse{
		UUID:    userBank.UUID,
		Name:    userBank.Name,
		Number:  userBank.Number,
		ExpDate: userBank.ExpDate.Format("2006-01-02"),
		User: struct {
			UUID     string `json:"uuid"`
			FullName string `json:"fullname"`
			Email    string `json:"email"`
		}{
			UUID:     userBank.User.UUID,
			FullName: fullname,
			Email:    userBank.User.Email,
		},
	}
	return userBankResponse, nil
}

func (s *userBankService) GetByUserUUID(uuid string) ([]response.UserBankResponse, error) {
	userBanks, err := s.repo.GetByUserUUID(uuid)
	if err != nil {
		return nil, err
	}

	var userBankResponse []response.UserBankResponse

	for _, bank := range userBanks {
		fullname := bank.User.FirstName + " " + bank.User.LastName
		userBankResponse = append(userBankResponse, response.UserBankResponse{
			UUID:    bank.UUID,
			Name:    bank.Name,
			Number:  bank.Number,
			ExpDate: bank.ExpDate.Format("2006-01-02"),
			User: struct {
				UUID     string `json:"uuid"`
				FullName string `json:"fullname"`
				Email    string `json:"email"`
			}{
				UUID:     bank.User.UUID,
				FullName: fullname,
				Email:    bank.User.Email,
			},
		})
	}

	return userBankResponse, nil
}

func (s *userBankService) Update(uuid string, req request.UserBankReq, actorID uint) (response.UserBankResponse, error) {
	userBank, err := s.repo.GetByUUID(uuid)
	if err != nil {
		return response.UserBankResponse{}, err
	}

	userBank.Name = req.Name
	userBank.Number = req.Number
	userBank.ExpDate = req.ExpDate
	userBank.ModifyBy = actorID

	if err := s.repo.Update(userBank); err != nil {
		return response.UserBankResponse{}, err
	}

	fullname := userBank.User.FirstName + " " + userBank.User.LastName

	userBankResponse := response.UserBankResponse{
		UUID:    userBank.UUID,
		Name:    userBank.Name,
		Number:  userBank.Number,
		ExpDate: userBank.ExpDate.Format("2006-01-02"),
		User: struct {
			UUID     string `json:"uuid"`
			FullName string `json:"fullname"`
			Email    string `json:"email"`
		}{
			UUID:     userBank.User.UUID,
			FullName: fullname,
			Email:    userBank.User.Email,
		},
	}

	return userBankResponse, nil
}

func (s *userBankService) Delete(uuid string) (response.UserBankResponse, error) {
	userBank, err := s.repo.GetByUUID(uuid)
	if err != nil {
		return response.UserBankResponse{}, err
	}

	if err := s.repo.Delete(uuid); err != nil {
		return response.UserBankResponse{}, err
	}

	fullname := userBank.User.FirstName + " " + userBank.User.LastName

	userBankResponse := response.UserBankResponse{
		UUID:    userBank.UUID,
		Name:    userBank.Name,
		Number:  userBank.Number,
		ExpDate: userBank.ExpDate.Format("2006-01-02"),
		User: struct {
			UUID     string `json:"uuid"`
			FullName string `json:"fullname"`
			Email    string `json:"email"`
		}{
			UUID:     userBank.User.UUID,
			FullName: fullname,
			Email:    userBank.User.Email,
		},
	}

	return userBankResponse, nil
}
