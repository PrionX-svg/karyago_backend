package handlers

import (
	"hris_backend/internal/request"
	"hris_backend/internal/services"
	"hris_backend/pkg"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type DepartmentGroupHandler interface {
	Create(c *fiber.Ctx) error
	GetAll(c *fiber.Ctx) error
	GetByUUID(c *fiber.Ctx) error
	GetAllDataTable(c *fiber.Ctx) error
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

func (h *departmentGroupHandler) GetAllDataTable(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	search := c.Query("search", "")
	companyUUID := c.Query("company_uuid")

	if companyUUID == "" {
		return pkg.Error(c, fiber.StatusBadRequest, "company_uuid is required")
	}

	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 10
	}

	data, total, filtered, err := h.departmentGroupServices.GetDataTable(page, limit, search, companyUUID)
	if err != nil {
		return pkg.Error(c, fiber.StatusInternalServerError, "Failed to get department group data table")
	}

	return c.JSON(fiber.Map{
		"data":     data,
		"page":     page,
		"limit":    limit,
		"total":    total,
		"filtered": filtered,
	})
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

	targetDepartmentGroup, err := h.departmentGroupServices.Delete(uuid)
	if err != nil {
		return pkg.Error(c, fiber.StatusInternalServerError, "Failed to delete department group")
	}

	return pkg.Success(c, targetDepartmentGroup, "Department Group Deleted Successfully")
}
