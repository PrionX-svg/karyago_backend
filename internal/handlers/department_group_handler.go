package handlers

import (
	"hris_backend/internal/request"
	"hris_backend/internal/services"
	"hris_backend/pkg"

	"github.com/gofiber/fiber/v2"
)

type DepartmentGroupHandler interface {
	Create(c *fiber.Ctx) error
	GetAll(c *fiber.Ctx) error
	GetByUUID(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
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
	var departmentGroupReq request.DepartmentGroupReq
	if err := c.BodyParser(&departmentGroupReq); err != nil {
		return pkg.Error(c, fiber.StatusBadRequest, "Failed to parse department group")
	}

	newDepartmentGroup, err := h.departmentGroupServices.Create(departmentGroupReq)
	if err != nil {
		return pkg.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return pkg.Created(c, newDepartmentGroup, "Department Group Created Succesfully")
}

func (h *departmentGroupHandler) GetAll(c *fiber.Ctx) error {
	departmentGroups, err := h.departmentGroupServices.GetAll()
	if err != nil {
		return pkg.Error(c, fiber.StatusInternalServerError, "Failed to get department groups")
	}
	return pkg.Success(c, departmentGroups, "Successfully get all department groups")
}

func (h *departmentGroupHandler) GetByUUID(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	if uuid == "" {
		return pkg.Error(c, fiber.StatusBadRequest, "UUID is required")
	}

	departmentGroup, err := h.departmentGroupServices.FindByUUID(uuid)
	if err != nil {
		return pkg.Error(c, fiber.StatusNotFound, "Failed to get department group")
	}
	return pkg.Success(c, departmentGroup, "Succesfully get department group")
}

func (h *departmentGroupHandler) Update(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	if uuid == "" {
		return pkg.Error(c, fiber.StatusBadRequest, "UUID is required")
	}

	var departmentGroupReq request.DepartmentGroupReq
	if err := c.BodyParser(&departmentGroupReq); err != nil {
		return pkg.Error(c, fiber.StatusBadRequest, "Failed to parse department group")
	}

	updatedDepartmentGroup, err := h.departmentGroupServices.Update(uuid, departmentGroupReq)
	if err != nil {
		return pkg.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return pkg.Success(c, updatedDepartmentGroup, "Department Group Updated Successfully")
}

func (h *departmentGroupHandler) Delete(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	if uuid == "" {
		return pkg.Error(c, fiber.StatusBadRequest, "UUID is required")
	}

	targetDepartmentGroup, err := h.departmentGroupServices.Delete(uuid); 
	if err != nil {
		return pkg.Error(c, fiber.StatusInternalServerError, "Failed to delete department group")
	}

	return pkg.Success(c, targetDepartmentGroup, "Department Group Deleted Successfully")
}