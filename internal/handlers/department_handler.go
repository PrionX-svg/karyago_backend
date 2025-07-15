package handlers

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"hris_backend/internal/request"
	"hris_backend/internal/services"
	"hris_backend/pkg"
)

type DepartmentHandler struct {
	service services.DepartmentServices
}

func NewDepartmentHandler(s services.DepartmentServices) *DepartmentHandler {
	return &DepartmentHandler{s}
}

func (h *DepartmentHandler) Create(c *fiber.Ctx) error {
	var req request.DepartmentReq

	actorID, ok := c.Locals("user_id").(uint)
	if !ok {
		fmt.Println("Invalid user_id:", c.Locals("user_id"))
		return pkg.Error(c, fiber.StatusUnauthorized, "Unauthorized")
	}

	if err := c.BodyParser(&req); err != nil {
		return pkg.Error(c, fiber.StatusBadRequest, "Invalid request body")
	}

	req.CreatedBy = actorID
	req.ModifyBy = actorID

	department, err := h.service.Create(req)
	if err != nil {
		return pkg.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return pkg.Created(c, department, "Department created")
}

func (h *DepartmentHandler) Get(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	department, err := h.service.FindByUUID(uuid)
	if err != nil {
		return pkg.Error(c, fiber.StatusNotFound, "Department not found")
	}
	return pkg.Success(c, department, "Successfully get department")
}

func (h *DepartmentHandler) GetAllDataTable(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	search := c.Query("search", "")
	companyUUID := c.Query("company_uuid")
	departmentGroupUUID := c.Query("department_group_uuid", "")

	if companyUUID == "" {
		return pkg.Error(c, fiber.StatusBadRequest, "company_uuid is required")
	}

	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 10
	}

	data, total, filtered, err := h.service.GetDataTable(page, limit, search, companyUUID, departmentGroupUUID)
	if err != nil {
		return pkg.Error(c, fiber.StatusInternalServerError, "Failed to get department data table")
	}

	return c.JSON(fiber.Map{
		"data":     data,
		"page":     page,
		"limit":    limit,
		"total":    total,
		"filtered": filtered,
	})
}

func (h *DepartmentHandler) Update(c *fiber.Ctx) error {
	uuid := c.Params("uuid")

	actorID, ok := c.Locals("user_id").(uint)
	if !ok {
		fmt.Println("Invalid user_id:", c.Locals("user_id"))
		return pkg.Error(c, fiber.StatusUnauthorized, "Unauthorized")
	}

	var req request.DepartmentReq
	if err := c.BodyParser(&req); err != nil {
		return pkg.Error(c, fiber.StatusBadRequest, "Invalid request body")
	}

	req.ModifyBy = actorID

	department, err := h.service.Update(uuid, req)
	if err != nil {
		return pkg.Error(c, fiber.StatusInternalServerError, err.Error())
	}
	return pkg.Success(c, department, "Department updated")
}

func (h *DepartmentHandler) Delete(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	department, err := h.service.Delete(uuid)
	if err != nil {
		return pkg.Error(c, fiber.StatusInternalServerError, err.Error())
	}
	return pkg.Success(c, department, "Department deleted")
}

func (h *DepartmentHandler) List(c *fiber.Ctx) error {
	departments, err := h.service.GetAll()
	if err != nil {
		return pkg.Error(c, fiber.StatusInternalServerError, err.Error())
	}
	return pkg.Success(c, departments, "Successfully get all departments")
}
