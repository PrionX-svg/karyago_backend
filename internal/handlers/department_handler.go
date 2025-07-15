package handlers

import (
	"fmt"

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
