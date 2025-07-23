package handlers

import (
	"fmt"
	"hris_backend/internal/request"
	"hris_backend/internal/response"
	"hris_backend/internal/services"
	"hris_backend/pkg"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type UserBankHandler struct {
	service services.UserBankService
}

func NewUserBankHandler(service services.UserBankService) *UserBankHandler {
	return &UserBankHandler{
		service: service,
	}
}

func (h *UserBankHandler) Create(c *fiber.Ctx) error {
	var req request.UserBankReq

	userID, ok := c.Locals("user_id").(uint)
	if !ok {
		fmt.Println("Invalid user_id:", c.Locals("user_id"))
		return pkg.Error(c, fiber.StatusUnauthorized, "Unauthorized")
	}

	if err := c.BodyParser(&req); err != nil {
		return pkg.Error(c, fiber.StatusBadRequest, "Invalid request body")
	}

	userBank, err := h.service.Create(req, userID)
	if err != nil {
		if err.Error() == "invalid user UUID" {
			return pkg.Error(c, fiber.StatusBadRequest, err.Error())
		}
		return pkg.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return pkg.Created(c, userBank, "User bank created successfully")
}

func (h *UserBankHandler) GetAll(c *fiber.Ctx) error {
	userBanks, err := h.service.GetAll()
	if err != nil {
		return pkg.Error(c, fiber.StatusInternalServerError, "Failed to retrieve user banks")
	}

	return pkg.Success(c, userBanks, "Successfully retrieved user banks")
}

func (h *UserBankHandler) GetByID(c *fiber.Ctx) error {
	idStr := c.Params("id")
	if idStr == "" {
		return pkg.Error(c, fiber.StatusBadRequest, "ID is required")
	}

	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return pkg.Error(c, fiber.StatusBadRequest, "Invalid ID format")
	}

	userBank, err := h.service.GetByID(uint(id))
	if err != nil {
		return pkg.Error(c, fiber.StatusInternalServerError, "Failed to retrieve user bank")
	}
	if (userBank == response.UserBankResponse{}) {
		return pkg.Error(c, fiber.StatusNotFound, "User bank not found")
	}
	return pkg.Success(c, userBank, "Successfully retrieved user bank")
}

func (h *UserBankHandler) GetByUserUUID(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	if uuid == "" {
		return pkg.Error(c, fiber.StatusBadRequest, "UUID is required")
	}

	userBanks, err := h.service.GetByUserUUID(uuid)
	if err != nil {
		return pkg.Error(c, fiber.StatusInternalServerError, "Failed to retrieve user banks")
	}

	if len(userBanks) == 0 {
		return pkg.Error(c, fiber.StatusNotFound, "User bank not found")
	}

	return pkg.Success(c, userBanks, "Successfully retrieved user banks")
}

func (h *UserBankHandler) GetByUUID(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	if uuid == "" {
		return pkg.Error(c, fiber.StatusBadRequest, "UUID is required")
	}

	userBank, err := h.service.GetByUUID(uuid)
	if err != nil {
		return pkg.Error(c, fiber.StatusInternalServerError, "Failed to retrieve user bank")
	}

	if (userBank == response.UserBankResponse{}) {
		return pkg.Error(c, fiber.StatusNotFound, "User bank not found")
	}

	return pkg.Success(c, userBank, "Successfully retrieved user bank")
}

func (h *UserBankHandler) Update(c *fiber.Ctx) error {
	var req request.UserBankReq

	uuid := c.Params("uuid")
	if uuid == "" {
		return pkg.Error(c, fiber.StatusBadRequest, "UUID is required")
	}

	userID, ok := c.Locals("user_id").(uint)
	if !ok {
		fmt.Println("Invalid user_id:", c.Locals("user_id"))
		return pkg.Error(c, fiber.StatusUnauthorized, "Unauthorized")
	}

	if err := c.BodyParser(&req); err != nil {
		return pkg.Error(c, fiber.StatusBadRequest, "Invalid request body")
	}

	userBank, err := h.service.Update(uuid, req, userID)
	if err != nil {
		return pkg.Error(c, fiber.StatusInternalServerError, "Failed to update user bank")
	}

	return pkg.Success(c, userBank, "User bank updated successfully")
}

func (h *UserBankHandler) Delete(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	if uuid == "" {
		return pkg.Error(c, fiber.StatusBadRequest, "UUID is required")
	}

	userBank, err := h.service.Delete(uuid)
	if err != nil {
		return pkg.Error(c, fiber.StatusInternalServerError, "Failed to delete user bank")
	}

	return pkg.Success(c, userBank, "User bank deleted successfully")
}
