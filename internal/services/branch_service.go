package services

import (
	"fmt"
	"hris_backend/internal/response"

	"hris_backend/internal/models"
	"hris_backend/internal/repositories"
	"hris_backend/internal/request"
	"hris_backend/pkg"

	"github.com/google/uuid"
)

type BranchService interface {
	Create(req request.BranchReq, actorID uint) (response.BranchResponse, error)
	Get(uuid string) (response.BranchResponse, error)
	Update(uuid string, req request.BranchReq, actorID uint) (response.BranchResponse, error)
	Delete(uuid string) (response.BranchResponse, error)
	List() ([]response.BranchResponse, error)
	GetByCompanyUUID(uuid string) ([]response.BranchResponse, error)
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

func (s *branchService) Create(req request.BranchReq, actorID uint) (response.BranchResponse, error) {
	if err := pkg.Validate.Struct(req); err != nil {
		return response.BranchResponse{}, err
	}

	company, err := s.companyRepo.GetByUUID(req.CompanyUUID)
	if err != nil {
		return response.BranchResponse{}, fmt.Errorf("invalid company UUID")
	}

	newBranch := models.Branch{
		UUID:      uuid.NewString(),
		CompanyID: company.ID,
		Image:     req.Image,
		Name:      req.Name,
		Address:   req.Address,
		Email:     req.Email,
		Phone:     req.Phone,
		CreatedBy: actorID,
		ModifyBy:  actorID,
	}

	if err := s.repo.Create(&newBranch); err != nil {
		return response.BranchResponse{}, err
	}

	branchModel, err := s.repo.FindByUUID(newBranch.UUID)
	if err != nil {
		return response.BranchResponse{}, err
	}

	branchResp := response.BranchResponse{
		UUID:    branchModel.UUID,
		Image:   branchModel.Image,
		Name:    branchModel.Name,
		Address: branchModel.Address,
		Email:   branchModel.Email,
		Phone:   branchModel.Phone,
		Company: struct {
			UUID    string `json:"uuid"`
			Logo    string `json:"logo"`
			Name    string `json:"name"`
			Address string `json:"address"`
			Email   string `json:"email"`
			Phone   string `json:"phone"`
		}{
			UUID:    branchModel.Company.UUID,
			Logo:    branchModel.Company.Logo,
			Name:    branchModel.Company.Name,
			Address: branchModel.Company.Address,
			Email:   branchModel.Company.Email,
			Phone:   branchModel.Company.Phone,
		},
	}

	return branchResp, nil
}

func (s *branchService) Get(uuid string) (response.BranchResponse, error) {
	branch, err := s.repo.FindByUUID(uuid)
	if err != nil {
		return response.BranchResponse{}, err
	}

	result := response.BranchResponse{
		UUID:    branch.UUID,
		Image:   branch.Image,
		Name:    branch.Name,
		Address: branch.Address,
		Email:   branch.Email,
		Phone:   branch.Phone,
		Company: struct {
			UUID    string `json:"uuid"`
			Logo    string `json:"logo"`
			Name    string `json:"name"`
			Address string `json:"address"`
			Email   string `json:"email"`
			Phone   string `json:"phone"`
		}{
			UUID:    branch.Company.UUID,
			Logo:    branch.Company.Logo,
			Name:    branch.Company.Name,
			Address: branch.Company.Address,
			Email:   branch.Company.Email,
			Phone:   branch.Company.Phone,
		},
	}

	return result, nil
}

func (s *branchService) GetByCompanyUUID(uuid string) ([]response.BranchResponse, error) {
	company, err := s.companyRepo.GetByUUID(uuid)
	if err != nil {
		return []response.BranchResponse{}, fmt.Errorf("company not found")
	}
	branches, err := s.repo.FindByCompanyID(company.ID)
	if err != nil {
		return []response.BranchResponse{}, fmt.Errorf("branch not found for company")
	}

	var results []response.BranchResponse
	for _, branch := range branches {
		results = append(results, response.BranchResponse{
			UUID:    branch.UUID,
			Image:   branch.Image,
			Name:    branch.Name,
			Address: branch.Address,
			Email:   branch.Email,
			Phone:   branch.Phone,
			Company: struct {
				UUID    string `json:"uuid"`
				Logo    string `json:"logo"`
				Name    string `json:"name"`
				Address string `json:"address"`
				Email   string `json:"email"`
				Phone   string `json:"phone"`
			}{
				UUID:    company.UUID,
				Logo:    company.Logo,
				Name:    company.Name,
				Address: company.Address,
				Email:   company.Email,
				Phone:   company.Phone,
			},
		})
	}

	return results, nil

}

func (s *branchService) Update(uuid string, req request.BranchReq, actorID uint) (response.BranchResponse, error) {
	if err := pkg.Validate.Struct(req); err != nil {
		return response.BranchResponse{}, err
	}

	branch, err := s.repo.FindByUUID(uuid)
	if err != nil {
		return response.BranchResponse{}, fmt.Errorf("branch not found")
	}

	company, err := s.companyRepo.GetByUUID(req.CompanyUUID)
	if err != nil {
		return response.BranchResponse{}, fmt.Errorf("invalid company UUID")
	}

	branch.Image = req.Image
	branch.Name = req.Name
	branch.Address = req.Address
	branch.Email = req.Email
	branch.Phone = req.Phone
	branch.CompanyID = company.ID
	branch.ModifyBy = actorID

	if err := s.repo.Update(branch); err != nil {
		return response.BranchResponse{}, err
	}

	updatedBranch, err := s.repo.FindByUUID(uuid)
	if err != nil {
		return response.BranchResponse{}, err
	}

	resp := response.BranchResponse{
		UUID:    updatedBranch.UUID,
		Image:   updatedBranch.Image,
		Name:    updatedBranch.Name,
		Address: updatedBranch.Address,
		Email:   updatedBranch.Email,
		Phone:   updatedBranch.Phone,
		Company: struct {
			UUID    string `json:"uuid"`
			Logo    string `json:"logo"`
			Name    string `json:"name"`
			Address string `json:"address"`
			Email   string `json:"email"`
			Phone   string `json:"phone"`
		}{
			UUID:    updatedBranch.Company.UUID,
			Logo:    updatedBranch.Company.Logo,
			Name:    updatedBranch.Company.Name,
			Address: updatedBranch.Company.Address,
			Email:   updatedBranch.Company.Email,
			Phone:   updatedBranch.Company.Phone,
		},
	}

	return resp, nil
}

func (s *branchService) Delete(uuid string) (response.BranchResponse, error) {
	branch, err := s.repo.FindByUUID(uuid)
	if err != nil {
		return response.BranchResponse{}, err
	}

	if err := s.repo.Delete(uuid); err != nil {
		return response.BranchResponse{}, err
	}

	resp := response.BranchResponse{
		UUID:    branch.UUID,
		Image:   branch.Image,
		Name:    branch.Name,
		Address: branch.Address,
		Email:   branch.Email,
		Phone:   branch.Phone,
		Company: struct {
			UUID    string `json:"uuid"`
			Logo    string `json:"logo"`
			Name    string `json:"name"`
			Address string `json:"address"`
			Email   string `json:"email"`
			Phone   string `json:"phone"`
		}{
			UUID:    branch.Company.UUID,
			Logo:    branch.Company.Logo,
			Name:    branch.Company.Name,
			Address: branch.Company.Address,
			Email:   branch.Company.Email,
			Phone:   branch.Company.Phone,
		},
	}

	return resp, nil
}

func (s *branchService) List() ([]response.BranchResponse, error) {
	branches, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}

	var result []response.BranchResponse
	for _, b := range branches {
		result = append(result, response.BranchResponse{
			UUID:    b.UUID,
			Image:   b.Image,
			Name:    b.Name,
			Address: b.Address,
			Email:   b.Email,
			Phone:   b.Phone,
			Company: struct {
				UUID    string `json:"uuid"`
				Logo    string `json:"logo"`
				Name    string `json:"name"`
				Address string `json:"address"`
				Email   string `json:"email"`
				Phone   string `json:"phone"`
			}{
				UUID:    b.Company.UUID,
				Logo:    b.Company.Logo,
				Name:    b.Company.Name,
				Address: b.Company.Address,
				Email:   b.Company.Email,
				Phone:   b.Company.Phone,
			},
		})
	}

	return result, nil
}
