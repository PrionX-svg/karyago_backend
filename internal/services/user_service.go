package services

import (
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"hris_backend/internal/models"
	"hris_backend/internal/repositories"
	"hris_backend/internal/request"
	"hris_backend/internal/response"
	"hris_backend/pkg"
)

type userService struct {
	db          *gorm.DB
	userRepo    repositories.UserRepository
	roleRepo    repositories.RoleRepositories
	companyRepo repositories.CompanyRepositories
	branchRepo  repositories.BranchRepository
}

type UserService interface {
	CreateUser(req request.UserEmployeeReq, creatorID uint) error
	GetMe(userID uint) (*response.UserWithEmployeeResponse, error)
	GetAllUsers() ([]response.UserWithEmployeeResponse, error)
	GetUserByUUID(uuid string) (*response.UserWithEmployeeResponse, error)
	GetUsersWithEmployeeDataTable(page, limit int, search, roleUUID, branchUUID string) ([]response.UserWithEmployeeResponse, int64, int64, error)
	UpdateUser(userUUID string, req request.UserEmployeeReq, modifierID uint) error
	DeleteUser(userUUID string) error
}

func NewUserService(
	db *gorm.DB,
	userRepo repositories.UserRepository,
	roleRepo repositories.RoleRepositories,
	branchRepo repositories.BranchRepository,
) UserService {
	return &userService{
		db:         db,
		userRepo:   userRepo,
		roleRepo:   roleRepo,
		branchRepo: branchRepo,
	}
}

func (s *userService) CreateUser(req request.UserEmployeeReq, creatorID uint) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		role, err := s.roleRepo.FindByUUID(req.RoleUUID)
		if err != nil {
			return fmt.Errorf("role not found: %w", err)
		}

		var branchID *uint
		if req.BranchUUID != "" {
			branch, err := s.branchRepo.FindByUUID(req.BranchUUID)
			if err != nil {
				return fmt.Errorf("branch not found: %w", err)
			}
			branchID = &branch.ID
		}

		hashedPassword, err := pkg.HashPassword(req.Password)
		if err != nil {
			return fmt.Errorf("failed to hash password: %w", err)
		}

		user := &models.User{
			UUID:        uuid.NewString(),
			RoleID:      role.ID,
			BranchID:    branchID,
			FirstName:   req.FirstName,
			LastName:    req.LastName,
			Phone:       req.Phone,
			Email:       req.Email,
			Password:    hashedPassword,
			DOB:         &req.DOB,
			Gender:      &req.Gender,
			IsFreelance: req.IsFreelance,
			CreatedBy:   creatorID,
			ModifyBy:    creatorID,
		}

		return s.userRepo.Create(user)
	})
}

func (s *userService) GetMe(userID uint) (*response.UserWithEmployeeResponse, error) {
	user, err := s.userRepo.GetByID(userID)
	if err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}

	var branchUUID, branchName string
	if user.BranchID != nil {
		branch, err := s.branchRepo.FindByID(*user.BranchID)
		if err == nil {
			branchUUID = branch.UUID
			branchName = branch.Name
		}
	}

	role, err := s.roleRepo.FindByID(user.RoleID)
	if err != nil {
		return nil, fmt.Errorf("role not found: %w", err)
	}

	return &response.UserWithEmployeeResponse{
		UserUUID:    user.UUID,
		FullName:    user.FirstName + " " + user.LastName,
		Email:       user.Email,
		Phone:       user.Phone,
		Gender:      user.Gender,
		DOB:         user.DOB,
		IsFreelance: user.IsFreelance,
		Role:        role.Name,
		Branch: struct {
			UUID string `json:"uuid"`
			Name string `json:"name"`
		}{UUID: branchUUID, Name: branchName},
	}, nil
}

func (s *userService) GetAllUsers() ([]response.UserWithEmployeeResponse, error) {
	users, err := s.userRepo.List()
	if err != nil {
		return nil, err
	}

	var result []response.UserWithEmployeeResponse

	for _, user := range users {
		role, _ := s.roleRepo.FindByID(user.RoleID)

		var branchUUID, branchName string
		if user.BranchID != nil {
			branch, err := s.branchRepo.FindByID(*user.BranchID)
			if err == nil {
				branchUUID = branch.UUID
				branchName = branch.Name
			}
		}

		result = append(result, response.UserWithEmployeeResponse{
			UserUUID:    user.UUID,
			FullName:    user.FirstName + " " + user.LastName,
			Email:       user.Email,
			Phone:       user.Phone,
			Gender:      user.Gender,
			DOB:         user.DOB,
			IsFreelance: user.IsFreelance,
			Role:        role.Name,
			Branch: struct {
				UUID string `json:"uuid"`
				Name string `json:"name"`
			}{UUID: branchUUID, Name: branchName},
		})
	}

	return result, nil
}

func (s *userService) GetUserByUUID(userUUID string) (*response.UserWithEmployeeResponse, error) {
	user, err := s.userRepo.GetByUUID(userUUID)
	if err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}

	var branchUUID, branchName string
	if user.BranchID != nil {
		branch, err := s.branchRepo.FindByID(*user.BranchID)
		if err == nil {
			branchUUID = branch.UUID
			branchName = branch.Name
		}
	}

	role, err := s.roleRepo.FindByID(user.RoleID)
	if err != nil {
		return nil, fmt.Errorf("role not found: %w", err)
	}

	return &response.UserWithEmployeeResponse{
		UserUUID:    user.UUID,
		FullName:    user.FirstName + " " + user.LastName,
		Email:       user.Email,
		Phone:       user.Phone,
		Gender:      user.Gender,
		DOB:         user.DOB,
		IsFreelance: user.IsFreelance,
		Role:        role.Name,
		Branch: struct {
			UUID string `json:"uuid"`
			Name string `json:"name"`
		}{UUID: branchUUID, Name: branchName},
	}, nil
}

func (s *userService) GetUsersWithEmployeeDataTable(page, limit int, search, roleUUID, branchUUID string) ([]response.UserWithEmployeeResponse, int64, int64, error) {
	var (
		users    []models.User
		result   []response.UserWithEmployeeResponse
		total    int64
		filtered int64
		offset   = (page - 1) * limit
	)

	query := s.db.Model(&models.User{}).Preload("Role").Preload("Branch")

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, 0, fmt.Errorf("failed to count total users: %w", err)
	}

	if search != "" {
		likeQuery := "%" + search + "%"
		query = query.Where("first_name LIKE ? OR last_name LIKE ? OR email LIKE ? OR phone LIKE ?", likeQuery, likeQuery, likeQuery, likeQuery)
	}
	if roleUUID != "" {
		role, err := s.roleRepo.FindByUUID(roleUUID)
		if err == nil {
			query = query.Where("role_id = ?", role.ID)
		}
	}
	if branchUUID != "" {
		branch, err := s.branchRepo.FindByUUID(branchUUID)
		if err == nil {
			query = query.Where("branch_id = ?", branch.ID)
		}
	}

	if err := query.Count(&filtered).Error; err != nil {
		return nil, 0, 0, fmt.Errorf("failed to count filtered users: %w", err)
	}

	if err := query.Offset(offset).Limit(limit).Order("created_at DESC").Find(&users).Error; err != nil {
		return nil, 0, 0, fmt.Errorf("failed to query users: %w", err)
	}

	for _, user := range users {
		role, _ := s.roleRepo.FindByID(user.RoleID)

		var branchUUIDStr, branchName string
		if user.BranchID != nil {
			branch, err := s.branchRepo.FindByID(*user.BranchID)
			if err == nil {
				branchUUIDStr = branch.UUID
				branchName = branch.Name
			}
		}

		result = append(result, response.UserWithEmployeeResponse{
			UserUUID:    user.UUID,
			FullName:    user.FirstName + " " + user.LastName,
			Email:       user.Email,
			Phone:       user.Phone,
			Gender:      user.Gender,
			DOB:         user.DOB,
			IsFreelance: user.IsFreelance,
			Role:        role.Name,
			Branch: struct {
				UUID string `json:"uuid"`
				Name string `json:"name"`
			}{UUID: branchUUIDStr, Name: branchName},
		})
	}

	return result, total, filtered, nil
}

func (s *userService) UpdateUser(userUUID string, req request.UserEmployeeReq, modifierID uint) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		user, err := s.userRepo.GetByUUID(userUUID)
		if err != nil {
			return fmt.Errorf("user not found: %w", err)
		}

		role, err := s.roleRepo.FindByUUID(req.RoleUUID)
		if err != nil {
			return fmt.Errorf("role not found: %w", err)
		}

		var branchID *uint
		if req.BranchUUID != "" {
			branch, err := s.branchRepo.FindByUUID(req.BranchUUID)
			if err != nil {
				return fmt.Errorf("branch not found: %w", err)
			}
			branchID = &branch.ID
		}

		user.RoleID = role.ID
		user.BranchID = branchID
		user.FirstName = req.FirstName
		user.LastName = req.LastName
		user.Phone = req.Phone
		user.Email = req.Email
		user.DOB = &req.DOB
		user.Gender = &req.Gender
		user.IsFreelance = req.IsFreelance
		user.ModifyBy = modifierID

		if req.Password != "" {
			hashedPassword, err := pkg.HashPassword(req.Password)
			if err != nil {
				return fmt.Errorf("failed to hash password: %w", err)
			}
			user.Password = hashedPassword
		}

		if err := s.userRepo.Update(&user); err != nil {
			return fmt.Errorf("failed to update user: %w", err)
		}

		return nil
	})
}

func (s *userService) DeleteUser(userUUID string) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		user, err := s.userRepo.GetByUUID(userUUID)
		if err != nil {
			return fmt.Errorf("user not found: %w", err)
		}

		if err := s.userRepo.Delete(user.ID); err != nil {
			return fmt.Errorf("failed to delete user: %w", err)
		}

		return nil
	})
}
