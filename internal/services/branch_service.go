package services

import (
	"fmt"

	"github.com/google/uuid"
	"hris_backend/internal/models"
	"hris_backend/internal/repositories"
	"hris_backend/internal/request"
	"hris_backend/pkg"
)

type BranchService interface {
	Create(request request.BranchReq, actorID uint) (models.Branch, error)
	Get(uuid string) (models.Branch, error)
	Update(uuid string, request request.BranchReq, actorID uint) (models.Branch, error)
	Delete(uuid string) (models.Branch, error)
	List() ([]models.Branch, error)
}

type branchService struct {
	repo        repositories.BranchRepository
	companyRepo repositories.CompanyRepositories
}

func NewBranchService(branchRepo repositories.BranchRepository, companyRepo repositories.CompanyRepositories) BranchService {
	return &branchService{
		repo:        branchRepo,
		companyRepo: companyRepo,
	}
}

func (s *branchService) Create(req request.BranchReq, actorID uint) (models.Branch, error) {
	if err := pkg.Validate.Struct(req); err != nil {
		return models.Branch{}, err
	}

	company, err := s.companyRepo.GetByUUID(req.CompanyUUID)
	if err != nil {
		return models.Branch{}, fmt.Errorf("invalid company UUID")
	}

	newBranch := models.Branch{
		UUID:      uuid.NewString(),
		CompanyID: company.ID,
		Name:      req.Name,
		Address:   req.Address,
		Email:     req.Email,
		Phone:     req.Phone,
		CreatedBy: actorID,
		ModifyBy:  actorID,
	}

	if err := s.repo.Create(&newBranch); err != nil {
		return models.Branch{}, err
	}
	return newBranch, nil
}

func (s *branchService) Get(uuid string) (models.Branch, error) {
	branch, err := s.repo.FindByUUID(uuid)
	if err != nil {
		return models.Branch{}, err
	}
	return *branch, nil
}

func (s *branchService) Update(uuid string, req request.BranchReq, actorID uint) (models.Branch, error) {
	if err := pkg.Validate.Struct(req); err != nil {
		return models.Branch{}, err
	}

	branch, err := s.repo.FindByUUID(uuid)
	if err != nil {
		return models.Branch{}, fmt.Errorf("branch not found")
	}

	company, err := s.companyRepo.GetByUUID(req.CompanyUUID)
	if err != nil {
		return models.Branch{}, fmt.Errorf("invalid company UUID")
	}

	branch.Name = req.Name
	branch.Address = req.Address
	branch.Email = req.Email
	branch.Phone = req.Phone
	branch.CompanyID = company.ID
	branch.ModifyBy = actorID

	if err := s.repo.Update(branch); err != nil {
		return models.Branch{}, err
	}

	return *branch, nil
}

func (s *branchService) Delete(uuid string) (models.Branch, error) {
	branch, err := s.repo.FindByUUID(uuid)
	if err != nil {
		return models.Branch{}, err
	}

	if err := s.repo.Delete(uuid); err != nil {
		return models.Branch{}, err
	}

	return *branch, nil
}

func (s *branchService) List() ([]models.Branch, error) {
	return s.repo.FindAll()
}
