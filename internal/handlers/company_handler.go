package handlers

import (
	"hris_backend/internal/request"
	"hris_backend/internal/services"
	"hris_backend/pkg"

	"github.com/gofiber/fiber/v2"
)

type CompanyHandler interface {
	Create(c *fiber.Ctx) error
	GetAll(c *fiber.Ctx) error
	GetByUUID(c *fiber.Ctx) error
	GetByUserUUID(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
}

type companyHandler struct {
	companyServices services.CompanyServices
}

func NewCompanyHandler(companyServices services.CompanyServices) CompanyHandler {
	return &companyHandler{companyServices}
}

func (h *companyHandler) Create(c *fiber.Ctx) error {
	var companyReq request.CompanyReq
	if err := c.BodyParser(&companyReq); err != nil {
		return pkg.Error(c, fiber.StatusBadRequest, "Failed to parse company")
	}

	newCompany, err := h.companyServices.Create(companyReq)
	if err != nil {
		return pkg.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return pkg.Created(c, newCompany, "Company Created Succesfully")
}

func (h *companyHandler) GetAll(c *fiber.Ctx) error {
	companies, err := h.companyServices.GetAll()
	if err != nil {
		return pkg.Error(c, fiber.StatusInternalServerError, "Failed to get companies")
	}
	return pkg.Success(c, companies, "Successfully get all companies")
}

func (h *companyHandler) GetByUUID(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	if uuid == "" {
		return pkg.Error(c, fiber.StatusBadRequest, "UUID is required")
	}

	company, err := h.companyServices.GetByUUID(uuid)
	if err != nil {
		return pkg.Error(c, fiber.StatusNotFound, "Failed to get company")
	}
	return pkg.Success(c, company, "Succesfully get company")
}

func (h *companyHandler) GetByUserUUID(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	if uuid == "" {
		return pkg.Error(c, fiber.StatusBadRequest, "UUID is required")
	}

	company, err := h.companyServices.GetByUserUUID(uuid)
	if err != nil {
		if err.Error() == "user-not-found" {
			return pkg.Error(c, fiber.StatusNotFound, "User not found")
		}
		return pkg.Error(c, fiber.StatusInternalServerError, "Something went wrong")
	}

	if company.UUID == "" {
		return pkg.Success(c, nil, "No company found for this user")
	}

	return pkg.Success(c, company, "Successfully retrieved company by user UUID")
}



func (h *companyHandler) Update(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	if uuid == "" {
		return pkg.Error(c, fiber.StatusBadRequest, "UUID is required")
	}

	var companyReq request.CompanyReq
	if err := c.BodyParser(&companyReq); err != nil {
		return pkg.Error(c, fiber.StatusBadRequest, "Failed to parse company")
	}

	updatedCompany, err := h.companyServices.Update(uuid, companyReq)
	if err != nil {
		return pkg.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return pkg.Success(c, updatedCompany, "Company updated succesfully")
}

func (h *companyHandler) Delete(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	if uuid == "" {
		return pkg.Error(c, fiber.StatusBadRequest, "UUID is required")
	}

	targetCompany, err := h.companyServices.Delete(uuid)
	if err != nil {
		return pkg.Error(c, fiber.StatusNotFound, err.Error())
	}

	return pkg.Success(c, targetCompany, "Company deleted successfully")

}
