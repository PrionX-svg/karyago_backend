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

type UserExperienceHandler struct {
	service services.UserExperienceService
}

func NewUserExperienceHandler(service services.UserExperienceService) *UserExperienceHandler {
	return &UserExperienceHandler{
		service: service,
	}
}

func (h *UserExperienceHandler) Create(c *fiber.Ctx) error {
	var req request.UserExperienceReq

	userID, ok := c.Locals("user_id").(uint)
	if !ok {
		fmt.Println("Invalid user_id:", c.Locals("user_id"))
		return pkg.Error(c, fiber.StatusUnauthorized, "Unauthorized")
	}

	if err := c.BodyParser(&req); err != nil {
		return pkg.Error(c, fiber.StatusBadRequest, "Invalid request body")
	}

	userExperience, err := h.service.Create(req, userID)
	if err != nil {

		if err.Error() == "invalid user UUID" {
			return pkg.Error(c, fiber.StatusBadRequest, err.Error())
		}
	
		return pkg.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return pkg.Created(c, userExperience, "User experience created successfully")
}

func (h *UserExperienceHandler) GetAll(c *fiber.Ctx) error {
	userEducations, err := h.service.GetAll()
	if err != nil {
		return pkg.Error(c, fiber.StatusInternalServerError, "Failed to retrieve user experience")
	}

	return pkg.Success(c, userEducations, "Successfully retrieved user experience")
}

func (h *UserExperienceHandler) GetByID(c *fiber.Ctx) error {
	idStr := c.Params("id")
	if idStr == "" {
		return pkg.Error(c, fiber.StatusBadRequest, "ID is required")
	}

	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return pkg.Error(c, fiber.StatusBadRequest, "Invalid ID format")
	}
	
	userExperience, err := h.service.GetByID(uint(id))
	if err != nil {
		return pkg.Error(c, fiber.StatusInternalServerError, "Failed to retrieve user experience")
	}
	if (userExperience == response.UserExperienceResponse{}) {
		return pkg.Error(c, fiber.StatusNotFound, "User experience not found")
	}
	return pkg.Success(c, userExperience, "Successfully retrieved user experience")
}

func (h *UserExperienceHandler) GetByUserUUID(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	if uuid == "" {
		return pkg.Error(c, fiber.StatusBadRequest, "UUID is required")
	}

	userEducations, err := h.service.GetByUserUUID(uuid)
	if err != nil {
		return pkg.Error(c, fiber.StatusInternalServerError, "Failed to retrieve user experience")
	}

	if len(userEducations) == 0 {
		return pkg.Error(c, fiber.StatusNotFound, "User experience not found")
	}

	return pkg.Success(c, userEducations, "Successfully retrieved user experience")
}

func (h *UserExperienceHandler) GetByUUID(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	if uuid == "" {
		return pkg.Error(c, fiber.StatusBadRequest, "UUID is required")
	}

	userExperience, err := h.service.GetByUUID(uuid)
	if err != nil {
		return pkg.Error(c, fiber.StatusInternalServerError, "Failed to retrieve user experience")
	}

	if (userExperience == response.UserExperienceResponse{}) {
		return pkg.Error(c, fiber.StatusNotFound, "User experience not found")
	}

	return pkg.Success(c, userExperience, "Successfully retrieved user experience")
}

func (h *UserExperienceHandler) Update(c *fiber.Ctx) error {
	var req request.UserExperienceReq

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

	userExperience, err := h.service.Update(uuid, req, userID)
	if err != nil {
		return pkg.Error(c, fiber.StatusInternalServerError, "Failed to update user experience")
	}

	return pkg.Success(c, userExperience, "User experience updated successfully")
}

func (h *UserExperienceHandler) Delete(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	if uuid == "" {
		return pkg.Error(c, fiber.StatusBadRequest, "UUID is required")
	}

	userExperience, err := h.service.Delete(uuid)
	if err != nil {
		return pkg.Error(c, fiber.StatusInternalServerError, "Failed to delete user experience")
	}

	return pkg.Success(c, userExperience, "User experience deleted successfully")
}
 