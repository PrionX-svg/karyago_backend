package services

import (
	"fmt"
	"github.com/google/uuid"
	"hris_backend/internal/models"
	"hris_backend/internal/repositories"
	"hris_backend/internal/request"
	"hris_backend/internal/response"
	"hris_backend/pkg"
)

type UserFamilyService interface {
	Create(req request.UserFamilyRequest, actorID uint) (response.UserFamilyResponse, error)
	GetAll() ([]response.UserFamilyResponse, error)
	GetByID(id uint) (response.UserFamilyResponse, error)
	GetByUUID(uuid string) (response.UserFamilyResponse, error)
	GetByUserUUID(uuid string) ([]response.UserFamilyResponse, error)
	Update(uuid string, req request.UserFamilyRequest, actorID uint) (response.UserFamilyResponse, error)
	Delete(uuid string) (response.UserFamilyResponse, error)
}

type userFamilyService struct {
	repo     repositories.UserFamilyRepository
	userRepo repositories.UserRepository
}

func NewUserFamilyService(repo repositories.UserFamilyRepository, userRepo repositories.UserRepository) UserFamilyService {
	return &userFamilyService{
		repo:     repo,
		userRepo: userRepo,
	}
}

func (s *userFamilyService) Create(req request.UserFamilyRequest, actorID uint) (response.UserFamilyResponse, error) {
	if err := pkg.Validate.Struct(req); err != nil {
		return response.UserFamilyResponse{}, err
	}

	user, err := s.userRepo.GetByUUID(req.UserUUID)
	if err != nil {
		return response.UserFamilyResponse{}, fmt.Errorf("invalid user UUID")
	}

	for _, member := range req.Members {
		userFamily := &models.UserFamily{
			UUID:     uuid.NewString(),
			UserID:   user.ID,
			Status:   *req.Status,
			Name:     member.Name,
			Relation: member.Relation,
			Phone:    member.Phone,
			CreateBy: actorID,
			ModifyBy: actorID,
		}

		if err := s.repo.Create(userFamily); err != nil {
			return response.UserFamilyResponse{}, err
		}
	}

	// Now return grouped response
	grouped, err := s.GetByUserUUID(req.UserUUID)
	if err != nil || len(grouped) == 0 {
		return response.UserFamilyResponse{}, err
	}

	return grouped[0], nil
}

func (s *userFamilyService) GetAll() ([]response.UserFamilyResponse, error) {
	families, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	// Map to group members by UserID
	grouped := make(map[uint]*response.UserFamilyResponse)

	for _, fam := range families {
		fullname := fam.User.FirstName + " " + fam.User.LastName

		member := struct {
			UUID     string `json:"uuid"`
			Name     string `json:"name"`
			Relation string `json:"relation"`
			Phone    string `json:"phone"`
		}{
			UUID:     fam.UUID,
			Name:     fam.Name,
			Relation: fam.Relation,
			Phone:    fam.Phone,
		}

		if entry, exists := grouped[fam.UserID]; exists {
			entry.Members = append(entry.Members, member)
			entry.Count += 1
		} else {
			grouped[fam.UserID] = &response.UserFamilyResponse{
				Status: fam.Status,
				Count:  1,
				Members: []struct {
					UUID     string `json:"uuid"`
					Name     string `json:"name"`
					Relation string `json:"relation"`
					Phone    string `json:"phone"`
				}{
					member,
				},
				User: struct {
					UUID     string `json:"uuid"`
					FullName string `json:"fullname"`
					Email    string `json:"email"`
				}{
					UUID:     fam.User.UUID,
					FullName: fullname,
					Email:    fam.User.Email,
				},
			}
		}
	}

	// Convert map to slice
	var responses []response.UserFamilyResponse
	for _, val := range grouped {
		responses = append(responses, *val)
	}

	return responses, nil
}

func (s *userFamilyService) GetByID(id uint) (response.UserFamilyResponse, error) {
	fam, err := s.repo.GetByID(id)
	if err != nil {
		return response.UserFamilyResponse{}, err
	}

	userUUID := fam.User.UUID
	families, err := s.repo.GetByUserUUID(userUUID)
	if err != nil {
		return response.UserFamilyResponse{}, err
	}

	if len(families) == 0 {
		return response.UserFamilyResponse{}, nil
	}

	user := families[0].User
	grouped := response.UserFamilyResponse{
		Status: families[0].Status,
		Count:  uint(len(families)),
		Members: []struct {
			UUID     string `json:"uuid"`
			Name     string `json:"name"`
			Relation string `json:"relation"`
			Phone    string `json:"phone"`
		}{},
		User: struct {
			UUID     string `json:"uuid"`
			FullName string `json:"fullname"`
			Email    string `json:"email"`
		}{
			UUID:     user.UUID,
			FullName: user.FirstName + " " + user.LastName,
			Email:    user.Email,
		},
	}

	for _, f := range families {
		member := struct {
			UUID     string `json:"uuid"`
			Name     string `json:"name"`
			Relation string `json:"relation"`
			Phone    string `json:"phone"`
		}{
			UUID:     f.UUID,
			Name:     f.Name,
			Relation: f.Relation,
			Phone:    f.Phone,
		}
		grouped.Members = append(grouped.Members, member)
	}

	return grouped, nil
}

func (s *userFamilyService) GetByUUID(uuid string) (response.UserFamilyResponse, error) {
	fam, err := s.repo.GetByUUID(uuid)
	if err != nil {
		return response.UserFamilyResponse{}, err
	}

	resp := response.UserFamilyResponse{
		Status: fam.Status,
		Count:  1,
		Members: []struct {
			UUID     string `json:"uuid"`
			Name     string `json:"name"`
			Relation string `json:"relation"`
			Phone    string `json:"phone"`
		}{
			{
				UUID:     fam.UUID,
				Name:     fam.Name,
				Relation: fam.Relation,
				Phone:    fam.Phone,
			},
		},
		User: struct {
			UUID     string `json:"uuid"`
			FullName string `json:"fullname"`
			Email    string `json:"email"`
		}{
			UUID:     fam.User.UUID,
			FullName: fam.User.FirstName + " " + fam.User.LastName,
			Email:    fam.User.Email,
		},
	}

	return resp, nil
}

func (s *userFamilyService) GetByUserUUID(userUUID string) ([]response.UserFamilyResponse, error) {
	_, err := s.userRepo.GetByUUID(userUUID)
	if err != nil {
		return nil, fmt.Errorf("invalid user UUID")
	}

	families, err := s.repo.GetByUserUUID(userUUID)
	if err != nil {
		return nil, err
	}

	if len(families) == 0 {
		return nil, nil
	}

	user := families[0].User
	grouped := response.UserFamilyResponse{
		Status: families[0].Status,
		Count:  uint(len(families)),
		Members: []struct {
			UUID     string `json:"uuid"`
			Name     string `json:"name"`
			Relation string `json:"relation"`
			Phone    string `json:"phone"`
		}{},
		User: struct {
			UUID     string `json:"uuid"`
			FullName string `json:"fullname"`
			Email    string `json:"email"`
		}{
			UUID:     user.UUID,
			FullName: user.FirstName + " " + user.LastName,
			Email:    user.Email,
		},
	}

	for _, fam := range families {
		grouped.Members = append(grouped.Members, struct {
			UUID     string `json:"uuid"`
			Name     string `json:"name"`
			Relation string `json:"relation"`
			Phone    string `json:"phone"`
		}{
			UUID:     fam.UUID,
			Name:     fam.Name,
			Relation: fam.Relation,
			Phone:    fam.Phone,
		})
	}

	return []response.UserFamilyResponse{grouped}, nil
}

func (s *userFamilyService) Update(userUUID string, req request.UserFamilyRequest, actorID uint) (response.UserFamilyResponse, error) {
	// Fetch the user by UUID
	fmt.Print("Updating family for user UUID: ", userUUID, "\n")
	user, err := s.userRepo.GetByUUID(userUUID)
	if err != nil {
		return response.UserFamilyResponse{}, fmt.Errorf("invalid user UUID")
	}

	// Validate request
	if err := pkg.Validate.Struct(req); err != nil {
		return response.UserFamilyResponse{}, err
	}

	fmt.Print("SERVICE PARAM: ", userUUID, "\n")
	fmt.Print("SERVICE REQ: ", req.UserUUID, "\n")
	if userUUID != req.UserUUID {
		return response.UserFamilyResponse{}, fmt.Errorf("user UUID in path does not match payload")
	}
	// Delete all family records for this user
	if err := s.repo.DeleteByUserID(user.ID); err != nil {
		return response.UserFamilyResponse{}, err
	}

	// Insert new family members
	for _, member := range req.Members {
		newFam := &models.UserFamily{
			UUID:     uuid.NewString(),
			UserID:   user.ID,
			Status:   *req.Status,
			Name:     member.Name,
			Relation: member.Relation,
			Phone:    member.Phone,
			CreateBy: actorID,
			ModifyBy: actorID,
		}

		if err := s.repo.Create(newFam); err != nil {
			return response.UserFamilyResponse{}, err
		}
	}

	// Return updated family response
	res, err := s.GetByUserUUID(userUUID)
	if err != nil || len(res) == 0 {
		return response.UserFamilyResponse{}, err
	}

	return res[0], nil
}

func (s *userFamilyService) Delete(uuid string) (response.UserFamilyResponse, error) {
	// Fetch the family member to delete
	fam, err := s.repo.GetByUUID(uuid)
	if err != nil {
		return response.UserFamilyResponse{}, err
	}

	// Delete the specific family record
	if err := s.repo.Delete(uuid); err != nil {
		return response.UserFamilyResponse{}, err
	}

	// Fetch remaining family members of that user
	families, err := s.repo.GetByUserUUID(fam.User.UUID)
	if err != nil {
		return response.UserFamilyResponse{}, err
	}

	// Build base response
	resp := response.UserFamilyResponse{
		Status: fam.Status,
		Count:  uint(len(families)),
		Members: []struct {
			UUID     string `json:"uuid"`
			Name     string `json:"name"`
			Relation string `json:"relation"`
			Phone    string `json:"phone"`
		}{},
		User: struct {
			UUID     string `json:"uuid"`
			FullName string `json:"fullname"`
			Email    string `json:"email"`
		}{
			UUID:     fam.User.UUID,
			FullName: fam.User.FirstName + " " + fam.User.LastName,
			Email:    fam.User.Email,
		},
	}

	if len(families) == 0 {
		return resp, nil
	}

	resp.Status = families[0].Status

	for _, f := range families {
		resp.Members = append(resp.Members, struct {
			UUID     string `json:"uuid"`
			Name     string `json:"name"`
			Relation string `json:"relation"`
			Phone    string `json:"phone"`
		}{
			UUID:     f.UUID,
			Name:     f.Name,
			Relation: f.Relation,
			Phone:    f.Phone,
		})
	}

	return resp, nil
}
