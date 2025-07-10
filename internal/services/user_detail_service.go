package services

import (
	"fmt"

	"github.com/google/uuid"
	"hris_backend/internal/models"
	"hris_backend/internal/repositories"
	"hris_backend/internal/request"
	"hris_backend/pkg"
)

type UserDetailService interface {
	Create(req request.UserDetailReq, actorID uint) (models.UserDetail, error)
	Get(userUUID string) (models.UserDetail, error)
	Update(uuid string, req request.UserDetailReq, actorID uint) (models.UserDetail, error)
	Delete(uuid string) (models.UserDetail, error)
}

type userDetailService struct {
	repo     repositories.UserDetailRepository
	userRepo repositories.UserRepository
}

func NewUserDetailService(repo repositories.UserDetailRepository, userRepo repositories.UserRepository) UserDetailService {
	return &userDetailService{
		repo:     repo,
		userRepo: userRepo,
	}
}

func (s *userDetailService) Create(req request.UserDetailReq, actorID uint) (models.UserDetail, error) {
	if err := pkg.Validate.Struct(req); err != nil {
		return models.UserDetail{}, err
	}

	user, err := s.userRepo.GetByUUID(req.UserUUID)
	if err != nil {
		return models.UserDetail{}, fmt.Errorf("invalid user UUID")
	}

	newDetail := models.UserDetail{
		UUID:      uuid.NewString(),
		UserID:    user.ID,
		TaxID:     req.TaxID,
		SocialID:  req.SocialID,
		CreatedBy: actorID,
		ModifyBy:  actorID,
	}

	if err := s.repo.Create(&newDetail); err != nil {
		return models.UserDetail{}, err
	}

	return newDetail, nil
}

func (s *userDetailService) Get(userUUID string) (models.UserDetail, error) {
	user, err := s.userRepo.GetByUUID(userUUID)
	if err != nil {
		return models.UserDetail{}, fmt.Errorf("user not found")
	}

	detail, err := s.repo.FindByUserID(user.ID)
	if err != nil {
		return models.UserDetail{}, fmt.Errorf("user detail not found")
	}

	return *detail, nil
}

func (s *userDetailService) Update(uuid string, req request.UserDetailReq, actorID uint) (models.UserDetail, error) {
	if err := pkg.Validate.Struct(req); err != nil {
		return models.UserDetail{}, err
	}

	detail, err := s.repo.FindByUUID(uuid)
	if err != nil {
		return models.UserDetail{}, fmt.Errorf("user detail not found")
	}

	user, err := s.userRepo.GetByUUID(req.UserUUID)
	if err != nil {
		return models.UserDetail{}, fmt.Errorf("invalid user UUID")
	}

	detail.TaxID = req.TaxID
	detail.SocialID = req.SocialID
	detail.UserID = user.ID
	detail.ModifyBy = actorID

	if err := s.repo.Update(detail); err != nil {
		return models.UserDetail{}, err
	}

	return *detail, nil
}

func (s *userDetailService) Delete(uuid string) (models.UserDetail, error) {
	detail, err := s.repo.FindByUUID(uuid)
	if err != nil {
		return models.UserDetail{}, err
	}

	if err := s.repo.Delete(detail.ID); err != nil {
		return models.UserDetail{}, err
	}

	return *detail, nil
}
