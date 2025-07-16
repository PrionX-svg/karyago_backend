package handlers

import (
	"fmt"
	"hris_backend/internal/request"
	"hris_backend/internal/services"
	"hris_backend/pkg"

	"github.com/gofiber/fiber/v2"
)

type BranchHandler struct {
	service services.BranchService
}

func NewBranchHandler(s services.BranchService) *BranchHandler {
	return &BranchHandler{s}
}

func (h *BranchHandler) Create(c *fiber.Ctx) error {
	var req request.BranchReq

	actorID, ok := c.Locals("user_id").(uint)
	if !ok {
		fmt.Println("Invalid user_id:", c.Locals("user_id"))
		return pkg.Error(c, fiber.StatusUnauthorized, "Unauthorized")
	}

	if err := c.BodyParser(&req); err != nil {
		return pkg.Error(c, fiber.StatusBadRequest, "Invalid request body")
	}

	branch, err := h.service.Create(req, actorID)
	if err != nil {
		return pkg.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return pkg.Created(c, branch, "Branch created")
}

func (h *BranchHandler) Get(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	branch, err := h.service.Get(uuid)
	if err != nil {
		return pkg.Error(c, fiber.StatusNotFound, "Branch not found")
	}
	return pkg.Success(c, branch, "Successfully get branch")
}

func (h *BranchHandler) GetByCompanyUUID(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	branch, err := h.service.GetByCompanyUUID(uuid)
	if err != nil {
		return pkg.Error(c, fiber.StatusNotFound, "Branch not found")
	}
	return pkg.Success(c, branch, "Successfully get branch by company UUID")
}

func (h *BranchHandler) Update(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	actorID, ok := c.Locals("user_id").(uint)
	if !ok {
		fmt.Println("Invalid user_id:", c.Locals("user_id"))
		return pkg.Error(c, fiber.StatusUnauthorized, "Unauthorized")
	}

	var req request.BranchReq
	if err := c.BodyParser(&req); err != nil {
		return pkg.Error(c, fiber.StatusBadRequest, "Invalid request body")
	}

	branch, err := h.service.Update(uuid, req, actorID)
	if err != nil {
		return pkg.Error(c, fiber.StatusInternalServerError, err.Error())
	}
	return pkg.Success(c, branch, "Branch updated")
}

func (h *BranchHandler) Delete(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	branch, err := h.service.Delete(uuid)
	if err != nil {
		return pkg.Error(c, fiber.StatusInternalServerError, err.Error())
	}
	return pkg.Success(c, branch, "Branch deleted")
}

func (h *BranchHandler) List(c *fiber.Ctx) error {
	branches, err := h.service.List()
	if err != nil {
		return pkg.Error(c, fiber.StatusInternalServerError, err.Error())
	}
	return pkg.Success(c, branches, "Successfully get all branches")
}
