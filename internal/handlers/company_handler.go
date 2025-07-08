package handlers

import (
	"hris_backend/internal/request"
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
	var request request.CompanyReq
	if err := c.BodyParser(&request); err != nil{
		return pkg.Error(c, fiber.StatusBadRequest, "Failed to parse company")
	}

	newCompany, err := h.companyServices.Create(request); 
	if err != nil{
		return pkg.Error(c, fiber.StatusInternalServerError, err.Error())
	}
	
	return pkg.Created(c, newCompany, "Company Created Succesfully")
}