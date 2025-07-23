package handlers

import (
	"hris_backend/internal/request"
	"hris_backend/internal/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type UserFamilyHandler interface {
	Create(c *fiber.Ctx) error
	GetAll(c *fiber.Ctx) error
	GetByID(c *fiber.Ctx) error
	GetByUUID(c *fiber.Ctx) error
	GetByUserUUID(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
}

type userFamilyHandler struct {
	service services.UserFamilyService
}

func NewUserFamilyHandler(service services.UserFamilyService) UserFamilyHandler {
	return &userFamilyHandler{service}
}

func (h *userFamilyHandler) Create(c *fiber.Ctx) error {
	var req request.UserFamilyRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid request body",
		})
	}

	actorID := c.Locals("user_id").(uint)

	res, err := h.service.Create(req, actorID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status": "success",
		"data":   res,
	})
}

func (h *userFamilyHandler) GetAll(c *fiber.Ctx) error {
	res, err := h.service.GetAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status": "success",
		"data":   res,
	})
}

func (h *userFamilyHandler) GetByID(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid ID",
		})
	}

	res, err := h.service.GetByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status": "success",
		"data":   res,
	})
}

func (h *userFamilyHandler) GetByUUID(c *fiber.Ctx) error {
	uuid := c.Params("uuid")

	res, err := h.service.GetByUUID(uuid)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status": "success",
		"data":   res,
	})
}

func (h *userFamilyHandler) GetByUserUUID(c *fiber.Ctx) error {
	uuid := c.Params("uuid")

	res, err := h.service.GetByUserUUID(uuid)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status": "success",
		"data":   res,
	})
}

func (h *userFamilyHandler) Update(c *fiber.Ctx) error {
	var req request.UserFamilyRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid request body",
		})
	}
	
	useruUUID := c.Params("user_uuid")
	actorID := c.Locals("user_id").(uint)

	res, err := h.service.Update(useruUUID, req, actorID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status": "success",
		"data":   res,
	})
}

func (h *userFamilyHandler) Delete(c *fiber.Ctx) error {
	uuid := c.Params("uuid")

	res, err := h.service.Delete(uuid)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status": "success",
		"data":   res,
	})
}
