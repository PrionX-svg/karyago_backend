package handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"hris_backend/internal/request"
	"hris_backend/internal/services"
	"hris_backend/pkg"
)

type UserDetailHandler struct {
	service services.UserDetailService
}

func NewUserDetailHandler(s services.UserDetailService) *UserDetailHandler {
	return &UserDetailHandler{s}
}

func (h *UserDetailHandler) Create(c *fiber.Ctx) error {
	var req request.UserDetailReq

	actorID, ok := c.Locals("user_id").(uint)
	if !ok {
		fmt.Println("Invalid user_id:", c.Locals("user_id"))
		return pkg.Error(c, fiber.StatusUnauthorized, "Unauthorized")
	}

	if err := c.BodyParser(&req); err != nil {
		return pkg.Error(c, fiber.StatusBadRequest, "Invalid request body")
	}

	detail, err := h.service.Create(req, actorID)
	if err != nil {
		return pkg.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return pkg.Created(c, detail, "User detail created")
}

func (h *UserDetailHandler) Get(c *fiber.Ctx) error {
	uuid := c.Params("uuid")

	detail, err := h.service.Get(uuid)
	if err != nil {
		return pkg.Error(c, fiber.StatusNotFound, "User detail not found")
	}

	return pkg.Success(c, detail, "Successfully retrieved user detail")
}

func (h *UserDetailHandler) Update(c *fiber.Ctx) error {
	uuid := c.Params("uuid")

	actorID, ok := c.Locals("user_id").(uint)
	if !ok {
		fmt.Println("Invalid user_id:", c.Locals("user_id"))
		return pkg.Error(c, fiber.StatusUnauthorized, "Unauthorized")
	}

	var req request.UserDetailReq
	if err := c.BodyParser(&req); err != nil {
		return pkg.Error(c, fiber.StatusBadRequest, "Invalid request body")
	}

	detail, err := h.service.Update(uuid, req, actorID)
	if err != nil {
		return pkg.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return pkg.Success(c, detail, "User detail updated")
}

func (h *UserDetailHandler) Delete(c *fiber.Ctx) error {
	uuid := c.Params("uuid")

	detail, err := h.service.Delete(uuid)
	if err != nil {
		return pkg.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return pkg.Success(c, detail, "User detail deleted")
}
