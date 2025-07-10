package handlers

import (
	"hris_backend/internal/request"
	"hris_backend/internal/services"
	"hris_backend/pkg"

	"github.com/gofiber/fiber/v2"
)

type DepartmentGroupHandler interface {
}

type departmentGroupHandler struct {
	departmentGroupServices services.DepartmentGroupServices
}

func NewDepartmentGroupHandler(departmentGroupsServices services.DepartmentGroupServices) DepartmentGroupHandler {
	return &departmentGroupHandler{
		departmentGroupServices: departmentGroupsServices,
	}
}

func (h *departmentGroupHandler) Create(c *fiber.Ctx) error {
	var departmentGroupReq request.DepartmentGroupCreateReq
	if err := c.BodyParser(&departmentGroupReq); err != nil {
		return pkg.Error(c, fiber.StatusBadRequest, "Failed to parse department group")
	}

	newDepartmentGroup, err := h.departmentGroupServices.Create(departmentGroupReq)
	if err != nil {
		return pkg.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return pkg.Created(c, newDepartmentGroup, "Department Group Created Succesfully")
}