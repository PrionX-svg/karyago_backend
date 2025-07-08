package handlers

import (
	"hris_backend/internal/models"
	"hris_backend/internal/services"
	"hris_backend/pkg"

	"github.com/gofiber/fiber/v2"
)


type CompanyHandler interface {
	Create(c * fiber.Ctx) error
}

type companyHandler struct {
	companyServices services.CompanyServices
}

func NewCompanyHandler(companyServices services.CompanyServices) CompanyHandler {
	return &companyHandler{companyServices}
}

func (h *companyHandler) Create(c * fiber.Ctx) error {
	var newCompany models.Company
	if err := c.BodyParser(&newCompany); err != nil{
		return pkg.Error(c, fiber.StatusBadRequest, "Failed to parse company")
	}
	
	if err := h.companyServices.Create(&newCompany); err != nil{
		return pkg.Error(c, fiber.StatusInternalServerError, err.Error())
	}
	
	return pkg.Created(c, newCompany, "Company Created Succesfully")
}