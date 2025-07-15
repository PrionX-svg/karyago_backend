package handlers

import (
	"github.com/gofiber/fiber/v2"
	"hris_backend/internal/models"
	"hris_backend/internal/repositories"
	"hris_backend/internal/request"
	"hris_backend/internal/services"
	"hris_backend/pkg"
)

type RoleHandler struct {
	service      services.RoleService
	companyRepo  repositories.CompanyRepositories
	employeeRepo repositories.EmployeeRepository
	roleRepo     repositories.RoleRepositories
}

func NewRoleHandler(service services.RoleService, companyRepo repositories.CompanyRepositories, employeeRepo repositories.EmployeeRepository, roleRepo repositories.RoleRepositories) *RoleHandler {
	return &RoleHandler{service, companyRepo, employeeRepo, roleRepo}
}

func (h *RoleHandler) GetByUUID(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	if uuid == "" {
		return pkg.Error(c, fiber.StatusBadRequest, "UUID is required")
	}

	role, err := h.service.GetByUUID(uuid)
	if err != nil {
		return pkg.Error(c, fiber.StatusNotFound, "Role not found")
	}

	return pkg.Success(c, role)
}

func (h *RoleHandler) GetByName(c *fiber.Ctx) error {
	name := c.Query("name")
	if name == "" {
		return pkg.Error(c, fiber.StatusBadRequest, "Name query param is required")
	}

	role, err := h.service.GetByName(name)
	if err != nil {
		return pkg.Error(c, fiber.StatusNotFound, "Role not found")
	}

	return pkg.Success(c, role)
}

func (h *RoleHandler) GetAll(c *fiber.Ctx) error {
	limit := c.QueryInt("limit", 10)
	page := c.QueryInt("page", 1)
	search := c.Query("search", "")
	sort := c.Query("sort", "id desc")
	companyUUID := c.Query("company_uuid", "")

	if limit <= 0 {
		limit = 10
	}
	offset := (page - 1) * limit

	roles, total, err := h.service.GetAll(limit, offset, search, sort, companyUUID)
	if err != nil {
		return pkg.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return pkg.Success(c, fiber.Map{
		"data":       roles,
		"total":      total,
		"page":       page,
		"limit":      limit,
		"totalPages": (total + int64(limit) - 1) / int64(limit),
	})
}

func (h *RoleHandler) Create(c *fiber.Ctx) error {
	var req request.CreateRoleRequest
	if err := c.BodyParser(&req); err != nil || req.Name == "" {
		return pkg.Error(c, fiber.StatusBadRequest, "Invalid request body")
	}

	actorID, ok := c.Locals("user_id").(uint)
	if !ok {
		return pkg.Error(c, fiber.StatusUnauthorized, "Unauthorized: user_id not found")
	}

	employee, err := h.employeeRepo.FindByUserID(actorID)
	if err != nil || employee.CompanyID == nil {
		return pkg.Error(c, fiber.StatusForbidden, "Failed to determine your company")
	}

	role, err := h.roleRepo.FindByID(employee.RoleID)
	if err != nil {
		return pkg.Error(c, fiber.StatusForbidden, "Failed to check user role")
	}

	var companyID *uint
	var companyUUID string

	if req.CompanyUUID != "" {
		company, err := h.companyRepo.GetByUUID(req.CompanyUUID)
		if err != nil {
			return pkg.Error(c, fiber.StatusBadRequest, "Invalid company UUID")
		}

		if role.Name != "Owner" && *employee.CompanyID != company.ID {
			return pkg.Error(c, fiber.StatusForbidden, "You are not allowed to create role for this company")
		}

		companyID = &company.ID
		companyUUID = company.UUID
	} else {
		if role.Name == "Owner" {
			return pkg.Error(c, fiber.StatusBadRequest, "Owner must provide company UUID")
		}
		companyID = employee.CompanyID
		company, err := h.companyRepo.GetByID(*companyID)
		if err != nil {
			return pkg.Error(c, fiber.StatusInternalServerError, "Failed to resolve company UUID")
		}
		companyUUID = company.UUID
	}

	roleModel := &models.Role{
		Name:      req.Name,
		CompanyID: companyID,
	}

	if err := h.service.Create(roleModel, actorID); err != nil {
		return pkg.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	response := fiber.Map{
		"uuid":         roleModel.UUID,
		"name":         roleModel.Name,
		"company_uuid": companyUUID,
	}

	return pkg.Created(c, response, "Role created successfully")
}

func (h *RoleHandler) Update(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	if uuid == "" {
		return pkg.Error(c, fiber.StatusBadRequest, "UUID is required")
	}

	existingRole, err := h.service.GetByUUID(uuid)
	if err != nil {
		return pkg.Error(c, fiber.StatusNotFound, "Role not found")
	}

	actorID, ok := c.Locals("user_id").(uint)
	if !ok {
		return pkg.Error(c, fiber.StatusUnauthorized, "Unauthorized")
	}

	employee, err := h.employeeRepo.FindByUserID(actorID)
	if err != nil || employee.CompanyID == nil {
		return pkg.Error(c, fiber.StatusForbidden, "You are not associated with any company")
	}

	role, err := h.roleRepo.FindByID(employee.RoleID)
	if err != nil {
		return pkg.Error(c, fiber.StatusForbidden, "Failed to check role")
	}

	if existingRole.CompanyID != nil && *existingRole.CompanyID != *employee.CompanyID && role.Name != "owner" {
		return pkg.Error(c, fiber.StatusForbidden, "You are not allowed to update this role")
	}

	var input map[string]interface{}
	if err := c.BodyParser(&input); err != nil {
		return pkg.Error(c, fiber.StatusBadRequest, "Invalid JSON body")
	}

	if name, ok := input["name"].(string); ok && name != "" {
		existingRole.Name = name
	}
	existingRole.ModifyBy = actorID

	if err := h.service.Update(existingRole, actorID); err != nil {
		return pkg.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return pkg.Success(c, fiber.Map{
		"uuid":         existingRole.UUID,
		"name":         existingRole.Name,
		"company_uuid": existingRole.CompanyID,
	}, "Role updated successfully")
}

func (h *RoleHandler) Delete(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	if uuid == "" {
		return pkg.Error(c, fiber.StatusBadRequest, "UUID is required")
	}

	userID, ok := c.Locals("user_id").(uint)
	if !ok {
		return pkg.Error(c, fiber.StatusUnauthorized, "Unauthorized: user_id not found")
	}

	employee, err := h.employeeRepo.FindByUserID(userID)
	if err != nil || employee.CompanyID == nil {
		return pkg.Error(c, fiber.StatusForbidden, "Failed to determine your company")
	}

	role, err := h.roleRepo.FindByID(employee.RoleID)
	if err != nil {
		return pkg.Error(c, fiber.StatusForbidden, "Failed to check user role")
	}

	if role.Name != "Owner" {
		if err := h.service.Delete(uuid, *employee.CompanyID); err != nil {
			return pkg.Error(c, fiber.StatusForbidden, err.Error())
		}
	} else {
		if err := h.service.Delete(uuid, 0); err != nil {
			return pkg.Error(c, fiber.StatusInternalServerError, err.Error())
		}
	}

	return pkg.Success(c, nil, "Role deleted successfully")
}
