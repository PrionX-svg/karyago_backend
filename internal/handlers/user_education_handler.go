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

type UserEducationHandler struct {
	service services.UserEducationService
}

func NewUserEducationHandler(service services.UserEducationService) *UserEducationHandler {
	return &UserEducationHandler{
		service: service,
	}
}

func (h *UserEducationHandler) Create(c *fiber.Ctx) error {
	var req request.UserEducationReq

	userID, ok := c.Locals("user_id").(uint)
	if !ok {
		fmt.Println("Invalid user_id:", c.Locals("user_id"))
		return pkg.Error(c, fiber.StatusUnauthorized, "Unauthorized")
	}

	if err := c.BodyParser(&req); err != nil {
		return pkg.Error(c, fiber.StatusBadRequest, "Invalid request body")
	}

	userEducation, err := h.service.Create(req, userID); 
	if err != nil {

		if err.Error() == "invalid user UUID" {
			return pkg.Error(c, fiber.StatusBadRequest, err.Error())
		}
	
		return pkg.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return pkg.Created(c, userEducation, "User education created successfully")
}

func (h *UserEducationHandler) GetAll(c *fiber.Ctx) error {
	userEducations, err := h.service.GetAll()
	if err != nil {
		return pkg.Error(c, fiber.StatusInternalServerError, "Failed to retrieve user education")
	}

	return pkg.Success(c, userEducations, "Successfully retrieved user education")
}

func (h *UserEducationHandler) GetByID(c *fiber.Ctx) error {
	idStr := c.Params("id")
	if idStr == "" {
		return pkg.Error(c, fiber.StatusBadRequest, "ID is required")
	}

	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return pkg.Error(c, fiber.StatusBadRequest, "Invalid ID format")
	}
	
	userEducation, err := h.service.GetByID(uint(id))
	if err != nil {
		return pkg.Error(c, fiber.StatusInternalServerError, "Failed to retrieve user education")
	}
	if (userEducation == response.UserEducationResponse{}) {
		return pkg.Error(c, fiber.StatusNotFound, "User education not found")
	}
	return pkg.Success(c, userEducation, "Successfully retrieved user education")
}

func (h *UserEducationHandler) GetByUserUUID(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	if uuid == "" {
		return pkg.Error(c, fiber.StatusBadRequest, "UUID is required")
	}

	userEducations, err := h.service.GetByUserUUID(uuid)
	if err != nil {
		return pkg.Error(c, fiber.StatusInternalServerError, "Failed to retrieve user education")
	}

	if len(userEducations) == 0 {
		return pkg.Error(c, fiber.StatusNotFound, "User education not found")
	}

	return pkg.Success(c, userEducations, "Successfully retrieved user education")
}

func (h *UserEducationHandler) GetByUUID(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	if uuid == "" {
		return pkg.Error(c, fiber.StatusBadRequest, "UUID is required")
	}

	userEducation, err := h.service.GetByUUID(uuid)
	if err != nil {
		return pkg.Error(c, fiber.StatusInternalServerError, "Failed to retrieve user education")
	}

	if (userEducation == response.UserEducationResponse{}) {
		return pkg.Error(c, fiber.StatusNotFound, "User education not found")
	}

	return pkg.Success(c, userEducation, "Successfully retrieved user education")
}

func (h *UserEducationHandler) Update(c *fiber.Ctx) error {
	var req request.UserEducationReq

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

	userEducation, err := h.service.Update(uuid, req, userID)
	if err != nil {
		return pkg.Error(c, fiber.StatusInternalServerError, "Failed to update user education")
	}

	return pkg.Success(c, userEducation, "User education updated successfully")
}

func (h *UserEducationHandler) Delete(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	if uuid == "" {
		return pkg.Error(c, fiber.StatusBadRequest, "UUID is required")
	}

	userEducation, err := h.service.Delete(uuid)
	if err != nil {
		return pkg.Error(c, fiber.StatusInternalServerError, "Failed to delete user education")
	}

	return pkg.Success(c, userEducation, "User education deleted successfully")
}
 